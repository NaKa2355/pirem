#go.modがあるディレクトリの名前
MOD_NAME:=pirem

#コマンドのインストール先
CMD_INSTALL:=/usr/local/bin

#コンフィグファイルのインストール先
CONFIG_FILE:=init/piremd.json
CONFIG_INSTALL:=/etc/piremd.json

#サービスファイルのインストール先
SERVICE_FILE:=init/piremd.service
SERVICE_INSTALL:=/lib/systemd/system/piremd.service

#プラグインのインストール先
PLUGIN_INSTALL:=/opt/piremd

#バイナリの出力先
CMD_BIN_DIR:=bin

#コマンドのパッケージ名
COMMAND_PACKAGES:=$(shell go list ./cmd/pirem)
#GOのファイル
GO_FILES:=$(shell find . -type f -name '*.go' -print)

CMD_BIN:=$(COMMAND_PACKAGES:$(MOD_NAME)/cmd/%=$(CMD_BIN_DIR)/%) 

BUILD_OPT := -ldflags="-s -w" -trimpath

BUILD_ENV := GOOS=linux CGO_ENABLED=1
.PHONY: clean
clean:
	rm bin/**

.PHONY: build
build: $(CMD_BIN)

$(CMD_BIN): $(GO_FILES)
	go build $(BUILD_OPT) -o $(CMD_BIN_DIR) $(@:$(CMD_BIN_DIR)/%=$(MOD_NAME)/cmd/%)

.PHONY: install
install:
	cp $(CMD_BIN_DIR)/* $(CMD_INSTALL)
	cp $(CONFIG_FILE) $(CONFIG_INSTALL)
	cp $(SERVICE_FILE) $(SERVICE_INSTALL)
	mkdir $(PLUGIN_INSTALL)

.PHONY: update
update:
	cp $(CMD_BIN_DIR)/* $(CMD_INSTALL)

.PHONY: remove
remove:
	rm $(CMD_BIN:$(CMD_BIN_DIR)/%=$(CMD_INSTALL)/%)

.PHONY: purge
purge: remove
	rm $(CONFIG_INSTALL)
	rm $(SERVICE_INSTALL)
	rm -rf $(PLUGIN_INSTALL)
