#####################################
#
# Copyright 2017 NXP
#
#####################################

INSTALL_DIR ?= /
INSTALL ?= install
ARCH ?= arm64

all:
	go get github.com/laurentluce/est-client-go
	go get gopkg.in/yaml.v2
	go run parse_config.go
	env GOOS=linux GOARCH=${ARCH} go build --ldflags="-w -s" -o ${ARCH}/bootstrap-enroll bootstrap-enroll.go config_tmp.go

install:
	$(INSTALL) -d --mode 755 $(INSTALL_DIR)/usr/local/edgescale/bin
	sudo cp -r ${ARCH}/bootstrap-enroll $(INSTALL_DIR)/usr/local/edgescale/bin/

