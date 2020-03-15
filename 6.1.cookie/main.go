package main

import (
	"fmt"
	"log"
	"net/http"
)

func setcookie(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{Name: "username", Value: "snail"}
	cookie.HttpOnly = true
	cookie.Secure = false
	http.SetCookie(w, &cookie)
	fmt.Fprintf(w, "cookie 设置成功")
}

func getcookie(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("username")
	fmt.Fprint(w, cookie.Value)
}

func main() {
	http.HandleFunc("/setcookie", setcookie)
	http.HandleFunc("/getcookie", getcookie)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
