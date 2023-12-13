package main

import (
	"context"
	"fmt"
	// "io"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Image struct {
	Title string
	File  []byte
}

func main() {
	arg := os.Args[1]

	file, err := os.ReadFile(arg)

	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %s\n", arg, err.Error())
	}

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println("Connected to MongoDB!")

	// создать коллекцию для картинок
	imagesColl := client.Database("test").Collection("images")

	// структура с метаданными и файлом картинки
	image := Image{
		Title: arg,
		File:  file,
	}

	// добавить данные картинки в коллекцию
	_, err = imagesColl.InsertOne(context.TODO(), image)

	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println("Inserted a single document: ", res.InsertedID)

	// найти картинку в бд и извлечь
	var resImage bson.M

	err = imagesColl.FindOne(context.TODO(), bson.D{{"title", arg}}).Decode(&resImage)

	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println("Found a single document: ", resImage)

	p, err := bson.Marshal(resImage)

	if err != nil {
		log.Fatal(err)
	}

	_, data := bson.Raw(p).Lookup("file").Binary()

	// записать в стандартный вывод двоичный файл картинки
	// io.Copy(os.Stdout, Data(data))

	_, err = os.Stdout.Write(data)

	if err != nil {
		log.Fatal(err)
	}
}

// type Data []byte
// 
// func (d Data) Read(bs []byte) (int, error) {
	// i := 0
	// for ; i < len(d); i++ {
		// bs[i] = d[i]
	// }
	// return i, io.EOF
// }
