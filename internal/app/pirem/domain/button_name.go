package domain

import (
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation"
)

type ButtonName string

func NewButtonName(name string) (ButtonName, error) {
	err := validation.Validate(name,
		validation.Required,
		validation.Length(1, 20),
	)
	if err != nil {
		return ButtonName(""), WrapError(
			CodeInvaildInput,
			fmt.Errorf("validation error at ButtonName: %w", err),
		)
	}
	return ButtonName(name), nil
}
