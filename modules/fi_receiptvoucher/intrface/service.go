package fi_receiptvoucher

import (
	domain "STACK-GO/modules/fi_receiptvoucher/domain"
)

type Services interface {
	FindByID(id string) (*domain.ReceiptvoucherPackage, error)
	List() (*[]domain.Receiptvoucher, error)
	Create(mdl domain.ReceiptvoucherPackage) (int, error)
	Update(id string, mdl domain.ReceiptvoucherPackage) (int, error)
	Delete(id string) (int, error)
}
