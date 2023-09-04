package main

import (
	"github.com/rogpeppe/go-internal/testenv"
	"log"
	"os"
	"os/exec"
	"syscall"
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

func TestErrProcessDone(t *testing.T) {
	testenv.MustHaveGoBuild(t)
	path, err := testenv.GoTool()
	if err != nil {
		t.Errorf("finding go tool: %v", err)
	}
	p, err := os.StartProcess(path, []string{"go"}, &os.ProcAttr{})
	if err != nil {
		t.Errorf("starting test process: %v", err)
	}
	p.Wait()
	if got := p.Signal(syscall.SIGKILL); got != os.ErrProcessDone {
		t.Errorf("got %v want %v", got, os.ErrProcessDone)
	}
}
