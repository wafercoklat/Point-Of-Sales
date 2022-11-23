package iv_item

import (
	"STACK-GO/adapter"
	domain "STACK-GO/modules/iv_item/domain"
	repo "STACK-GO/modules/iv_item/repo"
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

func (s *Services) FindByID(id string) *domain.Item {
	data, err := s.adapter.FindByID(id, &domain.Item{}, repo.QryFindByID())
	if err != nil {
		fmt.Println(err)
	}
	return data.(*domain.Item)
}

func (s *Services) List() *[]domain.Item {
	data, err := s.adapter.List(&domain.List, repo.QryList())
	if err != nil {
		fmt.Println(err)
	}
	return data.(*[]domain.Item)
}

func (s *Services) Create(data domain.Item) error {
	_, err := s.adapter.Create(data, repo.QryCreate())
	return err
}

func (s *Services) Update(id string, data domain.Item) error {
	_, err := s.adapter.Update(data, id, repo.QryUpdate())
	return err
}

func (s *Services) Delete(id string) error {
	_, err := s.adapter.Delete(id, repo.QryDelete())
	return err
}
