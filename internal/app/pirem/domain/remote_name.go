package domain

import (
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation"
)

type RemoteName string

func NewRemoteName(name string) (RemoteName, error) {
	err := validation.Validate(
		name,
		validation.Required,
		validation.RuneLength(0, 25),
	)
	if err != nil {
		return RemoteName(""), WrapError(
			CodeInvaildInput,
			fmt.Errorf("validation error at name: %w", err),
		)
	}
	return RemoteName(name), nil
}
