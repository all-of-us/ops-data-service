package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"runtime"
)

func main()  {
	//check for cloud-sdk install & pyenv
	//reqs := [2]string{"gcloud", "pyenv"}

	cloudSdkAvailable := isCommandAvailable("gcloud")
	if cloudSdkAvailable == false {
		installCloudSdk()
	}
	pyenvAvailable := isCommandAvailable("pyenv")
	if pyenvAvailable == false {
		installPyenv()
	}
	//check/create virtual env

	//pip install requirements

	//start jupyter notebook server
}

func installCloudSdk() {
	var url string
	var cloud_sdk_url = "https://dl.google.com/dl/cloudsdk/channels/rapid/downloads/google-cloud-sdk-233.0.0-"

	fmt.Printf("Installing google-cloud-sdk in home dir.")

	platform := runtime.GOOS
	if platform == "darwin" {
		cloud_sdk_url =  cloud_sdk_url + "darwin-x86_64.tar.gz"
	}
	if platform == "linux" {
		cloud_sdk_url = cloud_sdk_url + "linux-x86_64.tar.gz"
	}
	if platform == "darwin" || platform == "linux" {
		url = cloud_sdk_url
		downloadFromUrl(url)
		setupGcloudUnix()
	}
	if platform == "windows" {
		//windows 64 bit
		url = cloud_sdk_url + "windows-x86_64.zip"
		downloadFromUrl(url)
		setupGcloudWindows()
	}
}

func installPyenv() {
	platform := runtime.GOOS
	if platform == "darwin" || platform == "unix" {

	}
}


func runCommands(shCommand string, args string) {
	cmd := exec.Command(shCommand, args)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + out.String())
		return
	}
	fmt.Println("Setup result: " + out.String())
}

