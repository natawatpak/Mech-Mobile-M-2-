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
	http.HandleFunc("/create-user", controller.CustomerCreateProfile)
	http.HandleFunc("/update-user", controller.CustomerUpdateProfile)

	fmt.Println("Server listening on port 3000")
	log.Panic(
		http.ListenAndServe(":3000", nil),
	)

}

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "hello world")
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Panic(err)
	}
}
