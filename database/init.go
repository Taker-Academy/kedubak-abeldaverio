package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Init_database() *mongo.Client {
	godotenv.Load()
	var uri = fmt.Sprintf("mongodb+srv://%s:%s@keduback.lu29ses.mongodb.net/?retryWrites=true&w=majority&appName=KeDuBack",
		os.Getenv("DB_USER"), os.Getenv("DB_PSW"))
	ClientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.Background(), ClientOptions)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connected to mongoDB")
	}
	return (client)
}
