package test

import (
	"context"
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/Khan/genqlient/graphql"
	"github.com/natawatpak/Mech-Mobile-M-2-/backend/graph"
	"github.com/natawatpak/Mech-Mobile-M-2-/backend/util"
	"github.com/stretchr/testify/assert"
)

type ticketCreateInput struct {
	Problem      string
	CreateTime   time.Time
	AcceptedTime time.Time
	Status       string
}

var ticketCreateInputTest = []ticketCreateInput{
	{
		Problem:      "car run out of battery",
		CreateTime:   time.Now().Add(time.Second * -2),
		AcceptedTime: time.Now(),
		Status:       "Active",
	},
	{
		Problem:      "washer run out",
		CreateTime:   time.Now().Add(time.Second * -3),
		AcceptedTime: time.Now().Add(time.Second * 1),
		Status:       "Active",
	},
	{
		Problem:      "tie flat",
		CreateTime:   time.Now().Add(time.Hour * -3),
		AcceptedTime: time.Now().Add(time.Hour * -2),
		Status:       "InActive",
	},
}

type carCreateInput struct {
	OwnerID   string
	PlateNum  string
	PlateType string
	IssuedAt  string
	Color     string
	Type      string
	Brand     string
	Build     string
}

var carCreateInputTest = []carCreateInput{
	{
		PlateNum:  "9999",
		PlateType: "white",
		IssuedAt:  "Bangkok",
		Color:     "Blue",
		Type:      "SUV",
		Brand:     "Toyota",
		Build:     "Thailand",
	},
	{
		PlateNum:  "8888",
		PlateType: "Black",
		IssuedAt:  "London",
		Color:     "Red",
		Type:      "Sport",
		Brand:     "Honda",
		Build:     "USA",
	},
}

type customerCreateInput struct {
	FName string
	LName string
	Tel   string
	Email string
}

var customerCreateInputTest = []customerCreateInput{
	{
		FName: "Phone",
		LName: "Phum",
		Tel:   "098-278-9331",
		Email: "Chic@gmail.com",
	},
	{
		FName: "chicken",
		LName: "KFC",
		Tel:   "123456789",
		Email: "Chicken@gmail.com",
	},
}

func TestTicketCreate(t *testing.T) {
	ctx := context.Background()

	graphqlClient := graphql.NewClient("http://localhost:8081/", http.DefaultClient)

	respCus, err := graph.CustomerCreate(ctx, graphqlClient, &graph.CustomerCreateInput{
		FName: customerCreateInputTest[0].FName,
		LName: customerCreateInputTest[0].LName,
		Tel:   customerCreateInputTest[0].Tel,
		Email: customerCreateInputTest[0].Email,
	})
	if err != nil {
		log.Fatal(err)
	}
	respCar, err := graph.CarCreate(ctx, graphqlClient, &graph.CarCreateInput{
		OwnerID:   respCus.CustomerCreate.ID,
		PlateNum:  carCreateInputTest[0].PlateNum,
		PlateType: util.Ptr(carCreateInputTest[0].PlateType),
		IssuedAt:  util.Ptr(carCreateInputTest[0].IssuedAt),
		Color:     util.Ptr(carCreateInputTest[0].Color),
		Type:      util.Ptr(carCreateInputTest[0].Type),
		Brand:     carCreateInputTest[0].Brand,
		Build:     util.Ptr(carCreateInputTest[0].Build),
	})
	if err != nil {
		log.Fatal(err)
	}
	respShop, err := graph.ShopCreate(ctx, graphqlClient, &graph.ShopCreateInput{
		Name:    shopCreateInputTest[0].Name,
		Tel:     shopCreateInputTest[0].Tel,
		Email:   shopCreateInputTest[0].Email,
		Address: shopCreateInputTest[0].Address,
	})
	if err != nil {
		log.Fatal(err)
	}
	respTicket, err := graph.TicketCreate(ctx, graphqlClient, &graph.TicketCreateInput{
		CustomerID:   respCus.CustomerCreate.ID,
		CarID:        respCar.CarCreate.ID,
		Problem:      ticketCreateInputTest[0].Problem,
		CreateTime:   ticketCreateInputTest[0].CreateTime,
		ShopID:       respShop.ShopCreate.ID,
		AcceptedTime: &ticketCreateInputTest[0].AcceptedTime,
		Status:       &ticketCreateInputTest[0].Status,
	})
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, respCus.CustomerCreate.ID, respTicket.TicketCreate.CustomerID)
	assert.Equal(t, respCar.CarCreate.ID, respTicket.TicketCreate.CarID)
	assert.Equal(t, ticketCreateInputTest[0].Problem, respTicket.TicketCreate.Problem)
	assert.Equal(t, ticketCreateInputTest[0].CreateTime, respTicket.TicketCreate.CreateTime)
	assert.Equal(t, respShop.ShopCreate.ID, respTicket.TicketCreate.ShopID)
	assert.Equal(t, ticketCreateInputTest[0].AcceptedTime, *respTicket.TicketCreate.AcceptedTime)
	assert.Equal(t, ticketCreateInputTest[0].Status, *respTicket.TicketCreate.Status)

	graph.ShopDeleteAll(ctx, graphqlClient)
}
