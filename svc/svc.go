package svc

import "os/signal"

var signalNotify = signal.Notify

type Service interface {
	//
	Init(Environment) error

	// start
	Start() error

	// end
	Stop() error
}

type Environment interface {
	IsWindowsService() bool
}
