package ap_purchaseorder

import (
	domain "STACK-GO/modules/ap_purchaseorder/domain"
)

type Services interface {
	FindByID(id string) (*domain.PurchaseorderPackage, error)
	List() (*[]domain.Purchaseorder, error)
	Create(mdl domain.PurchaseorderPackage) (int, error)
	Update(id string, mdl domain.PurchaseorderPackage) (int, error)
	Delete(id string) (int, error)
}
