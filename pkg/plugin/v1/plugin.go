package plugin

import "encoding/json"

type Plugin interface {
	NewDriver(json.RawMessage) (Driver, error)
}
