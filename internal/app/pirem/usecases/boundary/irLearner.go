package boundary

import (
	"context"

	"github.com/NaKa2355/pirem/internal/app/pirem/domain"
)

type LearnIRInput struct {
	ButtonID domain.ButtonID
	IRData   domain.IRData
}

type IRLearner interface {
	LearnIR(context.Context, LearnIRInput) error
}
