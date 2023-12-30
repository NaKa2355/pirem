package domain

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid"
)

type RemoteID string

func NewRemoteID() RemoteID {
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	id := ulid.MustNew(ulid.Timestamp(t), entropy)
	return RemoteID(id.String())
}
