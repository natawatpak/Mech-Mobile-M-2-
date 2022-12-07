package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
	"github.com/Khan/genqlient/graphql"
	"github.com/natawatpak/Mech-Mobile-M-2-/backend/graph"
	"github.com/natawatpak/Mech-Mobile-M-2-/backend/util"
)

const GRAPHQL_CLIENT_URL = "https://a7okax4857.execute-api.us-east-1.amazonaws.com/default/carservice"


func AddHeader(w http.ResponseWriter) http.ResponseWriter {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	return w
}

func IsNil(strpt *string) string {
	if strpt == nil {
		return ""
	}
	return *strpt
}

func IsNilTime(strpt *time.Time) string {
	if strpt == nil {
		return ""
	}
	return strpt.String()
}

func StrToFloat(str string) float64{
	f,err := strconv.ParseFloat(str,64)
	if (err != nil){return -1}
	return f
}

func FloatToString(f float64) string{
	return fmt.Sprintf("%f", f)
}

func CustomerCreateProfile(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Fatal(err)
	}

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

func CustomerGetCar(w http.ResponseWriter, r *http.Request){
	r.ParseForm()

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	graphqlClient := graphql.NewClient(GRAPHQL_CLIENT_URL, http.DefaultClient)

	resp, err := graph.CarByID(ctx, graphqlClient, r.FormValue("carID"))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	fmt.Println(resp.CarByID.ID)
	data := map[string]string{
		"id": resp.CarByID.ID,
		"type": IsNil(resp.CarByID.Type),
		"brand": resp.CarByID.Brand,
		"plate": resp.CarByID.PlateNum,
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
		ShopID:       nil,
		AcceptedTime: nil,
		Status:       util.Ptr(r.FormValue("status")),
		Longitude:    StrToFloat(r.FormValue("lng")),
		Latitude: StrToFloat(r.FormValue("lat")),
		Description: util.Ptr(r.FormValue("description")),
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	fmt.Print("description")
	fmt.Println(r.FormValue("description"))
	fmt.Println(resp.TicketCreate.ID)

	ctx, _ = context.WithTimeout(context.Background(), 30*time.Second)
	graphqlClient = graphql.NewClient(GRAPHQL_CLIENT_URL, http.DefaultClient)

	_, err = graph.ActiveTicketCreate(ctx, graphqlClient, &graph.ActiveTicketCreateInput{
		ID: resp.TicketCreate.ID,
		CustomerID:   r.FormValue("cusID"),
		CarID:        r.FormValue("carID"),
		Problem:      r.FormValue("problem"),
		ShopID:       nil,
		Status:       util.Ptr(r.FormValue("status")),
		Longitude:    StrToFloat(r.FormValue("lng")),
		Latitude: StrToFloat(r.FormValue("lat")),
		Description: util.Ptr(r.FormValue("description")),
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	data := map[string]string{
		"ticketID": resp.TicketCreate.ID,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	sendUpdateSignal(WsShops)
	AddHeader(w).Write(jsonData)
}

func CustomerGetTicket(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	graphqlClient := graphql.NewClient(GRAPHQL_CLIENT_URL, http.DefaultClient)

	fmt.Println(r.FormValue("ticketID"))
	resp, err := graph.ActiveTicketByID(ctx, graphqlClient, r.FormValue("ticketID"))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	data := map[string]interface{}{
		"cusID": resp.ActiveTicketByID.CustomerID,
		"carID": resp.ActiveTicketByID.CarID,
		"shopID": IsNil(resp.ActiveTicketByID.ShopID),
		"problem": resp.ActiveTicketByID.Problem,
		"status": IsNil(resp.ActiveTicketByID.Status),
		"location": map[string]float64{
			"lat": resp.ActiveTicketByID.Latitude,
			"lng": resp.ActiveTicketByID.Longitude,
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

	data := make(map[int]map[string]interface{})
	for i, t := range resp.TicketByCustomer {
		tData := map[string]interface{}{
			"id":           t.ID,
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

	location := map[string]float64{
		"lat": resp.ShopByID.Latitude,
		"lng": resp.ShopByID.Longitude,
	}

	data := map[string]interface{}{
		"shopID":      resp.ShopByID.ID,
		"shopName":    resp.ShopByID.Name,
		"shopTel":     resp.ShopByID.Tel,
		"shopEmail":   resp.ShopByID.Email,
		"shopAddress": resp.ShopByID.Address,
		"location": location,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	AddHeader(w).Write(jsonData)

}