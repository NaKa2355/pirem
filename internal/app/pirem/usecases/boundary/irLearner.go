package boundary

import (
	"context"

	"github.com/NaKa2355/pirem/internal/app/pirem/domain/button"
	"github.com/NaKa2355/pirem/internal/app/pirem/domain/irdata"
)

type LearnIRInput struct {
	ButtonID button.ID
	IRData   irdata.IRData
}

type IRLearner interface {
	LearnIR(context.Context, LearnIRInput) error
}
