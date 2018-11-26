package service

// Mock is a fake implementation of service
type Mock struct {
	GetFn        func(ID int) (*Entity, error)
	GetFnInvoked bool
}

// Get mock
func (m *Mock) Get(ID int) (*Entity, error) {
	m.GetFnInvoked = true
	return &Entity{ID}, nil
}
