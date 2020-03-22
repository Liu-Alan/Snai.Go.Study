package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

type Friend struct {
	Fname string
}

type Person struct {
	UserName string
	Emails   []string
	Friends  []*Friend
}

func EmailDealWith(args ...interface{}) string {
	ok := false
	var s string
	if len(args) == 1 {
		s, ok = args[0].(string)
	}
	if !ok {
		s = fmt.Sprint(args...)
	}

	substrs := strings.Split(s, "@")
	if len(substrs) != 2 {
		return s
	}

	return (substrs[0] + " at " + substrs[1])
}

func index(w http.ResponseWriter, r *http.Request) {
	f1 := Friend{Fname: "minux.ma"}
	f2 := Friend{Fname: "xushiwei"}

	p := Person{
		UserName: "Alan",
		Emails:   []string{"Alan@qq.com", "snail@qq.com"},
		Friends:  []*Friend{&f1, &f2}}

	funcMap := template.FuncMap{"emailDeal": EmailDealWith}
	tmpl := template.New("index.html").Funcs(funcMap)
	tmpl, _ = tmpl.ParseFiles("index.html")
	tmpl.Execute(w, p)
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/index", index)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
