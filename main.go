package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/RegiAdi/venera/config"
	"github.com/RegiAdi/venera/kernel"
	"github.com/RegiAdi/venera/routes"
)

func main() {
	appKernel, err := kernel.NewAppKernel()
	if err != nil {
		log.Fatalf("App failed to initialize: %v", err)
	}

	routes.API(appKernel)

	// Channel to listen for OS signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	// Start server in a goroutine
	go func() {
		serverAddr := fmt.Sprintf(":%s", config.GetAppPort())
		log.Printf("Server starting on %s", serverAddr)
		if err := appKernel.Server.Listen(serverAddr); err != nil {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	// Block until a signal is received
	<-quit
	log.Println("Shutdown signal received, initiating graceful shutdown...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	appKernel.Shutdown(ctx)
}
