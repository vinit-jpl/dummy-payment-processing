package main

import (
	"context"
	"dummy-payment-processing/internal/handlers"
	"dummy-payment-processing/internal/service"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		time.Sleep(3 * time.Second)
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to gin server",
		})
	})

	transactionService := service.NewTransactionService()
	transactionHandler := handlers.NewTransactionHandler(transactionService)

	router.POST("/transaction/create", transactionHandler.CreateTransaction)
	router.GET("/transaction/status/:id", transactionHandler.GetTransactionStatus)
	router.GET("/transaction/stats", transactionHandler.GetTransactionStats)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down the server....")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Println("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
