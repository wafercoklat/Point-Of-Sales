package fi_receiptvoucher

import (
	"STACK-GO/adapter"
	domain "STACK-GO/modules/fi_receiptvoucher/domain"
	repo "STACK-GO/modules/fi_receiptvoucher/repo"
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

func (s *Services) FindByID(id string) (*domain.ReceiptvoucherPackage, error) {
	var rvpackage domain.ReceiptvoucherPackage

	// Get the Receiptvoucher
	receiptvoucher, err := s.adapter.FindByID(id, &domain.Receiptvoucher{}, repo.QryFindByID_Header())
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	// Get the Receiptvoucher Details
	receiptvoucherdetail, err := s.adapter.FindByID(id, &domain.ReceiptvoucherDetails{}, repo.QryFindByID_Detail())
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	rvpackage.Receiptvoucher = receiptvoucher.(domain.Receiptvoucher)
	rvpackage.ReceiptvoucherDetails = receiptvoucherdetail.([]domain.ReceiptvoucherDetails)

	return &rvpackage, nil
}

func (s *Services) List() (*[]domain.Receiptvoucher, error) {
	// Get List Receiptvoucher
	data, err := s.adapter.List(&domain.List, repo.QryList())
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return data.(*[]domain.Receiptvoucher), nil
}

func (s *Services) Create(data domain.ReceiptvoucherPackage) (int, error) {
	// Create Receiptvoucher
	headerID, err := s.adapter.Create(data.Receiptvoucher, repo.QryCreate_Header())
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	// Get SOID
	for i := range data.ReceiptvoucherDetails {
		data.ReceiptvoucherDetails[i].SOID = int(headerID)
	}

	// Update ReceiptvoucherDetail
	// Have to create pre-create check
	_, err = s.adapter.Create(data.ReceiptvoucherDetails, repo.QryCreate_Detail())
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	return 1, err
}

func (s *Services) Update(id string, data domain.ReceiptvoucherPackage) (int, error) {
	// Must to check if sales order have another transcation
	// Update Receiptvoucher
	_, err := s.adapter.Update(data.Receiptvoucher, id, repo.QryUpdate_Header())
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	// Update ReceiptvoucherDetail
	// Have to create pre-update check
	for i := range data.ReceiptvoucherDetails {
		_, err := s.adapter.Update(data.ReceiptvoucherDetails[i], strconv.Itoa(data.ReceiptvoucherDetails[i].DetailID), repo.QryCreate_Detail())
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
