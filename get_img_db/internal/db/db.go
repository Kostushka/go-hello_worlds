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
	Key string
	File []byte
}

// бд с методами
type DB struct {
	db *mongo.Collection
	mu sync.RWMutex
}

// добавить в бд
func (p *DB) Set(filename, key string, file []byte) error {
	// блокировка для всех: никому больше нельзя читать и писать
	p.mu.Lock()
	defer p.mu.Unlock()
	
	// экземпляр объекта коллекции
	image := ImageData{
		Filename: filename,
		Key: key,
		File: file,
	}
	
	// добавить картинку в коллекцию
	_, err := p.db.InsertOne(context.TODO(), image)

	if err != nil {
		log.Printf("The image %s has not been added to the database: %v\n", filename, err.Error())
		return err
	}

	return nil
}

// проверить, записан ли уже в бд файл по ключу картинки
func (p *DB) IsExist(key string) error {
	// блокировка на чтение: никому нельзя писать, но можно читать одновременно нескольким клиентам
	p.mu.RLock()
	defer p.mu.RUnlock()

	// bson.M — неупорядоченное представление документа BSON (порядок элементов не имеет значения)
	var resImage bson.M
	
	// найти картинку в бд и извлечь
	err := p.db.FindOne(context.TODO(), bson.D{{"key", key}}).Decode(&resImage) // bson.D - упорядоченное представление документа BSON (порядок элементов имеет значения)

	if err != nil {
		log.Printf("%v\n", err.Error())
		return err
	}

	return nil
}

// получить файл по ключу картинки
func (p *DB) Get(key string) ([]byte, error) {
	// блокировка на чтение: никому нельзя писать, но можно читать одновременно нескольким клиентам
	p.mu.RLock()
	defer p.mu.RUnlock()

	var resImage bson.M
	
	// найти картинку в бд и извлечь
	err := p.db.FindOne(context.TODO(), bson.D{{"key", key}}).Decode(&resImage)

	if err != nil {
		log.Printf("%v\n", err.Error())
		return nil, err
	}

	m, err := bson.Marshal(resImage)

	if err != nil {
		log.Printf("%v\n", err.Error())
		return nil, err
	}

	// bson.Raw(...) - полный документ bson
	// Lookup(...) - поиск в документе по ключу 
	// Binary() - возвращает двоичное знач bson
	_, data := bson.Raw(m).Lookup("file").Binary() 
	
	return data, nil
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
