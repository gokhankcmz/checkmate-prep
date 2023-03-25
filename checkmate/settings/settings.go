package settings

type Settings struct {
	StopAtFirstError     bool
	DefaultErrorMessages DefaultErrorMessages
}

type DefaultErrorMessages struct {
	NotEmpty     string
	GreaterThan  string
	LessThan     string
	Exactly      string
	Positive     string
	Negative     string
	InRange      string
	Length       string
	MinLength    string
	MaxLength    string
	ExactLength  string
	URL          string
	Email        string
	Contain      string
	EvenNumber   string
	OddNumber    string
	PrimeNumber  string
	Alphabetic   string
	Numeric      string
	NoWhiteSpace string
	UUID         string
	NotZero      string
}

func (d *DefaultErrorMessages) WithDefaults() DefaultErrorMessages {
	if d.NotEmpty == "" {
		d.NotEmpty = "The {fieldName} field is required and cannot be empty"
	}

	if d.GreaterThan == "" {
		d.GreaterThan = "The {fieldName} field must be greater than {n}"
	}

	if d.LessThan == "" {
		d.LessThan = "The {fieldName} field must be less than {n}"
	}

	if d.Exactly == "" {
		d.Exactly = "The {fieldName} field must be exactly {n}"
	}

	if d.Positive == "" {
		d.Positive = "The {fieldName} field must be a positive number"
	}

	if d.Negative == "" {
		d.Negative = "The {fieldName} field must be a negative number"
	}

	if d.InRange == "" {
		d.InRange = "The {fieldName} field must be between {min} and {max}"
	}

	if d.Length == "" {
		d.Length = "The length of {fieldName} must be {n}"
	}

	if d.MinLength == "" {
		d.MinLength = "The length of {fieldName} must be at least {n}"
	}

	if d.MaxLength == "" {
		d.MaxLength = "The length of {fieldName} must not exceed {n}"
	}

	if d.ExactLength == "" {
		d.ExactLength = "The length of {fieldName} must be exactly {n}"
	}

	if d.URL == "" {
		d.URL = "The {fieldName} field must be a valid URL"
	}

	if d.Email == "" {
		d.Email = "The {fieldName} field must be a valid email address"
	}

	if d.Contain == "" {
		d.Contain = "The {fieldName} field must contain '{substr}'"
	}

	if d.EvenNumber == "" {
		d.EvenNumber = "The {fieldName} field must be an even number"
	}

	if d.OddNumber == "" {
		d.OddNumber = "The {fieldName} field must be an odd number"
	}

	if d.PrimeNumber == "" {
		d.PrimeNumber = "The {fieldName} field must be a prime number"
	}

	if d.Alphabetic == "" {
		d.Alphabetic = "The {fieldName} field must contain only alphabetic characters"
	}

	if d.Numeric == "" {
		d.Numeric = "The {fieldName} field must contain only numeric characters"
	}

	if d.NoWhiteSpace == "" {
		d.NoWhiteSpace = "The {fieldName} field must not contain white space"
	}

	if d.UUID == "" {
		d.UUID = "The {fieldName} field must be a valid GUID err: {err}"
	}

	if d.NotZero == "" {
		d.NotZero = "The {fieldName} field must not be 0"
	}

	return *d
}
