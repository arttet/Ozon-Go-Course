package subdomain

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) List() []Subdomain {
	return allEntities
}

func (s *Service) Get(idx int) (*Subdomain, error) {
	return &allEntities[idx], nil
}
