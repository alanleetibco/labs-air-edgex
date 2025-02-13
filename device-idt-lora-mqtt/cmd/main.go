// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2018 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	device_generic_mqtt "github.com/alanleetibco/labs-air-edgex/device-idt-lora-mqtt"
	"github.com/alanleetibco/labs-air-edgex/device-idt-lora-mqtt/internal/driver"
	"github.com/edgexfoundry/device-sdk-go/v2/pkg/startup"
)

const (
	serviceName string = "device-idt-lora-mqtt"
)

func main() {
	sd := driver.NewProtocolDriver()
	startup.Bootstrap(serviceName, device_generic_mqtt.Version, sd)
}
