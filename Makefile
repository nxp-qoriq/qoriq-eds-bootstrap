#####################################
#
# Copyright 2017 NXP
#
#####################################

INSTALL_DIR ?= /
INSTALL ?= install
ARCH ?= arm64

install:
	$(INSTALL) -d --mode 755 $(INSTALL_DIR)/usr/local/edgescale/bin
	sudo cp -r ${ARCH}/bootstrap-enroll $(INSTALL_DIR)/usr/local/edgescale/bin/

