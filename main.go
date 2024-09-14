package main

import (
	"bytes"
	"fmt"
	"os"
	"text/template"
)

type NameFruits struct {
	Name   string
	Fruits []string
}

func renderTable(data []NameFruits) (string, error) {
	const tmpl = `
| name | fruits |
| ---- | ------ |
{{- range . }}
| {{ .Name }} | {{ index .Fruits 0 }} |
{{- range $index, $fruit := .Fruits -}}
	{{- if gt $index 0 }}
|  | {{ $fruit }} |
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
		{Name: "Alice", Fruits: []string{"Apple", "Banana"}},
		{Name: "Bob", Fruits: []string{"Orange"}},
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
