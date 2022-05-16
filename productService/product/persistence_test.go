package product

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var expected Product
var ctx context.Context

func TestCreateClient(t *testing.T) {
	client := Client
	NewClient()
	assert.NotEqual(t, Client, client)
}

func TestGetValue(t *testing.T) {
	teardown := addDummyProductToDb()
	value, _ := GetProduct(os.Getenv("MONGO_INITDB_TEST_DATABASE"), "products", "name", "Test")
	assert.Equal(t, expected, value)
	defer teardown()
}

func TestGetNonExistingValue(t *testing.T) {
	_, err := GetProduct(os.Getenv("MONGO_INITDB_TEST_DATABASE"), "products", "name", "NotExisting")
	assert.NotNil(t, err)
}

func TestDeleteMe(t *testing.T) {
	addDummyProductToDb()
	value, _ := DeleteProduct(os.Getenv("MONGO_INITDB_TEST_DATABASE"), "products", "name", "fish")
	assert.Equal(t, int64(1), value.DeletedCount)
}

func TestDeleteNoOne(t *testing.T) {
	value, _ := DeleteProduct(os.Getenv("MONGO_INITDB_TEST_DATABASE"), "products", "name", "NotExisting")
	assert.Equal(t, int64(0), value.DeletedCount)
}

func TestDeletionFail(t *testing.T) {
	Client.Disconnect(ctx)
	_, err := DeleteProduct(os.Getenv("MONGO_INITDB_TEST_DATABASE"), "products", "name", "Please")
	NewClient()
	assert.Error(t, err)

}

func TestPostProduct(t *testing.T) {
	var product Product
	product.ID = primitive.NewObjectID()
	product.UUID = primitive.NewObjectID()
	product.Name = "Post"
	product.Price = 123
	result, _ := PostProduct(os.Getenv("MONGO_INITDB_TEST_DATABASE"), "products", &product)
	DeleteProduct(os.Getenv("MONGO_INITDB_TEST_DATABASE"), "products", "name", "Post")
	assert.Equal(t, product.ID, result.InsertedID)
}

func TestPostProductFail(t *testing.T) {
	var product Product
	Client.Disconnect(ctx)
	product.ID = primitive.NewObjectID()
	product.UUID = primitive.NewObjectID()
	product.Name = "Post"
	product.Price = 123
	_, err := PostProduct(os.Getenv("MONGO_INITDB_TEST_DATABASE"), "products", &product)
	NewClient()
	assert.Error(t, err)
}
