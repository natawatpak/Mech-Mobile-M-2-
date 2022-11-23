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

const GRAPHQL_CLIENT_URL = "http://localhost:8081/"

func AddHeader(w http.ResponseWriter) http.ResponseWriter {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	return w
}

func IsNil(strpt *string) string{
	if (strpt == nil){return ""}
	return *strpt
}

func CustomerCreateProfile(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Fatal(err)
	}
	// fName := r.FormValue("fName")

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	graphqlClient := graphql.NewClient(GRAPHQL_CLIENT_URL, http.DefaultClient)

	resp, err := graph.CustomerCreate(ctx, graphqlClient, &graph.CustomerCreateInput{
		FName: r.FormValue("fName"),
		LName: r.FormValue("lName"),
		Tel:   r.FormValue("tel"),
		Email: r.FormValue("email"),
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	fmt.Println(resp.CustomerCreate.FName)

	data := map[string]string{
		"ID":    resp.CustomerCreate.ID,
		"fName": resp.CustomerCreate.FName,
		"lName": resp.CustomerCreate.LName,
		"tel":   resp.CustomerCreate.Tel,
		"email": resp.CustomerCreate.Email,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	AddHeader(w).Write(jsonData)
}

func CustomerUpdateProfile(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	graphqlClient := graphql.NewClient(GRAPHQL_CLIENT_URL, http.DefaultClient)

	resp, err := graph.CustomerUpdateMulti(ctx, graphqlClient, &graph.CustomerUpdateInput{
		ID:    r.FormValue("ID"),
		FName: r.FormValue("fName"),
		LName: r.FormValue("lName"),
		Tel:   r.FormValue("tel"),
		Email: r.FormValue("email"),
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	data := map[string]string{
		"ID":    resp.CustomerUpdateMulti.ID,
		"fName": resp.CustomerUpdateMulti.FName,
		"lName": resp.CustomerUpdateMulti.LName,
		"tel":   resp.CustomerUpdateMulti.Tel,
		"email": resp.CustomerUpdateMulti.Email,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	AddHeader(w).Write(jsonData)
}

func CustomerAddCar(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	graphqlClient := graphql.NewClient(GRAPHQL_CLIENT_URL, http.DefaultClient)

	resp, err := graph.CarCreate(ctx, graphqlClient, &graph.CarCreateInput{
		OwnerID:  r.FormValue("cusID"),
		PlateNum: r.FormValue("plateNum"),
		IssuedAt: util.Ptr(r.FormValue("issuedAt")),
		Color:    util.Ptr(r.FormValue("color")),
		Type:     util.Ptr(r.FormValue("type")),
		Brand:    r.FormValue("brand"),
		Build:    util.Ptr(r.FormValue("build")),
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	fmt.Println(resp.CarCreate.ID)

	data := map[string]string{
		"carID": resp.CarCreate.ID,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	AddHeader(w).Write(jsonData)
}

func CustomerGetCarList(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	graphqlClient := graphql.NewClient(GRAPHQL_CLIENT_URL, http.DefaultClient)

	resp, err := graph.CarByOwner(ctx, graphqlClient, r.FormValue("cusID"))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	data := make(map[int]map[string]string)
	for i, car := range resp.CarByOwner {
		carData := map[string]string{
			"id":       car.ID,
			"plate":    car.PlateNum,
			"issuedAt": IsNil(car.IssuedAt),
			"color":    IsNil(car.Color),
			"type":     IsNil(car.Type),
			"brand":    car.Brand,
			"build":    IsNil(car.Build),
		}
		data[i] = carData
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	AddHeader(w).Write(jsonData)
}

func CustomerRemoveCar(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	graphqlClient := graphql.NewClient(GRAPHQL_CLIENT_URL, http.DefaultClient)

	resp, err := graph.CarDelete(ctx, graphqlClient, &graph.DeleteIDInput{
		ID: r.FormValue("carID"),
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	fmt.Println(resp.CarDelete.ID)
	data := map[string]string{
		"CarID": resp.CarDelete.ID,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	AddHeader(w).Write(jsonData)
}

func CustomerAddTicket(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	graphqlClient := graphql.NewClient(GRAPHQL_CLIENT_URL, http.DefaultClient)

	resp, err := graph.TicketCreate(ctx, graphqlClient, &graph.TicketCreateInput{
		CustomerID:   r.FormValue("cusID"),
		CarID:        r.FormValue("carID"),
		Problem:      r.FormValue("problem"),
		CreateTime:   time.Now(),
		ShopID:       util.Ptr("no shopID"), // need to be optional
		AcceptedTime: nil,
		Status:       util.Ptr(r.FormValue("status")),
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	fmt.Println(resp.TicketCreate.ID)

	data := map[string]string{
		"ticketID": resp.TicketCreate.ID,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	AddHeader(w).Write(jsonData)
}

// func CustomerGetActiveTicket(w http.ResponseWriter, r *http.Request) []byte {
// 	r.ParseForm()

// 	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
// 	graphqlClient := graphql.NewClient(GRAPHQL_CLIENT_URL, http.DefaultClient)

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

// need improvise
func CustomerCancelTicket(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	graphqlClient := graphql.NewClient(GRAPHQL_CLIENT_URL, http.DefaultClient)

	resp, err := graph.ActiveTicketDelete(ctx, graphqlClient, &graph.DeleteIDInput{
		ID: r.FormValue("ticketID"),
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	fmt.Println(resp.ActiveTicketDelete.ID)

	data := map[string]string{
		"ticketID": resp.ActiveTicketDelete.ID,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	AddHeader(w).Write(jsonData)
}

func CustomerGetHistory(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	graphqlClient := graphql.NewClient(GRAPHQL_CLIENT_URL, http.DefaultClient)

	resp, err := graph.TicketByCustomer(ctx, graphqlClient, &graph.TicketByCustomerInput{
		CustomerID: r.FormValue("cusID"),
		FromTime:   time.Now().AddDate(0, -1, 0),
		ToTime:     time.Now(),
		Status:     util.Ptr(r.FormValue("status")),
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	fmt.Println(resp.TicketByCustomer)

	data := make(map[int]map[string]string)
	for i, t := range resp.TicketByCustomer {
		tData := map[string]string{
			"id":           t.ID,
			"cusID":        t.CustomerID,
			"carID":        t.CarID,
			"problem":      t.Problem,
			"createTime":   t.CreateTime.String(),
			"shopID":       *t.ShopID,
			"acceptedTime": t.AcceptedTime.String(),
			"status":       IsNil(t.Status),
		}
		data[i] = tData
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	AddHeader(w).Write(jsonData)
}

func CustomerGetShopProfile(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	graphqlClient := graphql.NewClient(GRAPHQL_CLIENT_URL, http.DefaultClient)

	resp, err := graph.ShopByID(ctx, graphqlClient, r.FormValue("shopID"))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	data := map[string]string{
		"shopID":      resp.ShopByID.ID,
		"shopName":    resp.ShopByID.Name,
		"shopTel":     resp.ShopByID.Tel,
		"shopEmail":   resp.ShopByID.Email,
		"shopAddress": resp.ShopByID.Address,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	AddHeader(w).Write(jsonData)

}
