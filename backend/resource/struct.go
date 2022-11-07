package resource

import (
	"context"

	"github.com/natawatpak/Mech-Mobile-M-2-/backend/graph/model"
	"github.com/uptrace/bun"
)

type SQLop struct {
	db           *bun.DB
	cusModel     *model.Customer
	carModel     *model.Car
	ticketModel  *model.Ticket
	shopModel    *model.Shop
	serviceModel *model.Service

	shopService   *model.ShopService
	activeTicket  *model.ActiveTicket
	ticketService *model.TicketService
}

type DatabaseOp interface {
	CustomerCreate(ctx context.Context, userInput *model.CustomerCreateInput) (*model.Customer, error)

	CustomerList(ctx context.Context) ([]*model.Customer, error)
}
