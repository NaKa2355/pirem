package remote

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid"
)

type ID string

func NewID() ID {
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	id := ulid.MustNew(ulid.Timestamp(t), entropy)
	return ID(id.String())
}
