package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
)

func main() {
	fs := http.FileServer(http.Dir("./frontend/dist"))
	http.Handle("/", fs)
	http.HandleFunc("/hello", hello)

	fmt.Println("Server listening on port 3000")
	log.Panic(
		http.ListenAndServe(":3000", nil),
	)
}

func hello(w http.ResponseWriter, r *http.Request) {
	txt := "hello world"
	json.NewEncoder(w).Encode(txt)
}

func checkError(err error) {
	if err != nil {
		log.Panic(err)
	}
}