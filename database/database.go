package database

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

var mongoConnection *mongo.Database = nil

func Initialize() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}

	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:27017", os.Getenv("DB_CONNECTION_STRING")))

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	mongoConnection = client.Database(os.Getenv("DB_NAME"))

	fmt.Println("Connected to MongoDB!")
}

func GetConnection() *mongo.Database {

	if mongoConnection == nil {
		Initialize()
	}

	return mongoConnection
}

/*
func ModelStructFieldTagResolver(tag string, a reflect.Value, t reflect.Type) map[string]interface{} {
	var structFields []string

	// Get current struct field names.
	for i := 0; i < a.Elem().NumField(); i++ {
		structFields = append(structFields, a.Elem().Type().Field(i).Name)
	}

	resolved := make(map[string]interface{})

	// Iterate each field name on struct. Then get "tag" struct tag.
	for _, fieldName := range structFields {
		field, found := t.FieldByName(fieldName)
		if !found {
			continue
		}

		val := reflect.Indirect(a).FieldByName(fieldName).Interface()

		if val == "" || val == nil {
			continue
		}

		resolved[field.Tag.Get(tag)] = val
	}

	return resolved
}*/
