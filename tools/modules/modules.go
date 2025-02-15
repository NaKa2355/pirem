package modules

/*
静的モジュールを設定する。
ハッシュマップにモジュール名とモジュールの構造体を入れることで追加できる
デフォルトでは、モックのモジュールが追加されている

*/

import (
	"github.com/NaKa2355/pirem/pkg/driver_module/v1"
)

var Modules map[string]driver_module.DriverModule = map[string]driver_module.DriverModule{
	//"mock": &mock.Module{},
}
