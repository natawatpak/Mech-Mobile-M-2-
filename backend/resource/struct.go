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
	CreateConnectTables(ctx context.Context) (bool, error)

	CustomerCreate(ctx context.Context, userInput *model.CustomerCreateInput) (*model.Customer, error)
	CustomerUpdateMulti(ctx context.Context, updateInput model.Customer) (*model.Customer, error)
	CustomerDelete(ctx context.Context, ID string) (*model.Customer, error)
	CustomerDeleteAll(ctx context.Context) ([]*model.Customer, error)
	CustomerFindByID(ctx context.Context, ID string) (*model.Customer, error)
	CustomerFindByEmail(ctx context.Context, Email string) (*model.Customer, error)
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

	ActiveTicketCreate(ctx context.Context, activeTicketInput *model.ActiveTicketCreateInput) (*model.ActiveTicket, error)
	ActiveTicketUpdateMulti(ctx context.Context, updateInput model.ActiveTicket) (*model.ActiveTicket, error)
	ActiveTicketDelete(ctx context.Context, ID string) (*model.ActiveTicket, error)
	ActiveTicketDeleteByStatus(ctx context.Context, input string) ([]*model.ActiveTicket, error)
	ActiveTicketDeleteAll(ctx context.Context) ([]*model.ActiveTicket, error)
	ActiveTicketFindByID(ctx context.Context, ID string) (*model.ActiveTicket, error)
	ActiveTicketFindByStatus(ctx context.Context, input string) ([]*model.ActiveTicket, error)
	ActiveTicketFindByCustomer(ctx context.Context, ID string) ([]*model.ActiveTicket, error)
	ActiveTicketFindByShop(ctx context.Context, ID string) ([]*model.ActiveTicket, error)
	ActiveTicketList(ctx context.Context) ([]*model.ActiveTicket, error)

	ShopCreate(ctx context.Context, shopInput *model.ShopCreateInput) (*model.Shop, error)
	ShopUpdateMulti(ctx context.Context, updateInput model.Shop) (*model.Shop, error)
	ShopDelete(ctx context.Context, ID string) (*model.Shop, error)
	ShopDeleteAll(ctx context.Context) ([]*model.Shop, error)
	ShopFindByID(ctx context.Context, ID string) (*model.Shop, error)
	ShopList(ctx context.Context) ([]*model.Shop, error)

	ServiceCreate(ctx context.Context, serviceInput *model.ServiceCreateInput) (*model.Service, error)
	ServiceUpdateMulti(ctx context.Context, updateInput model.ServiceUpdateInput) (*model.Service, error)
	ServiceDelete(ctx context.Context, ID string) (*model.Service, error)
	ServiceDeleteAll(ctx context.Context) ([]*model.Service, error)
	ServiceFindByID(ctx context.Context, ID string) (*model.Service, error)
	ServiceList(ctx context.Context) ([]*model.Service, error)

	ShopServiceCreate(ctx context.Context, shopServiceInput *model.ShopServiceCreateInput) (*model.ShopService, error)
	ShopServiceDelete(ctx context.Context, input *model.ShopServiceDeleteInput) (*model.ShopService, error)
	ShopServiceDeleteAll(ctx context.Context) ([]*model.ShopService, error)
	ShopServiceFindByID(ctx context.Context, input model.ShopServiceCreateInput) (*model.ShopService, error)
	ShopFindByEmail(ctx context.Context, Email string) (*model.Shop, error)
	ShopServiceList(ctx context.Context) ([]*model.ShopService, error)

	TicketServiceCreate(ctx context.Context, TicketServiceInput *model.TicketServiceCreateInput) (*model.TicketService, error)
	TicketServiceDelete(ctx context.Context, input *model.TicketServiceCreateInput) (*model.TicketService, error)
	TicketServiceDeleteAll(ctx context.Context) ([]*model.TicketService, error)
	TicketServiceFindByID(ctx context.Context, input model.TicketServiceCreateInput) (*model.TicketService, error)
	TicketServiceList(ctx context.Context) ([]*model.TicketService, error)

	TicketConnectCreate(ctx context.Context, input *model.TicketConnectCreateInput) (*model.TicketConnect, error)
	TicketConnectUpdateMulti(ctx context.Context, updateInput model.TicketConnectCreateInput) (*model.TicketConnect, error)
	TicketConnectDelete(ctx context.Context, ID string) (*model.TicketConnect, error)
	TicketConnectDeleteAll(ctx context.Context) ([]*model.TicketConnect, error)
	TicketConnectFindByID(ctx context.Context, ID string) (*model.TicketConnect, error)
	TicketConnectList(ctx context.Context) ([]*model.TicketConnect, error)

	ShopConnectCreate(ctx context.Context, input *model.ShopConnectCreateInput) (*model.ShopConnect, error)
	ShopConnectDelete(ctx context.Context, ID string) ([]*model.ShopConnect, error)
	ShopConnectDeleteAll(ctx context.Context) ([]*model.ShopConnect, error)
	ShopConnectFindByID(ctx context.Context, ID string) ([]*model.ShopConnect, error)
	ShopConnectList(ctx context.Context) ([]*model.ShopConnect, error)
}
