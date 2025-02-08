package usecases

import (
	"github.com/KathurimaKimathi/maybets/pkg/maybets/infrastructure"
)

// UsecaseMayBets represents an assemble of all use cases into a single object that can be instantiated anywhere
type UsecaseMayBets struct {
	Infrastructure infrastructure.Infrastructure
}

// NewUsecaseMayBetsImpl returns a new Maybets interactor
func NewUsecaseMayBetsImpl(
	infra infrastructure.Infrastructure,
) (*UsecaseMayBets, error) {
	return &UsecaseMayBets{
		Infrastructure: infra,
	}, nil
}
