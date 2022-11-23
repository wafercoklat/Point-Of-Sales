package fi_paymentvoucher

import (
	domain "STACK-GO/modules/fi_paymentvoucher/domain"
)

type Services interface {
	FindByID(id string) (*domain.PaymentvoucherPackage, error)
	List() (*[]domain.Paymentvoucher, error)
	Create(mdl domain.PaymentvoucherPackage) (int, error)
	Update(id string, mdl domain.PaymentvoucherPackage) (int, error)
	Delete(id string) (int, error)
}
