package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"regexp"
)

func validation(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("validation.html")
		log.Println(t.Execute(w, nil))
	} else {
		r.ParseForm()
		/*
			//必填，为空判断
			if len(r.Form["input"][0]) == 0 {
				fmt.Fprintf(w, "输入为空!")
				fmt.Println("输入为空!")
			} else {
				fmt.Fprintf(w, "输入的值:%s", r.Form["input"][0])
				fmt.Println("输入的值:", r.Form["input"][0])
			}
		*/
		/*
			//数字，判断大小
			getint, err := strconv.Atoi(r.Form.Get("input"))

			if err != nil {
				fmt.Fprintf(w, "输入的不为int类形!")
				fmt.Println("输入的不为int类形!")
			} else if getint > 100 {
				fmt.Fprintf(w, "输入的值太大!")
				fmt.Println("输入的值太大!")
			} else if getint < 0 {
				fmt.Fprintf(w, "输入的值太小!")
				fmt.Println("输入的值太小!")
			} else {
				fmt.Fprintf(w, "输入的值:%d", getint)
				fmt.Println("输入的值:", getint)
			}
		*/
		/*
			//中文
			if m, _ := regexp.MatchString("^\\p{Han}+$", r.Form.Get("input")); !m {
				fmt.Fprintf(w, "输入的值不是中文!")
				fmt.Println("输入的值不是中文!")
			} else {
				fmt.Fprintf(w, "输入的值:%s", r.Form.Get("input"))
				fmt.Println("输入的值:", r.Form.Get("input"))
			}
		*/
		/*
			//英文
			if m, _ := regexp.MatchString("^[a-zA-Z]+$", r.Form.Get("input")); !m {
				fmt.Fprintf(w, "输入的值不是英文!")
				fmt.Println("输入的值不是英文!")
			} else {
				fmt.Fprintf(w, "输入的值:%s", r.Form.Get("input"))
				fmt.Println("输入的值:", r.Form.Get("input"))
			}
		*/
		/*
			//电子邮件
			if m, _ := regexp.MatchString(`^([\w\.\_\-]{2,20})@([a-z]{1,20}).([a-z]{2,4})$`, r.Form.Get("input")); !m {
				fmt.Fprintf(w, "输入的值不是邮箱!")
				fmt.Println("输入的值不是邮箱!")
			} else {
				fmt.Fprintf(w, "输入的值:%s", r.Form.Get("input"))
				fmt.Println("输入的值:", r.Form.Get("input"))
			}
		*/
		/*
			//手机号码
			if m, _ := regexp.MatchString(`^(1[3|4|5|8][0-9]\d{8})$`, r.Form.Get("input")); !m {
				fmt.Fprintf(w, "输入的值不是手机号!")
				fmt.Println("输入的值不是手机号!")
			} else {
				fmt.Fprintf(w, "输入的值:%s", r.Form.Get("input"))
				fmt.Println("输入的值:", r.Form.Get("input"))
			}
		*/
		/*
			//下拉菜单
			slice := []string{"apple", "pear", "banane"}

			v := r.Form.Get("fruit")
			for _, item := range slice {
				if item == v {
					fmt.Fprintf(w, "选的值:%s", v)
					fmt.Println("选的值:", v)
					return
				}
			}

			fmt.Fprintf(w, "选的值不对")
			fmt.Println("选的值不对")
		*/
		/*
			//单选按钮
			slice := []int{1, 2}

			v, _ := strconv.Atoi(r.Form.Get("gender"))

			for _, item := range slice {
				if item == v {
					fmt.Fprintf(w, "选的值:%d", v)
					fmt.Println("选的值:", v)
					return
				}
			}

			fmt.Fprintf(w, "选的值不对")
			fmt.Println("选的值不对")
		*/
		/*
			//复选框
			slice := []string{"football", "basketball", "tennis"}
			v := r.Form["interest"]
			if len(v) <= 0 {
				fmt.Fprintf(w, "您没有选择")
				fmt.Println("您没有选择")
				return
			}

			si := make([]interface{}, len(slice))
			vi := make([]interface{}, len(v))
			for i, v := range slice {
				si[i] = v
			}
			for i, v := range v {
				vi[i] = v
			}

			a := beeku.Slice_diff(vi, si)

			if a == nil {
				fmt.Fprintf(w, "选的值:%s", v)
				fmt.Println("选的值:", v)
			} else {
				fmt.Fprintf(w, "选的值不对")
				fmt.Println("选的值不对")
			}
		*/
		/*
			//时间
			//t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
			//fmt.Fprintf(w, "Go launched at %s", t.Local())
			//fmt.Printf("Go launched at %s\n", t.Local())
			v := r.Form.Get("input")
			vt, err := time.ParseInLocation("2006-01-02 15:04:05", v, time.Local)
			if err != nil {
				fmt.Fprintf(w, "输入的时间不对")
				fmt.Println("输入的时间不对")
			} else {
				fmt.Fprintf(w, "输入的时间:%s\n", vt)
				fmt.Fprintf(w, "输入的时间:%s\n", vt.Format("2006-01-02 15:04:05"))
				fmt.Println("输入的时间:", vt)
				fmt.Println("输入的时间:", vt.Format("2006-01-02 15:04:05"))
			}
		*/
		//身份证号
		v := r.Form.Get("input")
		var matchString string
		if len(v) == 15 {
			matchString = `^(\d{15})$`
		} else if len(v) == 18 {
			matchString = `^(\d{17})([0-9]|X)$`
		} else {
			fmt.Fprintf(w, "输入的值不是身份证号!")
			fmt.Println("输入的值不是身份证号!")
			return
		}
		if m, _ := regexp.MatchString(matchString, v); !m {
			fmt.Fprintf(w, "输入的值不是身份证号!")
			fmt.Println("输入的值不是身份证号!")
		} else {
			fmt.Fprintf(w, "输入的值:%s\n", v)
			fmt.Println("输入的值:", v)
		}
	}
}

func main() {
	http.HandleFunc("/validation", validation)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
