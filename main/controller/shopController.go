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

func ShopGetActiveTicketList(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	graphqlClient := graphql.NewClient(GRAPHQL_CLIENT_URL, http.DefaultClient)

	resp, err := graph.ActiveTicketByStats(ctx, graphqlClient, "Active")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	fmt.Println(resp.ActiveTicketByStats)

	data := make(map[int]map[string]interface{},)
	for i, t := range resp.ActiveTicketByStats {
		tData := map[string]interface{}{
			"ticketID": t.ID,
			"cusID":    t.CustomerID,
			"carID":    t.CarID,
			"problem":  t.Problem,
			"shopID":   IsNil(t.ShopID),
			"status":   IsNil(t.Status),
			// "description": IsNil(t.Description),
			// "location": map[string]float64{
			// 	"lng": t.Longitude,
			// 	"lat": t.Latitude,
			// },
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

func ShopGetOngoingTicketList(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	graphqlClient := graphql.NewClient(GRAPHQL_CLIENT_URL, http.DefaultClient)

	resp, err := graph.ActiveTicketByShop(ctx, graphqlClient, r.FormValue("shopID"))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	fmt.Println(resp.ActiveTicketByShop)

	data := make(map[int]map[string]interface{})
	for i, t := range resp.ActiveTicketByShop {
		tData := map[string]interface{}{
			"ticketID": t.ID,
			"cusID":    t.CustomerID,
			"carID":    t.CarID,
			"problem":  t.Problem,
			"shopID":   IsNil(t.ShopID),
			"status":   IsNil(t.Status),
			// "description": IsNil(t.Description),
			// "location": map[string]float64{
			// 	"lng": t.Longitude,
			// 	"lat": t.Latitude,
			// },
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

func ShopAcceptTicket(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	graphqlClient := graphql.NewClient(GRAPHQL_CLIENT_URL, http.DefaultClient)

	resp, err := graph.ActiveTicketUpdateMulti(ctx, graphqlClient, &graph.ActiveTicketUpdateInput{
		ID:         r.FormValue("ticketID"),
		CarID:      r.FormValue("carID"),
		CustomerID: r.FormValue("cusID"),
		Problem:    r.FormValue("problem"),
		ShopID:     util.Ptr(r.FormValue("shopID")),
		Status:     util.Ptr("Accepted"),
		// Latitude: r.FormValue("lat"),
		// Longitude: r.FormValue("lng"),
		// Description: r.FormValue("description"),
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	fmt.Println(resp.ActiveTicketUpdateMulti.ID)

	data := map[string]interface{}{
		"ticketID": resp.ActiveTicketUpdateMulti.ID,
		"cusID":    resp.ActiveTicketUpdateMulti.CustomerID,
		"carID":    resp.ActiveTicketUpdateMulti.CarID,
		"problem":  resp.ActiveTicketUpdateMulti.Problem,
		"shopID":   IsNil(resp.ActiveTicketUpdateMulti.ShopID),
		"status":   IsNil(resp.ActiveTicketUpdateMulti.Status),
		// "description": resp.ActiveTicketUpdateMulti.Description,
		// "location": map[string]float64{
		// 	"lat": resp.ActiveTicketUpdateMulti.latitude,
		// 	"lng": resp.ActiveTicketUpdateMulti.longitude,
		// },
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	AddHeader(w).Write(jsonData)
}

func ShopCompleteTicket(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	graphqlClient := graphql.NewClient(GRAPHQL_CLIENT_URL, http.DefaultClient)

	resp, err := graph.ActiveTicketUpdateMulti(ctx, graphqlClient, &graph.ActiveTicketUpdateInput{
		ID:         r.FormValue("ticketID"),
		CarID:      r.FormValue("carID"),
		CustomerID: r.FormValue("cusID"),
		Problem:    r.FormValue("problem"),
		ShopID:     util.Ptr(r.FormValue("shopID")),
		Status:     util.Ptr("Completed"),
		// Latitude: r.FormValue("lat"),
		// Longitude: r.FormValue("lng"),
		// Description: r.FormValue("description"),
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	fmt.Println(resp.ActiveTicketUpdateMulti.ID)

	data := map[string]string{
		"ticketID": resp.ActiveTicketUpdateMulti.ID,
		"cusID":    resp.ActiveTicketUpdateMulti.CustomerID,
		"carID":    resp.ActiveTicketUpdateMulti.CarID,
		"problem":  resp.ActiveTicketUpdateMulti.Problem,
		"shopID":   IsNil(resp.ActiveTicketUpdateMulti.ShopID),
		"status":   IsNil(resp.ActiveTicketUpdateMulti.Status),
		// "description": resp.ActiveTicketUpdateMulti.Description,
		// "location": map[string]float64{
		// 	"lat": resp.ActiveTicketUpdateMulti.latitude,
		// 	"lng": resp.ActiveTicketUpdateMulti.longitude,
		// },
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	AddHeader(w).Write(jsonData)
}

// need improvise
func ShopCancelTicket(w http.ResponseWriter, r *http.Request) {
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

	AddHeader(w).Write(jsonData)
}

func ShopGetHistory(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	graphqlClient := graphql.NewClient(GRAPHQL_CLIENT_URL, http.DefaultClient)

	resp, err := graph.TicketByShop(ctx, graphqlClient, &graph.TicketByShopInput{
		ShopID:   r.FormValue("shopID"),
		FromTime: time.Now().AddDate(0, -1, 0),
		ToTime:   time.Now(),
		Status:   util.Ptr(r.FormValue("status")),
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	fmt.Println(resp.TicketByShop)

	data := make(map[int]map[string]interface{})
	for i, t := range resp.TicketByShop {
		tData := map[string]interface{}{
			"ticketID":     t.ID,
			"cusID":        t.CustomerID,
			"carID":        t.CarID,
			"problem":      t.Problem,
			"createTime":   t.CreateTime.String(),
			"shopID":       *t.ShopID,
			"acceptedTime": t.AcceptedTime.String(),
			"status":       IsNil(t.Status),
			"location": map[string]float64{
				"lat": t.Latitude,
				"lng": t.Longitude,
			},
			"description": t.Description,
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

func ShopGetTodayCompletedTicket(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	graphqlClient := graphql.NewClient(GRAPHQL_CLIENT_URL, http.DefaultClient)

	resp, err := graph.TicketByShop(ctx, graphqlClient, &graph.TicketByShopInput{
		ShopID:   r.FormValue("shopID"),
		FromTime: time.Now().Truncate(24 * time.Hour),
		ToTime:   time.Now().Truncate(24 * time.Hour).Add(time.Hour*24 - time.Second*1),
		Status:   util.Ptr("Finish"),
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	fmt.Println(resp.TicketByShop)

	data := make(map[int]map[string]interface{})
	for i, t := range resp.TicketByShop {
		tData := map[string]interface{}{
			"ticketID":     t.ID,
			"cusID":        t.CustomerID,
			"carID":        t.CarID,
			"problem":      t.Problem,
			"createTime":   t.CreateTime.String(),
			"shopID":       *t.ShopID,
			"acceptedTime": t.AcceptedTime.String(),
			"status":       IsNil(t.Status),
			"location": map[string]float64{
				"lat": t.Latitude,
				"lng": t.Longitude,
			},
			"description": t.Description,
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

func ShopCreateProfile(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	graphqlClient := graphql.NewClient(GRAPHQL_CLIENT_URL, http.DefaultClient)

	resp, err := graph.ShopCreate(ctx, graphqlClient, &graph.ShopCreateInput{
		Name:    r.FormValue("shopName"),
		Tel:     r.FormValue("shopTel"),
		Email:   r.FormValue("shopEmail"),
		Address: r.FormValue("shopAddress"),
		Latitude: StrToFloat(r.FormValue("lat")),
		Longitude: StrToFloat(r.FormValue("lng")),
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	data := map[string]interface{}{
		"shopID":      resp.ShopCreate.ID,
		"shopName":    resp.ShopCreate.Name,
		"shopTel":     resp.ShopCreate.Tel,
		"shopAddress": resp.ShopCreate.Address,
		"location": map[string]float64{
			"lat": resp.ShopCreate.Latitude,
			"lng": resp.ShopCreate.Longitude,
		},
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	AddHeader(w).Write(jsonData)
}

func ShopUpdateProfile(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	graphqlClient := graphql.NewClient(GRAPHQL_CLIENT_URL, http.DefaultClient)

	resp, err := graph.ShopUpdateMulti(ctx, graphqlClient, &graph.ShopUpdateInput{
		ID:      r.FormValue("shopID"),
		Name:    r.FormValue("shopName"),
		Tel:     r.FormValue("shopTel"),
		Email:   r.FormValue("shopEmail"),
		Address: r.FormValue("shopAddress"),
		Latitude: StrToFloat(r.FormValue("lat")),
		Longitude: StrToFloat(r.FormValue("lng")),
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	data := map[string]interface{}{
		"shopID":      resp.ShopUpdateMulti.ID,
		"shopName":    resp.ShopUpdateMulti.Name,
		"shopTel":     resp.ShopUpdateMulti.Tel,
		"shopEmail":   resp.ShopUpdateMulti.Email,
		"shopAddress": resp.ShopUpdateMulti.Address,
		"location": map[string]float64{
			"lat": resp.ShopUpdateMulti.Latitude,
			"lng": resp.ShopUpdateMulti.Longitude,
		},
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	AddHeader(w).Write(jsonData)
}

func ShopGetCustomerProfile(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	graphqlClient := graphql.NewClient(GRAPHQL_CLIENT_URL, http.DefaultClient)

	resp, err := graph.CustomerByID(ctx, graphqlClient, r.FormValue("cusID"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	data := map[string]string{
		"cusID": resp.CustomerByID.ID,
		"fName": resp.CustomerByID.FName,
		"lName": resp.CustomerByID.LName,
		"tel":   resp.CustomerByID.Tel,
		"email": resp.CustomerByID.Email,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	AddHeader(w).Write(jsonData)
}
