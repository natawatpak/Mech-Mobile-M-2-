package controller

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Khan/genqlient/graphql"
	"github.com/natawatpak/Mech-Mobile-M-2-/backend/graph"
)

func CreateCustomerProfile(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful\n")
	// fName := r.FormValue("fName")

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	graphqlClient := graphql.NewClient("http://localhost:8081/query", http.DefaultClient)

	resp, err := graph.CustomerCreate(ctx, graphqlClient, &graph.CustomerCreateInput{
		FName: r.FormValue("fName"),
		LName: r.FormValue("lName"),
		Tel:   r.FormValue("tel"),
		Email: r.FormValue("email"),
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.CustomerCreate.FName)

	// respquery, err := graph.CustomerByID(ctx, graphqlClient, "9fdd0d20-7d45-42b6-afe4-036189f5216b")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(*&respquery.CustomerByID.FName)

}

func UpdateCustomerProfile(w http.ResponseWriter, r *http.Request){
	r.ParseForm()

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	graphqlClient := graphql.NewClient("http://localhost:8081/query", http.DefaultClient)

	resp, err := graph.CustomerUpdateMulti(ctx, graphqlClient, &graph.CustomerUpdateInput{
		ID: r.FormValue("ID"),
		FName: r.FormValue("fName"),
		LName: r.FormValue("lName"),
		Tel:   r.FormValue("tel"),
		Email: r.FormValue("email"),
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.CustomerUpdateMulti.FName)
	println(resp.CustomerUpdateMulti.ID)

	respquery, err := graph.Customers(ctx, graphqlClient)
	if err != nil {
		log.Fatal(err)
	}
	println(*respquery)

}
