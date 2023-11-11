package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./htdocs")))

	fmt.Println("http://localhost:80")
	http.ListenAndServe(":80", nil)
}
