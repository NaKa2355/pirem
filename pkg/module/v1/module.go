package module

import "encoding/json"

type Module interface {
	NewDriver(json.RawMessage) (Driver, error)
}
