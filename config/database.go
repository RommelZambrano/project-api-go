package config
import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
const uri = "mongodb+srv://team6:team6@cluster0.y6tfgd9.mongodb.net/?retryWrites=true&w=majority"
func ConnectDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected to DB")
	return client
}
var DB = ConnectDB()
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection{
	collection := client.Database("librarydb").Collection(collectionName)
	return collection
}