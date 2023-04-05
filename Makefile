#go.modがあるディレクトリの名前
MOD_NAME:=$(shell go list -m)

#コマンドのインストール先
CMD_INSTALL:=/usr/local/bin

#コンフィグファイルのインストール先
CONFIG_FILE:=config/piremd.json
CONFIG_INSTALL:=/etc/piremd.json

#サービスファイルのインストール先
SERVICE_FILE:=config/piremd.service
SERVICE_INSTALL:=/etc/systemd/system/piremd.service

#プラグインのインストール先
PLUGIN_INSTALL:=/opt/piremd

#バイナリの出力先
CMD_BIN_DIR:=bin

#コマンドのパッケージ名
CMD_PACKAGES:=$(MOD_NAME)/cmd/pirem
#GOのファイル
GO_FILES:=$(shell find . -type f -name '*.go' -print)

CMD_BIN:=$(CMD_PACKAGES:$(MOD_NAME)/cmd/%=$(CMD_BIN_DIR)/%) 

BUILD_OPT := -ldflags="-s -w" -trimpath


all:
	@echo $(CMD_BIN)
.PHONY: clean
clean:
	rm $(CMD_BIN_DIR)/**

.PHONY: build
build: $(CMD_BIN)

$(CMD_BIN): $(GO_FILES)
	$(BUILD_ENV) go build $(BUILD_OPT) -o $(CMD_BIN_DIR) $(@:$(CMD_BIN_DIR)/%=$(MOD_NAME)/cmd/%)

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
