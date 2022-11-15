// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

package graph

import (
	"context"
	"encoding/json"

	"github.com/Khan/genqlient/graphql"
)

type CarCreateInput struct {
	ID        *string `json:"ID"`
	OwnerID   string  `json:"ownerID"`
	PlateNum  string  `json:"plateNum"`
	PlateType string  `json:"plateType"`
	IssuedAt  string  `json:"issuedAt"`
	Color     string  `json:"color"`
	Type      string  `json:"type"`
	Brand     string  `json:"brand"`
	Build     *string `json:"build"`
}

// GetID returns CarCreateInput.ID, and is useful for accessing the field via an interface.
func (v *CarCreateInput) GetID() *string { return v.ID }

// GetOwnerID returns CarCreateInput.OwnerID, and is useful for accessing the field via an interface.
func (v *CarCreateInput) GetOwnerID() string { return v.OwnerID }

// GetPlateNum returns CarCreateInput.PlateNum, and is useful for accessing the field via an interface.
func (v *CarCreateInput) GetPlateNum() string { return v.PlateNum }

// GetPlateType returns CarCreateInput.PlateType, and is useful for accessing the field via an interface.
func (v *CarCreateInput) GetPlateType() string { return v.PlateType }

// GetIssuedAt returns CarCreateInput.IssuedAt, and is useful for accessing the field via an interface.
func (v *CarCreateInput) GetIssuedAt() string { return v.IssuedAt }

// GetColor returns CarCreateInput.Color, and is useful for accessing the field via an interface.
func (v *CarCreateInput) GetColor() string { return v.Color }

// GetType returns CarCreateInput.Type, and is useful for accessing the field via an interface.
func (v *CarCreateInput) GetType() string { return v.Type }

// GetBrand returns CarCreateInput.Brand, and is useful for accessing the field via an interface.
func (v *CarCreateInput) GetBrand() string { return v.Brand }

// GetBuild returns CarCreateInput.Build, and is useful for accessing the field via an interface.
func (v *CarCreateInput) GetBuild() *string { return v.Build }

// CustomerByIDCustomerByIDCustomer includes the requested fields of the GraphQL type customer.
type CustomerByIDCustomerByIDCustomer struct {
	cusFragment `json:"-"`
}

// GetID returns CustomerByIDCustomerByIDCustomer.ID, and is useful for accessing the field via an interface.
func (v *CustomerByIDCustomerByIDCustomer) GetID() string { return v.cusFragment.ID }

// GetFName returns CustomerByIDCustomerByIDCustomer.FName, and is useful for accessing the field via an interface.
func (v *CustomerByIDCustomerByIDCustomer) GetFName() string { return v.cusFragment.FName }

// GetLName returns CustomerByIDCustomerByIDCustomer.LName, and is useful for accessing the field via an interface.
func (v *CustomerByIDCustomerByIDCustomer) GetLName() string { return v.cusFragment.LName }

// GetTel returns CustomerByIDCustomerByIDCustomer.Tel, and is useful for accessing the field via an interface.
func (v *CustomerByIDCustomerByIDCustomer) GetTel() string { return v.cusFragment.Tel }

// GetEmail returns CustomerByIDCustomerByIDCustomer.Email, and is useful for accessing the field via an interface.
func (v *CustomerByIDCustomerByIDCustomer) GetEmail() string { return v.cusFragment.Email }

func (v *CustomerByIDCustomerByIDCustomer) UnmarshalJSON(b []byte) error {

	if string(b) == "null" {
		return nil
	}

	var firstPass struct {
		*CustomerByIDCustomerByIDCustomer
		graphql.NoUnmarshalJSON
	}
	firstPass.CustomerByIDCustomerByIDCustomer = v

	err := json.Unmarshal(b, &firstPass)
	if err != nil {
		return err
	}

	err = json.Unmarshal(
		b, &v.cusFragment)
	if err != nil {
		return err
	}
	return nil
}

type __premarshalCustomerByIDCustomerByIDCustomer struct {
	ID string `json:"ID"`

	FName string `json:"fName"`

	LName string `json:"lName"`

	Tel string `json:"tel"`

	Email string `json:"email"`
}

func (v *CustomerByIDCustomerByIDCustomer) MarshalJSON() ([]byte, error) {
	premarshaled, err := v.__premarshalJSON()
	if err != nil {
		return nil, err
	}
	return json.Marshal(premarshaled)
}

func (v *CustomerByIDCustomerByIDCustomer) __premarshalJSON() (*__premarshalCustomerByIDCustomerByIDCustomer, error) {
	var retval __premarshalCustomerByIDCustomerByIDCustomer

	retval.ID = v.cusFragment.ID
	retval.FName = v.cusFragment.FName
	retval.LName = v.cusFragment.LName
	retval.Tel = v.cusFragment.Tel
	retval.Email = v.cusFragment.Email
	return &retval, nil
}

// CustomerByIDResponse is returned by CustomerByID on success.
type CustomerByIDResponse struct {
	CustomerByID *CustomerByIDCustomerByIDCustomer `json:"customerByID"`
}

// GetCustomerByID returns CustomerByIDResponse.CustomerByID, and is useful for accessing the field via an interface.
func (v *CustomerByIDResponse) GetCustomerByID() *CustomerByIDCustomerByIDCustomer {
	return v.CustomerByID
}

// CustomerCreateCustomerCreateCustomer includes the requested fields of the GraphQL type customer.
type CustomerCreateCustomerCreateCustomer struct {
	cusFragment `json:"-"`
}

// GetID returns CustomerCreateCustomerCreateCustomer.ID, and is useful for accessing the field via an interface.
func (v *CustomerCreateCustomerCreateCustomer) GetID() string { return v.cusFragment.ID }

// GetFName returns CustomerCreateCustomerCreateCustomer.FName, and is useful for accessing the field via an interface.
func (v *CustomerCreateCustomerCreateCustomer) GetFName() string { return v.cusFragment.FName }

// GetLName returns CustomerCreateCustomerCreateCustomer.LName, and is useful for accessing the field via an interface.
func (v *CustomerCreateCustomerCreateCustomer) GetLName() string { return v.cusFragment.LName }

// GetTel returns CustomerCreateCustomerCreateCustomer.Tel, and is useful for accessing the field via an interface.
func (v *CustomerCreateCustomerCreateCustomer) GetTel() string { return v.cusFragment.Tel }

// GetEmail returns CustomerCreateCustomerCreateCustomer.Email, and is useful for accessing the field via an interface.
func (v *CustomerCreateCustomerCreateCustomer) GetEmail() string { return v.cusFragment.Email }

func (v *CustomerCreateCustomerCreateCustomer) UnmarshalJSON(b []byte) error {

	if string(b) == "null" {
		return nil
	}

	var firstPass struct {
		*CustomerCreateCustomerCreateCustomer
		graphql.NoUnmarshalJSON
	}
	firstPass.CustomerCreateCustomerCreateCustomer = v

	err := json.Unmarshal(b, &firstPass)
	if err != nil {
		return err
	}

	err = json.Unmarshal(
		b, &v.cusFragment)
	if err != nil {
		return err
	}
	return nil
}

type __premarshalCustomerCreateCustomerCreateCustomer struct {
	ID string `json:"ID"`

	FName string `json:"fName"`

	LName string `json:"lName"`

	Tel string `json:"tel"`

	Email string `json:"email"`
}

func (v *CustomerCreateCustomerCreateCustomer) MarshalJSON() ([]byte, error) {
	premarshaled, err := v.__premarshalJSON()
	if err != nil {
		return nil, err
	}
	return json.Marshal(premarshaled)
}

func (v *CustomerCreateCustomerCreateCustomer) __premarshalJSON() (*__premarshalCustomerCreateCustomerCreateCustomer, error) {
	var retval __premarshalCustomerCreateCustomerCreateCustomer

	retval.ID = v.cusFragment.ID
	retval.FName = v.cusFragment.FName
	retval.LName = v.cusFragment.LName
	retval.Tel = v.cusFragment.Tel
	retval.Email = v.cusFragment.Email
	return &retval, nil
}

type CustomerCreateInput struct {
	ID    *string `json:"ID"`
	FName string  `json:"fName"`
	LName string  `json:"lName"`
	Tel   string  `json:"tel"`
	Email string  `json:"email"`
}

// GetID returns CustomerCreateInput.ID, and is useful for accessing the field via an interface.
func (v *CustomerCreateInput) GetID() *string { return v.ID }

// GetFName returns CustomerCreateInput.FName, and is useful for accessing the field via an interface.
func (v *CustomerCreateInput) GetFName() string { return v.FName }

// GetLName returns CustomerCreateInput.LName, and is useful for accessing the field via an interface.
func (v *CustomerCreateInput) GetLName() string { return v.LName }

// GetTel returns CustomerCreateInput.Tel, and is useful for accessing the field via an interface.
func (v *CustomerCreateInput) GetTel() string { return v.Tel }

// GetEmail returns CustomerCreateInput.Email, and is useful for accessing the field via an interface.
func (v *CustomerCreateInput) GetEmail() string { return v.Email }

// CustomerCreateResponse is returned by CustomerCreate on success.
type CustomerCreateResponse struct {
	CustomerCreate *CustomerCreateCustomerCreateCustomer `json:"customerCreate"`
}

// GetCustomerCreate returns CustomerCreateResponse.CustomerCreate, and is useful for accessing the field via an interface.
func (v *CustomerCreateResponse) GetCustomerCreate() *CustomerCreateCustomerCreateCustomer {
	return v.CustomerCreate
}

// CustomerDeleteAllCustomerDeleteAllCustomer includes the requested fields of the GraphQL type customer.
type CustomerDeleteAllCustomerDeleteAllCustomer struct {
	cusFragment `json:"-"`
}

// GetID returns CustomerDeleteAllCustomerDeleteAllCustomer.ID, and is useful for accessing the field via an interface.
func (v *CustomerDeleteAllCustomerDeleteAllCustomer) GetID() string { return v.cusFragment.ID }

// GetFName returns CustomerDeleteAllCustomerDeleteAllCustomer.FName, and is useful for accessing the field via an interface.
func (v *CustomerDeleteAllCustomerDeleteAllCustomer) GetFName() string { return v.cusFragment.FName }

// GetLName returns CustomerDeleteAllCustomerDeleteAllCustomer.LName, and is useful for accessing the field via an interface.
func (v *CustomerDeleteAllCustomerDeleteAllCustomer) GetLName() string { return v.cusFragment.LName }

// GetTel returns CustomerDeleteAllCustomerDeleteAllCustomer.Tel, and is useful for accessing the field via an interface.
func (v *CustomerDeleteAllCustomerDeleteAllCustomer) GetTel() string { return v.cusFragment.Tel }

// GetEmail returns CustomerDeleteAllCustomerDeleteAllCustomer.Email, and is useful for accessing the field via an interface.
func (v *CustomerDeleteAllCustomerDeleteAllCustomer) GetEmail() string { return v.cusFragment.Email }

func (v *CustomerDeleteAllCustomerDeleteAllCustomer) UnmarshalJSON(b []byte) error {

	if string(b) == "null" {
		return nil
	}

	var firstPass struct {
		*CustomerDeleteAllCustomerDeleteAllCustomer
		graphql.NoUnmarshalJSON
	}
	firstPass.CustomerDeleteAllCustomerDeleteAllCustomer = v

	err := json.Unmarshal(b, &firstPass)
	if err != nil {
		return err
	}

	err = json.Unmarshal(
		b, &v.cusFragment)
	if err != nil {
		return err
	}
	return nil
}

type __premarshalCustomerDeleteAllCustomerDeleteAllCustomer struct {
	ID string `json:"ID"`

	FName string `json:"fName"`

	LName string `json:"lName"`

	Tel string `json:"tel"`

	Email string `json:"email"`
}

func (v *CustomerDeleteAllCustomerDeleteAllCustomer) MarshalJSON() ([]byte, error) {
	premarshaled, err := v.__premarshalJSON()
	if err != nil {
		return nil, err
	}
	return json.Marshal(premarshaled)
}

func (v *CustomerDeleteAllCustomerDeleteAllCustomer) __premarshalJSON() (*__premarshalCustomerDeleteAllCustomerDeleteAllCustomer, error) {
	var retval __premarshalCustomerDeleteAllCustomerDeleteAllCustomer

	retval.ID = v.cusFragment.ID
	retval.FName = v.cusFragment.FName
	retval.LName = v.cusFragment.LName
	retval.Tel = v.cusFragment.Tel
	retval.Email = v.cusFragment.Email
	return &retval, nil
}

// CustomerDeleteAllResponse is returned by CustomerDeleteAll on success.
type CustomerDeleteAllResponse struct {
	CustomerDeleteAll []*CustomerDeleteAllCustomerDeleteAllCustomer `json:"customerDeleteAll"`
}

// GetCustomerDeleteAll returns CustomerDeleteAllResponse.CustomerDeleteAll, and is useful for accessing the field via an interface.
func (v *CustomerDeleteAllResponse) GetCustomerDeleteAll() []*CustomerDeleteAllCustomerDeleteAllCustomer {
	return v.CustomerDeleteAll
}

// CustomerDeleteCustomerDeleteCustomer includes the requested fields of the GraphQL type customer.
type CustomerDeleteCustomerDeleteCustomer struct {
	cusFragment `json:"-"`
}

// GetID returns CustomerDeleteCustomerDeleteCustomer.ID, and is useful for accessing the field via an interface.
func (v *CustomerDeleteCustomerDeleteCustomer) GetID() string { return v.cusFragment.ID }

// GetFName returns CustomerDeleteCustomerDeleteCustomer.FName, and is useful for accessing the field via an interface.
func (v *CustomerDeleteCustomerDeleteCustomer) GetFName() string { return v.cusFragment.FName }

// GetLName returns CustomerDeleteCustomerDeleteCustomer.LName, and is useful for accessing the field via an interface.
func (v *CustomerDeleteCustomerDeleteCustomer) GetLName() string { return v.cusFragment.LName }

// GetTel returns CustomerDeleteCustomerDeleteCustomer.Tel, and is useful for accessing the field via an interface.
func (v *CustomerDeleteCustomerDeleteCustomer) GetTel() string { return v.cusFragment.Tel }

// GetEmail returns CustomerDeleteCustomerDeleteCustomer.Email, and is useful for accessing the field via an interface.
func (v *CustomerDeleteCustomerDeleteCustomer) GetEmail() string { return v.cusFragment.Email }

func (v *CustomerDeleteCustomerDeleteCustomer) UnmarshalJSON(b []byte) error {

	if string(b) == "null" {
		return nil
	}

	var firstPass struct {
		*CustomerDeleteCustomerDeleteCustomer
		graphql.NoUnmarshalJSON
	}
	firstPass.CustomerDeleteCustomerDeleteCustomer = v

	err := json.Unmarshal(b, &firstPass)
	if err != nil {
		return err
	}

	err = json.Unmarshal(
		b, &v.cusFragment)
	if err != nil {
		return err
	}
	return nil
}

type __premarshalCustomerDeleteCustomerDeleteCustomer struct {
	ID string `json:"ID"`

	FName string `json:"fName"`

	LName string `json:"lName"`

	Tel string `json:"tel"`

	Email string `json:"email"`
}

func (v *CustomerDeleteCustomerDeleteCustomer) MarshalJSON() ([]byte, error) {
	premarshaled, err := v.__premarshalJSON()
	if err != nil {
		return nil, err
	}
	return json.Marshal(premarshaled)
}

func (v *CustomerDeleteCustomerDeleteCustomer) __premarshalJSON() (*__premarshalCustomerDeleteCustomerDeleteCustomer, error) {
	var retval __premarshalCustomerDeleteCustomerDeleteCustomer

	retval.ID = v.cusFragment.ID
	retval.FName = v.cusFragment.FName
	retval.LName = v.cusFragment.LName
	retval.Tel = v.cusFragment.Tel
	retval.Email = v.cusFragment.Email
	return &retval, nil
}

// CustomerDeleteResponse is returned by CustomerDelete on success.
type CustomerDeleteResponse struct {
	CustomerDelete *CustomerDeleteCustomerDeleteCustomer `json:"customerDelete"`
}

// GetCustomerDelete returns CustomerDeleteResponse.CustomerDelete, and is useful for accessing the field via an interface.
func (v *CustomerDeleteResponse) GetCustomerDelete() *CustomerDeleteCustomerDeleteCustomer {
	return v.CustomerDelete
}

type CustomerUpdateInput struct {
	ID    string `json:"ID"`
	FName string `json:"fName"`
	LName string `json:"lName"`
	Tel   string `json:"tel"`
	Email string `json:"email"`
}

// GetID returns CustomerUpdateInput.ID, and is useful for accessing the field via an interface.
func (v *CustomerUpdateInput) GetID() string { return v.ID }

// GetFName returns CustomerUpdateInput.FName, and is useful for accessing the field via an interface.
func (v *CustomerUpdateInput) GetFName() string { return v.FName }

// GetLName returns CustomerUpdateInput.LName, and is useful for accessing the field via an interface.
func (v *CustomerUpdateInput) GetLName() string { return v.LName }

// GetTel returns CustomerUpdateInput.Tel, and is useful for accessing the field via an interface.
func (v *CustomerUpdateInput) GetTel() string { return v.Tel }

// GetEmail returns CustomerUpdateInput.Email, and is useful for accessing the field via an interface.
func (v *CustomerUpdateInput) GetEmail() string { return v.Email }

// CustomerUpdateMultiCustomerUpdateMultiCustomer includes the requested fields of the GraphQL type customer.
type CustomerUpdateMultiCustomerUpdateMultiCustomer struct {
	cusFragment `json:"-"`
}

// GetID returns CustomerUpdateMultiCustomerUpdateMultiCustomer.ID, and is useful for accessing the field via an interface.
func (v *CustomerUpdateMultiCustomerUpdateMultiCustomer) GetID() string { return v.cusFragment.ID }

// GetFName returns CustomerUpdateMultiCustomerUpdateMultiCustomer.FName, and is useful for accessing the field via an interface.
func (v *CustomerUpdateMultiCustomerUpdateMultiCustomer) GetFName() string {
	return v.cusFragment.FName
}

// GetLName returns CustomerUpdateMultiCustomerUpdateMultiCustomer.LName, and is useful for accessing the field via an interface.
func (v *CustomerUpdateMultiCustomerUpdateMultiCustomer) GetLName() string {
	return v.cusFragment.LName
}

// GetTel returns CustomerUpdateMultiCustomerUpdateMultiCustomer.Tel, and is useful for accessing the field via an interface.
func (v *CustomerUpdateMultiCustomerUpdateMultiCustomer) GetTel() string { return v.cusFragment.Tel }

// GetEmail returns CustomerUpdateMultiCustomerUpdateMultiCustomer.Email, and is useful for accessing the field via an interface.
func (v *CustomerUpdateMultiCustomerUpdateMultiCustomer) GetEmail() string {
	return v.cusFragment.Email
}

func (v *CustomerUpdateMultiCustomerUpdateMultiCustomer) UnmarshalJSON(b []byte) error {

	if string(b) == "null" {
		return nil
	}

	var firstPass struct {
		*CustomerUpdateMultiCustomerUpdateMultiCustomer
		graphql.NoUnmarshalJSON
	}
	firstPass.CustomerUpdateMultiCustomerUpdateMultiCustomer = v

	err := json.Unmarshal(b, &firstPass)
	if err != nil {
		return err
	}

	err = json.Unmarshal(
		b, &v.cusFragment)
	if err != nil {
		return err
	}
	return nil
}

type __premarshalCustomerUpdateMultiCustomerUpdateMultiCustomer struct {
	ID string `json:"ID"`

	FName string `json:"fName"`

	LName string `json:"lName"`

	Tel string `json:"tel"`

	Email string `json:"email"`
}

func (v *CustomerUpdateMultiCustomerUpdateMultiCustomer) MarshalJSON() ([]byte, error) {
	premarshaled, err := v.__premarshalJSON()
	if err != nil {
		return nil, err
	}
	return json.Marshal(premarshaled)
}

func (v *CustomerUpdateMultiCustomerUpdateMultiCustomer) __premarshalJSON() (*__premarshalCustomerUpdateMultiCustomerUpdateMultiCustomer, error) {
	var retval __premarshalCustomerUpdateMultiCustomerUpdateMultiCustomer

	retval.ID = v.cusFragment.ID
	retval.FName = v.cusFragment.FName
	retval.LName = v.cusFragment.LName
	retval.Tel = v.cusFragment.Tel
	retval.Email = v.cusFragment.Email
	return &retval, nil
}

// CustomerUpdateMultiResponse is returned by CustomerUpdateMulti on success.
type CustomerUpdateMultiResponse struct {
	CustomerUpdateMulti *CustomerUpdateMultiCustomerUpdateMultiCustomer `json:"customerUpdateMulti"`
}

// GetCustomerUpdateMulti returns CustomerUpdateMultiResponse.CustomerUpdateMulti, and is useful for accessing the field via an interface.
func (v *CustomerUpdateMultiResponse) GetCustomerUpdateMulti() *CustomerUpdateMultiCustomerUpdateMultiCustomer {
	return v.CustomerUpdateMulti
}

// CustomersCustomersCustomer includes the requested fields of the GraphQL type customer.
type CustomersCustomersCustomer struct {
	cusFragment `json:"-"`
}

// GetID returns CustomersCustomersCustomer.ID, and is useful for accessing the field via an interface.
func (v *CustomersCustomersCustomer) GetID() string { return v.cusFragment.ID }

// GetFName returns CustomersCustomersCustomer.FName, and is useful for accessing the field via an interface.
func (v *CustomersCustomersCustomer) GetFName() string { return v.cusFragment.FName }

// GetLName returns CustomersCustomersCustomer.LName, and is useful for accessing the field via an interface.
func (v *CustomersCustomersCustomer) GetLName() string { return v.cusFragment.LName }

// GetTel returns CustomersCustomersCustomer.Tel, and is useful for accessing the field via an interface.
func (v *CustomersCustomersCustomer) GetTel() string { return v.cusFragment.Tel }

// GetEmail returns CustomersCustomersCustomer.Email, and is useful for accessing the field via an interface.
func (v *CustomersCustomersCustomer) GetEmail() string { return v.cusFragment.Email }

func (v *CustomersCustomersCustomer) UnmarshalJSON(b []byte) error {

	if string(b) == "null" {
		return nil
	}

	var firstPass struct {
		*CustomersCustomersCustomer
		graphql.NoUnmarshalJSON
	}
	firstPass.CustomersCustomersCustomer = v

	err := json.Unmarshal(b, &firstPass)
	if err != nil {
		return err
	}

	err = json.Unmarshal(
		b, &v.cusFragment)
	if err != nil {
		return err
	}
	return nil
}

type __premarshalCustomersCustomersCustomer struct {
	ID string `json:"ID"`

	FName string `json:"fName"`

	LName string `json:"lName"`

	Tel string `json:"tel"`

	Email string `json:"email"`
}

func (v *CustomersCustomersCustomer) MarshalJSON() ([]byte, error) {
	premarshaled, err := v.__premarshalJSON()
	if err != nil {
		return nil, err
	}
	return json.Marshal(premarshaled)
}

func (v *CustomersCustomersCustomer) __premarshalJSON() (*__premarshalCustomersCustomersCustomer, error) {
	var retval __premarshalCustomersCustomersCustomer

	retval.ID = v.cusFragment.ID
	retval.FName = v.cusFragment.FName
	retval.LName = v.cusFragment.LName
	retval.Tel = v.cusFragment.Tel
	retval.Email = v.cusFragment.Email
	return &retval, nil
}

// CustomersResponse is returned by Customers on success.
type CustomersResponse struct {
	Customers []*CustomersCustomersCustomer `json:"customers"`
}

// GetCustomers returns CustomersResponse.Customers, and is useful for accessing the field via an interface.
func (v *CustomersResponse) GetCustomers() []*CustomersCustomersCustomer { return v.Customers }

type DeleteIDInput struct {
	ID string `json:"ID"`
}

// GetID returns DeleteIDInput.ID, and is useful for accessing the field via an interface.
func (v *DeleteIDInput) GetID() string { return v.ID }

// __CustomerByIDInput is used internally by genqlient
type __CustomerByIDInput struct {
	CusInput string `json:"cusInput"`
}

// GetCusInput returns __CustomerByIDInput.CusInput, and is useful for accessing the field via an interface.
func (v *__CustomerByIDInput) GetCusInput() string { return v.CusInput }

// __CustomerCreateInput is used internally by genqlient
type __CustomerCreateInput struct {
	CusInput *CustomerCreateInput `json:"cusInput,omitempty"`
}

// GetCusInput returns __CustomerCreateInput.CusInput, and is useful for accessing the field via an interface.
func (v *__CustomerCreateInput) GetCusInput() *CustomerCreateInput { return v.CusInput }

// __CustomerDeleteInput is used internally by genqlient
type __CustomerDeleteInput struct {
	CusInput *DeleteIDInput `json:"cusInput,omitempty"`
}

// GetCusInput returns __CustomerDeleteInput.CusInput, and is useful for accessing the field via an interface.
func (v *__CustomerDeleteInput) GetCusInput() *DeleteIDInput { return v.CusInput }

// __CustomerUpdateMultiInput is used internally by genqlient
type __CustomerUpdateMultiInput struct {
	CusInput *CustomerUpdateInput `json:"cusInput,omitempty"`
}

// GetCusInput returns __CustomerUpdateMultiInput.CusInput, and is useful for accessing the field via an interface.
func (v *__CustomerUpdateMultiInput) GetCusInput() *CustomerUpdateInput { return v.CusInput }

// __carCreateInput is used internally by genqlient
type __carCreateInput struct {
	TagInput *CarCreateInput `json:"tagInput,omitempty"`
}

// GetTagInput returns __carCreateInput.TagInput, and is useful for accessing the field via an interface.
func (v *__carCreateInput) GetTagInput() *CarCreateInput { return v.TagInput }

// carCreateCarCreateCar includes the requested fields of the GraphQL type car.
type carCreateCarCreateCar struct {
	ID        string  `json:"ID"`
	OwnerID   string  `json:"ownerID"`
	PlateNum  string  `json:"plateNum"`
	PlateType string  `json:"plateType"`
	IssuedAt  string  `json:"issuedAt"`
	Color     string  `json:"color"`
	Type      string  `json:"type"`
	Brand     string  `json:"brand"`
	Build     *string `json:"build"`
}

// GetID returns carCreateCarCreateCar.ID, and is useful for accessing the field via an interface.
func (v *carCreateCarCreateCar) GetID() string { return v.ID }

// GetOwnerID returns carCreateCarCreateCar.OwnerID, and is useful for accessing the field via an interface.
func (v *carCreateCarCreateCar) GetOwnerID() string { return v.OwnerID }

// GetPlateNum returns carCreateCarCreateCar.PlateNum, and is useful for accessing the field via an interface.
func (v *carCreateCarCreateCar) GetPlateNum() string { return v.PlateNum }

// GetPlateType returns carCreateCarCreateCar.PlateType, and is useful for accessing the field via an interface.
func (v *carCreateCarCreateCar) GetPlateType() string { return v.PlateType }

// GetIssuedAt returns carCreateCarCreateCar.IssuedAt, and is useful for accessing the field via an interface.
func (v *carCreateCarCreateCar) GetIssuedAt() string { return v.IssuedAt }

// GetColor returns carCreateCarCreateCar.Color, and is useful for accessing the field via an interface.
func (v *carCreateCarCreateCar) GetColor() string { return v.Color }

// GetType returns carCreateCarCreateCar.Type, and is useful for accessing the field via an interface.
func (v *carCreateCarCreateCar) GetType() string { return v.Type }

// GetBrand returns carCreateCarCreateCar.Brand, and is useful for accessing the field via an interface.
func (v *carCreateCarCreateCar) GetBrand() string { return v.Brand }

// GetBuild returns carCreateCarCreateCar.Build, and is useful for accessing the field via an interface.
func (v *carCreateCarCreateCar) GetBuild() *string { return v.Build }

// carCreateResponse is returned by carCreate on success.
type carCreateResponse struct {
	CarCreate *carCreateCarCreateCar `json:"carCreate"`
}

// GetCarCreate returns carCreateResponse.CarCreate, and is useful for accessing the field via an interface.
func (v *carCreateResponse) GetCarCreate() *carCreateCarCreateCar { return v.CarCreate }

// cusFragment includes the GraphQL fields of customer requested by the fragment cusFragment.
type cusFragment struct {
	ID    string `json:"ID"`
	FName string `json:"fName"`
	LName string `json:"lName"`
	Tel   string `json:"tel"`
	Email string `json:"email"`
}

// GetID returns cusFragment.ID, and is useful for accessing the field via an interface.
func (v *cusFragment) GetID() string { return v.ID }

// GetFName returns cusFragment.FName, and is useful for accessing the field via an interface.
func (v *cusFragment) GetFName() string { return v.FName }

// GetLName returns cusFragment.LName, and is useful for accessing the field via an interface.
func (v *cusFragment) GetLName() string { return v.LName }

// GetTel returns cusFragment.Tel, and is useful for accessing the field via an interface.
func (v *cusFragment) GetTel() string { return v.Tel }

// GetEmail returns cusFragment.Email, and is useful for accessing the field via an interface.
func (v *cusFragment) GetEmail() string { return v.Email }

func CustomerByID(
	ctx context.Context,
	client graphql.Client,
	cusInput string,
) (*CustomerByIDResponse, error) {
	req := &graphql.Request{
		OpName: "CustomerByID",
		Query: `
query CustomerByID ($cusInput: ID!) {
	customerByID(input: $cusInput) {
		... cusFragment
	}
}
fragment cusFragment on customer {
	ID
	fName
	lName
	tel
	email
}
`,
		Variables: &__CustomerByIDInput{
			CusInput: cusInput,
		},
	}
	var err error

	var data CustomerByIDResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}

func CustomerCreate(
	ctx context.Context,
	client graphql.Client,
	cusInput *CustomerCreateInput,
) (*CustomerCreateResponse, error) {
	req := &graphql.Request{
		OpName: "CustomerCreate",
		Query: `
mutation CustomerCreate ($cusInput: customerCreateInput!) {
	customerCreate(input: $cusInput) {
		... cusFragment
	}
}
fragment cusFragment on customer {
	ID
	fName
	lName
	tel
	email
}
`,
		Variables: &__CustomerCreateInput{
			CusInput: cusInput,
		},
	}
	var err error

	var data CustomerCreateResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}

func CustomerDelete(
	ctx context.Context,
	client graphql.Client,
	cusInput *DeleteIDInput,
) (*CustomerDeleteResponse, error) {
	req := &graphql.Request{
		OpName: "CustomerDelete",
		Query: `
mutation CustomerDelete ($cusInput: DeleteIDInput!) {
	customerDelete(input: $cusInput) {
		... cusFragment
	}
}
fragment cusFragment on customer {
	ID
	fName
	lName
	tel
	email
}
`,
		Variables: &__CustomerDeleteInput{
			CusInput: cusInput,
		},
	}
	var err error

	var data CustomerDeleteResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}

func CustomerDeleteAll(
	ctx context.Context,
	client graphql.Client,
) (*CustomerDeleteAllResponse, error) {
	req := &graphql.Request{
		OpName: "CustomerDeleteAll",
		Query: `
mutation CustomerDeleteAll {
	customerDeleteAll {
		... cusFragment
	}
}
fragment cusFragment on customer {
	ID
	fName
	lName
	tel
	email
}
`,
	}
	var err error

	var data CustomerDeleteAllResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}

func CustomerUpdateMulti(
	ctx context.Context,
	client graphql.Client,
	cusInput *CustomerUpdateInput,
) (*CustomerUpdateMultiResponse, error) {
	req := &graphql.Request{
		OpName: "CustomerUpdateMulti",
		Query: `
mutation CustomerUpdateMulti ($cusInput: customerUpdateInput!) {
	customerUpdateMulti(input: $cusInput) {
		... cusFragment
	}
}
fragment cusFragment on customer {
	ID
	fName
	lName
	tel
	email
}
`,
		Variables: &__CustomerUpdateMultiInput{
			CusInput: cusInput,
		},
	}
	var err error

	var data CustomerUpdateMultiResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}

func Customers(
	ctx context.Context,
	client graphql.Client,
) (*CustomersResponse, error) {
	req := &graphql.Request{
		OpName: "Customers",
		Query: `
query Customers {
	customers {
		... cusFragment
	}
}
fragment cusFragment on customer {
	ID
	fName
	lName
	tel
	email
}
`,
	}
	var err error

	var data CustomersResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}

func carCreate(
	ctx context.Context,
	client graphql.Client,
	tagInput *CarCreateInput,
) (*carCreateResponse, error) {
	req := &graphql.Request{
		OpName: "carCreate",
		Query: `
mutation carCreate ($tagInput: carCreateInput!) {
	carCreate(input: $tagInput) {
		ID
		ownerID
		plateNum
		plateType
		issuedAt
		color
		type
		brand
		build
	}
}
`,
		Variables: &__carCreateInput{
			TagInput: tagInput,
		},
	}
	var err error

	var data carCreateResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}