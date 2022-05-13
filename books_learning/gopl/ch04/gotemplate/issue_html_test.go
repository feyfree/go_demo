package gotemplate

import (
	"go_demo/books_learning/gopl/ch04/github"
	"log"
	"os"
	"testing"
)

//!+gotemplate
import "html/template"

var issueList = template.Must(template.New("issuelist").Parse(`
<h1>{{.TotalCount}} gotemplate</h1>
<table>
<tr style='text-align: left'>
  <th>#</th>
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
</tr>
{{end}}
</table>
`))

//!-gotemplate

//!+
func TestHtml(t *testing.T) {
	//result, err := github.SearchIssues(os.Args[1:])
	terms := []string{"repo:golang/go", "is:open", "json", "decoder"}
	result, err := github.SearchIssues(terms)
	if err != nil {
		log.Fatal(err)
	}
	if err := issueList.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}
