package main

import (
	"fmt"
	"log"
	"net/http"
)

// MaxAge=0或未设置 表示未设置Max-Age属性，客户端会话期有效，关闭客户端删除
// MaxAge<0 表示立刻删除该cookie，等价于"Max-Age: 0"
// MaxAge>0 表示存在Max-Age属性，单位是秒

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
