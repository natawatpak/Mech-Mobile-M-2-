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

func TestTableCreateFunc(t *testing.T) {
	ctx := context.Background()

	graphqlClient := graphql.NewClient("http://localhost:8081/", http.DefaultClient)

	resp, err := graph.CreateTable(ctx, graphqlClient)
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, true, resp.CreateTable)
}
func TestTableDropFunc(t *testing.T) {
	ctx := context.Background()

	graphqlClient := graphql.NewClient("http://localhost:8081/", http.DefaultClient)

	resp, err := graph.DropTable(ctx, graphqlClient)
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, true, resp.DropTable)
}
