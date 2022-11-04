package resource

import (
	"context"

	"github.com/natawatpak/Mech-Mobile-M-2-/backend/graph/model"
	"github.com/uptrace/bun"
)

type SQLop struct {
	db            *bun.DB
	userModel     *model.User
	countryModel  *model.Country
	merchantModel *model.Merchant
	orderModel    *model.Order
	productModel  *model.Product
}

type DatabaseOp interface {
	UserCreate(ctx context.Context, userInput *model.UserCreateInput) (*model.User, error)

	UserList(ctx context.Context) ([]*model.User, error)
}
