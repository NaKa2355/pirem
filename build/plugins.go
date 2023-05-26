package build

/*
choose pluglins you want
when you add a structure which impliments plugin.Plugin and its name to the map, plugin will linked to pirem
*/

import (
	plugin "github.com/NaKa2355/pirem/pkg/plugin/v1"
	mockdevice "github.com/NaKa2355/pirem/third_party/mock_device"
)

var Plugins map[string]plugin.Plugin = map[string]plugin.Plugin{
	"mock": &mockdevice.Plugin{},
}
