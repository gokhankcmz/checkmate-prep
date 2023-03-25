package types

import (
	"cmDeneme/checkmate"
)

type User struct {
	Name    string
	Age     *int
	Surname string
	Address Address
}

type Address struct {
	Name    string
	Surname *string
}

func (u *User) Validate() []error {

	validator := checkmate.New()

	validator.ValidateString(&u.Name, "Name").NotEmpty().MinLength(3).MaxLength(8).Contain("this text").Alphabetic().NoWhiteSpace().UUID()
	validator.ValidateString(&u.Surname, "Surname").NotEmpty()
	validator.ValidateInt(u.Age, "Age").NotZero()

	return validator.Validate()

}
