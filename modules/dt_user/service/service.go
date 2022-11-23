package dt_user

import (
	"STACK-GO/adapter"
	domain "STACK-GO/modules/dt_user/domain"
	repo "STACK-GO/modules/dt_user/repo"
	"fmt"
)

type Services struct {
	adapter adapter.RepoAdapter
}

func NewService(port adapter.RepoAdapter) *Services {
	return &Services{
		adapter: port,
	}
}

func (s *Services) FindByID(id string) *domain.User {
	data, err := s.adapter.FindByID(id, &domain.User{}, repo.QryFindByID())
	if err != nil {
		fmt.Println(err)
	}
	return data.(*domain.User)
}

func (s *Services) List() *[]domain.User {
	data, err := s.adapter.List(&domain.List, repo.QryList())
	if err != nil {
		fmt.Println(err)
	}
	return data.(*[]domain.User)
}

func (s *Services) Create(data domain.User) error {
	_, err := s.adapter.Create(data, repo.QryCreate())
	return err
}

func (s *Services) Update(id string, data domain.User) error {
	_, err := s.adapter.Update(data, id, repo.QryUpdate())
	return err
}

func (s *Services) Delete(id string) error {
	_, err := s.adapter.Delete(id, repo.QryDelete())
	return err
}
