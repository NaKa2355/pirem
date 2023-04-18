#main.goがあるディレクトリ
MAIN:=third_party/airer/cmd/airer

#go.modがあるディレクトリの名前
MOD_NAME:=$(shell go list -m)

#コマンドのインストール先
CMD_INSTALL:=/usr/local/bin

#コンフィグファイルのインストール先とデフォルトのファイルのパス
CONFIG_FILE:=config/piremd.json
CONFIG_INSTALL:=/etc/piremd.json

#サービスファイルのインストール先とデフォルトのファイルのパス
SERVICE_FILE:=config/piremd.service
SERVICE_INSTALL:=/etc/systemd/system/piremd.service

#プラグインのインストール先
PLUGIN_INSTALL:=/opt/piremd

#バイナリの出力先
CMD_BIN_DIR:=bin

#コマンドのパッケージ名
CMD_PACKAGES:=$(MOD_NAME)/$(MAIN)

#GOのファイル
GO_FILES:=$(shell find . -type f -name '*.go' -print)

#
CMD_BIN:=$(CMD_PACKAGES:$(MOD_NAME)/cmd/%=$(CMD_BIN_DIR)/%) 

#ビルドオプションと環境変数
BUILD_OPT :=-buildmode=plugin -ldflags="-s -w" -trimpath
BUILD_ENV := CGO_ENABLED=1

.PHONY: clean
clean:
	rm $(CMD_BIN_DIR)/**

.PHONY: build
build: $(CMD_BIN)

$(CMD_BIN): $(GO_FILES)
	$(BUILD_ENV) go build $(BUILD_OPT) -o $(CMD_BIN_DIR) $(@:$(CMD_BIN_DIR)/%=$(MOD_NAME)/cmd/%)

.PHONY: install
install:
	cp $(CMD_BIN_DIR)/* $(PLUGIN_INSTALL)

.PHONY: update
update: install

.PHONY: remove
remove:
	@echo "To remove plugins, please delete appropriate *.so files at (/opt/piremd)"