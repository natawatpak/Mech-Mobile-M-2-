package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/natawatpak/Mech-Mobile-M-2-/backend/graph"
	"github.com/natawatpak/Mech-Mobile-M-2-/backend/graph/generated"
	"github.com/natawatpak/Mech-Mobile-M-2-/backend/graph/model"
	"github.com/natawatpak/Mech-Mobile-M-2-/backend/resource"
	"github.com/natawatpak/Mech-Mobile-M-2-/backend/util"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

var currentExpectCustomer model.Customer = model.Customer{
	ID:    "",
	FName: "phum",
	LName: "kitiphum",
	Tel:   "0123456789",
	Email: "phum@gmail.com",
}

var inputCustomer = model.CustomerCreateInput{
	FName: currentExpectCustomer.FName,
	LName: currentExpectCustomer.LName,
	Tel:   currentExpectCustomer.Tel,
	Email: currentExpectCustomer.Email,
}

func testSetup() (context.Context, *resource.SQLop) {
	r := chi.NewRouter()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	if cancel != nil {
		fmt.Printf("Context cancel msg : %v\n\n", cancel)
	}

	viper.SetConfigName("postgresConfig")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var host string = viper.GetString("connectionDetail.host")
	var port string = viper.GetString("connectionDetail.port")
	var user string = viper.GetString("connectionDetail.user")
	var password string = viper.GetString("connectionDetail.password")
	var dbname string = viper.GetString("connectionDetail.dbname")
	var goChiPort string = viper.GetString("connectionDetail.goChiPort")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	operator, err := resource.NewDBOperator(psqlInfo)
	if util.CheckErr(err) {
		println(err.Error())
	}

	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: &graph.Resolver{
					DB: operator,
				},
			},
		),
	)

	r.Handle("/", playground.Handler("GraphQL playground", "/query"))
	r.Handle("/query", srv)

	log.Printf("connect to http://%s:%s/ for GraphQL playground", host, goChiPort)
	log.Fatal(http.ListenAndServe(":"+goChiPort, r))

	return ctx, operator
}

func TestTestSetup(t *testing.T) {
	r := chi.NewRouter()

	viper.SetConfigName("postgresConfig")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var host string = viper.GetString("connectionDetail.host")
	var port string = viper.GetString("connectionDetail.port")
	var user string = viper.GetString("connectionDetail.user")
	var password string = viper.GetString("connectionDetail.password")
	var dbname string = viper.GetString("connectionDetail.dbname")
	var goChiPort string = viper.GetString("connectionDetail.goChiPort")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	operator, err := resource.NewDBOperator(psqlInfo)
	if util.CheckErr(err) {
		println(err.Error())
	}

	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: &graph.Resolver{
					DB: operator,
				},
			},
		),
	)

	r.Handle("/", playground.Handler("GraphQL playground", "/query"))
	r.Handle("/query", srv)

	log.Printf("connect to http://%s:%s/ for GraphQL playground", host, goChiPort)
	log.Fatal(http.ListenAndServe(":"+goChiPort, r))

}

func TestTableDropFunc(t *testing.T) {
	ctx, operator := testSetup()

	dropResult, err := operator.DropTable(ctx)
	assert.Equal(t, true, dropResult)
	assert.Equal(t, nil, err)
}
func TestTableCreateFunc(t *testing.T) {
	ctx, operator := testSetup()

	createResult, err := operator.CreateTables(ctx)
	assert.Equal(t, true, createResult)
	assert.Equal(t, nil, err)
}

func TestOpCreateFunc(t *testing.T) {
	ctx, operator := testSetup()

	returnCustomer, err := operator.CustomerCreate(ctx, &inputCustomer)
	currentExpectCustomer.ID = returnCustomer.ID
	assert.Equal(t, &currentExpectCustomer, returnCustomer)
	assert.Equal(t, nil, err)

}
