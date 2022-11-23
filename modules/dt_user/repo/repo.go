package dt_user

import "fmt"

const (
	ALL         = "*"
	TABLE       = "us_users"
	COLUMN      = "UID, Name, Mail, UName, Password, Location, Isdisabled, Verified, Level"
	VALUE       = ":UID, :Name, :Mail, :UName, :Password, :Location, :IsDisabled, :Verified, :Level"
	COLUMNVALUE = "UID = :UID, Name = :Name, Mail = :Mail, UName = :UName, Password = :Password, Location = :Location, IsDisabled = :IsDisabled, Verified =:Verified, Level =:Level"
	WHERE       = "ID"

	SELECT = "SELECT %s FROM %s WHERE %s = ?"
	CREATE = "INSERT INTO %s (%s) VALUES (%s)"
	UPDATE = "UPDATE %s SET %s WHERE %s = ?"
	LIST   = "SELECT %s FROM %s"
	DELETE = "DELETE FROM %s WHERE %s = ?"
)

func QryFindByID() string {
	return fmt.Sprintf(SELECT, ALL, TABLE, WHERE)
}

func QryList() string {
	return fmt.Sprintf(LIST, ALL, TABLE)
}

func QryCreate() string {
	return fmt.Sprintf(CREATE, TABLE, COLUMN, VALUE)
}

func QryUpdate() string {
	return fmt.Sprintf(UPDATE, TABLE, COLUMNVALUE, WHERE)
}

func QryDelete() string {
	return fmt.Sprintf(DELETE, TABLE, WHERE)
}
