GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
BASEPATH := $(shell pwd)
BUILDDIR=$(BASEPATH)/dist

KOTF_SRC=$(BASEPATH)/cmd
KOTF_SERVER_NAME=kotf-server
# KOTF_INVENTORY_NAME=kotf-inventory
# KOTF_CLIENT_NAME=kotf

BIN_DIR=usr/local/bin
CONFIG_DIR=etc/kotf
BASE_DIR=var/kotf
LIB_DIR=$(BASE_DIR)/lib
GOPROXY="https://goproxy.cn,direct"

build_linux:
	GOOS=linux GOARCH=amd64  $(GOBUILD) -o $(BUILDDIR)/$(BIN_DIR)/$(KOTF_SERVER_NAME) $(KOTF_SRC)/server/*.go
#     mkdir -p $(BUILDDIR)/$(LIB_DIR) && cp -r     $(BASEPATH)/resource $(BUILDDIR)/$(LIB_DIR)
	mkdir -p $(BUILDDIR)/$(CONFIG_DIR) && cp -r  $(BASEPATH)/conf/* $(BUILDDIR)/$(CONFIG_DIR)
build_darwin:
	GOOS=darwin GOARCH=amd64  $(GOBUILD) -o $(BUILDDIR)/$(BIN_DIR)/$(KOTF_SERVER_NAME) $(KOTF_SRC)/server/*.go
# 	mkdir -p $(BUILDDIR)/$(LIB_DIR) && cp -r     $(BASEPATH)/resource $(BUILDDIR)/$(LIB_DIR)
	mkdir -p $(BUILDDIR)/$(CONFIG_DIR) && cp -r  $(BASEPATH)/conf/* $(BUILDDIR)/$(CONFIG_DIR)
build_server_linux:
	GOOS=linux GOARCH=amd64  $(GOBUILD) -o $(BUILDDIR)/$(BIN_DIR)/$(KOTF_CLIENT_NAME) $(KOTF_SRC)/server/*.go
# 	mkdir -p $(BUILDDIR)/$(LIB_DIR) && cp -r     $(BASEPATH)/resource $(BUILDDIR)/$(LIB_DIR)
	mkdir -p $(BUILDDIR)/$(CONFIG_DIR) && cp -r  $(BASEPATH)/conf/* $(BUILDDIR)/$(CONFIG_DIR)


clean:
	$(GOCLEAN)
	rm -fr $(BUILDDIR)

docker:
	@echo "build docker images"
	docker build -t kotf-server --build-arg GOPROXY=$(GOPROXY) .
