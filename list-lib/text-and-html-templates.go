package main

import (
	"exercise410"
	"html/template"
	"log"
	"os"
	"time"
)

// var issueList = template.Must(template.New("issuelist").Parse(`
// <h1>{{.TotalCount}} issues</h1>
// <table>
// <tr style='text-align: left'>
// <th>#</th>
// <th>State</th>
// <th>User</th>
// <th>Title</th>
// </tr>
// {{range .Items}}
// <tr>
// <td><a href='{{.Number | ChangeNumber}} number</td>
// <td>{{.State | ChangeString}} states</td>
// <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
// <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
// </tr>
// {{end}}
// </table>
// `))

const templ = `{{.TotalCount}} issues:
{{range .Items}}----------------------------------------
Number: {{.Number}}
User:
{{.User.Login}}
Title: {{.Title | printf "%.64s"}}
Age:
{{.CreatedAt | daysAg}} 
{{end}}`

func main() {
	report, err := template.New("report").
		Funcs(template.FuncMap{"daysAg": daysAg}).
		Parse(templ)
	if err != nil {
		log.Fatal(err)
	}
	result, err := exercise410.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}

}

func daysAg(t time.Time) int {
	return 30
}
func ChangeString(s string) string {
	return s + "change"
}
func ChangeNumber(i int) int {
	return i * 2
}
