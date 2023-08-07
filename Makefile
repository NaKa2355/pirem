#---------------------ビルド系------------------
#モジュール名
MOD_NAME := $(shell go list -m)
#バイナリを格納するディレクトリ
BIN_DIR:=bin
#GOのファイル
GO_FILES:=$(shell find . -type f -name '*.go' -print)
#バイナリのファイルパス
BIN:=bin/pirem
#ビルドオプション
BUILD_OPT := -ldflags="-s -w" -trimpath


#------------------インストール系----------------
#コマンドのインストール先
CMD_INSTALL:=/usr/local/bin

#コンフィグファイルのインストール先
CONFIG_FILE:=config/piremd.json
CONFIG_INSTALL:=/etc/piremd.json

#サービスファイルのインストール先
SERVICE_FILE:=config/piremd.service
SERVICE_INSTALL:=/etc/systemd/system/piremd.service


#-----------------Makefile----------------------
.PHONY: all
all: build

.PHONY: clean
clean:
	rm $(BIN_DIR)/**

.PHONY: build
build: $(BIN)

$(BIN): $(GO_FILES)
	go build $(BUILD_OPT) -o $(BIN_DIR) $(@:$(BIN_DIR)/%=$(MOD_NAME)/cmd/%)

.PHONY: install
install:
	cp $(BIN_DIR)/* $(CMD_INSTALL)
	cp $(CONFIG_FILE) $(CONFIG_INSTALL)
	cp $(SERVICE_FILE) $(SERVICE_INSTALL)

.PHONY: update
update:
	cp $(BIN_DIR)/* $(CMD_INSTALL)

.PHONY: remove
remove:
	rm $(BIN:$(BIN_DIR)/%=$(CMD_INSTALL)/%)

.PHONY: purge
purge: remove
	rm $(CONFIG_INSTALL)
	rm $(SERVICE_INSTALL)
