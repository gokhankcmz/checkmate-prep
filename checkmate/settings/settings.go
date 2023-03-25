package settings

type Settings struct {
	NotEmpty     string
	BiggerThan   string
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
}

func (s *Settings) withDefaults() *Settings {
	if s.NotEmpty == "" {
		s.NotEmpty = "field is empty"
	}

	if s.BiggerThan == "" {
		s.BiggerThan = "field is not bigger than {n}"
	}

	if s.LessThan == "" {
		s.LessThan = "field is not less than {n}"
	}

	if s.Exactly == "" {
		s.Exactly = "field is not exactly {n}"
	}

	if s.Positive == "" {
		s.Positive = "field is not positive"
	}

	if s.Negative == "" {
		s.Negative = "field is not negative"
	}

	if s.InRange == "" {
		s.InRange = "field is not in range [{min}, {max}]"
	}

	if s.Length == "" {
		s.Length = "field length is not equal to {n}"
	}

	if s.MinLength == "" {
		s.MinLength = "field length is smaller than {n}"
	}

	if s.MaxLength == "" {
		s.MaxLength = "field length is larger than {n}"
	}

	if s.ExactLength == "" {
		s.ExactLength = "field length is not equal to {n}"
	}

	if s.URL == "" {
		s.URL = "field is not a valid URL"
	}

	if s.Email == "" {
		s.Email = "field is not a valid email address"
	}

	if s.Contain == "" {
		s.Contain = "field does not contain '{substr}'"
	}

	if s.EvenNumber == "" {
		s.EvenNumber = "field is not an even number"
	}

	if s.OddNumber == "" {
		s.OddNumber = "field is not an odd number"
	}

	if s.PrimeNumber == "" {
		s.PrimeNumber = "field is not a prime number"
	}

	if s.Alphabetic == "" {
		s.Alphabetic = "field contains non-alphabetic characters"
	}

	if s.Numeric == "" {
		s.Numeric = "field contains non-numeric characters"
	}

	if s.NoWhiteSpace == "" {
		s.NoWhiteSpace = "field contains white space"
	}

	return s
}
