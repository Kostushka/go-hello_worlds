package web

import (
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"text/template"
	"github.com/Kostushka/share-images/internal/db"
)

const maxUploadSize = 32 << 20 // 32 MB

type ErrorPage struct {
	Number int
	Text string
}

// структурами с данными
type Web struct {
	form []byte
	imgDir string
	db *db.DB
}

// метод, который записывает форму в сокет клиента
func (h *Web) Form(w http.ResponseWriter, r *http.Request) {
	log.Printf("PATH: %s\n", r.URL)

	// добавить иконку
	if r.URL.String() == "/favicon.ico" {
		// открыть файл с иконкой
		icon, err := os.Open("./web/ico.png")
		if err != nil {
			http.Error(w, "cannot open icon file: "+err.Error(), http.StatusBadRequest)
			log.Printf("cannot open icon file: %v\n", err.Error())
			return
		}
		defer icon.Close()

		// копировать иконку в клиентский сокет
		_, err = io.Copy(w, icon)
		if err != nil {
			http.Error(w, "cannot sent icon file to the client: "+err.Error(), http.StatusInternalServerError)
			log.Printf("cannot sent icon file to the client: %v\n", err.Error())
			return
		}
		
		log.Printf("The icon is written to the client's socket\n")
		return
	}
	
	// URL должен быть /
	if r.URL.String() != "/" {
		http.Redirect(w, r, "/", http.StatusFound)
		log.Printf("The request was redirected from address %q to address \"/\"\n", r.URL.String())
	}
	
	// записать содержимое формы в сокет клиента
	_, err := w.Write(h.form)
	if err != nil {
		http.Error(w, "cannot write form: "+err.Error(), http.StatusBadRequest)
		log.Printf("cannot write form: %v\n", err.Error())
		return
	}
	log.Printf("The contents of the form are written to the client's socket\n")
}

// метод, который записывает картинку в сокет клиента
func (h *Web) ServeImage(w http.ResponseWriter, r *http.Request) {
	// извлечь из пути ключ для поиска по бд
	key := path.Base(r.URL.String())
	
	log.Printf("A key for database search has been retrieved from the path: %s\n", key)

	// ключ должен быть в бд
	if !h.db.IsExist(key) {
		w.WriteHeader(http.StatusNotFound)
		// шаблон страницы с ошибкой
		error := ErrorPage{
			Number: http.StatusNotFound,
			Text: http.StatusText(http.StatusNotFound),
		}
		t := template.Must(template.ParseFiles("./web/error.html"))
		t.Execute(w, error)
		log.Printf("The key %q is not in the database\n", key)
		return
	}

	// получить имя файла из бд
	filename := h.db.Get(key)

	log.Printf("A filename was retrieved from the database: %s\n", filename)

	path := filepath.Join(h.imgDir, filename)

	// открыть файл
	file, err := os.Open(path)
	if err != nil {
		http.Error(w, "cannot open file: "+err.Error(), http.StatusBadRequest)
		log.Printf("cannot open file: %v\n", err.Error())
		return
	}
	defer file.Close()

	// копировать файл в клиентский сокет
	_, err = io.Copy(w, file)
	if err != nil {
		http.Error(w, "cannot sent file to the client: "+err.Error(), http.StatusInternalServerError)
		log.Printf("cannot sent file to the client: %v\n", err.Error())
		return
	}

	log.Printf("The contents of the %q have been sent to the client\n", path)
}

// функция-конструктор: создает экземпляр структуры с данными
func NewWeb(file, imgDir string, db *db.DB) (*Web, error) {
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
		http.Error(w, "cannot parse form: "+err.Error(), http.StatusBadRequest)
		log.Printf("cannot parse form: %v\n", err.Error())
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
		http.Error(w, "cannot get file from form: "+err.Error(), http.StatusBadRequest)
		log.Printf("cannot get file from form: %v\n", err.Error())
		return
	}

	// закрывает файл, делая его непригодным для ввода/вывода
	defer file.Close()

	// создать каталог для файлов, если отсутствует
	if err = os.MkdirAll(h.imgDir, 0o744); err != nil {
		http.Error(w, "cannot create dir for images: "+err.Error(), http.StatusInternalServerError)
		log.Printf("cannot create dir for images: %v\n", err.Error())
		return
	}

	log.Printf("Created a directory %q for files if missing\n", h.imgDir)

	// создать файл
	// func Create(name string) (*File, error)
	dst, err := os.Create(filepath.Join(h.imgDir, fileheader.Filename))
	if err != nil {
		http.Error(w, "cannot create file for image: "+err.Error(), http.StatusInternalServerError)
		log.Printf("cannot create file for image: %v\n", err.Error())
		return
	}
	defer dst.Close()

	log.Printf("A file %q has been created on disk\n", fileheader.Filename)


	// сохранить полученный файл на диске (записать содержимое в созданный на диске файл)
	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, "cannot copy images to file on disk: "+err.Error(), http.StatusInternalServerError)
		log.Printf("cannot copy images to file on disk: %v\n", err.Error())
		return
	}

	log.Printf("The image received from the client is written to a file on disk\n")

	// сгенирировать ключ
	key := createKey(fileheader.Filename, h.db)

	log.Printf("To store the file in the hash table, a key %q is generated\n", key)

	// записать в бд имя файла
	h.db.Set(key, fileheader.Filename)

	log.Printf("The file is written to the database\n")

	// сформировать ссылку для пользователя
	scheme := "http://"
	addr := r.Host
	dir := "/images"
	fileName := key
	userLink := scheme + addr + path.Join(dir, fileName)
	
	log.Printf("A link %q for the user has been generated\n", userLink)

	// шаблон файла с ссылкой
	t, err := template.ParseFiles("./web/link.html")
	if err != nil {
		http.Error(w, "cannot get template with link: "+err.Error(), http.StatusInternalServerError)
		log.Printf("cannot get template with link: %v\n", err.Error())
		return
	}
	// шаблон получает данные для обработки и записывается в сокет клиента
	t.Execute(w, userLink)
	
	log.Printf("A link %q has been sent to the client\n", userLink)
}

// генерирую ключ левой пяткой правой ноги (неоптимальный алгоритм с высокой вероятностью ошибок)
func createKey(str string, db *db.DB) string {
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
