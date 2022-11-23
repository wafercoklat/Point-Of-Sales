package dt_bank

import (
	domain "STACK-GO/modules/dt_bank/domain"
)

type Services interface {
	FindByID(id string) *domain.Bank
	List() *[]domain.Bank
	Create(mdl domain.Bank) error
	Update(id string, mdl domain.Bank) error
	Delete(id string) error
}
