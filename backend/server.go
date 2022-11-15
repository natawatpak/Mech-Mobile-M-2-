package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gorilla/mux"
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/natawatpak/Mech-Mobile-M-2-/backend/graph"
	"github.com/natawatpak/Mech-Mobile-M-2-/backend/graph/generated"
	"github.com/natawatpak/Mech-Mobile-M-2-/backend/resource"
	"github.com/natawatpak/Mech-Mobile-M-2-/backend/util"
)

var (
	QueryNameNotProvided = errors.New("no query was provided in the HTTP body")
	mainSchema           *graphql.Schema
)

func init() {
	// start the mux router
	r := mux.NewRouter()

	// viper.SetConfigName("postgresConfig")
	// viper.SetConfigType("json")
	// viper.AddConfigPath(".")
	// err := viper.ReadInConfig()
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	var host string = "localhost"
	var port string = "5432"
	var user string = "postgres"
	var password string = "Eauu0244"
	var dbname string = "postgres"

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	operator, err := resource.NewDBOperator(psqlInfo)

	util.CheckErr(err)
	schema := generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DB: operator}})
	server := handler.NewDefaultServer(schema)
	r.Handle("/query", server)
	r.Handle("/", playground.Handler("GraphQL playground", "/query"))
}

// Handler is a wrapper around lambda and mux so that we can continue to use mux via lambda
func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if len(request.Body) < 1 {
		return events.APIGatewayProxyResponse{}, QueryNameNotProvided
	}

	var params struct {
		Query         string                 `json:"query"`
		OperationName string                 `json:"operationName"`
		Variables     map[string]interface{} `json:"variables"`
	}
	if err := json.Unmarshal([]byte(request.Body), &params); err != nil {
		log.Print("Could not decode body", err)
	}

	response := mainSchema.Exec(ctx, params.Query, params.OperationName, params.Variables)
	responseJSON, err := json.Marshal(response)
	if err != nil {
		log.Print("Could not decode body")
	}

	return events.APIGatewayProxyResponse{
		Body:       string(responseJSON),
		StatusCode: 200,
	}, nil
}

func main() {
	// r := chi.NewRouter()

	// viper.SetConfigName("postgresConfig")
	// viper.SetConfigType("json")
	// viper.AddConfigPath(".")
	// err := viper.ReadInConfig()
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	// var host string = viper.GetString("connectionDetail.host")
	// var port string = viper.GetString("connectionDetail.port")
	// var user string = viper.GetString("connectionDetail.user")
	// var password string = viper.GetString("connectionDetail.password")
	// var dbname string = viper.GetString("connectionDetail.dbname")
	// var goChiPort string = viper.GetString("connectionDetail.goChiPort")

	// psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
	// 	"password=%s dbname=%s sslmode=disable",
	// 	host, port, user, password, dbname)

	// operator, err := resource.NewDBOperator(psqlInfo)
	// util.CheckErr(err)

	// srv := handler.NewDefaultServer(
	// 	generated.NewExecutableSchema(
	// 		generated.Config{
	// 			Resolvers: &graph.Resolver{
	// 				DB: operator,
	// 			},
	// 		},
	// 	),
	// )

	// r.Handle("/", playground.Handler("GraphQL playground", "/query"))
	// r.Handle("/query", srv)

	// log.Printf("connect to http://%s:%s/ for GraphQL playground", host, goChiPort)
	// log.Fatal(http.ListenAndServe(":"+goChiPort, r))

	lambda.Start(Handler)
}
