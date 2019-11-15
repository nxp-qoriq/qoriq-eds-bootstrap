#####################################
#
# Copyright 2017 NXP
#
#####################################

INSTALL_DIR ?= /
INSTALL ?= install
ARCH ?= arm64

GOROOT ?= $(HOME)/go
GOPATH ?= $(HOME)/gopathdir
GOVERSION ?= 1.9.4
GOFILE ?= go$(GOVERSION).linux-amd64.tar.gz
GO = $(GOROOT)/bin/go

.PHONY: clean all install

all: goenv
	$(shell GOPATH=$(GOPATH) GOROOT=$(GOROOT)  $(GO) get github.com/laurentluce/est-client-go)
	$(shell GOPATH=$(GOPATH) GOROOT=$(GOROOT)  $(GO) get gopkg.in/yaml.v2)
	$(shell GOPATH=$(GOPATH) GOROOT=$(GOROOT)  $(GO) run parse_config.go)
	$(shell GOPATH=$(GOPATH) GOROOT=$(GOROOT)  GOOS=linux GOARCH=${ARCH} $(GO) build --ldflags="-w -s" -o ${ARCH}/bootstrap-enroll bootstrap-enroll.go config_tmp.go)

goenv:
	$(GOROOT)/bin/go version | grep $(GOVERSION); \
	if [ "$$?" != "0" ]; then  \
		wget -c https://redirector.gvt1.com/edgedl/go/$(GOFILE); \
		rm -rf $(GOROOT) && tar -C $(HOME) -xzf $(GOFILE); \
	fi
	usr=`whoami`; \

install:
	$(INSTALL) -d --mode 755 $(INSTALL_DIR)/usr/local/edgescale/bin
	sudo cp -r ${ARCH}/bootstrap-enroll $(INSTALL_DIR)/usr/local/edgescale/bin/

