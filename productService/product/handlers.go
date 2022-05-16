package product

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// GetProduct is the handler for the GET api/product/* route
func GetProductHandler(c *gin.Context) {
	name := c.Param("name")
	log.Println("GetProduct: ", name)

	var err error
	product, err := GetProduct(os.Getenv("MONGO_INITDB_ROOT_DATABASE"), "products", "name", name)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, &product)
}

// PostProduct is the handler for the POST api/product/* route
func PostProductHandler(c *gin.Context) {
	var product Product
	c.Request.ParseForm()
	err := c.BindJSON(&product)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	_, err = PostProduct(os.Getenv("MONGO_INITDB_ROOT_DATABASE"), "products", &product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, &product)
}

func DeleteProductHandler(c *gin.Context) {
	name := c.Param("name")
	log.Println("DeleteProduct: ", name)

	result, err := DeleteProduct(os.Getenv("MONGO_INITDB_ROOT_DATABASE"), "products", "name", name)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}
