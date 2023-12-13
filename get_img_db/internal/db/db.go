package db

import (
	"context"
	"log"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// структура данных объекта коллекции
type ImageData struct {
	Filename string
	File []byte
}

// бд с методами
type DB struct {
	db *mongo.Collection
	mu sync.RWMutex
}

// добавить в бд
func (p *DB) Set(filename string, file []byte) {
	// блокировка для всех: никому больше нельзя читать и писать
	p.mu.Lock()
	defer p.mu.Unlock()
	
	// экземпляр объекта коллекции
	image := ImageData{
		Filename: filename,
		File: file,
	}
	
	// добавить картинку в коллекцию
	_, err := p.db.InsertOne(context.TODO(), image)

	if err != nil {
		log.Fatal(err)
	}
}

// проверить, записан ли уже в бд файл по имени картинки
func (p *DB) IsExist(filename string) bool {
	// блокировка на чтение: никому нельзя писать, но можно читать одновременно нескольким клиентам
	p.mu.RLock()
	defer p.mu.RUnlock()

	// найти картинку в бд и извлечь
	var resImage bson.M

	err := p.db.FindOne(context.TODO(), bson.D{{"filename", filename}}).Decode(&resImage)

	if err != nil {
		log.Fatal(err)
	}
	
	if resImage != nil {
		return true
	} else {
		return false
	}
}

// получить файл по имени картинки
func (p *DB) Get(filename string) []byte {
	// блокировка на чтение: никому нельзя писать, но можно читать одновременно нескольким клиентам
	p.mu.RLock()
	defer p.mu.RUnlock()

	// найти картинку в бд и извлечь
	var resImage bson.M

	err := p.db.FindOne(context.TODO(), bson.D{{"filename", filename}}).Decode(&resImage)

	if err != nil {
		log.Fatal(err)
	}

	m, err := bson.Marshal(resImage)

	if err != nil {
		log.Fatal(err)
	}

	_, data := bson.Raw(m).Lookup("file").Binary()
	
	return data
}

// функция-конструктор
func NewDB() *DB {
	collection := ConnectToDB()
	return &DB{
		db: collection,
	}
}

// создание подключения и инициализация коллекции бд
func ConnectToDB() *mongo.Collection {

	// установить параметры клиента
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	
	// подключиться к MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// проверить соединение
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	// создать коллекцию для картинок
	imagesCollection := client.Database("db").Collection("images")

	return imagesCollection
}
