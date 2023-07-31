# PiRem
## Install
Go 1.20以上が必要です。

リポジトリをクローンして移動
```
git clone https://github.com/NaKa2355/pirem.git && cd pirem
```
ビルド
```
make build
```

インストール
```
make install
```

アップデート
```
git pull && make build && make update
```

アンインストール
```
make purge
```

## デーモンの起動
デーモンの起動
systemd
```
systemctl start piremd
```
コマンドライン
```
pirem daemon
```

## Usage
デバイス一覧の取得
```
pirem info
```
モジュール一覧の取得
```
pirem modules
```



