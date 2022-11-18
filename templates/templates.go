package templates

import (
	"embed"
	"html/template"
)

//go:embed *.html
var Templates embed.FS

// parsing template diluar agar hanya sekali di eksekusi dan bisa digunakan dimana saja
var MyTemplates = template.Must(template.ParseFS(Templates, "*.html"))
