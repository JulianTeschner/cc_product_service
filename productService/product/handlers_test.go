package product

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

var r *gin.Engine

func setupHandlersTest() func() {
	r = gin.Default()

	productGroup := r.Group("/product")
	{
		productGroup.GET("/:name", GetProductHandler)
		productGroup.POST("", PostProductHandler)
		productGroup.DELETE("/:name", DeleteProductHandler)
	}

	return func() {
		log.Println("teardown suite")
	}
}

func TestGetProductHandler(t *testing.T) {
	teardown := addDummyProductToDb()
	defer teardown()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/product/fish", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetProductHandlerNotFound(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/product/NotFound", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestPostProductHandler(t *testing.T) {
	w := httptest.NewRecorder()
	dummyProduct := createDummyProduct()
	data, err := json.Marshal(&dummyProduct)
	if err != nil {
		log.Fatal(err)
	}
	req, _ := http.NewRequest("POST", "/product", bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	Client.Database("test").Collection("products").DeleteOne(context.Background(), bson.M{"name": dummyProduct.Name})
}

// func TestPostProductHandlerMarshallError(t *testing.T) {
// 	w := httptest.NewRecorder()
// 	data, err := json.Marshal(nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	req, _ := http.NewRequest("POST", "/product", bytes.NewBuffer(data))
// 	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
// 	r.ServeHTTP(w, req)
// 	assert.Equal(t, http.StatusBadRequest, w.Code)
// }
//
func TestPostProductHandlerNoConnection(t *testing.T) {
	w := httptest.NewRecorder()
	dummyProduct := createDummyProduct()
	data, err := json.Marshal(&dummyProduct)
	if err != nil {
		log.Fatal(err)
	}
	Client.Disconnect(ctx)
	req, _ := http.NewRequest("POST", "/product", bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	r.ServeHTTP(w, req)
	NewClient()
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestDeleteProductHandler(t *testing.T) {
	addDummyProductToDb()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/product/fish", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeleteProductHandlerNotFound(t *testing.T) {
	w := httptest.NewRecorder()
	Client.Disconnect(context.Background())
	req, _ := http.NewRequest("DELETE", "/product/NotFound", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
	NewClient()
}
