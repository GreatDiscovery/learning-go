package service

import "testing"

// impl service_interface

type runner struct{}

func (*runner) Add() error {
	//TODO implement me
	panic("implement me")
}

func (*runner) Delete() error {
	//TODO implement me
	panic("implement me")
}

func (*runner) Update() error {
	//TODO implement me
	panic("implement me")
}

func (*runner) Query() error {
	//TODO implement me
	panic("implement me")
}

func TestInterface(t *testing.T) {
	runner := &runner{}
	runner.Query()
}
