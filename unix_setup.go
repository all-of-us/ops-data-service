package main

import "os"

func setupGcloudUnix() {
	_ = os.Chdir("~/")
	shellCommand := "/bin/sh"
	var args string
	args = "./google-cloud-sdk/install.sh"
	runCommands(shellCommand, args)
	args = "./google-cloud-sdk/bin/gcloud init"
	runCommands(shellCommand, args)
}

func setupJupyterUnix() {
}

