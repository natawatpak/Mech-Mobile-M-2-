package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"bytes"
	"time"
)

type Customer struct {
	id    string
	fname string
	lname string
	tel   string
	email string
}

func CreateCustomerProfile(w http.ResponseWriter, r *http.Request){
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful\n")
	// fName := r.FormValue("fName")

	jsonData := map[string]string{
		"mutation":`{
			customerCreate(
				input: {
					fName: `+ r.FormValue("fName") +`,
					lName: `+ r.FormValue("lName") +`,
					tel: `+ r.FormValue("tel") +`,
					email: `+ r.FormValue("email")+`
				}
			){
				ID
				fName
				lName
				tel
				email
			}
		}`,
	}

	jsonValue,_ := json.Marshal(jsonData)
	request, err := http.NewRequest("POST", "localhost:8081/query", bytes.NewBuffer(jsonValue))
    client := &http.Client{Timeout: time.Second * 10}
    response, err := client.Do(request)
	defer response.Body.Close()
    if err != nil {
        fmt.Printf("The HTTP request failed with error %s\n", err)
    }
}
