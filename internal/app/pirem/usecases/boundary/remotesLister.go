package boundary

import (
	"context"

	"github.com/NaKa2355/pirem/internal/app/pirem/domain/remote"
)

type RemoteLister interface {
	ListRemotes(context.Context) ([]*remote.Remote, error)
}
