package test

import (
	"context"
	"log"
	"net/http"
	"testing"

	"github.com/Khan/genqlient/graphql"
	"github.com/natawatpak/Mech-Mobile-M-2-/backend/graph"
	"github.com/stretchr/testify/assert"
)

func TestCustomerCreate(t *testing.T) {
	ctx := context.Background()

	graphqlClient := graphql.NewClient("http://localhost:8081/", http.DefaultClient)

	resp, err := graph.CustomerCreate(ctx, graphqlClient, &graph.CustomerCreateInput{
		FName: "Phone",
		LName: "Phum",
		Tel:   "098-278-9331",
		Email: "Chic@gmail.com",
	})
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, "Phone", resp.CustomerCreate.FName)
	assert.Equal(t, "Phum", resp.CustomerCreate.LName)
	assert.Equal(t, "098-278-9331", resp.CustomerCreate.Tel)
	assert.Equal(t, "Chic@gmail.com", resp.CustomerCreate.Email)

	graph.CustomerDeleteAll(ctx, graphqlClient)
}

func TestCustomerUpdate(t *testing.T) {
	ctx := context.Background()
	graphqlClient := graphql.NewClient("http://localhost:8081/", http.DefaultClient)

	respCreate, err := graph.CustomerCreate(ctx, graphqlClient, &graph.CustomerCreateInput{
		FName: "Phone",
		LName: "Phum",
		Tel:   "098-278-9331",
		Email: "Chic@gmail.com",
	})
	if err != nil {
		log.Fatal(err)
	}

	createId := respCreate.CustomerCreate.ID

	respUpdate, err := graph.CustomerUpdateMulti(ctx, graphqlClient, &graph.CustomerUpdateInput{
		ID:    createId,
		FName: "ka ka ka",
		LName: "la la la",
		Tel:   "987654321",
		Email: "Update@hotmail.com",
	})
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, "ka ka ka", respUpdate.CustomerUpdateMulti.FName)
	assert.Equal(t, "la la la", respUpdate.CustomerUpdateMulti.LName)
	assert.Equal(t, "987654321", respUpdate.CustomerUpdateMulti.Tel)
	assert.Equal(t, "Update@hotmail.com", respUpdate.CustomerUpdateMulti.Email)

	graph.CustomerDeleteAll(ctx, graphqlClient)
}

func TestCustomerDelete(t *testing.T) {
	ctx := context.Background()
	graphqlClient := graphql.NewClient("http://localhost:8081/", http.DefaultClient)

	respCreate, err := graph.CustomerCreate(ctx, graphqlClient, &graph.CustomerCreateInput{
		FName: "Phone",
		LName: "Phum",
		Tel:   "098-278-9331",
		Email: "Chic@gmail.com",
	})
	if err != nil {
		log.Fatal(err)
	}

	createId := respCreate.CustomerCreate.ID

	respDelete, err := graph.CustomerDelete(ctx, graphqlClient, &graph.DeleteIDInput{
		ID: createId,
	})
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, "Phone", respDelete.CustomerDelete.FName)
	assert.Equal(t, "Phum", respDelete.CustomerDelete.LName)
	assert.Equal(t, "098-278-9331", respDelete.CustomerDelete.Tel)
	assert.Equal(t, "Chic@gmail.com", respDelete.CustomerDelete.Email)
}

func TestCustomerDeleteAll(t *testing.T) {
	ctx := context.Background()
	graphqlClient := graphql.NewClient("http://localhost:8081/", http.DefaultClient)

	type createInput struct {
		FName string
		LName string
		Tel   string
		Email string
	}
	createInputTest := []createInput{
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

	_, err := graph.CustomerCreate(ctx, graphqlClient, &graph.CustomerCreateInput{
		FName: createInputTest[0].FName,
		LName: createInputTest[0].LName,
		Tel:   createInputTest[0].Tel,
		Email: createInputTest[0].Email,
	})
	if err != nil {
		log.Fatal(err)
	}

	_, err = graph.CustomerCreate(ctx, graphqlClient, &graph.CustomerCreateInput{
		FName: createInputTest[1].FName,
		LName: createInputTest[1].LName,
		Tel:   createInputTest[1].Tel,
		Email: createInputTest[1].Email,
	})
	if err != nil {
		log.Fatal(err)
	}

	respDelete, err := graph.CustomerDeleteAll(ctx, graphqlClient)
	if err != nil {
		log.Fatal(err)
	}

	for i, v := range respDelete.CustomerDeleteAll {
		assert.Equal(t, createInputTest[i].FName, v.FName)
		assert.Equal(t, createInputTest[i].LName, v.LName)
		assert.Equal(t, createInputTest[i].Tel, v.Tel)
		assert.Equal(t, createInputTest[i].Email, v.Email)
	}

}

func TestCustomerByID(t *testing.T) {
	ctx := context.Background()
	graphqlClient := graphql.NewClient("http://localhost:8081/", http.DefaultClient)

	type createInput struct {
		FName string
		LName string
		Tel   string
		Email string
	}
	createInputTest := []createInput{
		{
			FName: "Phone",
			LName: "Phum",
			Tel:   "098-278-9331",
			Email: "Chic@gmail.com",
		},
	}

	respCreate, err := graph.CustomerCreate(ctx, graphqlClient, &graph.CustomerCreateInput{
		FName: createInputTest[0].FName,
		LName: createInputTest[0].LName,
		Tel:   createInputTest[0].Tel,
		Email: createInputTest[0].Email,
	})
	if err != nil {
		log.Fatal(err)
	}
	createID := respCreate.CustomerCreate.ID

	respDelete, err := graph.CustomerByID(ctx, graphqlClient, createID)
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, createInputTest[0].FName, respDelete.CustomerByID.FName)
	assert.Equal(t, createInputTest[0].LName, respDelete.CustomerByID.LName)
	assert.Equal(t, createInputTest[0].Tel, respDelete.CustomerByID.Tel)
	assert.Equal(t, createInputTest[0].Email, respDelete.CustomerByID.Email)

	graph.CustomerDeleteAll(ctx, graphqlClient)
}

func TestCustomers(t *testing.T) {
	ctx := context.Background()
	graphqlClient := graphql.NewClient("http://localhost:8081/", http.DefaultClient)

	type createInput struct {
		FName string
		LName string
		Tel   string
		Email string
	}
	createInputTest := []createInput{
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

	_, err := graph.CustomerCreate(ctx, graphqlClient, &graph.CustomerCreateInput{
		FName: createInputTest[0].FName,
		LName: createInputTest[0].LName,
		Tel:   createInputTest[0].Tel,
		Email: createInputTest[0].Email,
	})
	if err != nil {
		log.Fatal(err)
	}

	_, err = graph.CustomerCreate(ctx, graphqlClient, &graph.CustomerCreateInput{
		FName: createInputTest[1].FName,
		LName: createInputTest[1].LName,
		Tel:   createInputTest[1].Tel,
		Email: createInputTest[1].Email,
	})
	if err != nil {
		log.Fatal(err)
	}

	respDelete, err := graph.Customers(ctx, graphqlClient)
	if err != nil {
		log.Fatal(err)
	}

	for i, v := range respDelete.Customers {
		assert.Equal(t, createInputTest[i].FName, v.FName)
		assert.Equal(t, createInputTest[i].LName, v.LName)
		assert.Equal(t, createInputTest[i].Tel, v.Tel)
		assert.Equal(t, createInputTest[i].Email, v.Email)
	}

	graph.CustomerDeleteAll(ctx, graphqlClient)
}
