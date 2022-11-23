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
		CreateTime:   time.Date(2022, time.Month(6), 13, 10, 0, 3, 0, time.UTC),
		AcceptedTime: time.Date(2022, time.Month(6), 13, 15, 0, 3, 0, time.UTC),
		Status:       "Active",
	},
	{
		Problem:      "washer run out",
		CreateTime:   time.Date(2022, time.Month(7), 13, 13, 0, 2, 0, time.UTC),
		AcceptedTime: time.Date(2022, time.Month(7), 13, 13, 0, 3, 0, time.UTC),
		Status:       "Active",
	},
	{
		Problem:      "tie flat",
		CreateTime:   time.Date(2022, time.Month(2), 13, 13, 0, 2, 0, time.UTC),
		AcceptedTime: time.Date(2022, time.Month(2), 13, 13, 0, 3, 0, time.UTC),
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
		ShopID:       &respShop.ShopCreate.ID,
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
	assert.Equal(t, respShop.ShopCreate.ID, *respTicket.TicketCreate.ShopID)
	assert.Equal(t, ticketCreateInputTest[0].AcceptedTime, *respTicket.TicketCreate.AcceptedTime)
	assert.Equal(t, ticketCreateInputTest[0].Status, *respTicket.TicketCreate.Status)

	graph.CarDeleteAll(ctx, graphqlClient)
	graph.CustomerDeleteAll(ctx, graphqlClient)
	graph.ShopDeleteAll(ctx, graphqlClient)
	graph.TicketDeleteAll(ctx, graphqlClient)
}

func TestTicketUpdate(t *testing.T) {
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
	respCus2, err := graph.CustomerCreate(ctx, graphqlClient, &graph.CustomerCreateInput{
		FName: customerCreateInputTest[1].FName,
		LName: customerCreateInputTest[1].LName,
		Tel:   customerCreateInputTest[1].Tel,
		Email: customerCreateInputTest[1].Email,
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
	respTicketCreate, err := graph.TicketCreate(ctx, graphqlClient, &graph.TicketCreateInput{
		CustomerID:   respCus.CustomerCreate.ID,
		CarID:        respCar.CarCreate.ID,
		Problem:      ticketCreateInputTest[0].Problem,
		CreateTime:   ticketCreateInputTest[0].CreateTime,
		ShopID:       &respShop.ShopCreate.ID,
		AcceptedTime: &ticketCreateInputTest[0].AcceptedTime,
		Status:       &ticketCreateInputTest[0].Status,
	})
	if err != nil {
		log.Fatal(err)
	}
	respTicketUpdate, err := graph.TicketUpdateMulti(ctx, graphqlClient, &graph.TicketUpdateInput{
		ID:           respTicketCreate.TicketCreate.ID,
		CustomerID:   respCus2.CustomerCreate.ID,
		CarID:        respCar.CarCreate.ID,
		Problem:      ticketCreateInputTest[1].Problem,
		CreateTime:   ticketCreateInputTest[1].CreateTime,
		ShopID:       &respShop.ShopCreate.ID,
		AcceptedTime: &ticketCreateInputTest[1].AcceptedTime,
		Status:       &ticketCreateInputTest[1].Status,
	})
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, respCus2.CustomerCreate.ID, respTicketUpdate.TicketUpdateMulti.CustomerID)
	assert.Equal(t, respCar.CarCreate.ID, respTicketUpdate.TicketUpdateMulti.CarID)
	assert.Equal(t, ticketCreateInputTest[1].Problem, respTicketUpdate.TicketUpdateMulti.Problem)
	assert.Equal(t, ticketCreateInputTest[1].CreateTime, respTicketUpdate.TicketUpdateMulti.CreateTime)
	assert.Equal(t, respShop.ShopCreate.ID, *respTicketUpdate.TicketUpdateMulti.ShopID)
	assert.Equal(t, ticketCreateInputTest[1].AcceptedTime, *respTicketUpdate.TicketUpdateMulti.AcceptedTime)
	assert.Equal(t, ticketCreateInputTest[1].Status, *respTicketUpdate.TicketUpdateMulti.Status)

	graph.CarDeleteAll(ctx, graphqlClient)
	graph.CustomerDeleteAll(ctx, graphqlClient)
	graph.ShopDeleteAll(ctx, graphqlClient)
	graph.TicketDeleteAll(ctx, graphqlClient)
}

func TestTicketDelete(t *testing.T) {
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
	respTicketCreate, err := graph.TicketCreate(ctx, graphqlClient, &graph.TicketCreateInput{
		CustomerID:   respCus.CustomerCreate.ID,
		CarID:        respCar.CarCreate.ID,
		Problem:      ticketCreateInputTest[0].Problem,
		CreateTime:   ticketCreateInputTest[0].CreateTime,
		ShopID:       &respShop.ShopCreate.ID,
		AcceptedTime: &ticketCreateInputTest[0].AcceptedTime,
		Status:       &ticketCreateInputTest[0].Status,
	})
	if err != nil {
		log.Fatal(err)
	}

	respQuery, err := graph.TicketDelete(ctx, graphqlClient, &graph.DeleteIDInput{
		ID: respTicketCreate.TicketCreate.ID,
	})
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, respCus.CustomerCreate.ID, respQuery.TicketDelete.CustomerID)
	assert.Equal(t, respCar.CarCreate.ID, respQuery.TicketDelete.CarID)
	assert.Equal(t, ticketCreateInputTest[0].Problem, respQuery.TicketDelete.Problem)
	assert.Equal(t, ticketCreateInputTest[0].CreateTime, respQuery.TicketDelete.CreateTime)
	assert.Equal(t, respShop.ShopCreate.ID, *respQuery.TicketDelete.ShopID)
	assert.Equal(t, ticketCreateInputTest[0].AcceptedTime, *respQuery.TicketDelete.AcceptedTime)
	assert.Equal(t, ticketCreateInputTest[0].Status, *respQuery.TicketDelete.Status)

	graph.CarDeleteAll(ctx, graphqlClient)
	graph.CustomerDeleteAll(ctx, graphqlClient)
	graph.ShopDeleteAll(ctx, graphqlClient)
	graph.TicketDeleteAll(ctx, graphqlClient)
}

func TestTicketByID(t *testing.T) {
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
	respTicketCreate, err := graph.TicketCreate(ctx, graphqlClient, &graph.TicketCreateInput{
		CustomerID:   respCus.CustomerCreate.ID,
		CarID:        respCar.CarCreate.ID,
		Problem:      ticketCreateInputTest[0].Problem,
		CreateTime:   ticketCreateInputTest[0].CreateTime,
		ShopID:       &respShop.ShopCreate.ID,
		AcceptedTime: &ticketCreateInputTest[0].AcceptedTime,
		Status:       &ticketCreateInputTest[0].Status,
	})
	if err != nil {
		log.Fatal(err)
	}

	respQuery, err := graph.TicketByID(ctx, graphqlClient, respTicketCreate.TicketCreate.ID)
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, respCus.CustomerCreate.ID, respQuery.TicketByID.CustomerID)
	assert.Equal(t, respCar.CarCreate.ID, respQuery.TicketByID.CarID)
	assert.Equal(t, ticketCreateInputTest[0].Problem, respQuery.TicketByID.Problem)
	assert.Equal(t, ticketCreateInputTest[0].CreateTime, respQuery.TicketByID.CreateTime)
	assert.Equal(t, respShop.ShopCreate.ID, *respQuery.TicketByID.ShopID)
	assert.Equal(t, ticketCreateInputTest[0].AcceptedTime, *respQuery.TicketByID.AcceptedTime)
	assert.Equal(t, ticketCreateInputTest[0].Status, *respQuery.TicketByID.Status)

	graph.CarDeleteAll(ctx, graphqlClient)
	graph.CustomerDeleteAll(ctx, graphqlClient)
	graph.ShopDeleteAll(ctx, graphqlClient)
	graph.TicketDeleteAll(ctx, graphqlClient)
}

func TestTicketFindByCustomer(t *testing.T) {
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
	_, err = graph.TicketCreate(ctx, graphqlClient, &graph.TicketCreateInput{
		CustomerID:   respCus.CustomerCreate.ID,
		CarID:        respCar.CarCreate.ID,
		Problem:      ticketCreateInputTest[0].Problem,
		CreateTime:   ticketCreateInputTest[0].CreateTime,
		ShopID:       &respShop.ShopCreate.ID,
		AcceptedTime: &ticketCreateInputTest[0].AcceptedTime,
		Status:       &ticketCreateInputTest[0].Status,
	})
	if err != nil {
		log.Fatal(err)
	}

	respQuery, err := graph.TicketByCustomer(ctx, graphqlClient, &graph.TicketByCustomerInput{
		CustomerID: respCus.CustomerCreate.ID,
		FromTime:   time.Date(2022, time.Month(6), 12, 0, 0, 0, 0, time.UTC),
		ToTime:     time.Date(2022, time.Month(8), 12, 0, 0, 0, 0, time.UTC),
		Status:     &ticketCreateInputTest[0].Status,
	})
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, respCus.CustomerCreate.ID, respQuery.TicketByCustomer[0].CustomerID)
	assert.Equal(t, respCar.CarCreate.ID, respQuery.TicketByCustomer[0].CarID)
	assert.Equal(t, ticketCreateInputTest[0].Problem, respQuery.TicketByCustomer[0].Problem)
	assert.Equal(t, ticketCreateInputTest[0].CreateTime, respQuery.TicketByCustomer[0].CreateTime)
	assert.Equal(t, respShop.ShopCreate.ID, *respQuery.TicketByCustomer[0].ShopID)
	assert.Equal(t, ticketCreateInputTest[0].AcceptedTime, *respQuery.TicketByCustomer[0].AcceptedTime)
	assert.Equal(t, ticketCreateInputTest[0].Status, *respQuery.TicketByCustomer[0].Status)

	graph.CarDeleteAll(ctx, graphqlClient)
	graph.CustomerDeleteAll(ctx, graphqlClient)
	graph.ShopDeleteAll(ctx, graphqlClient)
	graph.TicketDeleteAll(ctx, graphqlClient)
}

func TestTicketFindByShop(t *testing.T) {
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
	_, err = graph.TicketCreate(ctx, graphqlClient, &graph.TicketCreateInput{
		CustomerID:   respCus.CustomerCreate.ID,
		CarID:        respCar.CarCreate.ID,
		Problem:      ticketCreateInputTest[0].Problem,
		CreateTime:   ticketCreateInputTest[0].CreateTime,
		ShopID:       &respShop.ShopCreate.ID,
		AcceptedTime: &ticketCreateInputTest[0].AcceptedTime,
		Status:       &ticketCreateInputTest[0].Status,
	})
	if err != nil {
		log.Fatal(err)
	}

	respQuery, err := graph.TicketByShop(ctx, graphqlClient, &graph.TicketByShopInput{
		ShopID:   respShop.ShopCreate.ID,
		FromTime: time.Date(2022, time.Month(6), 12, 0, 0, 0, 0, time.UTC),
		ToTime:   time.Date(2022, time.Month(8), 12, 0, 0, 0, 0, time.UTC),
		Status:   &ticketCreateInputTest[0].Status,
	})
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, respCus.CustomerCreate.ID, respQuery.TicketByShop[0].CustomerID)
	assert.Equal(t, respCar.CarCreate.ID, respQuery.TicketByShop[0].CarID)
	assert.Equal(t, ticketCreateInputTest[0].Problem, respQuery.TicketByShop[0].Problem)
	assert.Equal(t, ticketCreateInputTest[0].CreateTime, respQuery.TicketByShop[0].CreateTime)
	assert.Equal(t, respShop.ShopCreate.ID, *respQuery.TicketByShop[0].ShopID)
	assert.Equal(t, ticketCreateInputTest[0].AcceptedTime, *respQuery.TicketByShop[0].AcceptedTime)
	assert.Equal(t, ticketCreateInputTest[0].Status, *respQuery.TicketByShop[0].Status)

	graph.CarDeleteAll(ctx, graphqlClient)
	graph.CustomerDeleteAll(ctx, graphqlClient)
	graph.ShopDeleteAll(ctx, graphqlClient)
	graph.TicketDeleteAll(ctx, graphqlClient)
}
