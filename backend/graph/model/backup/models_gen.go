// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

type DeleteIDInput struct {
	ID string `json:"ID"`
}

type ActiveTicket struct {
	ID         string  `json:"ID" bun:"id,pk"`
	CarID      string  `json:"carID" bun:",unique"`
	CustomerID string  `json:"customerID" bun:",notnull"`
	Problem    string  `json:"problem" bun:",notnull"`
	Description *string `json:"description"`
	ShopID     *string `json:"shopID"`
	Status     *string `json:"status"`
	Longitude  float64 `json:"longitude" bun:",notnull"`
	Latitude   float64 `json:"latitude" bun:",notnull"`
}

type ActiveTicketCreateInput struct {
	ID          string  `json:"ID"`
	CarID       string  `json:"carID"`
	CustomerID  string  `json:"customerID"`
	Problem     string  `json:"problem"`
	Description *string `json:"description"`
	ShopID      *string `json:"shopID"`
	Status      *string `json:"status"`
	Longitude   float64 `json:"longitude"`
	Latitude    float64 `json:"latitude"`
}

type ActiveTicketUpdateInput struct {
	ID          string  `json:"ID"`
	CarID       string  `json:"carID"`
	CustomerID  string  `json:"customerID"`
	Problem     string  `json:"problem"`
	Description *string `json:"description"`
	ShopID      *string `json:"shopID"`
	Status      *string `json:"status"`
	Longitude   float64 `json:"longitude"`
	Latitude    float64 `json:"latitude"`
}

type Car struct {
	ID        string  `json:"ID" bun:"id,pk"`
	OwnerID   string  `json:"ownerID" bun:",notnull"`
	PlateNum  string  `json:"plateNum" bun:",notnull"`
	PlateType *string `json:"plateType"`
	IssuedAt  *string `json:"issuedAt"`
	Color     *string `json:"color"`
	Type      *string `json:"type"`
	Brand     string  `json:"brand" bun:",notnull"`
	Build     *string `json:"build"`
}

type CarCreateInput struct {
	ID        *string `json:"ID"`
	OwnerID   string  `json:"ownerID"`
	PlateNum  string  `json:"plateNum"`
	PlateType *string `json:"plateType"`
	IssuedAt  *string `json:"issuedAt"`
	Color     *string `json:"color"`
	Type      *string `json:"type"`
	Brand     string  `json:"brand"`
	Build     *string `json:"build"`
}

type CarUpdateInput struct {
	ID        string  `json:"ID"`
	OwnerID   string  `json:"ownerID"`
	PlateNum  string  `json:"plateNum"`
	PlateType *string `json:"plateType"`
	IssuedAt  *string `json:"issuedAt"`
	Color     *string `json:"color"`
	Type      *string `json:"type"`
	Brand     string  `json:"brand"`
	Build     *string `json:"build"`
}

type Customer struct {
	ID    string `json:"ID" bun:"id,pk"`
	FName string `json:"fName" bun:",notnull"`
	LName string `json:"lName" bun:",notnull"`
	Tel   string `json:"tel" bun:",notnull,unique"`
	Email string `json:"email" bun:",notnull,unique"`
}

type CustomerCreateInput struct {
	ID    *string `json:"ID"`
	FName string  `json:"fName"`
	LName string  `json:"lName"`
	Tel   string  `json:"tel"`
	Email string  `json:"email"`
}

type CustomerUpdateInput struct {
	ID    string `json:"ID"`
	FName string `json:"fName"`
	LName string `json:"lName"`
	Tel   string `json:"tel"`
	Email string `json:"email"`
}

type Service struct {
	ID   string `json:"ID" bun:"id,pk"`
	Name string `json:"name" bun:",notnull"`
}

type ServiceCreateInput struct {
	ID   *string `json:"ID"`
	Name string  `json:"name"`
}

type ServiceUpdateInput struct {
	ID   string `json:"ID"`
	Name string `json:"name"`
}

type Shop struct {
	ID      string `json:"ID" bun:"id,pk"`
	Name    string `json:"name" bun:",notnull"`
	Tel     string `json:"tel" bun:",notnull,unique"`
	Email   string `json:"email" bun:",notnull,unique"`
	Address string `json:"address" bun:",notnull"`
	Longitude float64 `json:"longitude" bun:",notnull"`
	Latitude  float64 `json:"latitude" bun:",notnull"`
}

type ShopConnect struct {
	ShopID       string `json:"shopID" bun:",pk"`
	ConnectionID string `json:"ConnectionID" bun:",pk"`
}

type ShopConnectCreateInput struct {
	ShopID       string `json:"shopID"`
	ConnectionID string `json:"ConnectionID"`
}

type ShopCreateInput struct {
	ID        *string `json:"ID"`
	Name      string  `json:"name"`
	Tel       string  `json:"tel"`
	Email     string  `json:"email"`
	Address   string  `json:"address"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

type ShopService struct {
	ShopID    string `json:"shopID" bun:",pk"`
	ServiceID string `json:"serviceID" bun:",pk"`
}

type ShopServiceCreateInput struct {
	ShopID    string `json:"shopID"`
	ServiceID string `json:"serviceID"`
}

type ShopServiceDeleteInput struct {
	ShopID    string `json:"shopID"`
	ServiceID string `json:"serviceID"`
}

type ShopUpdateInput struct {
	ID        string  `json:"ID"`
	Name      string  `json:"name"`
	Tel       string  `json:"tel"`
	Email     string  `json:"email"`
	Address   string  `json:"address"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

type Ticket struct {
	ID           string     `json:"ID" bun:"id,pk"`
	CustomerID   string     `json:"customerID" bun:",notnull"`
	CarID        string     `json:"carID" bun:",notnull"`
	Problem      string     `json:"problem" bun:",notnull"`
	Description  *string    `json:"description"`
	CreateTime   time.Time  `json:"createTime" bun:",notnull"`
	ShopID       *string    `json:"shopID"`
	AcceptedTime *time.Time `json:"acceptedTime"`
	Status       *string    `json:"status"`
	Longitude    float64    `json:"longitude" bun:",notnull"`
	Latitude     float64    `json:"latitude" bun:",notnull"`
}

type TicketByCustomerInput struct {
	CustomerID string    `json:"customerID"`
	FromTime   time.Time `json:"fromTime"`
	ToTime     time.Time `json:"toTime"`
	Status     *string   `json:"status"`
}

type TicketByShopInput struct {
	ShopID   string    `json:"shopID"`
	FromTime time.Time `json:"fromTime"`
	ToTime   time.Time `json:"toTime"`
	Status   *string   `json:"status"`
}

type TicketConnect struct {
	TicketID             string `json:"ticketID" bun:",pk"`
	CustomerConnectionID string `json:"customerConnectionID" bun:",notnull,unique"`
	ShopConnectionID     string `json:"shopConnectionID" bun:",notnull,unique"`
}

type TicketConnectCreateInput struct {
	TicketID             string `json:"ticketID"`
	CustomerConnectionID string `json:"customerConnectionID"`
	ShopConnectionID     string `json:"shopConnectionID"`
}

type TicketCreateInput struct {
	ID           *string    `json:"ID"`
	CustomerID   string     `json:"customerID"`
	CarID        string     `json:"carID"`
	Problem      string     `json:"problem"`
	Description  *string    `json:"description"`
	CreateTime   time.Time  `json:"createTime"`
	ShopID       *string    `json:"shopID"`
	AcceptedTime *time.Time `json:"acceptedTime"`
	Status       *string    `json:"status"`
	Longitude    float64    `json:"longitude"`
	Latitude     float64    `json:"latitude"`
}

type TicketService struct {
	TicketID  string `json:"ticketID" bun:",pk"`
	ServiceID string `json:"serviceID" bun:",pk"`
}

type TicketServiceCreateInput struct {
	TicketID  string `json:"ticketID"`
	ServiceID string `json:"serviceID"`
}

type TicketUpdateInput struct {
	ID           string     `json:"ID"`
	CustomerID   string     `json:"customerID"`
	CarID        string     `json:"carID"`
	Problem      string     `json:"problem"`
	Description  *string    `json:"description"`
	CreateTime   time.Time  `json:"createTime"`
	ShopID       *string    `json:"shopID"`
	AcceptedTime *time.Time `json:"acceptedTime"`
	Status       *string    `json:"status"`
	Longitude    float64    `json:"longitude"`
	Latitude     float64    `json:"latitude"`
}
