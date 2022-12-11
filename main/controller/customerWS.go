package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/apigatewaymanagementapi"
)

type WebSocketRequestBody struct {
	Action string          `json:"action"`
	Data   json.RawMessage `json:"data"`
}

type WebSocketResponseBody struct {
	Message string `json:"message"`
}

func Handler(context context.Context, request events.APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Get the message sent over the socket
	connectionID := request.RequestContext.ConnectionID
	// This the endpoint that is required to send a message using apigatewayManagementAPI
	callbackurl := fmt.Sprintf("https://%s/%s/@connections/%s", request.RequestContext.DomainName, request.RequestContext.Stage, connectionID)

	var wsrb WebSocketRequestBody
	log.Printf(request.Body)
	log.Printf(request.RequestContext.ConnectionID)
	log.Print(callbackurl)
	err := json.Unmarshal([]byte(request.Body), &wsrb)
	if err != nil {
		log.Print(0)
		log.Print(err)
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
			Body: "Error failed run",
		}, nil
	}


	// The body of the response will be returned to the client over the socket
	r := WebSocketResponseBody{
		Message: "via routeResponseSelectionExpression" + string(wsrb.Data),
	}
	b, err := json.Marshal(r)
	if err != nil {
		log.Print(1)
		log.Print(err)
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
			Body: "Error failed run",
		}, nil
	}

	//perform post request to callbackurl
	mySession := session.Must(session.NewSession())

	// Create a ApiGatewayManagementApi client
	apig := apigatewaymanagementapi.New(mySession)
	apig.Endpoint = "https://6eblsxltxc.execute-api.us-east-1.amazonaws.com/production"
	resp, err := apig.PostToConnection(&apigatewaymanagementapi.PostToConnectionInput{
		ConnectionId: aws.String(connectionID),
		Data:         b,
	})

	if err != nil {
		log.Print("manager broke")
		log.Print(err)
		log.Print(resp)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(b),
	}, nil
}

func main() {
	lambda.Start(Handler)
}