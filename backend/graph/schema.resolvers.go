package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
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
		return nil, fmt.Errorf("id must be null")
	}
	returnCustomer, err := r.DB.CustomerCreate(ctx, &input)
	return returnCustomer, err
}

// CustomerUpdateMulti is the resolver for the customerUpdateMulti field.
func (r *mutationResolver) CustomerUpdateMulti(ctx context.Context, input model.CustomerUpdateInput) ([]*model.Customer, error) {
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
	panic(fmt.Errorf("not implemented: CarCreate - carCreate"))
}

// TicketCreate is the resolver for the ticketCreate field.
func (r *mutationResolver) TicketCreate(ctx context.Context, input model.TicketCreateInput) (*model.Ticket, error) {
	panic(fmt.Errorf("not implemented: TicketCreate - ticketCreate"))
}

// ShopCreate is the resolver for the shopCreate field.
func (r *mutationResolver) ShopCreate(ctx context.Context, input model.ShopCreateInput) (*model.Shop, error) {
	panic(fmt.Errorf("not implemented: ShopCreate - shopCreate"))
}

// ServiceCreate is the resolver for the serviceCreate field.
func (r *mutationResolver) ServiceCreate(ctx context.Context, input model.ServiceCreateInput) (*model.Service, error) {
	panic(fmt.Errorf("not implemented: ServiceCreate - serviceCreate"))
}

// ShopServiceCreate is the resolver for the shopServiceCreate field.
func (r *mutationResolver) ShopServiceCreate(ctx context.Context, input model.ShopServiceCreateInput) (*model.ShopService, error) {
	panic(fmt.Errorf("not implemented: ShopServiceCreate - shopServiceCreate"))
}

// ActiveTicketCreate is the resolver for the activeTicketCreate field.
func (r *mutationResolver) ActiveTicketCreate(ctx context.Context, input model.ActiveTicketCreateInput) (*model.ActiveTicket, error) {
	panic(fmt.Errorf("not implemented: ActiveTicketCreate - activeTicketCreate"))
}

// TicketServiceCreate is the resolver for the ticketServiceCreate field.
func (r *mutationResolver) TicketServiceCreate(ctx context.Context, input model.TicketServiceCreateInput) (*model.TicketService, error) {
	panic(fmt.Errorf("not implemented: TicketServiceCreate - ticketServiceCreate"))
}

// Customer is the resolver for the customer field.
func (r *queryResolver) Customer(ctx context.Context) ([]*model.Customer, error) {
	result, err := r.DB.CustomerList(ctx)
	return result, err
}

// Car is the resolver for the car field.
func (r *queryResolver) Car(ctx context.Context) ([]*model.Car, error) {
	panic(fmt.Errorf("not implemented: Car - car"))
}

// Ticket is the resolver for the ticket field.
func (r *queryResolver) Ticket(ctx context.Context) ([]*model.Ticket, error) {
	panic(fmt.Errorf("not implemented: Ticket - ticket"))
}

// Shop is the resolver for the shop field.
func (r *queryResolver) Shop(ctx context.Context) ([]*model.Shop, error) {
	panic(fmt.Errorf("not implemented: Shop - shop"))
}

// Service is the resolver for the service field.
func (r *queryResolver) Service(ctx context.Context) ([]*model.Service, error) {
	panic(fmt.Errorf("not implemented: Service - service"))
}

// ShopService is the resolver for the shopService field.
func (r *queryResolver) ShopService(ctx context.Context) ([]*model.ShopService, error) {
	panic(fmt.Errorf("not implemented: ShopService - shopService"))
}

// ActiveTicket is the resolver for the activeTicket field.
func (r *queryResolver) ActiveTicket(ctx context.Context) ([]*model.ActiveTicket, error) {
	panic(fmt.Errorf("not implemented: ActiveTicket - activeTicket"))
}

// TicketService is the resolver for the ticketService field.
func (r *queryResolver) TicketService(ctx context.Context) ([]*model.TicketService, error) {
	panic(fmt.Errorf("not implemented: TicketService - ticketService"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
