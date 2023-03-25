package types

import (
	"cmDeneme/checkmate"
	"cmDeneme/checkmate/settings"
)

type User struct {
	Name    string
	Surname string
	Address Address
}

type Address struct {
	Name    string
	Surname string
}

func (u *User) Validate() []error {

	validator := checkmate.New(&settings.Settings{
		NotEmpty: "Boş olamaz",
	})

	validator.ValidateString(u.Name, "Name uygun değil").NotEmpty("Problem is: Name, Cause: field cannot be empty {err}")
	validator.ValidateString(u.Surname, "Problem is: Surname, Cause: {err}").NotEmpty()
	validator.ValidateStringPtr(&u.Name).LongerThan(5, "Problem is: Name's pointer, Cause: {err}")

	return validator.Validate()

}
