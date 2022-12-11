package resource

import (
	"context"

	"github.com/natawatpak/Mech-Mobile-M-2-/backend/graph/model"
	"github.com/natawatpak/Mech-Mobile-M-2-/backend/util"
)

func (op *SQLop) ShopConnectCreate(ctx context.Context, input *model.ShopConnectCreateInput) (*model.ShopConnect, error) {
	err := ValidateID(input.ShopID)
	if err != nil {
		return nil, err
	}
	shopConnectToBeAdd := model.ShopConnect{
		ShopID:       input.ShopID,
		ConnectionID: input.ConnectionID,
	}
	_, err = op.db.NewInsert().Model(&shopConnectToBeAdd).Exec(ctx)
	return &shopConnectToBeAdd, err
}

func (op *SQLop) ShopConnectDelete(ctx context.Context, ID string) ([]*model.ShopConnect, error) {
	err := ValidateID(ID)
	if err != nil {
		return nil, err
	}
	resultShopConnect, err := op.ShopConnectFindByID(ctx, ID)
	if util.CheckErr(err) {
		return nil, err
	}
	for _, v := range resultShopConnect {
		_, err = op.db.NewDelete().Model(op.cusModel).Where("shop_id = ?", v.ShopID).Exec(ctx)
	}
	return resultShopConnect, err
}

func (op *SQLop) ShopConnectDeleteAll(ctx context.Context) ([]*model.ShopConnect, error) {
	ticketArr, err := op.ShopConnectList(ctx)
	PrintIfErrorExist(err)
	for _, v := range ticketArr {
		_, err := op.ShopConnectDelete(ctx, v.ShopID)
		PrintIfErrorExist(err)
	}
	return ticketArr, err
}

func (op *SQLop) ShopConnectFindByID(ctx context.Context, ID string) ([]*model.ShopConnect, error) {
	err := ValidateID(ID)
	if err != nil {
		return nil, err
	}
	arrModel := new([]*model.ShopConnect)
	err = op.db.NewSelect().Model(op.cusModel).Where("shop_id = ?", ID).Scan(ctx, arrModel)
	return *arrModel, err
}

func (op *SQLop) ShopConnectList(ctx context.Context) ([]*model.ShopConnect, error) {
	ticketConnect := new([]*model.ShopConnect)
	err := op.db.NewSelect().Model(op.cusModel).Scan(ctx, ticketConnect)
	return *ticketConnect, err
}
