package mocks

import (
	domain "STACK-GO/modules/user/domain"

	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) FindByID(id string, data interface{}, query string) (interface{}, error) {
	ret := m.Called(id)

	var r0 = data.(*domain.User)
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(*domain.User)
	}

	var r1 error

	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}

	return r0, r1
}

func (m *MockUserRepository) List(mdl interface{}, query string) (interface{}, error) {
	ret := m.Called()

	var r0 = mdl.(*[]domain.User)
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(*[]domain.User)
	}

	var r1 error

	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}

	return r0, r1
}

// func (m *MockUserRepository) Create(data interface{}, query string) (int, error) {
// 	ret := m.Called(id)

// 	var r0 = data.(*domain.User)
// 	if ret.Get(0) != nil {
// 		r0 = ret.Get(0).(*domain.User)
// 	}

// 	var r1 error

// 	if ret.Get(1) != nil {
// 		r1 = ret.Get(1).(error)
// 	}

// 	return r0.ID, r1
// }

func (m *MockUserRepository) Update(data interface{}, id, query string) (int, error) {
	ret := m.Called(id)

	var r0 = data.(*domain.User)
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(*domain.User)
	}

	var r1 error

	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}

	return r0.ID, r1
}

func (m *MockUserRepository) Delete(id, query string) (int, error) {
	ret := m.Called(id)

	var r0 *domain.User
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(*domain.User)
	}

	var r1 error

	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}

	return r0.ID, r1
}
