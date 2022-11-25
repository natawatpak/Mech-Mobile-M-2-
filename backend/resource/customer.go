package resource

import (
	"context"

	"github.com/google/uuid"
	"github.com/natawatpak/Mech-Mobile-M-2-/backend/graph/model"
	"github.com/natawatpak/Mech-Mobile-M-2-/backend/util"
)

// CUSTOMER
func (op *SQLop) CustomerCreate(ctx context.Context, cusInput *model.CustomerCreateInput) (*model.Customer, error) {
	newID := uuid.New().String()
	cusToBeAdd := model.Customer{
		ID:    newID,
		FName: cusInput.FName,
		LName: cusInput.LName,
		Tel:   cusInput.Tel,
		Email: cusInput.Email,
	}
	_, err := op.db.NewInsert().Model(&cusToBeAdd).Exec(ctx)
	return &cusToBeAdd, err
}

func (op *SQLop) CustomerUpdateMulti(ctx context.Context, updateInput model.Customer) (*model.Customer, error) {
	err := ValidateID(updateInput.ID)
	if err != nil {
		return nil, err
	}
	_, err = op.db.NewUpdate().Model(&updateInput).Where("id = ?", updateInput.ID).Exec(ctx)
	if util.CheckErr(err) {
		return nil, err
	}
	resultCustomer, err := op.CustomerFindByID(ctx, updateInput.ID)
	return resultCustomer, err
}

func (op *SQLop) CustomerDelete(ctx context.Context, ID string) (*model.Customer, error) {
	err := ValidateID(ID)
	if err != nil {
		return nil, err
	}
	resultCustomer, err := op.CustomerFindByID(ctx, ID)
	if util.CheckErr(err) {
		return nil, err
	}
	_, err = op.db.NewDelete().Model(op.cusModel).Where("id = ?", ID).Exec(ctx)
	return resultCustomer, err
}

func (op *SQLop) CustomerDeleteAll(ctx context.Context) ([]*model.Customer, error) {
	cusArr, err := op.CustomerList(ctx)
	PrintIfErrorExist(err)
	for _, v := range cusArr {
		_, err := op.CustomerDelete(ctx, v.ID)
		PrintIfErrorExist(err)
	}
	return cusArr, err
}

func (op *SQLop) CustomerFindByID(ctx context.Context, ID string) (*model.Customer, error) {
	err := ValidateID(ID)
	if err != nil {
		return nil, err
	}
	arrModel := new(model.Customer)
	err = op.db.NewSelect().Model(op.cusModel).Where("id = ?", ID).Scan(ctx, arrModel)
	return arrModel, err
}

func (op *SQLop) CustomerList(ctx context.Context) ([]*model.Customer, error) {
	customer := new([]*model.Customer)
	err := op.db.NewSelect().Model(op.cusModel).Scan(ctx, customer)
	return *customer, err
}
