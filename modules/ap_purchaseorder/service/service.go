package ap_purchaseorder

import (
	"STACK-GO/adapter"
	domain "STACK-GO/modules/ap_purchaseorder/domain"
	repo "STACK-GO/modules/ap_purchaseorder/repo"
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

func (s *Services) FindByID(id string) (*domain.PurchaseorderPackage, error) {
	var popackage domain.PurchaseorderPackage

	// Get the Purchaseorder
	purchaseorder, err := s.adapter.FindByID(id, &domain.Purchaseorder{}, repo.QryFindByID_Header())
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	// Get the Purchaseorder Details
	purchaseorderdetail, err := s.adapter.FindByID(id, &domain.PurchaseorderDetails{}, repo.QryFindByID_Detail())
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	popackage.Purchaseorder = purchaseorder.(domain.Purchaseorder)
	popackage.PurchaseorderDetails = purchaseorderdetail.([]domain.PurchaseorderDetails)

	return &popackage, nil
}

func (s *Services) List() (*[]domain.Purchaseorder, error) {
	// Get List Purchaseorder
	data, err := s.adapter.List(&domain.List, repo.QryList())
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return data.(*[]domain.Purchaseorder), nil
}

func (s *Services) Create(data domain.PurchaseorderPackage) (int, error) {
	// Create Purchaseorder
	headerID, err := s.adapter.Create(data.Purchaseorder, repo.QryCreate_Header())
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	// Get SOID
	for i := range data.PurchaseorderDetails {
		data.PurchaseorderDetails[i].SOID = int(headerID)
	}

	// Update PurchaseorderDetail
	// Have to create pre-create check
	_, err = s.adapter.Create(data.PurchaseorderDetails, repo.QryCreate_Detail())
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	return 1, err
}

func (s *Services) Update(id string, data domain.PurchaseorderPackage) (int, error) {
	// Must to check if sales order have another transcation
	// Update Purchaseorder
	_, err := s.adapter.Update(data.Purchaseorder, id, repo.QryUpdate_Header())
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	// Update PurchaseorderDetail
	// Have to create pre-update check
	for i := range data.PurchaseorderDetails {
		_, err := s.adapter.Update(data.PurchaseorderDetails[i], strconv.Itoa(data.PurchaseorderDetails[i].DetailID), repo.QryCreate_Detail())
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
