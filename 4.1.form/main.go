package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.html")
		log.Println(t.Execute(w, nil))
	} else {
		r.ParseForm()
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])

		/*
			fmt.Println("username:", r.FormValue("username"))
			fmt.Println("username:", r.FormValue("password"))
		*/
		/*
			v := url.Values{}
			v.Set("name", "Ava")
			v.Add("friend", "Jess")
			v.Add("friend", "Sarah")
			v.Add("friend", "Zoe")
			fmt.Println(v.Get("name"))
			fmt.Println(v.Get("friend"))
			fmt.Println(v["friend"])
		*/
	}
}

func main() {
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
