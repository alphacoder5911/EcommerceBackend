// Package main is the entry point (commenting for lint purposes)
package main

import (
	"context"
	"fmt"

	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	server "github.com/alphacoder5911/EcommerceBackend/internal/Server"
	"github.com/alphacoder5911/EcommerceBackend/internal/config"
	"github.com/alphacoder5911/EcommerceBackend/internal/database"
	"github.com/alphacoder5911/EcommerceBackend/internal/logger"
	"github.com/gin-gonic/gin"
	
)

func main() {
	log := logger.New()
	cfg, err := config.Load()
	
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load config")
	}

	db, err := database.New(&cfg.Database)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
	}

	mainDB, err := db.DB()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to get database connection")
	}

	defer func() {
		if err := mainDB.Close(); err != nil {
			log.Fatal().Err(err).Msg("Error closing db:")
		}
	}()

	gin.SetMode(cfg.Server.GinMode)

	srv:=server.NewServer(cfg,db,log)
	router:=srv.SetupRoutes()
	fmt.Printf("%s is the port ",cfg.Server.Port)
	// log.Info().Msg("%s is the port ",cfg.Server.Port)
	httpServer:=&http.Server{
		Addr: cfg.Server.Port,
		Handler: router,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func(){
		log.Info().Str("Port",cfg.Server.Port).Msg("Starting http server")
		if err:=httpServer.ListenAndServe();err !=nil && err != http.ErrServerClosed {
			log.Fatal().Err(err).Msg("Failed to start http server")
		}
	}()

	quit:=make(chan os.Signal,1)
	signal.Notify(quit,os.Interrupt,syscall.SIGINT,syscall.SIGTERM)

	<-quit

	log.Info().Msg("Shutting down http server")

	ctx,cancel:=context.WithTimeout(context.Background(),5 * time.Second)
	defer cancel()


	if err:=httpServer.Shutdown(ctx);err!=nil{
		log.Fatal().Err(err).Msg("Server failed to shutdown")

	}

	log.Info().Msg("Shutting down database ")



	

}
