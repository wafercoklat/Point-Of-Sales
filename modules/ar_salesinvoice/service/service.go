package ar_salesinvoice

import (
	"STACK-GO/adapter"
	domain "STACK-GO/modules/ar_salesinvoice/domain"
	repo "STACK-GO/modules/ar_salesinvoice/repo"
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

func (s *Services) FindByID(id string) (*domain.SalesinvoicePackage, error) {
	var sipackage domain.SalesinvoicePackage

	// Get the Salesinvoice
	Salesinvoice, err := s.adapter.FindByID(id, &domain.Salesinvoice{}, repo.QryFindByID_Header())
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	// Get the Salesinvoice Details
	Salesinvoicedetail, err := s.adapter.FindByID(id, &domain.SalesinvoiceDetails{}, repo.QryFindByID_Detail())
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	sipackage.Salesinvoice = Salesinvoice.(domain.Salesinvoice)
	sipackage.SalesinvoiceDetails = Salesinvoicedetail.([]domain.SalesinvoiceDetails)

	return &sipackage, nil
}

func (s *Services) List() (*[]domain.Salesinvoice, error) {
	// Get List Salesinvoice
	data, err := s.adapter.List(&domain.List, repo.QryList())
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return data.(*[]domain.Salesinvoice), nil
}

func (s *Services) Create(data domain.SalesinvoicePackage) (int, error) {
	// Create Salesinvoice
	headerID, err := s.adapter.Create(data.Salesinvoice, repo.QryCreate_Header())
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	// Get SOID
	for i := range data.SalesinvoiceDetails {
		data.SalesinvoiceDetails[i].SOID = int(headerID)
	}

	// Update SalesinvoiceDetail
	// Have to create pre-create check
	_, err = s.adapter.Create(data.SalesinvoiceDetails, repo.QryCreate_Detail())
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	return 1, err
}

func (s *Services) Update(id string, data domain.SalesinvoicePackage) (int, error) {
	// Must to check if sales order have another transcation
	// Update Salesinvoice
	_, err := s.adapter.Update(data.Salesinvoice, id, repo.QryUpdate_Header())
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	// Update SalesinvoiceDetail
	// Have to create pre-update check
	for i := range data.SalesinvoiceDetails {
		_, err := s.adapter.Update(data.SalesinvoiceDetails[i], strconv.Itoa(data.SalesinvoiceDetails[i].DetailID), repo.QryCreate_Detail())
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
