package modules

/*
静的モジュールを設定する。
ハッシュマップにモジュール名とモジュールの構造体を入れることで追加できる
デフォルトでは、モックのモジュールが追加されている
*/

import (
	mock "github.com/NaKa2355/pirem-mock-module"
	"github.com/NaKa2355/pirem/pkg/module/v1"
)

var Modules map[string]module.Module = map[string]module.Module{
	"mock": &mock.Module{},
}
