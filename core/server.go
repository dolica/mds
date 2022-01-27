package core

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dolica/mds/global"
	"github.com/gin-gonic/gin"
)

func RunServer() {
	g := gin.Default()
	httpConfig := global.HTF_Config.Http
	addr := fmt.Sprintf(":%s", httpConfig.Port)
	server := &http.Server{
		Addr:         addr,
		WriteTimeout: time.Millisecond * time.Duration(httpConfig.WriteTimeout),
		ReadTimeout:  time.Millisecond * time.Duration(httpConfig.ReadTimeout),
		Handler:      g,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Printf("listen:%s\n", err)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("shutting down server....")

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server force to shutdown:", err)
	}
	log.Printf("Server existing...")
}
