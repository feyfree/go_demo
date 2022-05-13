package gotemplate

import (
	"html/template"
	"log"
	"os"
	"testing"
)

func TestAutoEscape(t *testing.T) {
	const templ = `<p>A: {{.A}}</p><p>B: {{.B}}</p>`
	temp := template.Must(template.New("escape").Parse(templ))
	var data struct {
		A string        // untrusted plain text
		B template.HTML // trusted HTML
	}
	data.A = "<b>Hello!</b>"
	data.B = "<b>Hello!</b>"
	if err := temp.Execute(os.Stdout, data); err != nil {
		log.Fatal(err)
	}
}
