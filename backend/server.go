package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/core"
	"github.com/awslabs/aws-lambda-go-api-proxy/gorillamux"
	routerProxy "github.com/awslabs/aws-lambda-go-api-proxy/gorillamux"
	"github.com/gorilla/mux"
	"github.com/natawatpak/Mech-Mobile-M-2-/backend/graph"
	"github.com/natawatpak/Mech-Mobile-M-2-/backend/graph/generated"
	"github.com/natawatpak/Mech-Mobile-M-2-/backend/resource"
	"github.com/natawatpak/Mech-Mobile-M-2-/backend/util"
)

var gorillaLambda *routerProxy.GorillaMuxAdapter

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

func main() {
	r := mux.NewRouter()

	// config
	var host string = "localhost"    //IP
	var port string = "5432"         //port
	var user string = "postgres"     //root user
	var password string = "Eauu0244" //root pass
	var dbname string = "postgres"   //dbname
	var goChiPort string = "8081"

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	operator, err := resource.NewDBOperator(psqlInfo)
	util.CheckErr(err)
	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Not found", r.RequestURI)
		http.Error(w, fmt.Sprintf("Not found: %s", r.RequestURI), http.StatusNotFound)
	})

	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: &graph.Resolver{
					DB: operator,
				},
			},
		),
	)
	r.Handle("/carservice", srv)
	r.Handle("/playground", playground.Handler("GraphQL playground", "/"))

	if runtime_api, _ := os.LookupEnv("AWS_LAMBDA_RUNTIME_API"); runtime_api != "" {
		gorillaLambda = gorillamux.New(r)
		lambda.Start(Handler)
	} else {

		log.Printf("connect to http://%s:%s/ for GraphQL playground", host, goChiPort)
		log.Fatal(http.ListenAndServe(":"+goChiPort, r))
	}
}
