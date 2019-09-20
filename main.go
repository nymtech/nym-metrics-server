package main

import (
	_ "github.com/nymtech/directory-server/docs"
	"github.com/nymtech/directory-server/server"
)

// @title Nym Directory API
// @version 1.0
// @description This is a temporarily centralized directory/PKI/metrics API to allow us to get the other Nym node types running. Its functionality will eventually be folded into other parts of Nym.
// @termsOfService http://swagger.io/terms/

// @license.name Apache 2.0
// @license.url https://github.com/nymtech/directory-server/license
func main() {
	router := server.New()
	router.Run(":8080")
}
