package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/core"
	"github.com/awslabs/aws-lambda-go-api-proxy/gorillamux"
	"github.com/gorilla/mux"
	"github.com/natawatpak/Mech-Mobile-M-2-/main/controller"
)

var gorillaLambda *gorillamux.GorillaMuxAdapter

func main() {
	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Not found", r.RequestURI)
		http.Error(w, fmt.Sprintf("Not found: %s", r.RequestURI), http.StatusNotFound)
	})
	if runtime_api, _ := os.LookupEnv("AWS_LAMBDA_RUNTIME_API"); runtime_api != "" {
		r.HandleFunc("/create-user", controller.CustomerCreateProfile)
		r.HandleFunc("/update-user", controller.CustomerUpdateProfile)
		gorillaLambda = gorillamux.New(r)
		lambda.Start(Handler)
	} else {
		r.HandleFunc("/create-user", controller.CustomerCreateProfile)
		r.HandleFunc("/update-user", controller.CustomerUpdateProfile)

		fmt.Println("Server listening on port 3000")
		log.Fatal(http.ListenAndServe(":"+"3000", r))
	}

}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// If no name is provided in the HTTP request body, throw an error
	log.Println(ctx)
	log.Println(req.Body)
	log.Println(req.Headers)
	log.Println(req.HTTPMethod)
	log.Println(req.Path)
	switchAble, err := gorillaLambda.ProxyWithContext(ctx, *core.NewSwitchableAPIGatewayRequestV1(&req))
	resp := switchAble.Version1()

	//create apigate way response and add cors headers
	return events.APIGatewayProxyResponse{
		Body:       resp.Body,
		StatusCode: resp.StatusCode,
		Headers: map[string]string{
			"Access-Control-Allow-Origin":  "*",
			"Access-Control-Allow-Methods": "GET, POST, OPTIONS",
			"Access-Control-Allow-Headers": "Content-Type",
			"Content-Type":                 "application/json",
		},
	}, err
}

func checkError(err error) {
	if err != nil {
		log.Panic(err)
	}
}
