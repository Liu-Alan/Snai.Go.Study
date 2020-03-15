package main

import (
	"fmt"
	"log"
	"net/http"

	_ "Snai.Go.Study/6.2.session/memory"
	"Snai.Go.Study/6.2.session/session"
)

var globalSession *session.Manager

func init() {
	var err error
	globalSession, err = session.NewSessionManager("memory", "goSessionID", 3600)
	if err != nil {
		fmt.Println(err)
		return
	}
	go globalSession.GC()
	fmt.Println("fd")
}

func sayHelloHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("name")
	if err == nil {
		fmt.Println(cookie.Value)
		fmt.Println(cookie.Domain)
		fmt.Println(cookie.Expires)
	}
	fmt.Fprintf(w, "Hello world!\n")
}

func login(w http.ResponseWriter, r *http.Request) {
	sess := globalSession.SessionStart(w, r)
	val := sess.Get("username")
	if val != nil {
		fmt.Println(val)
	} else {
		sess.Set("username", "Alan")
		fmt.Println("set session")
	}
}

func loginout(w http.ResponseWriter, r *http.Request) {
	globalSession.SessionDestroy(w, r)
	fmt.Println("session destroy")
}

func main() {
	http.HandleFunc("/", sayHelloHandler)
	http.HandleFunc("/login", login)
	http.HandleFunc("/loginout", loginout)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
