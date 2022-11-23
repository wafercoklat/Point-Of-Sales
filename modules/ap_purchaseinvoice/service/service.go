package ap_purchaseinvoice

import (
	"STACK-GO/adapter"
	domain "STACK-GO/modules/ap_purchaseinvoice/domain"
	repo "STACK-GO/modules/ap_purchaseinvoice/repo"
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

func (s *Services) FindByID(id string) (*domain.PurchaseinvoicePackage, error) {
	var pipackage domain.PurchaseinvoicePackage

	// Get the Purchaseinvoice
	purchaseinvoice, err := s.adapter.FindByID(id, &domain.Purchaseinvoice{}, repo.QryFindByID_Header())
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	// Get the Purchaseinvoice Details
	purchaseinvoicedetail, err := s.adapter.FindByID(id, &domain.PurchaseinvoiceDetails{}, repo.QryFindByID_Detail())
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	pipackage.Purchaseinvoice = purchaseinvoice.(domain.Purchaseinvoice)
	pipackage.PurchaseinvoiceDetails = purchaseinvoicedetail.([]domain.PurchaseinvoiceDetails)

	return &pipackage, nil
}

func (s *Services) List() (*[]domain.Purchaseinvoice, error) {
	// Get List Purchaseinvoice
	data, err := s.adapter.List(&domain.List, repo.QryList())
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return data.(*[]domain.Purchaseinvoice), nil
}

func (s *Services) Create(data domain.PurchaseinvoicePackage) (int, error) {
	// Create Purchaseinvoice
	headerID, err := s.adapter.Create(data.Purchaseinvoice, repo.QryCreate_Header())
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	// Get SOID
	for i := range data.PurchaseinvoiceDetails {
		data.PurchaseinvoiceDetails[i].SOID = int(headerID)
	}

	// Update PurchaseinvoiceDetail
	// Have to create pre-create check
	_, err = s.adapter.Create(data.PurchaseinvoiceDetails, repo.QryCreate_Detail())
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	return 1, err
}

func (s *Services) Update(id string, data domain.PurchaseinvoicePackage) (int, error) {
	// Must to check if sales order have another transcation
	// Update Purchaseinvoice
	_, err := s.adapter.Update(data.Purchaseinvoice, id, repo.QryUpdate_Header())
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	// Update PurchaseinvoiceDetail
	// Have to create pre-update check
	for i := range data.PurchaseinvoiceDetails {
		_, err := s.adapter.Update(data.PurchaseinvoiceDetails[i], strconv.Itoa(data.PurchaseinvoiceDetails[i].DetailID), repo.QryCreate_Detail())
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
