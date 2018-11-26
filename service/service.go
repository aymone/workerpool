package service

type (
	// Entity ...
	Entity struct {
		ID int
	}

	service struct{}
)

// JobService interface
type JobService interface {
	Get(ID int) (*Entity, error)
}

// New service
func New() *service {
	return &service{}
}

func (s *service) Get(ID int) (*Entity, error) {
	return &Entity{ID}, nil
}
