package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	"strings"
)

type Customer struct {
	id    string `json:ID`
	fname string `json:fName`
	lname string `json:lName`
	tel   string `json: tel`
	email string `json: email`
}

func CreateCustomerProfile(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful\n")
	// fName := r.FormValue("fName")

	jsonData := map[string]string{
		"query": `mutation{
			customerCreate(
				input: {
					fName: "` + r.FormValue("fName") + `",
					lName: "` + r.FormValue("lName") + `",
					tel: "` + r.FormValue("tel") + `",
					email: "` + r.FormValue("email") + `"
				}
			)
		{ID}}`,
	}

	jsonValue, _ := json.Marshal(jsonData)

	request, err := http.NewRequest("POST", "http://localhost:8081/query", bytes.NewBuffer(jsonValue))
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{Timeout: time.Second * 10}
	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}
	
	body,_ := ioutil.ReadAll(response.Body)
	response.Body.Close()
	sbody := strings.Split(string(body), `ID":"`)
	id := strings.Trim(sbody[1], `""}`)
	fmt.Println(string(id))

}
