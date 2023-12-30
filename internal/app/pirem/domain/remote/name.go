package remote

import (
	"fmt"

	"github.com/NaKa2355/pirem/internal/app/pirem/domain"
	validation "github.com/go-ozzo/ozzo-validation"
)

type Name string

func NewName(name string) (Name, error) {
	err := validation.Validate(
		name,
		validation.Required,
		validation.RuneLength(0, 25),
	)
	if err != nil {
		return Name(""), domain.WrapError(
			domain.CodeInvaildInput,
			fmt.Errorf("validation error at name: %w", err),
		)
	}
	return Name(name), nil
}
