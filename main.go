package main

import (
	"fmt"
	"os"

	_ "github.com/nymtech/nym-directory/docs"
	"github.com/nymtech/nym-directory/metrics"
	"github.com/nymtech/nym-directory/server"
)

// @title Nym Metrics API
// @version 0.9.0
// @description This is a temporarily centralized metrics API to allow us to get the other Nym node types running. Its functionality will eventually be folded into other parts of Nym.
// @termsOfService http://swagger.io/terms/

// @license.name Apache 2.0
// @license.url https://github.com/nymtech/nym-metrics-server/
func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Fprint(os.Stderr, "Expected single argument to be passed - address of the validator server")
		return
	}
	validatorAddress := args[0]

	router := server.New()
	go metrics.DynamicallyUpdateReportDelay(validatorAddress)
	router.Run(":8080")
}
