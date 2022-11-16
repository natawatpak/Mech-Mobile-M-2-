package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Khan/genqlient/graphql"
	"github.com/natawatpak/Mech-Mobile-M-2-/backend/graph"
	"github.com/natawatpak/Mech-Mobile-M-2-/backend/util"
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

func UpdateCustomerProfile(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	graphqlClient := graphql.NewClient("http://localhost:8081/query", http.DefaultClient)

	resp, err := graph.CustomerUpdateMulti(ctx, graphqlClient, &graph.CustomerUpdateInput{
		ID:    r.FormValue("ID"),
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

func AddCar(w http.ResponseWriter, r *http.Request) []byte {
	r.ParseForm()

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	graphqlClient := graphql.NewClient("http://localhost:8081/query", http.DefaultClient)

	resp, err := graph.CarCreate(ctx, graphqlClient, &graph.CarCreateInput{
		OwnerID:  r.FormValue("cusID"),
		PlateNum: r.FormValue("plateNum"),
		IssuedAt: r.FormValue("issuedAt"),
		Color:    r.FormValue("color"),
		Type:     r.FormValue("type"),
		Brand:    r.FormValue("brand"),
		Build:    util.Ptr(r.FormValue("build")),
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.CarCreate.ID)

	data := map[string]string{
		"carID": resp.CarCreate.ID,
	}
	jsonData,err := json.Marshal(data)
	if (err!=nil){
		log.Fatal(err)
	}
	return jsonData
}

func GetCarList(w http.ResponseWriter, r *http.Request) []byte {
	r.ParseForm()

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	graphqlClient := graphql.NewClient("http://localhost:8081/query", http.DefaultClient)

	resp, err := graph.CarByOwner(ctx, graphqlClient, r.FormValue("cusID"))

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.CarByOwner)
	data := make(map[int]map[string]string)
	for i, car := range resp.CarByOwner {
		carData := map[string]string{
			"id":       car.ID,
			"plate":    car.PlateNum,
			"issuedAt": car.IssuedAt,
			"color":    car.Color,
			"type":     car.Type,
			"brand":    car.Brand,
			"build":    *car.Build,
		}
		data[i] = carData
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	return jsonData
}

func RemoveCar(w http.ResponseWriter, r *http.Request) []byte {
	r.ParseForm()

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	graphqlClient := graphql.NewClient("http://localhost:8081/query", http.DefaultClient)

	resp, err := graph.CarDelete(ctx, graphqlClient, &graph.DeleteIDInput{
		ID: r.FormValue("carID"),
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp.CarDelete.ID)
	data := map[string]string{
		"CarID": resp.CarDelete.ID,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	return jsonData
}

func AddTicket(w http.ResponseWriter, r *http.Request) []byte {
	r.ParseForm()

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	graphqlClient := graphql.NewClient("http://localhost:8081/query", http.DefaultClient)

	resp, err := graph.TicketCreate(ctx, graphqlClient, &graph.TicketCreateInput{
		CustomerID:   r.FormValue("cusID"),
		CarID:        r.FormValue("carID"),
		Problem:      r.FormValue("problem"),
		CreateTime:   time.Now().String(),
		ShopID:       r.FormValue("shopID"),
		AcceptedTime: time.Now().String(),
		Status:       r.FormValue("status"),
	})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.TicketCreate.ID)

	data := map[string]string{
		"ticketID": resp.TicketCreate.ID,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	return jsonData
}

// func GetCustomerActiveTicket(w http.ResponseWriter, r *http.Request) []byte {
// 	r.ParseForm()

// 	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
// 	graphqlClient := graphql.NewClient("http://localhost:8081/query", http.DefaultClient)

// 	resp, err := graph.ActiveTicketByCustomer(ctx, graphqlClient, r.FormValue("cusID"))

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	data := make(map[int]map[string]string)

// 	jsonData, err := json.Marshal(data)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return jsonData
// }
