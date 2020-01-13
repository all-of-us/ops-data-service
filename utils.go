package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

func ungzip(source, target string) error {
	reader, err := os.Open(source)
	if err != nil {
		return err
	}
	defer reader.Close()

	archive, err := gzip.NewReader(reader)
	if err != nil {
		return err
	}
	defer archive.Close()

	target = filepath.Join(target, archive.Name)
	fmt.Println("target is : " + target)
	writer, err := os.Create(target)
	if err != nil {
		return err
	}
	defer writer.Close()

	_, err = io.Copy(writer, archive)
	return err
}

func downloadFromUrl(url string) {
	tokens := strings.Split(url, "/")
	fileName := tokens[len(tokens)-1]
	fmt.Println("Downloading", url, "to", fileName)

	output, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error while creating", fileName, "-", err)
		return
	}
	defer output.Close()

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error while downloading", url, "-", err)
		return
	}
	defer response.Body.Close()

	n, err := io.Copy(output, response.Body)
	if err != nil {
		fmt.Println("Error while downloading", url, "-", err)
		return
	}

	fmt.Println(n, "bytes downloaded.")

	wd, _ := os.Getwd()
	fmt.Println("Working directory: " + wd)
	err = ungzip(fileName, wd)
	if err != nil {
		fmt.Println(err)
		fmt.Printf("Can not install Google cloud sdk, try installing it and run this script again. \n")
	}
}

func isCommandAvailable(name string) bool {
	//checks if the given command is available in path
	platform := runtime.GOOS
	switch {
	case platform == "darwin" || platform == "linux":
		cmd := exec.Command("/bin/sh", "-c", "command -v "+name)
		if err := cmd.Run(); err != nil {
			fmt.Printf("%s not found.", name)
			return false
		}
		return true
	case platform == "windows":
		if name == "python3" {
			name = "python"
		}
		cmd := exec.Command("where", name)
		if err := cmd.Run(); err != nil {
			fmt.Printf("%s not found.", name)
			return false
		}
		return true
	}
	//default
	return false
}

func startJupyter(cmd string) {
	//cmd := "venv/bin/jupyter"
	//Can not pass in ipynb notebook directly in jupyter notebook == 5.7.4 on python 2
	args := []string{"notebook", "ops_data_api.ipynb"}
	fmt.Println("Running jupyter server in browser. Press Ctrl + C to quit.")
	runCommands(cmd, args)
}

func setupGcloud() {
	_ = os.Chdir("HOME")
	shellCommand := ""
	param := ""
	var args []string
	var filetype string
	if runtime.GOOS == "windows" {
		filetype = ".bat"
		shellCommand = "cmd"
		param = "/C"
	} else {
		filetype = ".sh"
		shellCommand = "/bin/sh"
	}
	// may need to also pass in /C for windows. @TODO
	argString := filepath.FromSlash("./google-cloud-sdk/install" + filetype)
	args = append(args, param, argString)
	fmt.Printf("Attempting to install Google cloud sdk...")
	runCommands(shellCommand, args)
	init := []string{param, filepath.FromSlash("./google-cloud-sdk/bin/gcloud"), "init"}
	fmt.Printf("Attempting to initialize Google cloud sdk in %s", home)
	runCommands(shellCommand, init)
}

func installCloudSdk() {
	var url string
	var cloud_sdk_url = "https://dl.google.com/dl/cloudsdk/channels/rapid/downloads/google-cloud-sdk-233.0.0-"

	fmt.Printf("Installing google-cloud-sdk in home directory. \n")

	platform := runtime.GOOS
	if platform == "darwin" {
		url = cloud_sdk_url + "darwin-x86_64.tar.gz"
	}
	if platform == "linux" {
		url = cloud_sdk_url + "linux-x86_64.tar.gz"
	}
	if platform == "windows" {
		//windows 64 bit
		url = cloud_sdk_url + "windows-x86_64.zip"
	}
	downloadFromUrl(url)
	setupGcloud()
}
