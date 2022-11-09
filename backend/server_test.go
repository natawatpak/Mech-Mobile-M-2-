package main

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/natawatpak/Mech-Mobile-M-2-/backend/graph/model"
	"github.com/natawatpak/Mech-Mobile-M-2-/backend/resource"
	"github.com/stretchr/testify/assert"
)

var currentExpectCustomer model.Customer = model.Customer{
	ID:    "",
	FName: "phum kitiphum",
	Email: "phum@gmail.com",
}

var inputCustomer = model.CustomerCreateInput{
	FName: currentExpectCustomer.FName,
	Email: currentExpectCustomer.Email,
}

func testSetup(dbName string) (context.Context, *resource.SQLop) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	if cancel != nil {
		fmt.Printf("Context cancel msg : %v\n\n", cancel)
	}
	operator, err := resource.NewDBOperator("sql.DB")
	resource.PrintIfErrorExist(err)
	return ctx, operator
}

func TestTableCreateFunc(t *testing.T) {
	ctx, operator := testSetup("sql.DB")

	createResult, err := operator.CreateTables(ctx)
	println(createResult)
	println(err)
}

func TestOpCreateFunc(t *testing.T) {
	ctx, operator := testSetup("sql.DB")

	returnCustomer, err := operator.CustomerCreate(ctx, &inputCustomer)
	currentExpectCustomer.ID = returnCustomer.ID
	assert.Equal(t, &currentExpectCustomer, returnCustomer)
	assert.Equal(t, nil, err)

}
