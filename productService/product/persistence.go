package product

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func NewClient() {
	log.Println("Connecting to database...")
	var err error
	var url string
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)

	if os.Getenv("MONGO_INITDB_ROOT_PORT") != "" {
		url = fmt.Sprintf("mongodb://%s:%s@%s:%s", os.Getenv("MONGO_INITDB_ROOT_USERNAME"), os.Getenv("MONGO_INITDB_ROOT_PASSWORD"), os.Getenv("MONGO_INITDB_ROOT_HOST"), os.Getenv("MONGO_INITDB_ROOT_PORT"))
	} else {
		url = fmt.Sprintf("mongodb+srv://%s:%s@%s", os.Getenv("MONGO_INITDB_ROOT_USERNAME"), os.Getenv("MONGO_INITDB_ROOT_PASSWORD"), os.Getenv("MONGO_INITDB_ROOT_HOST"))
	}

	clientOptions := options.Client().ApplyURI(url).SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	Client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
}

// GetProduct returns a product from the database. If the product does not exist, it returns an empty product.
func GetProduct(database string,
	collection string,
	key string,
	value string) (Product, error) {

	var product Product
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err := Client.Database(database).Collection(collection).FindOne(ctx, bson.M{key: value}).Decode(&product)
	if err != nil {
		log.Println("Product not found: ", err)
		return product, err
	}
	return product, nil
}

// PostProduct adds a product to the database.
func PostProduct(database string,
	collection string,
	product *Product) (*mongo.InsertOneResult, error) {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	result, err := Client.Database(database).Collection(collection).InsertOne(ctx, product)
	if err != nil {
		log.Println("Product could not be added: ", err)
		return result, err
	}
	return result, nil
}

// DeleteProduct deletes a product from the database.
func DeleteProduct(database string,
	collection string,
	key string,
	value string) (*mongo.DeleteResult, error) {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	result, err := Client.Database(database).Collection(collection).DeleteOne(ctx, bson.M{key: value})
	if err != nil {
		log.Println("Deletion failed: ", err)
		return result, err
	}
	return result, nil
}
