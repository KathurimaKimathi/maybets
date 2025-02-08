package enums

type Environment string

const (
	Prod    Environment = "PROD"
	Test    Environment = "TEST"
	Staging Environment = "STAGING"
	Local   Environment = "LOCAL"
)

// IsValid checks whether the outcome is a valid enum
func (e Environment) IsValid() bool {
	switch e {
	case Prod, Test, Staging, Local:
		return true
	default:
		return false
	}
}

// String converts enum to string
func (e Environment) String() string {
	return string(e)
}
