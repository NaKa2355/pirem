package boundary

import (
	"context"

	"github.com/NaKa2355/pirem/internal/app/pirem/domain"
)

type RemoteLister interface {
	ListRemotes(context.Context) ([]*domain.Remote, error)
}
