package main

import (
	"log"

	"github.com/RegiAdi/venera/config"
	"github.com/RegiAdi/venera/kernel"
	"github.com/RegiAdi/venera/routes"
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
