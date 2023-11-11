package main

import (
	"time"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"os"
	"log"
	// "text/template"
	"html/template"
)

const IssuesURL = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items []*Issue
}

type Issue struct {
	Number int
	HTMLURL string `json:"html_url"`
	Title string
	State string
	User *User
	CreatedAt time.Time `json:"created_at"`
	Body string
}

type User struct {
	Login string
	HTMLURL string `json:"html_url"`
}

func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("%s", resp.Status)
	}

	var result IssuesSearchResult

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}

	resp.Body.Close()
	return &result, nil
}

func main() {
	res, err := SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d\n", res.TotalCount)
	for _, item := range res.Items {
		fmt.Printf("%d %s %s\n", item.Number, item.User.Login, item.Title)
	}

	const templ = `{{.TotalCount}} тем
	{{range .Items}}
	Number: {{.Number}}
	User: {{.User.Login}}
	Title: {{.Title | printf "%.64s"}}
	Age: {{.CreatedAt | daysAgo}} days
	{{end}}`

	// создает шаблон, добавляет к набору доступных функций daysAgo, применяет функцию Parse
	// report, err := template.New("report").Funcs(template.FuncMap{"daysAgo": daysAgo}).Parse(templ)
	// if err != nil {
	// log.Fatal(err)
	// }
	// более удобная обработка ошибок, создает шаблон, добавляет к набору доступных функций daysAgo, применяет функцию Parse
	report := template.Must(template.New("report").Funcs(template.FuncMap{"daysAgo": daysAgo}).Parse(templ))
	// данные из структуры, вывод в stdout
	if err := report.Execute(os.Stdout, res); err != nil {
		log.Fatal(err)
	}

	htmlreport := template.Must(template.New("htmlreport").Parse(`
	<h1>{{.TotalCount}} тем</h1>
	<table>
	<tr style='text-align: left'>
		<th>State</th>
		<th>User</th>
		<th>Title</th>
	</tr>
	{{range .Items}}
	<tr>
		<td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
		<td>{{.State}}</td>
		<td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
		<td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
	<\tr>
	{{end}}
	</table>
	`))
	if err := htmlreport.Execute(os.Stdout, res); err != nil {
		log.Fatal(err)
	}
}

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

