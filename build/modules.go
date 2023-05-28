package build

/*
choose modules you want
when you add a structure which impliments module.Module and its name to the map, module will linked to pirem
*/

import (
	"github.com/NaKa2355/pirem/pkg/module/v1"
	"github.com/NaKa2355/pirem/third_party/airer"
)

var Modules map[string]module.Module = map[string]module.Module{
	"airer": &airer.Module{},
}
