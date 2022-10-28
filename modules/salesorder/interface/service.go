package user

import (
	domain "STACK-GO/modules/salesorder/domain"
)

type Services interface {
	FindByID(id string) (*domain.SalesorderPackage, error)
	List() (*[]domain.Salesorder, error)
	Create(mdl domain.SalesorderPackage) (int, error)
	Update(id string, mdl domain.SalesorderPackage) (int, error)
	Delete(id string) (int, error)
}
