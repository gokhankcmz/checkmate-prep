package example

import (
	"cmDeneme/checkmate"
	"errors"
)

type User struct {
	Name        string
	Surname     string
	YearOfBirth int
	Address     Address
	Email       string
	WebPage     string
}

func (u *User) Validate() []error {

	validator := checkmate.New()
	validator.ValidateString(&u.Name, "Name").NotEmpty().MinLength(3).MaxLength(8).Contain("this text").Alphabetic().NoWhiteSpace().UUID()
	validator.ValidateString(&u.Surname, "Surname").NotEmpty()
	validator.ValidateInt(&u.YearOfBirth, "YearOfBirth").NotZero().GreaterThan(1900)
	validator.ValidateString(&u.Email, "Email").Email()
	validator.ValidateStruct(&u.Address)

	validator.Custom(func() []error {
		if u.Name == "Gökhan" && u.Surname == "Kocamaz" {
			return []error{errors.New("Gökhan Kocamaz cannot pass through validation")}
		}
		return nil
	})

	return validator.Validate()

}

type Address struct {
	City    string
	Country *string
}

func (a *Address) Validate() []error {

	validator := checkmate.New()
	validator.ValidateString(&a.City, "City").NotEmpty().MinLength(3).MaxLength(20).Alphabetic().NoWhiteSpace()
	validator.ValidateString(a.Country, "Country").NotEmpty().MinLength(3).MaxLength(20).Alphabetic().NoWhiteSpace()

	return validator.Validate()

}
