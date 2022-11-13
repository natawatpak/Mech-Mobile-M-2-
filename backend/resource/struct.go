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
	DropTable(ctx context.Context) (bool, error)
	CreateTables(ctx context.Context) (bool, error)

	CustomerCreate(ctx context.Context, userInput *model.CustomerCreateInput) (*model.Customer, error)
	CustomerUpdateMulti(ctx context.Context, updateInput model.Customer) (*model.Customer, error)
	CustomerDelete(ctx context.Context, ID string) (*model.Customer, error)
	CustomerDeleteAll(ctx context.Context) ([]*model.Customer, error)
	CustomerFindByID(ctx context.Context, ID string) (*model.Customer, error)
	CustomerList(ctx context.Context) ([]*model.Customer, error)

	CarCreate(ctx context.Context, carInput *model.CarCreateInput) (*model.Car, error)
	CarUpdateMulti(ctx context.Context, updateInput model.Car) (*model.Car, error)
	CarDelete(ctx context.Context, ID string) (*model.Car, error)
	CarDeleteAll(ctx context.Context) ([]*model.Car, error)
	CarFindByID(ctx context.Context, ID string) (*model.Car, error)
	CarFindByOwner(ctx context.Context, ID string) ([]*model.Car, error)
	CarList(ctx context.Context) ([]*model.Car, error)

	TicketCreate(ctx context.Context, ticketInput *model.TicketCreateInput) (*model.Ticket, error)
	TicketUpdateMulti(ctx context.Context, updateInput model.Ticket) (*model.Ticket, error)
	TicketDelete(ctx context.Context, ID string) (*model.Ticket, error)
	TicketDeleteAll(ctx context.Context) ([]*model.Ticket, error)
	TicketFindByID(ctx context.Context, ID string) (*model.Ticket, error)
	TicketFindByCustomer(ctx context.Context, input model.TicketByCustomerInput) ([]*model.Ticket, error)
	TicketFindByShop(ctx context.Context, input model.TicketByShopInput) ([]*model.Ticket, error)
	TicketList(ctx context.Context) ([]*model.Ticket, error)
}
