package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Khan/genqlient/graphql"
	"github.com/natawatpak/Mech-Mobile-M-2-/backend/graph"
)

func main() {
	// fs := http.FileServer(http.Dir("./frontend/dist"))
	// http.Handle("/", fs)

	// fmt.Println("Server listening on port 3000")
	// log.Panic(
	// 	http.ListenAndServe(":3000", nil),
	// )

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	graphqlClient := graphql.NewClient("http://localhost:8081/query", http.DefaultClient)


	resp, err := graph.CustomerCreate(ctx, graphqlClient, &graph.CustomerCreateInput{
		FName: "hello",
		LName: "world",
		Tel:   "098123456789",
		Email: "hello@gmail.com",
	})
	if err != nil {
		log.Fatal(err)
	}
	println(resp.CustomerCreate.FName)

	respquery, err := graph.Customers(ctx, graphqlClient)
	if err != nil {
		log.Fatal(err)
	}
	println(*respquery)
}

func checkError(err error) {
	if err != nil {
		log.Panic(err)
	}
}
