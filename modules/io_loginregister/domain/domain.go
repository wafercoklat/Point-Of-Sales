package io_registerlogin

type Login struct {
	UName    string `form:"username" binding:"required" json:"username,omitempty" db:"UName"`
	Password string `form:"password" binding:"required" json:"password,omitempty" db:"Password"`
	Level    string `json:"level,omitempty" db:"Level"`
	UID      string `json:"uid,omitempty" db:"UID"`
}
