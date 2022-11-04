package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/natawatpak/Mech-Mobile-M-2-/backend/graph/generated"
	"github.com/natawatpak/Mech-Mobile-M-2-/backend/graph/model"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.UserCreateInput) (*model.User, error) {
	panic(fmt.Errorf("not implemented: CreateUser - createUser"))
}

// CreateCountry is the resolver for the createCountry field.
func (r *mutationResolver) CreateCountry(ctx context.Context, input model.CountryCreateInput) (*model.Country, error) {
	panic(fmt.Errorf("not implemented: CreateCountry - createCountry"))
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented: User - user"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
