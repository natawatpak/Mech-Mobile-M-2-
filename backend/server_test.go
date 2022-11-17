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
	"github.com/Khan/genqlient/graphql"
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

func testSetup1() (context.Context, *resource.SQLop) {
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

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	operator, err := resource.NewDBOperator(psqlInfo)
	if util.CheckErr(err) {
		println(err.Error())
	}

	return ctx, operator
}

func testSetup2() {
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

	r.Handle("/playground", playground.Handler("GraphQL playground", "/"))
	r.Handle("/", srv)

}

func TestTableDropFunc(t *testing.T) {
	ctx, operator := testSetup1()

	dropResult, err := operator.DropTable(ctx)
	assert.Equal(t, true, dropResult)
	assert.Equal(t, nil, err)
}
func TestTableCreateFunc(t *testing.T) {
	ctx, operator := testSetup1()

	createResult, err := operator.CreateTables(ctx)
	assert.Equal(t, true, createResult)
	assert.Equal(t, nil, err)
}

func TestOpCreateFunc(t *testing.T) {
	ctx, operator := testSetup1()

	returnCustomer, err := operator.CustomerCreate(ctx, &inputCustomer)
	currentExpectCustomer.ID = returnCustomer.ID
	assert.Equal(t, &currentExpectCustomer, returnCustomer)
	assert.Equal(t, nil, err)

}

func TestTicketCreate(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	graphqlClient := graphql.NewClient("http://localhost:8081/", http.DefaultClient)
	curentTime := time.Now()
	log.Println(curentTime)

	resp, err := graph.TicketCreate(ctx, graphqlClient, &graph.TicketCreateInput{
		CustomerID:   "f4209ba4-9e88-4b9b-a43b-7f71c96c43a8",
		CarID:        "dd13c45c-1e15-4d7d-a96a-94b161086b4b",
		Problem:      "bad battery",
		CreateTime:   curentTime,
		ShopID:       "d9ef6fb7-837c-4433-be81-f24f5943b3d0",
		AcceptedTime: &curentTime,
		Status:       util.Ptr("active"),
	})
	if err != nil {
		log.Fatal(err)
	}
	ticketCreate := resp.TicketCreate
	assert.Equal(t, "bad battery", ticketCreate.Problem)
	log.Println(ticketCreate.ID)
}

func TestTicketByCustomer(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	graphqlClient := graphql.NewClient("http://localhost:8081/", http.DefaultClient)
	curentTime := time.Now()
	log.Println(curentTime)

	resp, err := graph.TicketByCustomer(ctx, graphqlClient, &graph.TicketByCustomerInput{
		CustomerID: "f4209ba4-9e88-4b9b-a43b-7f71c96c43a8",
		FromTime:   time.Now().AddDate(0, 0, -1),
		ToTime:     time.Now(),
	})
	if err != nil {
		log.Fatal(err)
	}
	ticket := resp.TicketByCustomer[0].ID
	log.Println(ticket)
}
