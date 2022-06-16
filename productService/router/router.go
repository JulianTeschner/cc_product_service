package router

import (
	"log"

	_ "github.com/JulianTeschner/cloudcomputing/service/productService/middleware"
	"github.com/JulianTeschner/cloudcomputing/service/productService/product"
	"github.com/gin-gonic/gin"

	_ "github.com/gwatts/gin-adapter"
)

// New returns a new router
func New() *gin.Engine {
	log.Println("Setting up router")
	gin.ForceConsoleColor()

	r := gin.Default()

	// Wrap the http handler with gin adapter
	productGroup := r.Group("/product")
	productGroup.GET("/:name", product.GetProductHandler)
	productGroup.POST("", product.PostProductHandler)
	productGroup.DELETE("/:name", product.DeleteProductHandler)
	return r
}
