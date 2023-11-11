package main

import (
	"fmt"
	"text/template"
	"log"
)

func main() {
	const templ = `{{.TotalCount}} тем
	{{range .Items}}
	Number: {{.Number}}
	User: {{.User}}
	Title: {{.Title | printf "%.64"}}
	Age: {{.CreatedAt | daysAgo}} days
	{{end}}`

	// создает шаблон, добавляет к набору доступных функций daysAgo, применяет функцию Parse
	// report, err := template.New("report").Funcs(template.FuncMap{"daysAgo": daysAgo}).Parse(templ)
	// if err != nil {
		// log.Fatal(err)
	// }
	report := template.Must(template.New("report").Funcs(template.FuncMap{"daysAgo": daysAgo}).Parse(templ))
	
}

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}
