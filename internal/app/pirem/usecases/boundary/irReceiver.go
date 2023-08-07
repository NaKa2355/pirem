package boundary

import "context"

type ReceiveIRInput struct {
	ID string
}

type IRReceiver interface {
	ReceiveIR(context.Context, ReceiveIRInput) (IRData, error)
}
