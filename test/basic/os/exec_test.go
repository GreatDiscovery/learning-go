package main

import (
	"log"
	"os/exec"
	"testing"
)

// example from os/exec/example_test.go

func TestCommand(t *testing.T) {
	cmd := exec.Command("sleep", "1")
	log.Printf("Running command and waiting for it to finish...")
	err := cmd.Run()
	log.Printf("Command finished with error: %v", err)

	cmd = exec.Command("sleep", "5")
	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Waiting for command to finish...")
	err = cmd.Wait()
	log.Printf("Command finished with error: %v", err)
}
