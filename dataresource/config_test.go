package dataresource_test

import (
	"STACK-GO/adapter"
	"STACK-GO/dataresource"
	domain "STACK-GO/modules/user/domain"
	repo "STACK-GO/modules/user/repo"
	"fmt"
	"log"
	"os"
	"strconv"
	"testing"

	"github.com/joho/godotenv"
	. "github.com/smartystreets/goconvey/convey"
)

var (
	dt_ID     = 2
	database  adapter.RepoAdapter
	dt_create = domain.User{
		UID:        "0123",
		Name:       "Testing",
		Mail:       "test@testing.com",
		UName:      "testing",
		Password:   "1234567",
		Location:   "Medan",
		IsDisabled: 0,
		Verified:   0,
		Level:      "Administrator",
	}
	dt_create_fail = domain.User{
		UID:        "0123",
		Name:       "Testing",
		Mail:       "test@testing.com",
		UName:      "testing",
		Password:   "1234567",
		Location:   "Medan",
		IsDisabled: 0,
		Verified:   0,
		Level:      "Administrator",
	}
)

func TestMain(m *testing.M) {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	DSN := os.Getenv("ADM_USER") + ":" + os.Getenv("ADM_PASS") + "@tcp(" + os.Getenv("IP_WHITELIST") + ":" + os.Getenv("PORT") + ")/" + os.Getenv("DATABASE") + "?parseTime=true"
	db, err := dataresource.New(os.Getenv("DIALECT"), DSN, 10, 10)
	if err != nil {
		fmt.Printf("%s", err)
	}
	database = db

	exitVal := m.Run()

	os.Exit(exitVal)
}

func TestUser_Create_Success(t *testing.T) {
	info := "Test Success Create Method"
	t.Run(info, func(t *testing.T) {
		Convey(info, t, func() {
			_, err := database.Create(dt_create, repo.QryCreate())
			So(err, ShouldBeNil)
		})
	})
}

func TestUser_Create_Fail(t *testing.T) {
	info := "Test Fail Create Method"
	t.Run(info, func(t *testing.T) {
		Convey(info, t, func() {
			_, err := database.Create(dt_create_fail, repo.QryCreate())
			So(err, ShouldBeError)
		})
	})
}

func TestUser_FindByID_Success(t *testing.T) {
	info := "Test Success FindbyID Method"
	t.Run(info, func(t *testing.T) {
		Convey(info, t, func() {
			data, err := database.FindByID(strconv.Itoa(dt_ID), &domain.User{}, repo.QryFindByID())
			So(err, ShouldBeNil)
			So(data.(*domain.User).ID, ShouldEqual, dt_ID)
			So(data.(*domain.User).Mail, ShouldEqual, dt_create.Mail)
		})
	})
}

func TestUser_FindByID_Fail(t *testing.T) {
	info := "Test Fail FindbyID Method"
	t.Run(info, func(t *testing.T) {
		Convey(info, t, func() {
			_, err := database.Create(dt_create_fail, repo.QryCreate())
			So(err, ShouldBeError)
		})
	})
}
