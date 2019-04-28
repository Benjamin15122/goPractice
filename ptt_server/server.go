package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type Commits struct {
	Commit1 string `json:"commit1"`
	Commit2 string `json:"commit2"`
}

type Profile struct {
	Name    string   `json:"name"`
	Hobbies []string `json:"hobbies"`
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析参数，默认是不会解析的
	// fmt.Println(r.Form)  //这些信息是输出到服务器端的打印信息
	// fmt.Println("path", r.URL.Path)
	// fmt.Println("scheme", r.URL.Scheme)
	// fmt.Println(r.Form["url_long"])
	var c1 string
	var c2 string
	for k, v := range r.Form {
		fmt.Println(k)
		if k == "commit1" {
			c1 = strings.Join(v, "")
			fmt.Println(c1)
		}
		if k == "commit2" {
			c2 = strings.Join(v, "")
			fmt.Println(c2)
		}
	}

	res := Commits{c1, c2}

	js, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// fmt.Fprintf(w, "Hello astaxie!") //这个写入到w的是输出到客户端的
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func main() {
	http.HandleFunc("/", sayhelloName)       //设置访问的路由
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
