package main

import (
	"encoding/json"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	txt := "hello world"
	json.NewEncoder(w).Encode(txt)
}
