package svc

type mockProgram struct {
	start func() error
	stop  func() error
	init  func(Environment) error
}

func (p *mockProgram) Start() error {
	return p.start()
}

func (p *mockProgram) Stop() error {
	return p.stop()
}
func (p *mockProgram) Init(env Environment) error {
	return p.init(env)
}

func makeProgram(startCalled, stopCalled, initCalled *int) *mockProgram {
	return &mockProgram{
		start: func() error {
			*startCalled++
			return nil
		},
		stop: func() error {
			*stopCalled++
			return nil
		},
		init: func(env Environment) error {
			*initCalled++
			return nil
		},
	}
}
