package resource

import (
	"context"

	"github.com/google/uuid"
	"github.com/natawatpak/Mech-Mobile-M-2-/backend/graph/model"
	"github.com/natawatpak/Mech-Mobile-M-2-/backend/util"
)

// Service
func (op *SQLop) ServiceCreate(ctx context.Context, serviceInput *model.ServiceCreateInput) (*model.Service, error) {
	newID := uuid.New().String()
	serviceToBeAdd := model.Service{
		ID:   newID,
		Name: serviceInput.Name,
	}
	_, err := op.db.NewInsert().Model(&serviceToBeAdd).Exec(ctx)
	return &serviceToBeAdd, err
}

func (op *SQLop) ServiceUpdateMulti(ctx context.Context, updateInput model.ServiceUpdateInput) (*model.Service, error) {
	_, err := op.db.NewUpdate().Model(&updateInput).Where("id = ?", updateInput.ID).Exec(ctx)
	if util.CheckErr(err) {
		return nil, err
	}
	resultService, err := op.ServiceFindByID(ctx, updateInput.ID)
	return resultService, err
}

func (op *SQLop) ServiceDelete(ctx context.Context, ID string) (*model.Service, error) {
	resultService, err := op.ServiceFindByID(ctx, ID)
	if util.CheckErr(err) {
		return nil, err
	}
	_, err = op.db.NewDelete().Model(op.serviceModel).Where("id = ?", ID).Exec(ctx)
	return resultService, err
}

func (op *SQLop) ServiceDeleteAll(ctx context.Context) ([]*model.Service, error) {
	serviceArr, err := op.ServiceList(ctx)
	PrintIfErrorExist(err)
	for _, v := range serviceArr {
		_, err := op.ServiceDelete(ctx, v.ID)
		PrintIfErrorExist(err)
	}
	return serviceArr, err
}

func (op *SQLop) ServiceFindByID(ctx context.Context, ID string) (*model.Service, error) {
	arrModel := new(model.Service)
	err := op.db.NewSelect().Model(op.serviceModel).Where("id = ?", ID).Scan(ctx, arrModel)
	return arrModel, err
}

func (op *SQLop) ServiceList(ctx context.Context) ([]*model.Service, error) {
	Service := new([]*model.Service)
	err := op.db.NewSelect().Model(op.serviceModel).Scan(ctx, Service)
	return *Service, err
}
