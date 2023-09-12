package main

import (
	"log"

	"github.com/RegiAdi/hatchet/config"
	"github.com/RegiAdi/hatchet/kernel"
	"github.com/RegiAdi/hatchet/routes"
)

func main() {
	appKernel := kernel.NewAppKernel()

	routes.API(appKernel)

	err := appKernel.Server.Listen(":" + config.GetAppPort())

	if err != nil {
		log.Fatal("App failed to start.")
		panic(err)
	}
}
