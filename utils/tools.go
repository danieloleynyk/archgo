package utils

import (
	"bytes"
	"io"
	"log"
	"os"
	"os/exec"
)

// RunBashScript is a function that runs a bash script with go
func RunBashScript(scriptPath string) error {
	cmd := exec.Command("/bin/sh", scriptPath)
	var stdBuffer bytes.Buffer
	mw := io.MultiWriter(os.Stdout, &stdBuffer)

	cmd.Stdout = mw
	cmd.Stderr = mw

	if err := cmd.Run(); err != nil {
		return err
	}

	log.Println(stdBuffer.String())
	return nil
}

// GetScriptName returns the name of the script to use for the package
func GetScriptName(packageName, distribution string) (string, error) {
	return packageName + ".sh", nil
}

// Catch is used for catching errors
// err: The error
// message: A message to log if there is an error
func Catch(err error, message string) {
	if err != nil {
		log.Fatal(message)
		panic(message)
	}
}
