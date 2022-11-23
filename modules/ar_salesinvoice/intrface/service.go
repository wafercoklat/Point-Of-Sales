package ar_salesinvoice

import (
	domain "STACK-GO/modules/ar_salesinvoice/domain"
)

type Services interface {
	FindByID(id string) (*domain.SalesinvoicePackage, error)
	List() (*[]domain.Salesinvoice, error)
	Create(mdl domain.SalesinvoicePackage) (int, error)
	Update(id string, mdl domain.SalesinvoicePackage) (int, error)
	Delete(id string) (int, error)
}
