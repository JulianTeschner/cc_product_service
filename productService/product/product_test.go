package product

import (
	// "io/ioutil"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
)

func TestMain(m *testing.M) {
	path, _ := os.Getwd()
	path = path + "/../cmd/.env"
	log.Println(path)
	err := godotenv.Load(path)
	if err != nil {
		log.Println("No .env file found, using default values")
	}

	log.SetOutput(ioutil.Discard)
	NewClient()
	teardownHandlers := setupHandlersTest()
	// Run the tests
	code := m.Run()
	// Exit with the code
	log.Println("Tear down persistence tests")
	teardownHandlers()
	defer Client.Disconnect(ctx)

	os.Exit(code)
}

func addDummyProductToDb() func() {
	dummyProduct := createDummyProduct()
	Client.Database(os.Getenv("MONGO_INITDB_ROOT_DATABASE")).Collection("products").InsertOne(ctx, &dummyProduct)
	return func() {
		Client.Database(os.Getenv("MONGO_INITDB_ROOT_DATABASE")).Collection("products").DeleteOne(ctx, bson.M{"name": "fish"})
	}
}
func createDummyProduct() Product {

	dummyProduct := Product{
		Name:  "fish",
		Price: 10,
	}
	return dummyProduct
}
