package enums

type Outcome string

const (
	Win  Outcome = "win"
	Lose Outcome = "lose"
)

// IsValid checks whether the outcome is a valid enum
func (o Outcome) IsValid() bool {
	switch o {
	case Win, Lose:
		return true
	default:
		return false
	}
}

// String converts enum to string
func (o Outcome) String() string {
	return string(o)
}
