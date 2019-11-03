package main

import (
	"html/template"
	"os"
)

func main() {
	tEmpty := template.New("template test")
	tEmpty = template.Must(tEmpty.Parse("Empty pipeline if demo: {{if ``}} Nothing output {{end}}"))
	tEmpty.Execute(os.Stdout, nil)

	tWithValue := template.New("templateValue")
	tWithValue = template.Must(tWithValue.Parse("value content: {{if `values`}} I will show {{end}}"))
	tWithValue.Execute(os.Stdout, nil)
}
