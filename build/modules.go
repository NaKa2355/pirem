package build

/*
静的モジュールを設定する。
ハッシュマップにモジュール名とモジュールの構造体を入れることで追加できる
デフォルトでは、モックのモジュールが追加されている
*/

import (
	"github.com/NaKa2355/pirem/pkg/module/v1"
	mock "github.com/NaKa2355/pirem_mock_module"
)

var Modules map[string]module.Module = map[string]module.Module{
	"mock": &mock.Module{},
}
