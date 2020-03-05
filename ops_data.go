package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

var home = os.Getenv("HOME")
var workingDirectory, _ = os.Getwd()

func main() {
	checkAppStatus()
	exitApp()
}

func checkAppStatus() {
	skip := false
	//check for cloud-sdk and python install.
	cloudSdkAvailable := isCommandAvailable("gcloud")
	if cloudSdkAvailable == false {
		//installCloudSdk() // It's just not worth it
		fmt.Printf("\n Google-cloud-sdk not found on system. Install from https://cloud.google.com/sdk/install")
		os.Exit(1)
	}
	pythonAvailable := isCommandAvailable("python3")
	if pythonAvailable == false {
		//if its windows it will be named python
		if runtime.GOOS == "windows" {
			fmt.Println("Even though we checked for python it still failed. Ignore!")
			if isCommandAvailable("python") == true {
				skip = true
			}
		}
		if !skip {
			fmt.Printf("Python not found on system. Install Python version > 3")
			os.Exit(1)
		}
	}
	makeVirtualEnv()

	cmd := "venv/"
	if runtime.GOOS == "darwin" || runtime.GOOS == "linux" {
		cmd += "bin/jupyter"
	} else {
		cmd += "Scripts/jupyter"
	}
	startJupyter(cmd)
}

func exitApp() {
	fmt.Println("Run the following commands from ", workingDirectory)
	fmt.Println()
	if runtime.GOOS == "windows" {
		fmt.Println("source venv/Scripts/activate")
	} else {
		fmt.Println("source venv/bin/activate")
	}
	fmt.Println("jupyter notebook ops_data_api.ipynb")
	fmt.Println("when you're done working, type deactivate")
	os.Exit(0)
}

func makeVirtualEnv() {
	var args []string
	err := os.Chdir(workingDirectory)
	if err != nil {
		err := os.MkdirAll(workingDirectory, 0777)
		if err != nil {
			fmt.Printf("Can not create directory. Try creating %s directory from home folder then run again. \n", workingDirectory)
			os.Exit(1)
		}
	}
	python := "python"
	venv := workingDirectory + "/venv"
	if _, err := os.Stat(venv); os.IsNotExist(err) {
		fmt.Println("creating virtual environment")
		//venv will create parent directories
		args = append(args, "-m", "venv", venv)
		if runtime.GOOS == "darwin" || runtime.GOOS == "linux" {
			python += "3"
		}
		runCommands(python, args)
		//install project requirements using the explicit path to pip. No need to "activate".
		pip := filepath.FromSlash(workingDirectory + "/venv/")
		if runtime.GOOS == "darwin" || runtime.GOOS == "linux" {
			pip += "bin/pip3"
		} else {
			pip += "Scripts/pip3"
		}
		pipArgs := []string{"install", "-r", "requirements.txt"}
		fmt.Printf("Installing requirements in %s/venv \n", workingDirectory)
		runCommands(pip, pipArgs)
	}
}

func runCommands(shCommand string, args []string) {
	cmd := exec.Command(shCommand, args...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return
	}
	fmt.Println("Success")
}
