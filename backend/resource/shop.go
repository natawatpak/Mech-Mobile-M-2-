package resource

import (
	"context"

	"github.com/google/uuid"
	"github.com/natawatpak/Mech-Mobile-M-2-/backend/graph/model"
	"github.com/natawatpak/Mech-Mobile-M-2-/backend/util"
)

// Shop
func (op *SQLop) ShopCreate(ctx context.Context, shopInput *model.ShopCreateInput) (*model.Shop, error) {
	newID := uuid.New().String()
	shopToBeAdd := model.Shop{
		ID:      newID,
		Name:    shopInput.Name,
		Tel:     shopInput.Tel,
		Email:   shopInput.Email,
		Address: shopInput.Address,
	}
	_, err := op.db.NewInsert().Model(&shopToBeAdd).Exec(ctx)
	return &shopToBeAdd, err
}

func (op *SQLop) ShopUpdateMulti(ctx context.Context, updateInput model.Shop) (*model.Shop, error) {
	_, err := op.db.NewUpdate().Model(&updateInput).Where("id = ?", updateInput.ID).Exec(ctx)
	if util.CheckErr(err) {
		return nil, err
	}
	resultShop, err := op.ShopFindByID(ctx, updateInput.ID)
	return resultShop, err
}

func (op *SQLop) ShopDelete(ctx context.Context, ID string) (*model.Shop, error) {
	resultShop, err := op.ShopFindByID(ctx, ID)
	if util.CheckErr(err) {
		return nil, err
	}
	_, err = op.db.NewDelete().Model(op.shopModel).Where("id = ?", ID).Exec(ctx)
	return resultShop, err
}

func (op *SQLop) ShopDeleteAll(ctx context.Context) ([]*model.Shop, error) {
	shopArr, err := op.ShopList(ctx)
	PrintIfErrorExist(err)
	for _, v := range shopArr {
		_, err := op.ShopDelete(ctx, v.ID)
		PrintIfErrorExist(err)
	}
	return shopArr, err
}

func (op *SQLop) ShopFindByID(ctx context.Context, ID string) (*model.Shop, error) {
	arrModel := new(model.Shop)
	err := op.db.NewSelect().Model(op.shopModel).Where("id = ?", ID).Scan(ctx, arrModel)
	return arrModel, err
}

func (op *SQLop) ShopList(ctx context.Context) ([]*model.Shop, error) {
	Shop := new([]*model.Shop)
	err := op.db.NewSelect().Model(op.shopModel).Scan(ctx, Shop)
	return *Shop, err
}
