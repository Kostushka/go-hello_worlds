package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
)

const maxUploadSize = 32 << 20 // 32 MB

// структурами с данными
type Web struct {
	form []byte
	imgDir string
	db *DB
}

// метод, который записывает форму в сокет клиента
func (h *Web) Form(w http.ResponseWriter, r *http.Request) {
	log.Printf("PATH: %s\n", r.URL)
	// записать содержимое формы в сокет клиента
	w.Write(h.form)
	log.Printf("The contents of the form are written to the client's socket\n")
}

// метод, который записывает картинку в сокет клиента
func (h *Web) ServeImage(w http.ResponseWriter, r *http.Request) {
	// извлечь из пути ключ для поиска по бд
	key := path.Base(r.URL.String())

	log.Printf("A key for database search has been retrieved from the path: %s\n", key)

	file := h.db.Get(key)

	log.Printf("A file was retrieved from the database: %s\n", file)

	path := filepath.Join(h.imgDir, file)

	// отвечает на запрос содержимым файла
	http.ServeFile(w, r, path)

	log.Printf("The contents of the \"%s\" have been sent to the client\n", path)
}

// функция-конструктор: создает экземпляр структуры с данными
func NewWeb(file, imgDir string, db *DB) (*Web, error) {
	// прочитать содержимое файла с формой в срез байт
	f, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	return &Web{
		form: f,
		imgDir: imgDir,
		db: db,
	}, nil
}

// метод, который обрабатывет POST запрос, сохраняет картинку в бд
func (h *Web) Upload(w http.ResponseWriter, r *http.Request) {
	// от клиента должен прийти POST запрос
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		log.Printf("Method %v not allowed\n", r.Method)
		return
	}
	log.Printf("Received a POST request from a client\n")

	// задает размер входящего буфера, больше которого из сети считывать не надо
	// ограничивает попытку вычитать серверу слишком большой запрос от клиента, который может уронить сервер
	r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)

	// метод, задает буфер для обработки maxMemory байт тела запроса в памяти, остальное временно хранит на диске
	// нужно, чтобы не исчерпать лимиты памяти
	err := r.ParseMultipartForm(maxUploadSize)
	if err != nil {
		http.Error(w, "can not parse form: "+err.Error(), http.StatusBadRequest)
		log.Printf("can not parse form: "+err.Error())
		return
	}
	log.Printf("Processed no more than %d bytes of the request body\n", maxUploadSize)

	// func (r *Request) FormFile(key string) (multipart.File, *multipart.FileHeader, error)
	// file - интерфейс для доступа к файлу
	// fileheader - структура с полями:
	// type FileHeader struct {
	// Filename string
	// Header   textproto.MIMEHeader // заголовок в стиле MIME
	// Size     int64
	// }
	file, fileheader, err := r.FormFile("image") // FormFile возвращает первый файл для ключа формы
	if err != nil {
		http.Error(w, "can not get file from form: "+err.Error(), http.StatusBadRequest)
		log.Printf("can not get file from form: "+err.Error())
		return
	}

	// закрывает файл, делая его непригодным для ввода/вывода
	defer file.Close()

	// создать каталог для файлов, если отсутствует
	if err = os.MkdirAll(h.imgDir, 0o744); err != nil {
		http.Error(w, "can not create dir for images: "+err.Error(), http.StatusInternalServerError)
		log.Printf("can not create dir for images: "+err.Error())
		return
	}

	log.Printf("Created a directory \"%s\" for files if missing\n", h.imgDir)

	// создать файл
	// func Create(name string) (*File, error)
	dst, err := os.Create(filepath.Join(h.imgDir, fileheader.Filename))
	if err != nil {
		http.Error(w, "can not create file for image: "+err.Error(), http.StatusInternalServerError)
		log.Printf("can not create file for image: "+err.Error())
		return
	}
	defer dst.Close()

	log.Printf("A file \"%s\" has been created on disk\n", fileheader.Filename)


	// сохранить полученный файл на диске (записать содержимое в созданный на диске файл)
	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, "can not copy images to file on disk: "+err.Error(), http.StatusInternalServerError)
		log.Printf("can not copy images to file on disk: "+err.Error())
		return
	}

	log.Printf("The image received from the client is written to a file on disk\n")

	// сгенирировать ключ
	key := createKey(fileheader.Filename, h.db)

	log.Printf("To store the file in the hash table, a key \"%s\" is generated\n", key)

	// записать в бд имя файла
	h.db.Add(key, fileheader.Filename)

	log.Printf("The file is written to the database\n")

	// сформировать ссылку для пользователя
	scheme := "http://"
	addr := r.Host
	dir := "/" + h.imgDir + "/"
	fileName := key
	userLink := scheme + addr + filepath.Join(dir, fileName)

	log.Printf("A link \"%s\" for the user has been generated\n", userLink)

	fmt.Fprintf(w, "Upload successful\n")
	fmt.Fprintf(w, "Вот по этой ссылке можно показать вашу картинку: %s", userLink)
}

// генерирую ключ левой пяткой правой ноги (неоптимальный алгоритм с высокой вероятностью ошибок)
func createKey(str string, db *DB) string {
	res := ""
	sum := 0
	for _, val := range str {
		for _, v := range str {
			sum += int(v)
		}
		c := int(val) + 6*sum
		// случайное числовое значение должно соответствовать букве английского алфавита
		for c < 65 || c > 90 && c < 97 || c > 122 {
			if c < 122 {
				c += 5
			} else {
				c /= 2
			}
		}
		res += string(byte(c))
	}

	// итоговый результат должен включать 6 букв английского алфавита
	x, y := 0, 6
	for {
		// в бд не должно быть одинаковых ключей
		if db.IsExist(res[x:y]) {
			x++
			y++
		}
		break
	}
	// вернуть готовый ключ, по которому будет храниться имя файла
	return res[x:y]
}
