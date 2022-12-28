package main

import (
	"context"
	"fmt"
	"github.com/Phonevilai/assessment/expense"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func init() {
	if err := godotenv.Load("dev.env"); err != nil {
		fmt.Printf("please consider environment variables: %s\n", err)
	}
	gin.SetMode(os.Getenv("GIN_MODE"))
}

func main() {

	r := gin.Default()
	db := expense.NewDB(os.Getenv("DATABASE_URL"))
	mydb := expense.NewStore(db)
	service := expense.NewService(mydb)

	r.GET("/healthz", healthCheck())
	r.POST("/expenses", expense.CreateExpense(service))
	r.GET("/expenses/:id", expense.GetExpense(service))

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	s := &http.Server{
		Addr:           ":" + os.Getenv("PORT"),
		Handler:        r,
		ReadTimeout:    15 * time.Second,
		WriteTimeout:   15 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	<-ctx.Done()
	stop()
	fmt.Println("shutting down gracefully, press Ctrl+C again to force")

	timeoutCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := s.Shutdown(timeoutCtx); err != nil {
		fmt.Println(err)
	}

}

func healthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Status(http.StatusOK)
	}
}
