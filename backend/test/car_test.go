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

func TestCarCreate(t *testing.T) {
	ctx := context.Background()

	graphqlClient := graphql.NewClient("http://localhost:8081/", http.DefaultClient)

	respCus, err := graph.CustomerCreate(ctx, graphqlClient, &graph.CustomerCreateInput{
		FName: "Phone",
		LName: "Phum",
		Tel:   "098-278-9331",
		Email: "Chic@gmail.com",
	})
	if err != nil {
		log.Fatal(err)
	}
	respCar, err := graph.CarCreate(ctx, graphqlClient, &graph.CarCreateInput{
		OwnerID:   respCus.CustomerCreate.ID,
		PlateNum:  "9999",
		PlateType: util.Ptr("white"),
		IssuedAt:  util.Ptr("Bangkok"),
		Color:     util.Ptr("Blue"),
		Type:      util.Ptr("SUV"),
		Brand:     "Toyota",
		Build:     util.Ptr("Thailand"),
	})
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, respCus.CustomerCreate.ID, respCar.CarCreate.OwnerID)
	assert.Equal(t, "9999", respCar.CarCreate.PlateNum)
	assert.Equal(t, "white", *respCar.CarCreate.PlateType)
	assert.Equal(t, "Bangkok", *respCar.CarCreate.IssuedAt)
	assert.Equal(t, "Blue", *respCar.CarCreate.Color)
	assert.Equal(t, "SUV", *respCar.CarCreate.Type)
	assert.Equal(t, "Toyota", respCar.CarCreate.Brand)
	assert.Equal(t, "Thailand", *respCar.CarCreate.Build)

	graph.CustomerDeleteAll(ctx, graphqlClient)
	graph.CarDeleteAll(ctx, graphqlClient)
}

func TestCarUpdate(t *testing.T) {
	ctx := context.Background()

	graphqlClient := graphql.NewClient("http://localhost:8081/", http.DefaultClient)

	respCus, err := graph.CustomerCreate(ctx, graphqlClient, &graph.CustomerCreateInput{
		FName: "Phone",
		LName: "Phum",
		Tel:   "098-278-9331",
		Email: "Chic@gmail.com",
	})
	if err != nil {
		log.Fatal(err)
	}
	respCarCreate, err := graph.CarCreate(ctx, graphqlClient, &graph.CarCreateInput{
		OwnerID:   respCus.CustomerCreate.ID,
		PlateNum:  "9999",
		PlateType: util.Ptr("white"),
		IssuedAt:  util.Ptr("Bangkok"),
		Color:     util.Ptr("Blue"),
		Type:      util.Ptr("SUV"),
		Brand:     "Toyota",
		Build:     util.Ptr("Thailand"),
	})
	if err != nil {
		log.Fatal(err)
	}
	respCarUpdate, err := graph.CarUpdateMulti(ctx, graphqlClient, &graph.CarUpdateInput{
		ID:        respCarCreate.CarCreate.ID,
		OwnerID:   respCus.CustomerCreate.ID,
		PlateNum:  "8888",
		PlateType: util.Ptr("Black"),
		IssuedAt:  util.Ptr("London"),
		Color:     util.Ptr("Red"),
		Type:      util.Ptr("Sport"),
		Brand:     "Honda",
		Build:     util.Ptr("USA"),
	})
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, respCus.CustomerCreate.ID, respCarUpdate.CarUpdateMulti.OwnerID)
	assert.Equal(t, "8888", respCarUpdate.CarUpdateMulti.PlateNum)
	assert.Equal(t, "Black", *respCarUpdate.CarUpdateMulti.PlateType)
	assert.Equal(t, "London", *respCarUpdate.CarUpdateMulti.IssuedAt)
	assert.Equal(t, "Red", *respCarUpdate.CarUpdateMulti.Color)
	assert.Equal(t, "Sport", *respCarUpdate.CarUpdateMulti.Type)
	assert.Equal(t, "Honda", respCarUpdate.CarUpdateMulti.Brand)
	assert.Equal(t, "USA", *respCarUpdate.CarUpdateMulti.Build)

	graph.CustomerDeleteAll(ctx, graphqlClient)
	graph.CarDeleteAll(ctx, graphqlClient)
}

func TestCarDelete(t *testing.T) {
	ctx := context.Background()

	graphqlClient := graphql.NewClient("http://localhost:8081/", http.DefaultClient)

	respCus, err := graph.CustomerCreate(ctx, graphqlClient, &graph.CustomerCreateInput{
		FName: "Phone",
		LName: "Phum",
		Tel:   "098-278-9331",
		Email: "Chic@gmail.com",
	})
	if err != nil {
		log.Fatal(err)
	}
	respCarCreate, err := graph.CarCreate(ctx, graphqlClient, &graph.CarCreateInput{
		OwnerID:   respCus.CustomerCreate.ID,
		PlateNum:  "9999",
		PlateType: util.Ptr("white"),
		IssuedAt:  util.Ptr("Bangkok"),
		Color:     util.Ptr("Blue"),
		Type:      util.Ptr("SUV"),
		Brand:     "Toyota",
		Build:     util.Ptr("Thailand"),
	})
	if err != nil {
		log.Fatal(err)
	}
	respCarDelete, err := graph.CarDelete(ctx, graphqlClient, &graph.DeleteIDInput{
		ID: respCarCreate.CarCreate.ID,
	})
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, respCus.CustomerCreate.ID, respCarDelete.CarDelete.OwnerID)
	assert.Equal(t, "9999", respCarDelete.CarDelete.PlateNum)
	assert.Equal(t, "white", *respCarDelete.CarDelete.PlateType)
	assert.Equal(t, "Bangkok", *respCarDelete.CarDelete.IssuedAt)
	assert.Equal(t, "Blue", *respCarDelete.CarDelete.Color)
	assert.Equal(t, "SUV", *respCarDelete.CarDelete.Type)
	assert.Equal(t, "Toyota", respCarDelete.CarDelete.Brand)
	assert.Equal(t, "Thailand", *respCarDelete.CarDelete.Build)

	graph.CustomerDeleteAll(ctx, graphqlClient)
	graph.CarDeleteAll(ctx, graphqlClient)
}

func TestCarDeleteAll(t *testing.T) {
	ctx := context.Background()
	graphqlClient := graphql.NewClient("http://localhost:8081/", http.DefaultClient)

	respCusCreate, err := graph.CustomerCreate(ctx, graphqlClient, &graph.CustomerCreateInput{
		FName: "Phone",
		LName: "Phum",
		Tel:   "098-278-9331",
		Email: "Chic@gmail.com",
	})
	if err != nil {
		log.Fatal(err)
	}

	type createInput struct {
		OwnerID   string
		PlateNum  string
		PlateType string
		IssuedAt  string
		Color     string
		Type      string
		Brand     string
		Build     string
	}
	createInputTest := []createInput{
		{
			OwnerID:   respCusCreate.CustomerCreate.ID,
			PlateNum:  "9999",
			PlateType: "white",
			IssuedAt:  "Bangkok",
			Color:     "Blue",
			Type:      "SUV",
			Brand:     "Toyota",
			Build:     "Thailand",
		},
		{
			OwnerID:   respCusCreate.CustomerCreate.ID,
			PlateNum:  "8888",
			PlateType: "Black",
			IssuedAt:  "London",
			Color:     "Red",
			Type:      "Sport",
			Brand:     "Honda",
			Build:     "USA",
		},
	}

	for _, v := range createInputTest {
		_, err := graph.CarCreate(ctx, graphqlClient, &graph.CarCreateInput{
			OwnerID:   v.OwnerID,
			PlateNum:  v.PlateNum,
			PlateType: util.Ptr(v.PlateType),
			IssuedAt:  util.Ptr(v.IssuedAt),
			Color:     util.Ptr(v.Color),
			Type:      util.Ptr(v.Type),
			Brand:     v.Brand,
			Build:     util.Ptr(v.Build),
		})
		if err != nil {
			log.Fatal(err)
		}
	}

	respDelete, err := graph.CarDeleteAll(ctx, graphqlClient)
	if err != nil {
		log.Fatal(err)
	}

	for i, v := range respDelete.CarDeleteAll {
		assert.Equal(t, createInputTest[i].OwnerID, v.OwnerID)
		assert.Equal(t, createInputTest[i].PlateNum, v.PlateNum)
		assert.Equal(t, createInputTest[i].PlateType, *v.PlateType)
		assert.Equal(t, createInputTest[i].IssuedAt, *v.IssuedAt)
		assert.Equal(t, createInputTest[i].Color, *v.Color)
		assert.Equal(t, createInputTest[i].Type, *v.Type)
		assert.Equal(t, createInputTest[i].Brand, v.Brand)
		assert.Equal(t, createInputTest[i].Build, *v.Build)
	}

	graph.CustomerDeleteAll(ctx, graphqlClient)

}

func TestCarByID(t *testing.T) {
	ctx := context.Background()

	graphqlClient := graphql.NewClient("http://localhost:8081/", http.DefaultClient)

	respCus, err := graph.CustomerCreate(ctx, graphqlClient, &graph.CustomerCreateInput{
		FName: "Phone",
		LName: "Phum",
		Tel:   "098-278-9331",
		Email: "Chic@gmail.com",
	})
	if err != nil {
		log.Fatal(err)
	}
	respCar, err := graph.CarCreate(ctx, graphqlClient, &graph.CarCreateInput{
		OwnerID:   respCus.CustomerCreate.ID,
		PlateNum:  "9999",
		PlateType: util.Ptr("white"),
		IssuedAt:  util.Ptr("Bangkok"),
		Color:     util.Ptr("Blue"),
		Type:      util.Ptr("SUV"),
		Brand:     "Toyota",
		Build:     util.Ptr("Thailand"),
	})
	if err != nil {
		log.Fatal(err)
	}
	createID := respCar.CarCreate.ID

	respQuery, err := graph.CarByID(ctx, graphqlClient, createID)
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, respCus.CustomerCreate.ID, respQuery.CarByID.OwnerID)
	assert.Equal(t, "9999", respQuery.CarByID.PlateNum)
	assert.Equal(t, "white", *respQuery.CarByID.PlateType)
	assert.Equal(t, "Bangkok", *respQuery.CarByID.IssuedAt)
	assert.Equal(t, "Blue", *respQuery.CarByID.Color)
	assert.Equal(t, "SUV", *respQuery.CarByID.Type)
	assert.Equal(t, "Toyota", respQuery.CarByID.Brand)
	assert.Equal(t, "Thailand", *respQuery.CarByID.Build)

	graph.CustomerDeleteAll(ctx, graphqlClient)
	graph.CarDeleteAll(ctx, graphqlClient)
}

func TestCarByOwner(t *testing.T) {
	ctx := context.Background()
	graphqlClient := graphql.NewClient("http://localhost:8081/", http.DefaultClient)

	type cusCreateInput struct {
		FName string
		LName string
		Tel   string
		Email string
	}
	cusCreateInputTest := []cusCreateInput{
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

	cusID := [2]string{}
	for i, v := range cusCreateInputTest {
		cusCreateResp, err := graph.CustomerCreate(ctx, graphqlClient, &graph.CustomerCreateInput{
			FName: v.FName,
			LName: v.LName,
			Tel:   v.Tel,
			Email: v.Email,
		})
		if err != nil {
			log.Fatal(err)
		}
		cusID[i] = cusCreateResp.CustomerCreate.ID
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
	createInputTest := []carCreateInput{
		{
			OwnerID:   cusID[0],
			PlateNum:  "9999",
			PlateType: "white",
			IssuedAt:  "Bangkok",
			Color:     "Blue",
			Type:      "SUV",
			Brand:     "Toyota",
			Build:     "Thailand",
		},
		{
			OwnerID:   cusID[0],
			PlateNum:  "8888",
			PlateType: "Black",
			IssuedAt:  "London",
			Color:     "Red",
			Type:      "Sport",
			Brand:     "Honda",
			Build:     "USA",
		},
		{
			OwnerID:   cusID[1],
			PlateNum:  "7777",
			PlateType: "Green",
			IssuedAt:  "New york",
			Color:     "yellow",
			Type:      "Bus",
			Brand:     "Volvo",
			Build:     "Germany",
		},
	}

	for _, v := range createInputTest {
		_, err := graph.CarCreate(ctx, graphqlClient, &graph.CarCreateInput{
			OwnerID:   v.OwnerID,
			PlateNum:  v.PlateNum,
			PlateType: util.Ptr(v.PlateType),
			IssuedAt:  util.Ptr(v.IssuedAt),
			Color:     util.Ptr(v.Color),
			Type:      util.Ptr(v.Type),
			Brand:     v.Brand,
			Build:     util.Ptr(v.Build),
		})
		if err != nil {
			log.Fatal(err)
		}
	}

	respQuery, err := graph.CarByOwner(ctx, graphqlClient, cusID[0])
	if err != nil {
		log.Fatal(err)
	}

	for i, v := range respQuery.CarByOwner {
		assert.Equal(t, createInputTest[i].OwnerID, v.OwnerID)
		assert.Equal(t, createInputTest[i].PlateNum, v.PlateNum)
		assert.Equal(t, createInputTest[i].PlateType, *v.PlateType)
		assert.Equal(t, createInputTest[i].IssuedAt, *v.IssuedAt)
		assert.Equal(t, createInputTest[i].Color, *v.Color)
		assert.Equal(t, createInputTest[i].Type, *v.Type)
		assert.Equal(t, createInputTest[i].Brand, v.Brand)
		assert.Equal(t, createInputTest[i].Build, *v.Build)
	}

	graph.CustomerDeleteAll(ctx, graphqlClient)
	graph.CarDeleteAll(ctx, graphqlClient)
}

func TestCars(t *testing.T) {
	ctx := context.Background()
	graphqlClient := graphql.NewClient("http://localhost:8081/", http.DefaultClient)

	respCusCreate, err := graph.CustomerCreate(ctx, graphqlClient, &graph.CustomerCreateInput{
		FName: "Phone",
		LName: "Phum",
		Tel:   "098-278-9331",
		Email: "Chic@gmail.com",
	})
	if err != nil {
		log.Fatal(err)
	}

	type createInput struct {
		OwnerID   string
		PlateNum  string
		PlateType string
		IssuedAt  string
		Color     string
		Type      string
		Brand     string
		Build     string
	}
	createInputTest := []createInput{
		{
			OwnerID:   respCusCreate.CustomerCreate.ID,
			PlateNum:  "9999",
			PlateType: "white",
			IssuedAt:  "Bangkok",
			Color:     "Blue",
			Type:      "SUV",
			Brand:     "Toyota",
			Build:     "Thailand",
		},
		{
			OwnerID:   respCusCreate.CustomerCreate.ID,
			PlateNum:  "8888",
			PlateType: "Black",
			IssuedAt:  "London",
			Color:     "Red",
			Type:      "Sport",
			Brand:     "Honda",
			Build:     "USA",
		},
	}

	for _, v := range createInputTest {
		_, err := graph.CarCreate(ctx, graphqlClient, &graph.CarCreateInput{
			OwnerID:   v.OwnerID,
			PlateNum:  v.PlateNum,
			PlateType: util.Ptr(v.PlateType),
			IssuedAt:  util.Ptr(v.IssuedAt),
			Color:     util.Ptr(v.Color),
			Type:      util.Ptr(v.Type),
			Brand:     v.Brand,
			Build:     util.Ptr(v.Build),
		})
		if err != nil {
			log.Fatal(err)
		}
	}

	respQuery, err := graph.Cars(ctx, graphqlClient)
	if err != nil {
		log.Fatal(err)
	}

	for i, v := range respQuery.Cars {
		assert.Equal(t, createInputTest[i].OwnerID, v.OwnerID)
		assert.Equal(t, createInputTest[i].PlateNum, v.PlateNum)
		assert.Equal(t, createInputTest[i].PlateType, *v.PlateType)
		assert.Equal(t, createInputTest[i].IssuedAt, *v.IssuedAt)
		assert.Equal(t, createInputTest[i].Color, *v.Color)
		assert.Equal(t, createInputTest[i].Type, *v.Type)
		assert.Equal(t, createInputTest[i].Brand, v.Brand)
		assert.Equal(t, createInputTest[i].Build, *v.Build)
	}

	graph.CustomerDeleteAll(ctx, graphqlClient)
	graph.CarDeleteAll(ctx, graphqlClient)
}
