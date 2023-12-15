package main

import (
	"flag"
	"github.com/Kostushka/share-images/internal/db"
	"github.com/Kostushka/share-images/internal/web"
	"log"
	"os"
)

func main() {

	// получить конфигурационные данные
	port, imgDir, formFile, URIDb, nameDb, nameCollection := configParse()

	// определили пустую бд с коллекцией
	db, err := db.NewDB(URIDb, nameDb, nameCollection)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("An empty db is defined")

	// объявили экземпляр структуры с данными формы, каталога для картинок, бд
	webServer, err := web.NewWeb(*formFile, *imgDir, db)
	if err != nil {
		log.Fatalf("cannot init webServer: %v", err)
	}

	// запуск слушателя и обработчика клиентских запросов
	log.Fatal(webServer.Run(*port))
}

func configParse() (portPtr, imgDirPtr, formFilePtr, URIDb, nameDb, nameCollection *string) {
	// флаг порта, на котором будет слушать запущенный сервер
	portPtr = flag.String("port", "5000", "port for listen")

	// флаг каталога для изображений
	imgDirPtr = flag.String("images-dir", "./images", "catalog for images")

	// флаг файла с формой
	formFilePtr = flag.String("form-file", "", "form file")

	// адрес для запуска процесса работы с бд
	URIDb = flag.String("URI-db", "mongodb://localhost:27017", "URI for database")

	// название бд
	nameDb = flag.String("name-db", "service", "database name")

	// название коллекции в бд
	nameCollection = flag.String("name-collection", "images", "collection name")

	flag.Parse()

	log.Printf("Received command-line arguments: port %q\na directory for images %q\n"+
		"a file with a form %q\nURI for database %q\ndatabase name %q\ncollection name %q",
		*portPtr, *imgDirPtr, *formFilePtr, *URIDb, *nameDb, *nameCollection)

	// порт должен быть корректным
	if (*portPtr)[0] != ':' {
		*portPtr = ":" + *portPtr
	}

	// файл с формой должен быть указан в аргументах командной строки при запуске сервера
	if len(*formFilePtr) == 0 {
		log.Printf("There is no html file with the form in the command line args")
		os.Exit(1)
	}

	return
}
