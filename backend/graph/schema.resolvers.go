package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"

	"github.com/natawatpak/Mech-Mobile-M-2-/backend/graph/generated"
	"github.com/natawatpak/Mech-Mobile-M-2-/backend/graph/model"
	"github.com/natawatpak/Mech-Mobile-M-2-/backend/util"
)

// DropTable is the resolver for the dropTable field.
func (r *mutationResolver) DropTable(ctx context.Context) (bool, error) {
	_, err := r.DB.DropTable(ctx)
	if util.CheckErr(err) {
		return false, err
	}
	return true, err
}

// CreateTable is the resolver for the createTable field.
func (r *mutationResolver) CreateTable(ctx context.Context) (bool, error) {
	_, err := r.DB.CreateTables(ctx)
	if util.CheckErr(err) {
		return false, err
	}
	return true, err
}

// CustomerCreate is the resolver for the customerCreate field.
func (r *mutationResolver) CustomerCreate(ctx context.Context, input model.CustomerCreateInput) (*model.Customer, error) {
	if input.ID != nil {
		return nil, errors.New("id must be null")
	}
	returnCustomer, err := r.DB.CustomerCreate(ctx, &input)
	return returnCustomer, err
}

// CustomerUpdateMulti is the resolver for the customerUpdateMulti field.
func (r *mutationResolver) CustomerUpdateMulti(ctx context.Context, input model.CustomerUpdateInput) (*model.Customer, error) {
	result, err := r.DB.CustomerUpdateMulti(ctx, model.Customer(input))
	return result, err
}

// CustomerDelete is the resolver for the customerDelete field.
func (r *mutationResolver) CustomerDelete(ctx context.Context, input model.DeleteIDInput) (*model.Customer, error) {
	result, err := r.DB.CustomerDelete(ctx, input.ID)
	return result, err
}

// CustomerDeleteAll is the resolver for the customerDeleteAll field.
func (r *mutationResolver) CustomerDeleteAll(ctx context.Context) ([]*model.Customer, error) {
	result, err := r.DB.CustomerDeleteAll(ctx)
	return result, err
}

// CarCreate is the resolver for the carCreate field.
func (r *mutationResolver) CarCreate(ctx context.Context, input model.CarCreateInput) (*model.Car, error) {
	if input.ID != nil {
		return nil, errors.New("id must be null")
	}
	result, err := r.DB.CarCreate(ctx, &input)
	return result, err
}

// CarUpdateMulti is the resolver for the carUpdateMulti field.
func (r *mutationResolver) CarUpdateMulti(ctx context.Context, input model.CarUpdateInput) (*model.Car, error) {
	result, err := r.DB.CarUpdateMulti(ctx, model.Car(input))
	return result, err
}

// CarDelete is the resolver for the carDelete field.
func (r *mutationResolver) CarDelete(ctx context.Context, input model.DeleteIDInput) (*model.Car, error) {
	result, err := r.DB.CarDelete(ctx, input.ID)
	return result, err
}

// CarDeleteAll is the resolver for the carDeleteAll field.
func (r *mutationResolver) CarDeleteAll(ctx context.Context) ([]*model.Car, error) {
	result, err := r.DB.CarDeleteAll(ctx)
	return result, err
}

// TicketCreate is the resolver for the ticketCreate field.
func (r *mutationResolver) TicketCreate(ctx context.Context, input model.TicketCreateInput) (*model.Ticket, error) {
	if input.ID != nil {
		return nil, errors.New("id must be null")
	}
	returnCustomer, err := r.DB.TicketCreate(ctx, &input)
	return returnCustomer, err
}

// TicketUpdateMulti is the resolver for the ticketUpdateMulti field.
func (r *mutationResolver) TicketUpdateMulti(ctx context.Context, input model.TicketUpdateInput) (*model.Ticket, error) {
	result, err := r.DB.TicketUpdateMulti(ctx, model.Ticket(input))
	return result, err
}

// TicketDelete is the resolver for the ticketDelete field.
func (r *mutationResolver) TicketDelete(ctx context.Context, input model.DeleteIDInput) (*model.Ticket, error) {
	result, err := r.DB.TicketDelete(ctx, input.ID)
	return result, err
}

// TicketDeleteAll is the resolver for the ticketDeleteAll field.
func (r *mutationResolver) TicketDeleteAll(ctx context.Context) ([]*model.Ticket, error) {
	result, err := r.DB.TicketDeleteAll(ctx)
	return result, err
}

// ActiveTicketCreate is the resolver for the activeTicketCreate field.
func (r *mutationResolver) ActiveTicketCreate(ctx context.Context, input model.ActiveTicketCreateInput) (*model.ActiveTicket, error) {
	panic(fmt.Errorf("not implemented: ActiveTicketCreate - activeTicketCreate"))
}

// ActiveTicketUpdateMulti is the resolver for the activeTicketUpdateMulti field.
func (r *mutationResolver) ActiveTicketUpdateMulti(ctx context.Context, input model.ActiveTicketUpdateInput) (*model.ActiveTicket, error) {
	panic(fmt.Errorf("not implemented: ActiveTicketUpdateMulti - activeTicketUpdateMulti"))
}

// ActiveTicketDelete is the resolver for the activeTicketDelete field.
func (r *mutationResolver) ActiveTicketDelete(ctx context.Context, input model.DeleteIDInput) (*model.ActiveTicket, error) {
	panic(fmt.Errorf("not implemented: ActiveTicketDelete - activeTicketDelete"))
}

// ActiveTicketDeleteStatus is the resolver for the activeTicketDeleteStatus field.
func (r *mutationResolver) ActiveTicketDeleteStatus(ctx context.Context, input model.Status) ([]*model.ActiveTicket, error) {
	panic(fmt.Errorf("not implemented: ActiveTicketDeleteStatus - activeTicketDeleteStatus"))
}

// ActiveTicketDeleteAll is the resolver for the activeTicketDeleteAll field.
func (r *mutationResolver) ActiveTicketDeleteAll(ctx context.Context) ([]*model.ActiveTicket, error) {
	panic(fmt.Errorf("not implemented: ActiveTicketDeleteAll - activeTicketDeleteAll"))
}

// ShopCreate is the resolver for the shopCreate field.
func (r *mutationResolver) ShopCreate(ctx context.Context, input model.ShopCreateInput) (*model.Shop, error) {
	panic(fmt.Errorf("not implemented: ShopCreate - shopCreate"))
}

// ShopCUpdateMulti is the resolver for the shopCUpdateMulti field.
func (r *mutationResolver) ShopCUpdateMulti(ctx context.Context, input model.ShopUpdateInput) (*model.Shop, error) {
	panic(fmt.Errorf("not implemented: ShopCUpdateMulti - shopCUpdateMulti"))
}

// ShopCDelete is the resolver for the shopCDelete field.
func (r *mutationResolver) ShopCDelete(ctx context.Context, input model.DeleteIDInput) (*model.Shop, error) {
	panic(fmt.Errorf("not implemented: ShopCDelete - shopCDelete"))
}

// ShopCDeleteAll is the resolver for the shopCDeleteAll field.
func (r *mutationResolver) ShopCDeleteAll(ctx context.Context) ([]*model.Shop, error) {
	panic(fmt.Errorf("not implemented: ShopCDeleteAll - shopCDeleteAll"))
}

// ServiceCreate is the resolver for the serviceCreate field.
func (r *mutationResolver) ServiceCreate(ctx context.Context, input model.ServiceCreateInput) (*model.Service, error) {
	panic(fmt.Errorf("not implemented: ServiceCreate - serviceCreate"))
}

// ServiceDelete is the resolver for the serviceDelete field.
func (r *mutationResolver) ServiceDelete(ctx context.Context, input model.DeleteIDInput) (*model.Service, error) {
	panic(fmt.Errorf("not implemented: ServiceDelete - serviceDelete"))
}

// ServiceDeleteAll is the resolver for the serviceDeleteAll field.
func (r *mutationResolver) ServiceDeleteAll(ctx context.Context) ([]*model.Service, error) {
	panic(fmt.Errorf("not implemented: ServiceDeleteAll - serviceDeleteAll"))
}

// ShopServiceCreate is the resolver for the shopServiceCreate field.
func (r *mutationResolver) ShopServiceCreate(ctx context.Context, input model.ShopServiceCreateInput) (*model.ShopService, error) {
	panic(fmt.Errorf("not implemented: ShopServiceCreate - shopServiceCreate"))
}

// ShopServiceDelete is the resolver for the shopServiceDelete field.
func (r *mutationResolver) ShopServiceDelete(ctx context.Context, input model.DeleteIDInput) (*model.ShopService, error) {
	panic(fmt.Errorf("not implemented: ShopServiceDelete - shopServiceDelete"))
}

// ShopServiceDeleteAll is the resolver for the shopServiceDeleteAll field.
func (r *mutationResolver) ShopServiceDeleteAll(ctx context.Context) ([]*model.ShopService, error) {
	panic(fmt.Errorf("not implemented: ShopServiceDeleteAll - shopServiceDeleteAll"))
}

// TicketServiceCreate is the resolver for the ticketServiceCreate field.
func (r *mutationResolver) TicketServiceCreate(ctx context.Context, input model.TicketServiceCreateInput) (*model.TicketService, error) {
	panic(fmt.Errorf("not implemented: TicketServiceCreate - ticketServiceCreate"))
}

// TicketServiceDelete is the resolver for the ticketServiceDelete field.
func (r *mutationResolver) TicketServiceDelete(ctx context.Context, input model.DeleteIDInput) (*model.TicketService, error) {
	panic(fmt.Errorf("not implemented: TicketServiceDelete - ticketServiceDelete"))
}

// TicketServiceDeleteAll is the resolver for the ticketServiceDeleteAll field.
func (r *mutationResolver) TicketServiceDeleteAll(ctx context.Context) ([]*model.TicketService, error) {
	panic(fmt.Errorf("not implemented: TicketServiceDeleteAll - ticketServiceDeleteAll"))
}

// CustomerByID is the resolver for the customerByID field.
func (r *queryResolver) CustomerByID(ctx context.Context, input string) (*model.Customer, error) {
	result, err := r.DB.CustomerFindByID(ctx, input)
	return result, err
}

// Customers is the resolver for the customers field.
func (r *queryResolver) Customers(ctx context.Context) ([]*model.Customer, error) {
	result, err := r.DB.CustomerList(ctx)
	return result, err
}

// CarByID is the resolver for the carByID field.
func (r *queryResolver) CarByID(ctx context.Context, input string) (*model.Car, error) {
	result, err := r.DB.CarFindByID(ctx, input)
	return result, err
}

// CarByOwner is the resolver for the carByOwner field.
func (r *queryResolver) CarByOwner(ctx context.Context, input string) ([]*model.Car, error) {
	result, err := r.DB.CarFindByOwner(ctx, input)
	return result, err
}

// Cars is the resolver for the cars field.
func (r *queryResolver) Cars(ctx context.Context) ([]*model.Car, error) {
	result, err := r.DB.CarList(ctx)
	return result, err
}

// TicketByID is the resolver for the ticketByID field.
func (r *queryResolver) TicketByID(ctx context.Context, input string) (*model.Ticket, error) {
	result, err := r.DB.TicketFindByID(ctx, input)
	return result, err
}

// TicketByCustomer is the resolver for the ticketByCustomer field.
func (r *queryResolver) TicketByCustomer(ctx context.Context, input model.TicketByCustomerInput) ([]*model.Ticket, error) {
	result, err := r.DB.TicketFindByCustomer(ctx, input)
	return result, err
}

// TicketByShop is the resolver for the ticketByShop field.
func (r *queryResolver) TicketByShop(ctx context.Context, input model.TicketByShopInput) ([]*model.Ticket, error) {
	result, err := r.DB.TicketFindByShop(ctx, input)
	return result, err
}

// Tickets is the resolver for the tickets field.
func (r *queryResolver) Tickets(ctx context.Context) ([]*model.Ticket, error) {
	result, err := r.DB.TicketList(ctx)
	return result, err
}

// ShopByID is the resolver for the shopByID field.
func (r *queryResolver) ShopByID(ctx context.Context, input string) (*model.Shop, error) {
	panic(fmt.Errorf("not implemented: ShopByID - shopByID"))
}

// Shops is the resolver for the shops field.
func (r *queryResolver) Shops(ctx context.Context) ([]*model.Shop, error) {
	panic(fmt.Errorf("not implemented: Shops - shops"))
}

// ServiceByID is the resolver for the serviceByID field.
func (r *queryResolver) ServiceByID(ctx context.Context, input string) (*model.Service, error) {
	panic(fmt.Errorf("not implemented: ServiceByID - serviceByID"))
}

// Services is the resolver for the services field.
func (r *queryResolver) Services(ctx context.Context) ([]*model.Service, error) {
	panic(fmt.Errorf("not implemented: Services - services"))
}

// ShopServiceByID is the resolver for the shopServiceByID field.
func (r *queryResolver) ShopServiceByID(ctx context.Context, input string) (*model.ShopService, error) {
	panic(fmt.Errorf("not implemented: ShopServiceByID - shopServiceByID"))
}

// ShopServices is the resolver for the shopServices field.
func (r *queryResolver) ShopServices(ctx context.Context) ([]*model.ShopService, error) {
	panic(fmt.Errorf("not implemented: ShopServices - shopServices"))
}

// ActiveTicketByID is the resolver for the activeTicketByID field.
func (r *queryResolver) ActiveTicketByID(ctx context.Context, input string) (*model.ActiveTicket, error) {
	panic(fmt.Errorf("not implemented: ActiveTicketByID - activeTicketByID"))
}

// ActiveTicketByCustomer is the resolver for the activeTicketByCustomer field.
func (r *queryResolver) ActiveTicketByCustomer(ctx context.Context, input string) ([]*model.ActiveTicket, error) {
	panic(fmt.Errorf("not implemented: ActiveTicketByCustomer - activeTicketByCustomer"))
}

// ActiveTicketByShop is the resolver for the activeTicketByShop field.
func (r *queryResolver) ActiveTicketByShop(ctx context.Context, input string) ([]*model.ActiveTicket, error) {
	panic(fmt.Errorf("not implemented: ActiveTicketByShop - activeTicketByShop"))
}

// ActiveTickets is the resolver for the activeTickets field.
func (r *queryResolver) ActiveTickets(ctx context.Context) ([]*model.ActiveTicket, error) {
	panic(fmt.Errorf("not implemented: ActiveTickets - activeTickets"))
}

// TicketServiceByID is the resolver for the ticketServiceByID field.
func (r *queryResolver) TicketServiceByID(ctx context.Context, input string) (*model.TicketService, error) {
	panic(fmt.Errorf("not implemented: TicketServiceByID - ticketServiceByID"))
}

// TicketServices is the resolver for the ticketServices field.
func (r *queryResolver) TicketServices(ctx context.Context) ([]*model.TicketService, error) {
	panic(fmt.Errorf("not implemented: TicketServices - ticketServices"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *mutationResolver) ActiveTicketDeleteActive(ctx context.Context) ([]*model.ActiveTicket, error) {
	panic(fmt.Errorf("not implemented: ActiveTicketDeleteActive - activeTicketDeleteActive"))
}
