package main

import (
	"context"
	"fmt"
	"log"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/core"
	"github.com/awslabs/aws-lambda-go-api-proxy/gorillamux"
	"github.com/gorilla/mux"
	"github.com/natawatpak/Mech-Mobile-M-2-/backend/graph"
	"github.com/natawatpak/Mech-Mobile-M-2-/backend/graph/generated"
	"github.com/natawatpak/Mech-Mobile-M-2-/backend/resource"
	"github.com/natawatpak/Mech-Mobile-M-2-/backend/util"
)

var muxAdapter *gorillamux.GorillaMuxAdapter

// var apiGatewayAdapter *handlerfunc.HandlerFuncAdapter

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
	// var goChiPort string = viper.GetString("connectionDetail.goChiPort")

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
func Handler(ctx context.Context, req core.SwitchableAPIGatewayRequest) (*core.SwitchableAPIGatewayResponse, error) {
	rsp, err := muxAdapter.Proxy(req)
	if err != nil {
		log.Println(err)
	}
	return rsp, err
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
