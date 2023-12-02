package main

import (
	"fmt"
	"os"
	"io"
	"net/http"
	"log"
	"path"
	"flag"
	"sync"
)

var mu sync.Mutex

const maxUploadSize = 32 << 20 // 32 MB

// условная бд с методами
type BD map[string]string

// добавить в бд
func (p *BD) Add(key, value string) {
	(*p)[key] = value
}

// проверить, записано ли уже в бд значение по этому ключу
func (p *BD) IsExist(key string) bool {
	if _, ok := (*p)[key]; ok {
		return true
	}
	return false
}

// получить значение по ключу
func (p *BD) Get(key string) string {
	return (*p)[key]
}

// объявить пустую бд
var bd = BD{}

type Form struct {
	Form []byte
}

type ImgDir struct {
	ImgDir string
}

var dirPath = ""

func main() {

	// флаг каталога для изображений
	imgDirPtr := flag.String("images-dir", "./images", "catalog for images")
	
	// флаг файла с формой
	formFilePtr := flag.String("form-file", "", "form file")

	flag.Parse()

	log.Printf("Received command-line arguments: a directory for images \"%s\" and a file with a form \"%s\"\n", *imgDirPtr, *formFilePtr)

	imgDir := *imgDirPtr
	// в начале имени каталога не должно быть . или /
	if imgDir[0] == '.' {
		imgDir = imgDir[1:len(imgDir)]
	}
	if imgDir[0] == '/' {
		imgDir = imgDir[1:len(imgDir)]
	}
	// в конце имени каталога не должно быть /
	if imgDir[len(imgDir) - 1] == '/' {
		imgDir = imgDir[:len(imgDir) - 2]
	}

	log.Printf("Processed directory path for images: %s -> %s\n", *imgDirPtr, imgDir)
	
	dirPath = imgDir
	
	formFile := *formFilePtr
	// файл с формой должен быть указан в аргументах командной строки при запуске сервера
	if len(formFile) == 0 {
		fmt.Fprintf(os.Stderr, "there is no html file with the form in the command line args\n")
		os.Exit(1)
	}

	// прочитать содержимое файла с формой в срез байт
	f, err := os.ReadFile(formFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	log.Printf("Read the content of the form in a byte slice\n")
	
	// записать срез с формой в структуру
	formHandler := &Form{Form: f}
	http.Handle("/", formHandler)
	
	http.HandleFunc("/upload", upload)

	// получаем доступ к содержимому файловой системы сервера
	// getter := http.StripPrefix("/images/", http.FileServer(http.Dir("./images")))
	// http.Handle("/images/", getter)

	imgHandler := &ImgDir{ImgDir: imgDir}
	http.Handle("/images/", imgHandler)
	
	log.Fatal(http.ListenAndServe("10.0.2.4:5000", nil))
}

func (h *ImgDir) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// извлечь из пути ключ для поиска по бд
	key := path.Base(r.URL.String())
	
	log.Printf("A key for database search has been retrieved from the path: %s\n", key)
	
	file := bd.Get(key)
	
	log.Printf("A file was retrieved from the database: %s\n", file)

	path := (*h).ImgDir + "/" + file
	
	// отвечает на запрос содержимым файла
	http.ServeFile(w, r, path)
	
	log.Printf("The contents of the \"%s\" have been sent to the client\n", path)
}

func (h *Form) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("PATH: %s\n", r.URL.String())
	// записать содержимое формы в сокет клиента
	w.Write((*h).Form)
	log.Printf("The contents of the form are written to the client's socket\n")
}

func upload(w http.ResponseWriter, r *http.Request) {
	// от клиента должен прийти POST запрос
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
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
		http.Error(w, "can not parse form: " + err.Error(), http.StatusBadRequest)
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
		http.Error(w, "can not get file from form: " + err.Error(), http.StatusBadRequest)
		return
	}

	// закрывает файл, делая его непригодным для ввода/вывода
	defer file.Close()

	// создать каталог для файлов, если отсутствует
	err = os.MkdirAll("./" + dirPath, 0o744)
	if err != nil {
		http.Error(w, "can not create dir for images: " + err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("Created a directory \"%s\" for files if missing\n", dirPath)
	
	// создать файл
	// func Create(name string) (*File, error)
	dst, err := os.Create(fmt.Sprintf("./" + dirPath + "/" + "%s", fileheader.Filename))
	if err != nil {
		http.Error(w, "can not create file for image: " + err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("A file \"%s\" has been created on disk\n", fileheader.Filename)
	
	defer dst.Close()

	// сохранить полученный файл на диске (записать содержимое в созданный на диске файл)
	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, "can not copy images to file on disk: " + err.Error(), http.StatusInternalServerError)
		return	
	}

	log.Printf("The image received from the client is written to a file on disk\n")

	// сгенирировать ключ
	key := createKey(fileheader.Filename)

	log.Printf("To store the file in the hash table, a key \"%s\" is generated\n", key)

	// во время записи в хеш, в созданную структуру записывается указатель на весь связанный список, сама же структура записывается в массив по индексу
	// при одновременной записи структур по одному индексу можно потерять какую-либо из новых структур
	mu.Lock()
	// записать в бд имя файла
	bd.Add(key, fileheader.Filename)
	mu.Unlock()

	log.Printf("The file is written to the database\n")
	
	// сформировать ссылку для пользователя
	scheme := "http://"
	addr := r.Host
	dir := "/" + dirPath + "/"
	fileName := key
	userLink := scheme + addr + dir + fileName

	log.Printf("A link \"%s\" for the user has been generated\n", userLink)
	
	fmt.Println("Вот по этой ссылке можно показать вашу картинку:", userLink)

	fmt.Fprintf(w, "Upload successful")
}

// генерирую ключ левой пяткой правой ноги (неоптимальный алгоритм с высокой вероятностью ошибок)
func createKey(str string) string {
	res := ""
	sum := 0
	for _, val := range str {
		for _, v := range str {
			sum += int(v)
		}
		c := int(val) + 6 * sum
		// случайное числовое значение должно соответствовать букве английского алфавита
		for c < 65 || c > 90 && c < 97 || c > 122 {
			if c < 122 {
				c += 5
			} else {
				c /= 2
			}
		}
		res += string(c)
	}

	// итоговый результат должен включать 6 букв английского алфавита
	x, y := 0, 6
	for {
		// в бд не должно быть одинаковых ключей
		if bd.IsExist(res[x:y]) {
			x++
			y++
		}
		break
	}
	// вернуть готовый ключ, по которому будет храниться имя файла
	return res[x:y]
}
