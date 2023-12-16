package main

import (
	"strconv"
	"flag"
	"fmt"
	"github.com/Kostushka/share-images/internal/db"
	"github.com/Kostushka/share-images/internal/web"
	"log"
)

func main() {

	var conf config
	// получить конфигурационные данные
	if err := configParse(&conf); err != nil {
		log.Fatal("cannot get config data: %v", err)
	}
	
	// определили пустую бд с коллекцией
	db, err := db.NewDB(conf.URIDb, conf.nameDb, conf.nameCollection)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("db %s is defined", conf.nameDb)

	// объявили экземпляр структуры с данными формы, каталога для картинок, бд
	webServer, err := web.NewWeb(conf.formFile, conf.imgDir, db)
	if err != nil {
		log.Fatalf("cannot init webServer: %v", err)
	}

	// запуск слушателя и обработчика клиентских запросов
	log.Fatal(webServer.Run(conf.port))
}

type config struct {
	port string
	imgDir string
	formFile string
	URIDb string
	nameDb string
	nameCollection string
}

func configParse(conf *config) error {
	
	// флаг порта, на котором будет слушать запущенный сервер
	var port int
	flag.IntVar(&port, "port", 5000, "port for listen")

	// флаг каталога для изображений
	flag.StringVar(&conf.imgDir, "images-dir", "./images", "catalog for images")

	// флаг файла с формой
	flag.StringVar(&conf.formFile, "form-file", "", "form file")

	// адрес для запуска процесса работы с бд
	flag.StringVar(&conf.URIDb, "URI-db", "mongodb://localhost:27017", "URI for database")

	// название бд
	flag.StringVar(&conf.nameDb, "name-db", "service", "database name")

	// название коллекции в бд
	flag.StringVar(&conf.nameCollection, "name-collection", "images", "collection name")

	flag.Parse()

	// порт должен быть корректным
	if port < 0 || port > 65535 {
		return fmt.Errorf("port invalid")
	}
	conf.port = ":" + strconv.Itoa(port)

	// файл с формой должен быть указан в аргументах командной строки при запуске сервера
	if len(conf.formFile) == 0 {
		fmt.Errorf("There is no html file with the form in the command line args")
	}

	log.Printf("Received command-line arguments: port %q\na directory for images %q\n"+
			"a file with a form %q\nURI for database %q\ndatabase name %q\ncollection name %q",
			conf.port, conf.imgDir, conf.formFile, conf.URIDb, conf.nameDb, conf.nameCollection)

	return nil
}
