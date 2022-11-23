package dt_user

import (
	"time"
)

type User struct {
	ID         int       `json:"id,omitempty" db:"ID"`
	UID        string    `json:"uid,omitempty" db:"UID"`
	Name       string    `json:"name,omitempty" db:"Name"`
	Mail       string    `json:"mail,omitempty" db:"Mail"`
	UName      string    `json:"username,omitempty" db:"UName"`
	Password   string    `json:"password,omitempty" db:"Password"`
	Location   string    `json:"location,omitempty" db:"Location"`
	IsDisabled int       `json:"isdisabled,omitempty" db:"IsDisabled"`
	Verified   int       `json:"verified,omitempty" db:"Verified"`
	Level      string    `json:"level,omitempty" db:"Level"`
	CreateDate time.Time `json:"createdate,omitempty" db:"CreateDate"`
}

var List []User

type LoginUser struct {
	UName    string `json:"username,omitempty" db:"UName"`
	Password string `json:"password,omitempty" db:"Password"`
	Level    string `json:"level,omitempty" db:"Level"`
}
