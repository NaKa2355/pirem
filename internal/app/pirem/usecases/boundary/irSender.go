package boundary

import "context"

type SendIRInput struct {
	ID     string
	IRData IRData
}

type IRSender interface {
	SendIR(context.Context, SendIRInput) error
}
