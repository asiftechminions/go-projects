package user

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{repo: r}
}

func (s *Service) GetUserName(id int) (string, error) {
	u, err := s.repo.FindByID(id)
	if err != nil {
		return "", err
	}
	return u.Name, nil
}
