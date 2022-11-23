package test

import (
	"context"
	"log"
	"net/http"
	"testing"

	"github.com/Khan/genqlient/graphql"
	"github.com/natawatpak/Mech-Mobile-M-2-/backend/graph"
	"github.com/natawatpak/Mech-Mobile-M-2-/backend/util"
	"github.com/stretchr/testify/assert"
)

func TestActiveTicketCreate(t *testing.T) {
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
		Name:      shopCreateInputTest[0].Name,
		Tel:       shopCreateInputTest[0].Tel,
		Email:     shopCreateInputTest[0].Email,
		Address:   shopCreateInputTest[0].Address,
		Longitude: shopCreateInputTest[0].longitude,
		Latitude:  shopCreateInputTest[0].latitude,
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
	respActiveTicket, err := graph.ActiveTicketCreate(ctx, graphqlClient, &graph.ActiveTicketCreateInput{
		ID:         respTicket.TicketCreate.ID,
		CustomerID: respTicket.TicketCreate.CustomerID,
		CarID:      respTicket.TicketCreate.CarID,
		Problem:    respTicket.TicketCreate.Problem,
		ShopID:     respTicket.TicketCreate.ShopID,
		Status:     respTicket.TicketCreate.Status,
	})
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, respCus.CustomerCreate.ID, respActiveTicket.ActiveTicketCreate.CustomerID)
	assert.Equal(t, respCar.CarCreate.ID, respActiveTicket.ActiveTicketCreate.CarID)
	assert.Equal(t, ticketCreateInputTest[0].Problem, respActiveTicket.ActiveTicketCreate.Problem)
	assert.Equal(t, respShop.ShopCreate.ID, *respActiveTicket.ActiveTicketCreate.ShopID)
	assert.Equal(t, ticketCreateInputTest[0].Status, *respActiveTicket.ActiveTicketCreate.Status)

	graph.CarDeleteAll(ctx, graphqlClient)
	graph.CustomerDeleteAll(ctx, graphqlClient)
	graph.ShopDeleteAll(ctx, graphqlClient)
	graph.ActiveTicketDeleteAll(ctx, graphqlClient)
}

func TestActiveTicketUpdate(t *testing.T) {
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
	respActiveTicketCreate, err := graph.ActiveTicketCreate(ctx, graphqlClient, &graph.ActiveTicketCreateInput{
		ID:         respTicket.TicketCreate.ID,
		CustomerID: respTicket.TicketCreate.CustomerID,
		CarID:      respTicket.TicketCreate.CarID,
		Problem:    respTicket.TicketCreate.Problem,
		ShopID:     respTicket.TicketCreate.ShopID,
		Status:     respTicket.TicketCreate.Status,
	})
	if err != nil {
		log.Fatal(err)
	}
	respActiveTicketUpdate, err := graph.ActiveTicketUpdateMulti(ctx, graphqlClient, &graph.ActiveTicketUpdateInput{
		ID:         respActiveTicketCreate.ActiveTicketCreate.ID,
		CustomerID: respTicket.TicketCreate.CustomerID,
		CarID:      respTicket.TicketCreate.CarID,
		Problem:    ticketCreateInputTest[1].Problem,
		ShopID:     respTicket.TicketCreate.ShopID,
		Status:     &ticketCreateInputTest[2].Status,
	})
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, respCus.CustomerCreate.ID, respActiveTicketUpdate.ActiveTicketUpdateMulti.CustomerID)
	assert.Equal(t, respCar.CarCreate.ID, respActiveTicketUpdate.ActiveTicketUpdateMulti.CarID)
	assert.Equal(t, ticketCreateInputTest[1].Problem, respActiveTicketUpdate.ActiveTicketUpdateMulti.Problem)
	assert.Equal(t, respShop.ShopCreate.ID, *respActiveTicketUpdate.ActiveTicketUpdateMulti.ShopID)
	assert.Equal(t, ticketCreateInputTest[2].Status, *respActiveTicketUpdate.ActiveTicketUpdateMulti.Status)

	graph.CarDeleteAll(ctx, graphqlClient)
	graph.CustomerDeleteAll(ctx, graphqlClient)
	graph.ShopDeleteAll(ctx, graphqlClient)
	graph.ActiveTicketDeleteAll(ctx, graphqlClient)
}

func TestActiveTicketDelete(t *testing.T) {
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
	respActiveTicketCreate, err := graph.ActiveTicketCreate(ctx, graphqlClient, &graph.ActiveTicketCreateInput{
		ID:         respTicket.TicketCreate.ID,
		CustomerID: respTicket.TicketCreate.CustomerID,
		CarID:      respTicket.TicketCreate.CarID,
		Problem:    respTicket.TicketCreate.Problem,
		ShopID:     respTicket.TicketCreate.ShopID,
		Status:     respTicket.TicketCreate.Status,
	})
	if err != nil {
		log.Fatal(err)
	}
	respActiveTicketDelete, err := graph.ActiveTicketDelete(ctx, graphqlClient, &graph.DeleteIDInput{
		ID: respActiveTicketCreate.ActiveTicketCreate.ID,
	})
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, respCus.CustomerCreate.ID, respActiveTicketDelete.ActiveTicketDelete.CustomerID)
	assert.Equal(t, respCar.CarCreate.ID, respActiveTicketDelete.ActiveTicketDelete.CarID)
	assert.Equal(t, ticketCreateInputTest[0].Problem, respActiveTicketDelete.ActiveTicketDelete.Problem)
	assert.Equal(t, respShop.ShopCreate.ID, *respActiveTicketDelete.ActiveTicketDelete.ShopID)
	assert.Equal(t, ticketCreateInputTest[0].Status, *respActiveTicketDelete.ActiveTicketDelete.Status)

	graph.CarDeleteAll(ctx, graphqlClient)
	graph.CustomerDeleteAll(ctx, graphqlClient)
	graph.ShopDeleteAll(ctx, graphqlClient)
	graph.ActiveTicketDeleteAll(ctx, graphqlClient)
}

func TestActiveTicketDeleteActive(t *testing.T) {
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
	respActiveTicketCreate, err := graph.ActiveTicketCreate(ctx, graphqlClient, &graph.ActiveTicketCreateInput{
		ID:         respTicket.TicketCreate.ID,
		CustomerID: respTicket.TicketCreate.CustomerID,
		CarID:      respTicket.TicketCreate.CarID,
		Problem:    respTicket.TicketCreate.Problem,
		ShopID:     respTicket.TicketCreate.ShopID,
		Status:     respTicket.TicketCreate.Status,
	})
	if err != nil {
		log.Fatal(err)
	}
	respActiveTicketDelete, err := graph.ActiveTicketDeleteStatus(ctx, graphqlClient, *respActiveTicketCreate.ActiveTicketCreate.Status)
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, respCus.CustomerCreate.ID, respActiveTicketDelete.ActiveTicketDeleteStatus[0].CustomerID)
	assert.Equal(t, respCar.CarCreate.ID, respActiveTicketDelete.ActiveTicketDeleteStatus[0].CarID)
	assert.Equal(t, ticketCreateInputTest[0].Problem, respActiveTicketDelete.ActiveTicketDeleteStatus[0].Problem)
	assert.Equal(t, respShop.ShopCreate.ID, *respActiveTicketDelete.ActiveTicketDeleteStatus[0].ShopID)
	assert.Equal(t, ticketCreateInputTest[0].Status, *respActiveTicketDelete.ActiveTicketDeleteStatus[0].Status)

	graph.CarDeleteAll(ctx, graphqlClient)
	graph.CustomerDeleteAll(ctx, graphqlClient)
	graph.ShopDeleteAll(ctx, graphqlClient)
	graph.TicketDeleteAll(ctx, graphqlClient)
	graph.ActiveTicketDeleteAll(ctx, graphqlClient)
}

func TestActiveTicketByID(t *testing.T) {
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
	respActiveTicketCreate, err := graph.ActiveTicketCreate(ctx, graphqlClient, &graph.ActiveTicketCreateInput{
		ID:         respTicket.TicketCreate.ID,
		CustomerID: respTicket.TicketCreate.CustomerID,
		CarID:      respTicket.TicketCreate.CarID,
		Problem:    respTicket.TicketCreate.Problem,
		ShopID:     respTicket.TicketCreate.ShopID,
		Status:     respTicket.TicketCreate.Status,
	})
	if err != nil {
		log.Fatal(err)
	}
	respActiveTicketQuery, err := graph.ActiveTicketByID(ctx, graphqlClient, respActiveTicketCreate.ActiveTicketCreate.ID)
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, respCus.CustomerCreate.ID, respActiveTicketQuery.ActiveTicketByID.CustomerID)
	assert.Equal(t, respCar.CarCreate.ID, respActiveTicketQuery.ActiveTicketByID.CarID)
	assert.Equal(t, ticketCreateInputTest[0].Problem, respActiveTicketQuery.ActiveTicketByID.Problem)
	assert.Equal(t, respShop.ShopCreate.ID, *respActiveTicketQuery.ActiveTicketByID.ShopID)
	assert.Equal(t, ticketCreateInputTest[0].Status, *respActiveTicketQuery.ActiveTicketByID.Status)

	graph.CarDeleteAll(ctx, graphqlClient)
	graph.CustomerDeleteAll(ctx, graphqlClient)
	graph.ShopDeleteAll(ctx, graphqlClient)
	graph.TicketDeleteAll(ctx, graphqlClient)
	graph.ActiveTicketDeleteAll(ctx, graphqlClient)
}

func TestActiveTicketByCustomer(t *testing.T) {
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
	respActiveTicketCreate, err := graph.ActiveTicketCreate(ctx, graphqlClient, &graph.ActiveTicketCreateInput{
		ID:         respTicket.TicketCreate.ID,
		CustomerID: respTicket.TicketCreate.CustomerID,
		CarID:      respTicket.TicketCreate.CarID,
		Problem:    respTicket.TicketCreate.Problem,
		ShopID:     respTicket.TicketCreate.ShopID,
		Status:     respTicket.TicketCreate.Status,
	})
	if err != nil {
		log.Fatal(err)
	}
	respActiveTicketQuery, err := graph.ActiveTicketByCustomer(ctx, graphqlClient, respActiveTicketCreate.ActiveTicketCreate.CustomerID)
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, respCus.CustomerCreate.ID, respActiveTicketQuery.ActiveTicketByCustomer[0].CustomerID)
	assert.Equal(t, respCar.CarCreate.ID, respActiveTicketQuery.ActiveTicketByCustomer[0].CarID)
	assert.Equal(t, ticketCreateInputTest[0].Problem, respActiveTicketQuery.ActiveTicketByCustomer[0].Problem)
	assert.Equal(t, respShop.ShopCreate.ID, *respActiveTicketQuery.ActiveTicketByCustomer[0].ShopID)
	assert.Equal(t, ticketCreateInputTest[0].Status, *respActiveTicketQuery.ActiveTicketByCustomer[0].Status)

	graph.CarDeleteAll(ctx, graphqlClient)
	graph.CustomerDeleteAll(ctx, graphqlClient)
	graph.ShopDeleteAll(ctx, graphqlClient)
	graph.TicketDeleteAll(ctx, graphqlClient)
	graph.ActiveTicketDeleteAll(ctx, graphqlClient)
}

func TestActiveTicketTicketFindByShop(t *testing.T) {
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
	respActiveTicketCreate, err := graph.ActiveTicketCreate(ctx, graphqlClient, &graph.ActiveTicketCreateInput{
		ID:         respTicket.TicketCreate.ID,
		CustomerID: respTicket.TicketCreate.CustomerID,
		CarID:      respTicket.TicketCreate.CarID,
		Problem:    respTicket.TicketCreate.Problem,
		ShopID:     respTicket.TicketCreate.ShopID,
		Status:     respTicket.TicketCreate.Status,
	})
	if err != nil {
		log.Fatal(err)
	}
	respActiveTicketQuery, err := graph.ActiveTicketByShop(ctx, graphqlClient, *respActiveTicketCreate.ActiveTicketCreate.ShopID)
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, respCus.CustomerCreate.ID, respActiveTicketQuery.ActiveTicketByShop[0].CustomerID)
	assert.Equal(t, respCar.CarCreate.ID, respActiveTicketQuery.ActiveTicketByShop[0].CarID)
	assert.Equal(t, ticketCreateInputTest[0].Problem, respActiveTicketQuery.ActiveTicketByShop[0].Problem)
	assert.Equal(t, respShop.ShopCreate.ID, *respActiveTicketQuery.ActiveTicketByShop[0].ShopID)
	assert.Equal(t, ticketCreateInputTest[0].Status, *respActiveTicketQuery.ActiveTicketByShop[0].Status)

	graph.CarDeleteAll(ctx, graphqlClient)
	graph.CustomerDeleteAll(ctx, graphqlClient)
	graph.ShopDeleteAll(ctx, graphqlClient)
	graph.TicketDeleteAll(ctx, graphqlClient)
	graph.ActiveTicketDeleteAll(ctx, graphqlClient)
}
