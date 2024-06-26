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
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/aws/aws-sdk-go/service/apigatewaymanagementapi"
)

func getCar(id string) (*graph.CarByIDResponse, error) {
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	graphqlClient := graphql.NewClient(GRAPHQL_CLIENT_URL, http.DefaultClient)

	resp, err := graph.CarByID(ctx, graphqlClient, id)
	return resp, err
}

func getCus(id string) (*graph.CustomerByIDResponse, error) {
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	graphqlClient := graphql.NewClient(GRAPHQL_CLIENT_URL, http.DefaultClient)

	resp, err := graph.CustomerByID(ctx, graphqlClient, id)
	return resp, err
}

func getTicket(id string) (*graph.TicketByIDResponse, error) {
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	graphqlClient := graphql.NewClient(GRAPHQL_CLIENT_URL, http.DefaultClient)

	resp, err := graph.TicketByID(ctx, graphqlClient, id)
	return resp, err
}

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

	data := make(map[int]map[string]interface{})
	for i, t := range resp.ActiveTicketByStats {
		car, err := getCar(t.CarID)
		cus, err := getCus(t.CustomerID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		tData := map[string]interface{}{
			"ticketID": t.ID,
			"cus": map[string]interface{}{
				"cusID": cus.CustomerByID.ID,
				"fName": cus.CustomerByID.FName,
				"lName": cus.CustomerByID.LName,
			},
			"car": map[string]interface{}{
				"carID": car.CarByID.ID,
				"plate": car.CarByID.PlateNum,
				"type":  car.CarByID.Type,
				"brand": car.CarByID.Brand,
			},
			"problem":     t.Problem,
			"shopID":      IsNil(t.ShopID),
			"status":      IsNil(t.Status),
			"description": IsNil(t.Description),
			"location": map[string]float64{
				"lng": t.Longitude,
				"lat": t.Latitude,
			},
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
			"ticketID":    t.ID,
			"cusID":       t.CustomerID,
			"carID":       t.CarID,
			"problem":     t.Problem,
			"shopID":      IsNil(t.ShopID),
			"status":      IsNil(t.Status),
			"description": IsNil(t.Description),
			"location": map[string]float64{
				"lng": t.Longitude,
				"lat": t.Latitude,
			},
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

func ShopGetTicket(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	graphqlClient := graphql.NewClient(GRAPHQL_CLIENT_URL, http.DefaultClient)

	resp, err := graph.ActiveTicketByID(ctx, graphqlClient, r.FormValue("ticketID"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	cus, err := getCus(resp.ActiveTicketByID.CustomerID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	car, err := getCar(resp.ActiveTicketByID.CarID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	data := map[string]interface{}{
		"ticketID": resp.ActiveTicketByID.ID,
		"cus": map[string]interface{}{
			"cusID": cus.CustomerByID.ID,
			"fName": cus.CustomerByID.FName,
			"lName": cus.CustomerByID.LName,
		},
		"car": map[string]interface{}{
			"carID": car.CarByID.ID,
			"plate": car.CarByID.PlateNum,
			"type":  car.CarByID.Type,
			"brand": car.CarByID.Brand,
		},
		"problem":     resp.ActiveTicketByID.Problem,
		"shopID":      IsNil(resp.ActiveTicketByID.ShopID),
		"status":      IsNil(resp.ActiveTicketByID.Status),
		"description": IsNil(resp.ActiveTicketByID.Description),
		"location": map[string]float64{
			"lng": resp.ActiveTicketByID.Longitude,
			"lat": resp.ActiveTicketByID.Latitude,
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

func ShopAcceptTicket(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	graphqlClient := graphql.NewClient(GRAPHQL_CLIENT_URL, http.DefaultClient)
	t, err := getTicket(r.FormValue("ticketID"))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		println(err.Error())
		return
	}

	_, err = graph.TicketUpdateMulti(ctx, graphqlClient, &graph.TicketUpdateInput{
		ID:           r.FormValue("ticketID"),
		CarID:        r.FormValue("carID"),
		CustomerID:   r.FormValue("cusID"),
		Problem:      r.FormValue("problem"),
		ShopID:       util.Ptr(r.FormValue("shopID")),
		Status:       util.Ptr("Accepted"),
		Latitude:     StrToFloat(r.FormValue("lat")),
		Longitude:    StrToFloat(r.FormValue("lng")),
		Description:  util.Ptr(r.FormValue("description")),
		AcceptedTime: toTimePtr(time.Now()),
		CreateTime:   t.TicketByID.CreateTime,
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		println(err.Error())
		return
	}

	ctx, _ = context.WithTimeout(context.Background(), 30*time.Second)
	graphqlClient = graphql.NewClient(GRAPHQL_CLIENT_URL, http.DefaultClient)

	resp, err := graph.ActiveTicketUpdateMulti(ctx, graphqlClient, &graph.ActiveTicketUpdateInput{
		ID:          r.FormValue("ticketID"),
		CarID:       r.FormValue("carID"),
		CustomerID:  r.FormValue("cusID"),
		Problem:     r.FormValue("problem"),
		ShopID:      util.Ptr(r.FormValue("shopID")),
		Status:      util.Ptr("Accepted"),
		Latitude:    StrToFloat(r.FormValue("lat")),
		Longitude:   StrToFloat(r.FormValue("lng")),
		Description: util.Ptr(r.FormValue("description")),
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		println(err.Error())
		return
	}
	fmt.Println(resp.ActiveTicketUpdateMulti.ID)

	conn, _ := graph.TicketConnectsByID(ctx, graphqlClient,r.FormValue("ticketID"))
	mySession := session.Must(session.NewSession())

	apig := apigatewaymanagementapi.New(mySession)
	apig.Endpoint = "https://6eblsxltxc.execute-api.us-east-1.amazonaws.com/production"
	_, err = apig.PostToConnection(&apigatewaymanagementapi.PostToConnectionInput{
		ConnectionId: aws.String(conn.TicketConnectByID.CustomerConnectionID),
		Data:         []byte(`"data" : "Accepted"`),
	})
	if err != nil {
		fmt.Println(err)
	}

	data := map[string]interface{}{
		"ticketID":    resp.ActiveTicketUpdateMulti.ID,
		"cusID":       resp.ActiveTicketUpdateMulti.CustomerID,
		"carID":       resp.ActiveTicketUpdateMulti.CarID,
		"problem":     resp.ActiveTicketUpdateMulti.Problem,
		"shopID":      IsNil(resp.ActiveTicketUpdateMulti.ShopID),
		"status":      IsNil(resp.ActiveTicketUpdateMulti.Status),
		"description": IsNil(resp.ActiveTicketUpdateMulti.Description),
		"location": map[string]float64{
			"lat": resp.ActiveTicketUpdateMulti.Latitude,
			"lng": resp.ActiveTicketUpdateMulti.Longitude,
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

func ShopUpdateTicket(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	graphqlClient := graphql.NewClient(GRAPHQL_CLIENT_URL, http.DefaultClient)
	t, err := getTicket(r.FormValue("ticketID"))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	_, err = graph.TicketUpdateMulti(ctx, graphqlClient, &graph.TicketUpdateInput{
		ID:           r.FormValue("ticketID"),
		CarID:        r.FormValue("carID"),
		CustomerID:   r.FormValue("cusID"),
		Problem:      r.FormValue("problem"),
		ShopID:       util.Ptr(r.FormValue("shopID")),
		Status:       util.Ptr(r.FormValue("status")),
		Latitude:     StrToFloat(r.FormValue("lat")),
		Longitude:    StrToFloat(r.FormValue("lng")),
		Description:  util.Ptr(r.FormValue("description")),
		AcceptedTime: t.TicketByID.AcceptedTime,
		CreateTime:   t.TicketByID.CreateTime,
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	ctx, _ = context.WithTimeout(context.Background(), 30*time.Second)
	graphqlClient = graphql.NewClient(GRAPHQL_CLIENT_URL, http.DefaultClient)

	resp, err := graph.ActiveTicketUpdateMulti(ctx, graphqlClient, &graph.ActiveTicketUpdateInput{
		ID:          r.FormValue("ticketID"),
		CarID:       r.FormValue("carID"),
		CustomerID:  r.FormValue("cusID"),
		Problem:     r.FormValue("problem"),
		ShopID:      util.Ptr(r.FormValue("shopID")),
		Status:      util.Ptr(r.FormValue("status")),
		Latitude:    StrToFloat(r.FormValue("lat")),
		Longitude:   StrToFloat(r.FormValue("lng")),
		Description: util.Ptr(r.FormValue("description")),
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	fmt.Println(resp.ActiveTicketUpdateMulti.ID)

	conn, _ := graph.TicketConnectsByID(ctx, graphqlClient,r.FormValue("ticketID"))
	mySession := session.Must(session.NewSession())

	apig := apigatewaymanagementapi.New(mySession)
	apig.Endpoint = "https://6eblsxltxc.execute-api.us-east-1.amazonaws.com/production"
	_, err = apig.PostToConnection(&apigatewaymanagementapi.PostToConnectionInput{
		ConnectionId: aws.String(conn.TicketConnectByID.CustomerConnectionID),
		Data:         []byte(`"data" : `+ r.FormValue("status")),
	})
	if err != nil {
		fmt.Println(err)
	}

	data := map[string]interface{}{
		"ticketID":    resp.ActiveTicketUpdateMulti.ID,
		"cusID":       resp.ActiveTicketUpdateMulti.CustomerID,
		"carID":       resp.ActiveTicketUpdateMulti.CarID,
		"problem":     resp.ActiveTicketUpdateMulti.Problem,
		"shopID":      IsNil(resp.ActiveTicketUpdateMulti.ShopID),
		"status":      IsNil(resp.ActiveTicketUpdateMulti.Status),
		"description": IsNil(resp.ActiveTicketUpdateMulti.Description),
		"location": map[string]float64{
			"lat": resp.ActiveTicketUpdateMulti.Latitude,
			"lng": resp.ActiveTicketUpdateMulti.Longitude,
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
		car, err := getCar(t.CarID)
		cus, err := getCus(t.CustomerID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		tData := map[string]interface{}{
			"ticketID": t.ID,
			"cus": map[string]interface{}{
				"cusID": cus.CustomerByID.ID,
				"fName": cus.CustomerByID.FName,
				"lName": cus.CustomerByID.LName,
			},
			"car": map[string]interface{}{
				"carID": car.CarByID.ID,
				"plate": car.CarByID.PlateNum,
				"type":  car.CarByID.Type,
				"brand": car.CarByID.Brand,
			},
			"problem":      t.Problem,
			"createTime":   t.CreateTime.String(),
			"shopID":       *t.ShopID,
			"acceptedTime": t.AcceptedTime.String(),
			"status":       IsNil(t.Status),
			"location": map[string]float64{
				"lat": t.Latitude,
				"lng": t.Longitude,
			},
			"description": IsNil(t.Description),
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
		Name:      r.FormValue("shopName"),
		Tel:       r.FormValue("shopTel"),
		Email:     r.FormValue("shopEmail"),
		Address:   r.FormValue("shopAddress"),
		Latitude:  StrToFloat(r.FormValue("lat")),
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
		ID:        r.FormValue("shopID"),
		Name:      r.FormValue("shopName"),
		Tel:       r.FormValue("shopTel"),
		Email:     r.FormValue("shopEmail"),
		Address:   r.FormValue("shopAddress"),
		Latitude:  StrToFloat(r.FormValue("lat")),
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

func ShopGetProfile(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	graphqlClient := graphql.NewClient(GRAPHQL_CLIENT_URL, http.DefaultClient)

	resp, err := graph.ShopByEmail(ctx, graphqlClient, r.FormValue("shopEmail"))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	data := map[string]interface{}{
		"shopID":      resp.ShopByEmail.ID,
		"shopName":    resp.ShopByEmail.Name,
		"shopTel":     resp.ShopByEmail.Tel,
		"shopEmail":   resp.ShopByEmail.Email,
		"shopAddress": resp.ShopByEmail.Address,
		"location": map[string]float64{
			"lat": resp.ShopByEmail.Latitude,
			"lng": resp.ShopByEmail.Longitude,
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
