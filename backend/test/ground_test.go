package test

import (
	"context"
	"log"
	"net/http"
	"testing"

	"github.com/Khan/genqlient/graphql"
	"github.com/natawatpak/Mech-Mobile-M-2-/backend/graph"
)

func TestTicketConnectCreate(t *testing.T) {
	ctx := context.Background()

	graphqlClient := graphql.NewClient("http://localhost:8081/", http.DefaultClient)

	resp, err := graph.ShopConnects(ctx, graphqlClient)
	if err != nil {
		log.Fatal(err)
	}

	println(resp)

	graph.CustomerDeleteAll(ctx, graphqlClient)
}
