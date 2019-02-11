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

	// TODO: check file existence first with io.IsExist
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

	err = ungzip(fileName, "~/code/playground/test") //@TODO: set path to home
	if err != nil {
		fmt.Printf("Error unzipping file: %s", err)
	}
}

func isCommandAvailable(name string) bool {
	//checks if the given command is available in path
	platform := runtime.GOOS
	switch {
	case platform == "darwin" || platform == "linux":
		cmd := exec.Command("/bin/sh", "-c", "command -v "+name)
		if err := cmd.Run(); err != nil {
			return false
		}
		return true
	case platform == "windows":
		cmd := exec.Command("where", name)
		if err := cmd.Run(); err != nil {
			return false
		}
		return true
	}
	//default
	return false
}
