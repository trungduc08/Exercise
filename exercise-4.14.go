package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"text/template"
)

const URL = "https://api.github.com/repos/trungduc08/GoTour/issues"

type IssueGitHubResult struct {
	Items []*IssueGitHub
}
type IssueGitHub struct {
	Url           string `json:"url"`
	RepositoryUrl string `json:"repository_url"`
	User          *UserGitHub
	Milestone     string `json:"milestone"`
	Labels        []*Label
}
type UserGitHub struct {
	Login   string `json:"login"`
	HTMLURL string `json:"html_url"`
}
type Label struct {
	URL  string `json:"url"`
	Name string `json:"name"`
}

var result *IssueGitHubResult
var issueListTemplate *template.Template
var a [4]int

func init() {

	var err error
	result, err = GetIssues()

	if err != nil {
		log.Fatal(err)
	}
	issueListTemplate = template.Must(template.ParseFiles("/home/hblab/work/src/github.com/user/Exercise/layout.html"))
}

func issueListHandler(w http.ResponseWriter, r *http.Request) {
	if err := issueListTemplate.Execute(w, result); err != nil {
		log.Fatal(err)
	}
}
func main() {
	for _, v := range result.Items {
		fmt.Printf("%s\n", v.Labels)
	}
	http.HandleFunc("/", issueListHandler)
	fmt.Printf("%s\n", issueListHandler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))

}
func GetIssues() (*IssueGitHubResult, error) {

	resp, err := http.Get(URL)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}
	var result IssueGitHubResult
	if err := json.NewDecoder(resp.Body).Decode(&result.Items); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil

}

var issueList = template.Must(template.New("issuelist").Parse(`
<table>
<tr style='text-align: left'>
<th>Url</th>
<th>RepositoryUrl</th>
<th>User</th>
<th>Milestone</th>
</tr>
{{range .Items}}
<tr>
<td>{{.Url}}</td>
<td>{{.RepositoryUrl}}</td>
<td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>

<td>{{.Milestone}}</td>
</tr>
{{end}}
</table>
`))
