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

	graphqlClient := graphql.NewClient("https://a7okax4857.execute-api.us-east-1.amazonaws.com/default/carservice", http.DefaultClient)

	resp, err := graph.TicketConnectsUpdateMulti(ctx, graphqlClient, &graph.TicketConnectCreateInput{
		TicketID:             "5c4e2ade-e9d7-4d37-b77e-0d27e523c8aa",
		CustomerConnectionID: "c_pKGeSGoAMCEdQ=",
		ShopConnectionID:     "123456789",
	})
	if err != nil {
		log.Fatal(err)
	}

	println(resp)

	graph.CustomerDeleteAll(ctx, graphqlClient)
}
