package main

import (
	"log"
	"net/http"
	"text/template"
)

func index(w http.ResponseWriter, r *http.Request) {
	/*
		fmt.Println(template.HTMLEscapeString("<script>alert()</script>"))
		template.HTMLEscape(w, []byte("<script>alert()</script>"))
	*/
	t, _ := template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
	t.ExecuteTemplate(w, "T", "<script>alert('you have been pwned')</script>")
}

func main() {
	http.HandleFunc("/index", index)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
