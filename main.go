package main

import (
	"flag"

	"github.com/atmom/datagent/common"
	"github.com/atmom/datagent/config"
	"github.com/atmom/datagent/server"
)

var (
	configPath = flag.String("config", "", "The data agent server configuration path")
)

func main() {
	// parse command flags
	flag.Parse()

	// load configuration from json file.
	c := config.Load(*configPath)

	// init log.
	common.InitLog(c)

	// start data agent server.
	server.Start(c)
}
