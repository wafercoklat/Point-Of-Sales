package ar_salesorder

import (
	"STACK-GO/adapter"
	domain "STACK-GO/modules/ar_salesorder/domain"
	repo "STACK-GO/modules/ar_salesorder/repo"
	"fmt"
	"strconv"
)

type Services struct {
	adapter adapter.RepoAdapter
}

func NewService(port adapter.RepoAdapter) *Services {
	return &Services{
		adapter: port,
	}
}

func (s *Services) FindByID(id string) (*domain.SalesorderPackage, error) {
	var sopackage domain.SalesorderPackage

	// Get the SalesOrder
	salesorder, err := s.adapter.FindByID(id, &domain.Salesorder{}, repo.QryFindByID_Header())
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	// Get the SalesOrder Details
	salesorderdetail, err := s.adapter.FindByID(id, &domain.SalesorderDetails{}, repo.QryFindByID_Detail())
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	sopackage.Salesorder = salesorder.(domain.Salesorder)
	sopackage.SalesorderDetails = salesorderdetail.([]domain.SalesorderDetails)

	return &sopackage, nil
}

func (s *Services) List() (*[]domain.Salesorder, error) {
	// Get List SalesOrder
	data, err := s.adapter.List(&domain.List, repo.QryList())
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return data.(*[]domain.Salesorder), nil
}

func (s *Services) Create(data domain.SalesorderPackage) (int, error) {
	// Create SalesOrder
	headerID, err := s.adapter.Create(data.Salesorder, repo.QryCreate_Header())
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	// Get SOID
	for i := range data.SalesorderDetails {
		data.SalesorderDetails[i].SOID = int(headerID)
	}

	// Update SalesOrderDetail
	// Have to create pre-create check
	_, err = s.adapter.Create(data.SalesorderDetails, repo.QryCreate_Detail())
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	return 1, err
}

func (s *Services) Update(id string, data domain.SalesorderPackage) (int, error) {
	// Must to check if sales order have another transcation
	// Update SalesOrder
	_, err := s.adapter.Update(data.Salesorder, id, repo.QryUpdate_Header())
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	// Update SalesOrderDetail
	// Have to create pre-update check
	for i := range data.SalesorderDetails {
		_, err := s.adapter.Update(data.SalesorderDetails[i], strconv.Itoa(data.SalesorderDetails[i].DetailID), repo.QryCreate_Detail())
		if err != nil {
			fmt.Println(err)
			return 0, err
		}
	}

	return 1, err
}

func (s *Services) Delete(id string) (int, error) {
	var err error
	// Must to check if sales order have another transcation
	_, err = s.adapter.Delete(id, repo.QryDelete_Header())
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	_, err = s.adapter.Delete(id, repo.QryDelete_Detail())
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	return 1, err
}
