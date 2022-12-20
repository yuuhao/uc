package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"uc/boot"
	"uc/routes"

	"github.com/gin-gonic/gin"
)

var server *gin.Engine

func main() {

	boot.Boot()

	server = gin.New()

	server.MaxMultipartMemory = 8 << 20 //8 MiB
	routes.Route(server)

	srv := &http.Server{
		Addr:    ":8087",
		Handler: server,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	quit := make(chan os.Signal)

	signal.Notify(quit, os.Interrupt)

	<-quit

	log.Println("Shutdown Server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("server shutdown:", err)
	}

	log.Println("server exiting")
}
