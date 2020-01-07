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

//var opsDataPath = filepath.FromSlash(home + "/ops_data_api")

func main() {
	//check for cloud-sdk and python install.

	cloudSdkAvailable := isCommandAvailable("gcloud")
	if cloudSdkAvailable == false {
		installCloudSdk() //@TODO: FIX THE INSTALL AND RUN FUNCTIONS. UNTIL THEN, EXIT.
		//fmt.Printf("\n Google-cloud-sdk not found on system. Install from https://cloud.google.com/sdk/install")
		//os.Exit(1)
	}
	pythonAvailable := isCommandAvailable("python3")
	if pythonAvailable == false {
		fmt.Printf("Python not found on system. Install Python version > 3")
		os.Exit(1)
	}
	makeVirtualEnv()

	startJupyter() //TODO: GCLOUD COMMANDS CURRENTLY DON'T WORK UNLESS THE USER EXPLICITLY ACTIVATES VENV. AUTOMATE THIS.
	exitApp()

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
	python := "python3"
	venv := workingDirectory + "/venv"
	if _, err := os.Stat(venv); os.IsNotExist(err) {
		fmt.Println("creating virtual environment")
		//venv will create parent directories
		args = append(args, "-m", "venv", venv)
		//args = append(args, "venv")
		runCommands(python, args)
		//install project requirements using the explicit path to pip. No need to "activate".
		pip := filepath.FromSlash(workingDirectory + "/venv/bin/pip3")
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
	} else {
		fmt.Println("Success")
	}
}
