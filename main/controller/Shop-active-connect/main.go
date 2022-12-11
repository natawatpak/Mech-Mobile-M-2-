package main

import (
	"context"
	"time"
	"net/http"
	"github.com/Khan/genqlient/graphql"
	"github.com/natawatpak/Mech-Mobile-M-2-/backend/graph"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

const GRAPHQL_CLIENT_URL = "https://a7okax4857.execute-api.us-east-1.amazonaws.com/default/carservice"


func main() {
	lambda.Start(handler)
}

func handler(_ context.Context, request events.APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyResponse, error) {

  ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	graphqlClient := graphql.NewClient(GRAPHQL_CLIENT_URL, http.DefaultClient)

	resp, err := graph.TicketConnectsByID(ctx, graphqlClient, request.QueryStringParameters["ticketID"])

	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       err.Error(),
		}, err
	}
  
  ctx, _ = context.WithTimeout(context.Background(), 30*time.Second)
	graphqlClient = graphql.NewClient(GRAPHQL_CLIENT_URL, http.DefaultClient)

	_, err = graph.TicketConnectsCreate(ctx, graphqlClient, &graph.TicketConnectCreateInput{
        TicketID: request.QueryStringParameters["ticketID"],
        CustomerConnectionID: resp.TicketConnectByID.CustomerConnectionID,
        ShopConnectionID: request.RequestContext.ConnectionID,
    })

	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       err.Error(),
		}, err
	}
  
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "OK",
	}, nil
}