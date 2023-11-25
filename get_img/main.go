package main

import (
	"fmt"
	"os"
	"io"
	"net/http"
	"log"
)

const MAX_UPLOAD_SIZE = 32 << 20 // 32 MB

// условная бд
var bd = map[string]string{}

func main() {

	// файл с формой должен быть указан в аргументах командной строки при запуске сервера
	if len(os.Args) == 1 {
		fmt.Fprintf(os.Stderr, "there is no html file with the form in the command line args\n")
		os.Exit(1)
	}
	http.HandleFunc("/", handler)
	http.HandleFunc("/upload", upload)

	// получаем доступ к содержимому файловой системы сервера
	// getter := http.StripPrefix("/images/", http.FileServer(http.Dir("./images")))
	// http.Handle("/images/", getter)

	http.HandleFunc("/images/", getter)
	
	log.Fatal(http.ListenAndServe("10.0.2.4:5000", nil))

}

func getter(w http.ResponseWriter, r *http.Request) {
	// извлечь из пути ключ для поиска по бд
	lastSlash := 0
	for i, val := range r.URL.String() {
		if val == '/' {
			lastSlash = i
		}
	}
	key := r.URL.String()[lastSlash + 1:]

	// отвечает на запрос содержимым файла
	http.ServeFile(w, r, "images/" + bd[key])
}

func handler(w http.ResponseWriter, r *http.Request) {
	// прочитать содержимое файла в срез байт
	f, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	// записать содержимое среза в сокет клиента
	w.Write(f)
}

func upload(w http.ResponseWriter, r *http.Request) {
	// от клиента должен прийти POST запрос
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// задает размер входящего буфера, больше которого из сети считывать не надо
	r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE)

	// метод, задает буфер для обработки maxMemory байт тела запроса в памяти, остальное временно хранит на диске
	err := r.ParseMultipartForm(MAX_UPLOAD_SIZE)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

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
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// закрывает файл, делая его непригодным для ввода/вывода
	defer file.Close()

	// создать каталог для файлов, если отсутствует
	err = os.MkdirAll("./images", 0744)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	// создать файл
	// func Create(name string) (*File, error)
	dst, err := os.Create(fmt.Sprintf("./images/%s", fileheader.Filename))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	// сохранить полученный файл на диске (записать содержимое в созданный на диске файл)
	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return	
	}

	// сгенирировать ключ
	key := createKey(fileheader.Filename)
	
	// записать в бд имя файла
	bd[key] = fileheader.Filename
	
	// сформировать ссылку для пользователя
	userLink := "http://127.0.0.1:5000/images/" + key
	
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
		if _, ok := bd[res[x:y]]; ok {
			x++
			y++
		}
		break
	}
	// вернуть готовый ключ, по которому будет храниться имя файла
	return res[x:y]
}
