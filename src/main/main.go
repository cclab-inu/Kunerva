package main

import (
	core "github.com/cclab-inu/Kunerva/src/core"
	libs "github.com/cclab-inu/Kunerva/src/libs"
)

func main() {
	operationMode := libs.GetEnv("OPERATION_MODE", "cronjob")
	if operationMode == "cronjob" { // timer
		core.StartCronJob()
	} else { // one-time generation
		core.StartToDiscoverNetworkPolicies()
	}
}
