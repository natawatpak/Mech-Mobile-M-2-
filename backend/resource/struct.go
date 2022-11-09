package resource

import (
	"context"
	"database/sql"

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
	CreateTables(ctx context.Context) (sql.Result, error)

	CustomerCreate(ctx context.Context, userInput *model.CustomerCreateInput) (*model.Customer, error)

	CustomerList(ctx context.Context) ([]*model.Customer, error)
}
