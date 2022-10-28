package user_test

import (
	domain "STACK-GO/modules/user/domain"
	"math/rand"
)

var (
	data = domain.User{
		ID:         rand.Int(),
		UID:        "1234",
		Name:       "Testing",
		Mail:       "testing@test.com",
		UName:      "test",
		Password:   "54321",
		Location:   "",
		IsDisabled: 0,
		Verified:   0,
		Level:      "Admin",
	}
)

// func TestMain(m *testing.M) {
// 	mockRepo := new(mocks.MockUserRepository)
// 	userservice := user.NewService(adapter.RepoAdapter{
// 		// m
// 	})
// }
