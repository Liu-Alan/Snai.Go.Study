package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello,%s!\n", ps.ByName("name"))
}

func GetUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	uid := ps.ByName("uid")
	fmt.Fprintf(w, "you are get user %s", uid)
}

func AddUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	uid := ps.ByName("uid")
	fmt.Fprintf(w, "you are add user %s", uid)
}

func DeleteUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	uid := ps.ByName("uid")
	fmt.Fprintf(w, "you are delete user %s", uid)
}

func ModifyUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	uid := ps.ByName("uid")
	fmt.Fprintf(w, "you are mmodify user %s", uid)
}

func main() {
	router := httprouter.New()

	router.GET("/", Index)
	router.GET("/hello/:name", Hello)

	router.GET("/user/:uid", GetUser)
	router.POST("/adduser/:uid", AddUser)
	router.DELETE("/deluser/:uid", DeleteUser)
	router.PUT("/moduser/:uid", ModifyUser)

	err := http.ListenAndServe(":9090", router)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
