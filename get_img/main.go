package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	// флаг каталога для изображений
	imgDirPtr := flag.String("images-dir", "./images", "catalog for images")

	// флаг файла с формой
	formFilePtr := flag.String("form-file", "", "form file")

	flag.Parse()

	log.Printf("Received command-line arguments: a directory for images \"%s\" and a file with a form \"%s\"\n", *imgDirPtr, *formFilePtr)

	imgDir := *imgDirPtr
	// в начале имени каталога не должно быть . или /
	// if imgDir[0] == '.' {
		// imgDir = imgDir[1:len(imgDir)]
	// }
	// if imgDir[0] == '/' {
		// imgDir = imgDir[1:len(imgDir)]
	// }
	// в конце имени каталога не должно быть /
	// if imgDir[len(imgDir)-1] == '/' {
		// imgDir = imgDir[:len(imgDir)-2]
	// }

	log.Printf("Processed directory path for images: %s -> %s\n", *imgDirPtr, imgDir)

	formFile := *formFilePtr
	// файл с формой должен быть указан в аргументах командной строки при запуске сервера
	if len(formFile) == 0 {
		fmt.Fprintf(os.Stderr, "there is no html file with the form in the command line args\n")
		os.Exit(1)
	}

	// объявили пустую бд
	db := NewDB()

	// объявили экземпляр структуры с данными формы, каталога для картинок, бд
	webServer, err := NewWeb(formFile, imgDir, db)
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

	log.Fatal(http.ListenAndServe(":5000", nil))
}

