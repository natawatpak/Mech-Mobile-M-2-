package test

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

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

func testSetupLocal() (context.Context, *resource.SQLop) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	if cancel != nil {
		fmt.Printf("Context cancel msg : %v\n\n", cancel)
	}

	viper.SetConfigName("../postgresConfig")
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

func TestOpCreateFunc(t *testing.T) {
	ctx, operator := testSetupLocal()

	returnCustomer, err := operator.CustomerCreate(ctx, &inputCustomer)
	currentExpectCustomer.ID = returnCustomer.ID
	assert.Equal(t, &currentExpectCustomer, returnCustomer)
	assert.Equal(t, nil, err)

}
