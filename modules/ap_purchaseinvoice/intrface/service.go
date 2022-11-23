package ap_purchaseinvoice

import (
	domain "STACK-GO/modules/ap_purchaseinvoice/domain"
)

type Services interface {
	FindByID(id string) (*domain.PurchaseinvoicePackage, error)
	List() (*[]domain.Purchaseinvoice, error)
	Create(mdl domain.PurchaseinvoicePackage) (int, error)
	Update(id string, mdl domain.PurchaseinvoicePackage) (int, error)
	Delete(id string) (int, error)
}
