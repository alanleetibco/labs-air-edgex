#!/bin/bash
#
# Copyright (c) 2018
# Mainflux
#
# SPDX-License-Identifier: Apache-2.0
#

###
# Launches all EdgeX Go binaries (must be previously built).
#
# Expects that Consul and MongoDB are already installed and running.
#
###

DIR=$PWD
CMD=../cmd

function cleanup {
	pkill edgex-device-idt-lora-mqtt
}

cd $CMD
exec -a edgex-device-idt-lora-mqtt ./device-idt-lora-mqtt &
cd $DIR


trap cleanup EXIT

while : ; do sleep 1 ; done