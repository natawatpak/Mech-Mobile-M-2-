package resource

import (
	"context"

	"github.com/natawatpak/Mech-Mobile-M-2-/backend/graph/model"
	"github.com/natawatpak/Mech-Mobile-M-2-/backend/util"
)

// ShopService
func (op *SQLop) ShopServiceCreate(ctx context.Context, shopServiceInput *model.ShopServiceCreateInput) (*model.ShopService, error) {
	err := ValidateID(shopServiceInput.ShopID)
	if err != nil {
		return nil, err
	}
	err = ValidateID(shopServiceInput.ServiceID)
	if err != nil {
		return nil, err
	}
	shopServiceToBeAdd := model.ShopService{
		ShopID:    shopServiceInput.ShopID,
		ServiceID: shopServiceInput.ServiceID,
	}
	_, err = op.db.NewInsert().Model(&shopServiceToBeAdd).Exec(ctx)
	return &shopServiceToBeAdd, err
}

func (op *SQLop) ShopServiceDelete(ctx context.Context, input *model.ShopServiceDeleteInput) (*model.ShopService, error) {
	err := ValidateID(input.ShopID)
	if err != nil {
		return nil, err
	}
	err = ValidateID(input.ServiceID)
	if err != nil {
		return nil, err
	}
	resultShopService, err := op.ShopServiceFindByID(ctx, model.ShopServiceCreateInput{ShopID: input.ShopID, ServiceID: input.ServiceID})
	if util.CheckErr(err) {
		return nil, err
	}
	_, err = op.db.NewDelete().Model(op.shopService).
		Where("shop_id = ? AND service_id = ?", input.ShopID, input.ServiceID).Exec(ctx)
	return resultShopService, err
}

func (op *SQLop) ShopServiceDeleteAll(ctx context.Context) ([]*model.ShopService, error) {
	shopServiceArr, err := op.ShopServiceList(ctx)
	PrintIfErrorExist(err)
	for _, v := range shopServiceArr {
		_, err := op.ShopServiceDelete(ctx, &model.ShopServiceDeleteInput{ShopID: v.ShopID, ServiceID: v.ServiceID})
		PrintIfErrorExist(err)
	}
	return shopServiceArr, err
}

func (op *SQLop) ShopServiceFindByID(ctx context.Context, input model.ShopServiceCreateInput) (*model.ShopService, error) {
	err := ValidateID(input.ShopID)
	if err != nil {
		return nil, err
	}
	err = ValidateID(input.ServiceID)
	if err != nil {
		return nil, err
	}
	arrModel := new(model.ShopService)
	err = op.db.NewSelect().Model(op.shopService).Where("shop_id = ? AND service_id = ?", input.ShopID, input.ServiceID).Scan(ctx, arrModel)
	return arrModel, err
}

func (op *SQLop) ShopServiceList(ctx context.Context) ([]*model.ShopService, error) {
	ShopService := new([]*model.ShopService)
	err := op.db.NewSelect().Model(op.shopService).Scan(ctx, ShopService)
	return *ShopService, err
}
