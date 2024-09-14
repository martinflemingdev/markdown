package main

import (
	"bytes"
	"fmt"
	"os"
	"text/template"
)

type NameFruits struct {
	Name   string
	Fruits []Fruit
}

type Fruit struct {
	Fruit  string
	Number string
}

func renderTable(data []NameFruits) (string, error) {
	const tmpl = `
| name    | fruit | number |
| ------- | ----- | ------ |
{{- range . }}
	{{- $name := .Name -}}
	{{- if eq (len .Fruits) 0 }}
| {{ $name }} | N/A   | N/A   |
	{{- else }}
		{{- range $index, $fruit := .Fruits }}
| {{ if eq $index 0 }}{{ $name }}{{ else }}   {{ end }} | {{ $fruit.Fruit }} | {{ $fruit.Number }} |
		{{- end }}
	{{- end }}
{{- end }}
`
	t := template.Must(template.New("table").Parse(tmpl))

	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func main() {
	data := []NameFruits{
		{Name: "Alice", Fruits: []Fruit{{Fruit: "Apple", Number: "3"}, {Fruit: "Banana", Number: "5"}}},
		{Name: "Bob", Fruits: []Fruit{{Fruit: "Orange", Number: "2"}}},
		{Name: "Charlie", Fruits: []Fruit{}},
		{Name: "Martin", Fruits: []Fruit{
			{Fruit: "Mango", Number: "1"},
			{Fruit: "Pineapple", Number: "2"},
			{Fruit: "Strawberry", Number: "3"},
			{Fruit: "Blueberry", Number: "4"},
			{Fruit: "Raspberry", Number: "5"},
		}},
	}

	result, err := renderTable(data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err = os.WriteFile("result.md", []byte(result), 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("Markdown table successfully written to result.md")
}
