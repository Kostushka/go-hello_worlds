package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/Kostushka/share-images/internal/web"
	"github.com/Kostushka/share-images/internal/db"
)

func main() {

	// флаг порта, на котором будет слушать запущенный сервер
	portPtr := flag.String("port", "5000", "port for listen")

	// флаг каталога для изображений
	imgDirPtr := flag.String("images-dir", "./images", "catalog for images")

	// флаг файла с формой
	formFilePtr := flag.String("form-file", "", "form file")

	flag.Parse()

	log.Printf("Received command-line arguments: port %q a directory for images %q and a file with a form %q\n", *portPtr, *imgDirPtr, *formFilePtr)

	port := ":" + *portPtr

	imgDir := *imgDirPtr

	formFile := *formFilePtr
	// файл с формой должен быть указан в аргументах командной строки при запуске сервера
	if len(formFile) == 0 {
		fmt.Fprintf(os.Stderr, "there is no html file with the form in the command line args\n")
		os.Exit(1)
	}

	// определили пустую бд с коллекцией
	db := db.NewDB()

	// объявили экземпляр структуры с данными формы, каталога для картинок, бд
	webServer, err := web.NewWeb(formFile, imgDir, db)
	if err != nil {
		log.Fatalf("cannot init webServer: %v\n", err)
	}
	
	log.Printf("Read the content of the form in a byte slice\n")

	http.HandleFunc("/", webServer.Form)

	http.HandleFunc("/upload", webServer.Upload)

	// получаем доступ к содержимому файловой системы сервера
	// getter := http.StripPrefix("/images/", http.FileServer(http.Dir("./images")))
	// http.Handle("/images/", getter)

	http.HandleFunc("/images/", webServer.ServeImage)

	log.Printf("Server listen and serve on port%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

