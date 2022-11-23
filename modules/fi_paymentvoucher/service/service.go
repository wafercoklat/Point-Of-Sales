package fi_paymentvoucher

import (
	"STACK-GO/adapter"
	domain "STACK-GO/modules/fi_paymentvoucher/domain"
	repo "STACK-GO/modules/fi_paymentvoucher/repo"
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

func (s *Services) FindByID(id string) (*domain.PaymentvoucherPackage, error) {
	var pvpackage domain.PaymentvoucherPackage

	// Get the Paymentvoucher
	paymentvoucher, err := s.adapter.FindByID(id, &domain.Paymentvoucher{}, repo.QryFindByID_Header())
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	// Get the Paymentvoucher Details
	paymentvoucherdetail, err := s.adapter.FindByID(id, &domain.PaymentvoucherDetails{}, repo.QryFindByID_Detail())
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	pvpackage.Paymentvoucher = paymentvoucher.(domain.Paymentvoucher)
	pvpackage.PaymentvoucherDetails = paymentvoucherdetail.([]domain.PaymentvoucherDetails)

	return &pvpackage, nil
}

func (s *Services) List() (*[]domain.Paymentvoucher, error) {
	// Get List Paymentvoucher
	data, err := s.adapter.List(&domain.List, repo.QryList())
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return data.(*[]domain.Paymentvoucher), nil
}

func (s *Services) Create(data domain.PaymentvoucherPackage) (int, error) {
	// Create Paymentvoucher
	headerID, err := s.adapter.Create(data.Paymentvoucher, repo.QryCreate_Header())
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	// Get SOID
	for i := range data.PaymentvoucherDetails {
		data.PaymentvoucherDetails[i].SOID = int(headerID)
	}

	// Update PaymentvoucherDetail
	// Have to create pre-create check
	_, err = s.adapter.Create(data.PaymentvoucherDetails, repo.QryCreate_Detail())
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	return 1, err
}

func (s *Services) Update(id string, data domain.PaymentvoucherPackage) (int, error) {
	// Must to check if sales order have another transcation
	// Update Paymentvoucher
	_, err := s.adapter.Update(data.Paymentvoucher, id, repo.QryUpdate_Header())
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	// Update PaymentvoucherDetail
	// Have to create pre-update check
	for i := range data.PaymentvoucherDetails {
		_, err := s.adapter.Update(data.PaymentvoucherDetails[i], strconv.Itoa(data.PaymentvoucherDetails[i].DetailID), repo.QryCreate_Detail())
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
