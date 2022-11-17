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
	"github.com/natawatpak/Mech-Mobile-M-2-/backend/graph/model"
)

func ShopGetActiveTicketList(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	graphqlClient := graphql.NewClient("http://localhost:8081/", http.DefaultClient)

	resp, err := graph.ActiveTicketByStats(ctx, graphqlClient, graph.Status(model.StatusActive))

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.ActiveTicketByStats)

	data := make(map[int]map[string]string)
	for i, t := range resp.ActiveTicketByStats {
		tData := map[string]string{
			"ticketID": t.ID,
			"cusID":    t.CustomerID,
			"carID":    t.CarID,
			"problem":  t.Problem,
			"shopID":   t.ShopID,
			"status":   t.Status,
		}
		data[i] = tData
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

// func GetOngoingTicketList(w http.ResponseWriter, r *http.Request) []byte {

// }

func ShopGetActiveTicket(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	graphqlClient := graphql.NewClient("http://localhost:8081/", http.DefaultClient)

	resp, err := graph.ActiveTicketByShop(ctx, graphqlClient, r.FormValue("shopID"))

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.ActiveTicketByShop)

	data := make(map[int]map[string]string)
	for i, t := range resp.ActiveTicketByShop {
		tData := map[string]string{
			"ticketID": t.ID,
			"cusID":    t.CustomerID,
			"carID":    t.CarID,
			"problem":  t.Problem,
			"shopID":   t.ShopID,
			"status":   t.Status,
		}
		data[i] = tData
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func ShopAcceptTicket(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	graphqlClient := graphql.NewClient("http://localhost:8081/", http.DefaultClient)

	resp, err := graph.ActiveTicketUpdateMulti(ctx, graphqlClient, &graph.ActiveTicketUpdateInput{
		ID:         r.FormValue("ticketID"),
		CarID:      r.FormValue("carID"),
		CustomerID: r.FormValue("cusID"),
		Problem:    r.FormValue("problem"),
		ShopID:     r.FormValue("shopID"),
		Status:     "Accepted",
	})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.ActiveTicketUpdateMulti.ID)

	data := map[string]string{
		"ticketID": resp.ActiveTicketUpdateMulti.ID,
		"cusID":    resp.ActiveTicketUpdateMulti.CustomerID,
		"carID":    resp.ActiveTicketUpdateMulti.CarID,
		"problem":  resp.ActiveTicketUpdateMulti.Problem,
		"shopID":   resp.ActiveTicketUpdateMulti.ShopID,
		"status":   resp.ActiveTicketUpdateMulti.Status,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func ShopCompleteTicket(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	graphqlClient := graphql.NewClient("http://localhost:8081/", http.DefaultClient)

	resp, err := graph.ActiveTicketUpdateMulti(ctx, graphqlClient, &graph.ActiveTicketUpdateInput{
		ID:         r.FormValue("ticketID"),
		CarID:      r.FormValue("carID"),
		CustomerID: r.FormValue("cusID"),
		Problem:    r.FormValue("problem"),
		ShopID:     r.FormValue("shopID"),
		Status:     "Completed",
	})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.ActiveTicketUpdateMulti.ID)

	data := map[string]string{
		"ticketID": resp.ActiveTicketUpdateMulti.ID,
		"cusID":    resp.ActiveTicketUpdateMulti.CustomerID,
		"carID":    resp.ActiveTicketUpdateMulti.CarID,
		"problem":  resp.ActiveTicketUpdateMulti.Problem,
		"shopID":   resp.ActiveTicketUpdateMulti.ShopID,
		"status":   resp.ActiveTicketUpdateMulti.Status,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

// need improvise
func ShopCancelTicket(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	graphqlClient := graphql.NewClient("http://localhost:8081/", http.DefaultClient)

	resp, err := graph.ActiveTicketDelete(ctx, graphqlClient, &graph.DeleteIDInput{
		ID: r.FormValue("ticketID"),
	})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.ActiveTicketDelete.ID)

	data := map[string]string{
		"ticketID": resp.ActiveTicketDelete.ID,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func ShopGetHistory(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	graphqlClient := graphql.NewClient("http://localhost:8081/", http.DefaultClient)

	resp, err := graph.TicketByShop(ctx, graphqlClient, &graph.TicketByShopInput{
		ShopID:   r.FormValue("shopID"),
		FromTime: time.Now().AddDate(0, -1, 0).String(),
		ToTime:   time.Now().String(),
		Status:   r.FormValue("status"),
	})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.TicketByShop)

	data := make(map[int]map[string]string)
	for i, t := range resp.TicketByShop {
		tData := map[string]string{
			"ticketID":     t.ID,
			"cusID":        t.CustomerID,
			"carID":        t.CarID,
			"problem":      t.Problem,
			"createTime":   t.CreateTime,
			"shopID":       t.ShopID,
			"acceptedTime": t.AcceptedTime,
			"status":       t.Status,
		}
		data[i] = tData
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func ShopGetTodayCompletedTicket(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	graphqlClient := graphql.NewClient("http://localhost:8081/", http.DefaultClient)

	resp, err := graph.TicketByShop(ctx, graphqlClient, &graph.TicketByShopInput{
		ShopID:   r.FormValue("shopID"),
		FromTime: time.Now().Truncate(24 * time.Hour).String(),
		ToTime:   time.Now().Truncate(24 * time.Hour).Add(time.Hour*24 - time.Second*1).String(),
		Status:   "Finish",
	})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.TicketByShop)

	data := make(map[int]map[string]string)
	for i, t := range resp.TicketByShop {
		tData := map[string]string{
			"ticketID":     t.ID,
			"cusID":        t.CustomerID,
			"carID":        t.CarID,
			"problem":      t.Problem,
			"createTime":   t.CreateTime,
			"shopID":       t.ShopID,
			"acceptedTime": t.AcceptedTime,
			"status":       t.Status,
		}
		data[i] = tData
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func ShopCreateProfile(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	graphqlClient := graphql.NewClient("http://localhost:8081/", http.DefaultClient)

	resp, err := graph.ShopCreate(ctx, graphqlClient, &graph.ShopCreateInput{
		Name:    r.FormValue("shopName"),
		Tel:     r.FormValue("shopTel"),
		Email:   r.FormValue("shopEmail"),
		Address: r.FormValue("shopAddress"),
	})
	if err != nil {
		log.Fatal(err)
	}

	data := map[string]string{
		"shopID":      resp.ShopCreate.ID,
		"shopName":    resp.ShopCreate.Name,
		"shopTel":     resp.ShopCreate.Tel,
		"shopAddress": resp.ShopCreate.Address,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func ShopUpdateProfile(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	graphqlClient := graphql.NewClient("http://localhost:8081/", http.DefaultClient)

	resp, err := graph.ShopUpdateMulti(ctx, graphqlClient, &graph.ShopUpdateInput{
		ID:      r.FormValue("shopID"),
		Name:    r.FormValue("shopName"),
		Tel:     r.FormValue("shopTel"),
		Email:   r.FormValue("shopEmail"),
		Address: r.FormValue("shopAddress"),
	})
	if err != nil {
		log.Fatal(err)
	}

	data := map[string]string{
		"shopID":      resp.ShopUpdateMulti.ID,
		"shopName":    resp.ShopUpdateMulti.Name,
		"shopTel":     resp.ShopUpdateMulti.Tel,
		"shopEmail":   resp.ShopUpdateMulti.Email,
		"shopAddress": resp.ShopUpdateMulti.Address,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)

}
