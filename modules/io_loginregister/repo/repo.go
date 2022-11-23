package io_registerlo

import "fmt"

const (
	DESC        = "UName, UID, Level"
	TABLE       = "us_users"
	COLUMN      = "UID, Name, Mail, UName, Password, Location, Isdisabled, Verified, Level"
	VALUE       = ":UID, :Name, :Mail, :UName, :Password, :Location, :IsDisabled, :Verified, :Level"
	COLUMNVALUE = "UID = :UID, Name = :Name, Mail = :Mail, UName = :UName, Password = :Password, Location = :Location, IsDisabled = :IsDisabled, Verified =:Verified, Level =:Level"
	WHERE       = "UName = ? AND Password = ?"

	SELECT = "SELECT %s FROM %s WHERE %s"
)

func QryFindByUser() string {
	return fmt.Sprintf(SELECT, DESC, TABLE, WHERE)
}
