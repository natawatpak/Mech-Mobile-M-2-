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

var currentExpectUser model.User = model.User{
	ID:          "",
	FName:       "phum kitiphum",
	Email:       "phum@gmail.com",
	Gender:      "m",
	DoB:         "05-02-2001",
	CountryCode: 66,
	CreateTime:  "05-02-2021",
}

var inputUser = model.UserCreateInput{
	FName:       currentExpectUser.FName,
	Email:       currentExpectUser.Email,
	Gender:      currentExpectUser.Gender,
	DoB:         currentExpectUser.DoB,
	CountryCode: currentExpectUser.CountryCode,
	CreateTime:  currentExpectUser.CreateTime,
}

var countryInput = model.CountryCreateInput{
	Code:      66,
	Name:      "Thailand",
	Continent: "Asia",
}

func testSetup(dbName string) (context.Context, *resource.SQLop) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	if cancel != nil {
		fmt.Printf("Context cancel msg : %v\n\n", cancel)
	}
	operator, err := resource.NewDBOperator("sql.DB")
	resource.PrintIfErrorExist(err)
	return ctx, operator
}

func TestTableCreateFunc(t *testing.T) {
	ctx, operator := testSetup("sql.DB")

	operator.CreateCountryTable(ctx)
	operator.CountryCreate(ctx, &countryInput)
	re, err := operator.CreateUserTable(ctx)
	println(re)
	println(err)
}

func TestOpCreateFunc(t *testing.T) {
	ctx, operator := testSetup("sql.DB")

	returnPokemon, err := operator.UserCreate(ctx, &inputUser)
	currentExpectUser.ID = returnPokemon.ID
	assert.Equal(t, &currentExpectUser, returnPokemon)
	assert.Equal(t, nil, err)

}
