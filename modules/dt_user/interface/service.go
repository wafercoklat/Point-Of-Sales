package dt_user

import (
	domain "STACK-GO/modules/dt_user/domain"
)

type Services interface {
	FindByID(id string) *domain.User
	List() *[]domain.User
	Create(mdl domain.User) error
	Update(id string, mdl domain.User) error
	Delete(id string) error
}
