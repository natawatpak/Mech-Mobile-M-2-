package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/natawatpak/Mech-Mobile-M-2-/main/controller"
)

func main() {
	fs := http.FileServer(http.Dir("./frontend/dist"))
	http.Handle("/", fs)
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/create-user", controller.CreateCustomerProfile)

	fmt.Println("Server listening on port 3000")
	log.Panic(
		http.ListenAndServe(":3000", nil),
	)
}

func checkError(err error) {
	if err != nil {
		log.Panic(err)
	}
}
