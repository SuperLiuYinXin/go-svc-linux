package svc

import (
	"os"
	"syscall"
	"testing"
)

func TestDefaultSignalHandling(t *testing.T) {

	signals := []os.Signal{syscall.SIGINT, syscall.SIGTERM}
	for _, signal := range signals {
		testSignalNotify(t, signal)
	}
}

func testSignalNotify(t *testing.T, signal os.Signal, sig ...os.Signal) {

	sigChan := make(chan os.Signal)

	var startCalled, stopCalled, initCalled int

	signalNotify = func(c chan<- os.Signal, sig ...os.Signal) {
		if c == nil {
			panic("os/signal: Notify using nil channel")
		}
		go func() {
			for val := range sigChan {
				for _, registeredSig := range sig {
					if val == registeredSig {
						c <- val
					}
				}
			}
		}()
	}

	prg := makeProgram(&startCalled, &stopCalled, &initCalled)

	go func() {
		sigChan <- signal
	}()

	if err := Run(prg, sig...); err != nil {
		t.Fatal(err)
	}

	if startCalled != 1 {
		t.Errorf("startCalled, want:1 got: %d", startCalled)
	}
	if stopCalled != 1 {
		t.Errorf("stopCalled, want:1 got: %d", stopCalled)
	}
	if initCalled != 1 {
		t.Errorf("initCalled, want:1 got: %d", initCalled)
	}
}
