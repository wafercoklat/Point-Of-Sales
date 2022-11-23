package iv_item

import (
	domain "STACK-GO/modules/iv_item/domain"
)

type Services interface {
	FindByID(id string) *domain.Item
	List() *[]domain.Item
	Create(mdl domain.Item) error
	Update(id string, mdl domain.Item) error
	Delete(id string) error
}
