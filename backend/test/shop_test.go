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

type shopCreateInput struct {
	Name      string
	Tel       string
	Email     string
	Address   string
	longitude string
	latitude  string
}

var shopCreateInputTest = []shopCreateInput{
	{
		Name:      "Phone's auto",
		Tel:       "098-278-9331",
		Email:     "Chic@gmail.com",
		Address:   "second street",
		longitude: "1000",
		latitude:  "1000",
	},
	{
		Name:      "chicken's auto",
		Tel:       "123456789",
		Email:     "Chicken@gmail.com",
		Address:   "3/4 street",
		longitude: "1500",
		latitude:  "1500",
	},
}

func TestShopCreate(t *testing.T) {
	ctx := context.Background()

	graphqlClient := graphql.NewClient("http://localhost:8081/", http.DefaultClient)

	resp, err := graph.ShopCreate(ctx, graphqlClient, &graph.ShopCreateInput{
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
	assert.Equal(t, shopCreateInputTest[0].Name, resp.ShopCreate.Name)
	assert.Equal(t, shopCreateInputTest[0].Tel, resp.ShopCreate.Tel)
	assert.Equal(t, shopCreateInputTest[0].Email, resp.ShopCreate.Email)
	assert.Equal(t, shopCreateInputTest[0].Address, resp.ShopCreate.Address)

	graph.ShopDeleteAll(ctx, graphqlClient)
}

func TestShopUpdate(t *testing.T) {
	ctx := context.Background()

	graphqlClient := graphql.NewClient("http://localhost:8081/", http.DefaultClient)

	respCreate, err := graph.ShopCreate(ctx, graphqlClient, &graph.ShopCreateInput{
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
	respUpdate, err := graph.ShopUpdateMulti(ctx, graphqlClient, &graph.ShopUpdateInput{
		ID:        respCreate.ShopCreate.ID,
		Name:      shopCreateInputTest[1].Name,
		Tel:       shopCreateInputTest[1].Tel,
		Email:     shopCreateInputTest[1].Email,
		Address:   shopCreateInputTest[1].Address,
		Longitude: shopCreateInputTest[1].longitude,
		Latitude:  shopCreateInputTest[1].latitude,
	})
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, shopCreateInputTest[1].Name, respUpdate.ShopUpdateMulti.Name)
	assert.Equal(t, shopCreateInputTest[1].Tel, respUpdate.ShopUpdateMulti.Tel)
	assert.Equal(t, shopCreateInputTest[1].Email, respUpdate.ShopUpdateMulti.Email)
	assert.Equal(t, shopCreateInputTest[1].Address, respUpdate.ShopUpdateMulti.Address)

	graph.ShopDeleteAll(ctx, graphqlClient)
}

func TestShopDelete(t *testing.T) {
	ctx := context.Background()

	graphqlClient := graphql.NewClient("http://localhost:8081/", http.DefaultClient)

	respCreate, err := graph.ShopCreate(ctx, graphqlClient, &graph.ShopCreateInput{
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
	respDelete, err := graph.ShopDelete(ctx, graphqlClient, &graph.DeleteIDInput{
		ID: respCreate.ShopCreate.ID,
	})
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, shopCreateInputTest[0].Name, respDelete.ShopDelete.Name)
	assert.Equal(t, shopCreateInputTest[0].Tel, respDelete.ShopDelete.Tel)
	assert.Equal(t, shopCreateInputTest[0].Email, respDelete.ShopDelete.Email)
	assert.Equal(t, shopCreateInputTest[0].Address, respDelete.ShopDelete.Address)

	graph.ShopDeleteAll(ctx, graphqlClient)
}

func TestShopDeleteAll(t *testing.T) {
	ctx := context.Background()
	graphqlClient := graphql.NewClient("http://localhost:8081/", http.DefaultClient)

	for _, v := range shopCreateInputTest {
		_, err := graph.ShopCreate(ctx, graphqlClient, &graph.ShopCreateInput{
			Name:    v.Name,
			Tel:     v.Tel,
			Email:   v.Email,
			Address: v.Address,
		})
		if err != nil {
			log.Fatal(err)
		}
	}

	respDelete, err := graph.ShopDeleteAll(ctx, graphqlClient)
	if err != nil {
		log.Fatal(err)
	}

	for i, v := range respDelete.ShopDeleteAll {
		assert.Equal(t, shopCreateInputTest[i].Name, v.Name)
		assert.Equal(t, shopCreateInputTest[i].Tel, v.Tel)
		assert.Equal(t, shopCreateInputTest[i].Email, v.Email)
		assert.Equal(t, shopCreateInputTest[i].Address, v.Address)
	}

	graph.ShopDeleteAll(ctx, graphqlClient)
}

func TestShopByID(t *testing.T) {
	ctx := context.Background()

	graphqlClient := graphql.NewClient("http://localhost:8081/", http.DefaultClient)

	resp, err := graph.ShopCreate(ctx, graphqlClient, &graph.ShopCreateInput{
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

	respQuery, err := graph.ShopByID(ctx, graphqlClient, resp.ShopCreate.ID)
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, shopCreateInputTest[0].Name, respQuery.ShopByID.Name)
	assert.Equal(t, shopCreateInputTest[0].Tel, respQuery.ShopByID.Tel)
	assert.Equal(t, shopCreateInputTest[0].Email, respQuery.ShopByID.Email)
	assert.Equal(t, shopCreateInputTest[0].Address, respQuery.ShopByID.Address)

	graph.ShopDeleteAll(ctx, graphqlClient)
}
