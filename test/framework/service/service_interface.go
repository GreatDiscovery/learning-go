package service

//how to define service and impl

type Interface interface {
	Add() error
	Delete() error
	Update() error
	Query() error
}
