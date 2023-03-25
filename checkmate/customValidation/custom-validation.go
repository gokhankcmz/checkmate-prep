package customValidation

type CustomValidation struct {
	F func() []error
}

func (c *CustomValidation) Validate() []error {
	return c.F()
}
