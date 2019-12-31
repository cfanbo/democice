package main

import (
    "net/http"
	"log"  
	"fmt"
)

func myHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello there!\n")
}
func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world\n")
}

func main(){
	http.HandleFunc("/", myHandler)		//	设置访问路由
	http.HandleFunc("/hello", sayHello)

    log.Fatal(http.ListenAndServe(":8080", nil))
}
