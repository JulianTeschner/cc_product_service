package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/JulianTeschner/cc_product_service/product"
	"github.com/JulianTeschner/cc_product_service/router"
	"github.com/joho/godotenv"
)

// @title        Product Service
// @version      1.0
// @description  This is a simple API for the cloudcomputing Project.

// @host      localhost:8080
// @BasePath  https://localhost:8080/

var ctx context.Context

func init() {
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
}

func main() {
	path, _ := os.Getwd()
	path = path + "/.env"
	err := godotenv.Load(path)
	if err != nil {
		log.Println("No .env file found, using default values")
	}

	product.NewClient()
	defer product.Client.Disconnect(ctx)

	r := router.New()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	err = r.Run(":" + port)
	if err != nil {
		log.Fatalln(err)
	}

}
