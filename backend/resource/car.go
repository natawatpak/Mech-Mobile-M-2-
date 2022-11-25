package resource

import (
	"context"

	"github.com/google/uuid"
	"github.com/natawatpak/Mech-Mobile-M-2-/backend/graph/model"
	"github.com/natawatpak/Mech-Mobile-M-2-/backend/util"
)

// CAR

func (op *SQLop) CarCreate(ctx context.Context, carInput *model.CarCreateInput) (*model.Car, error) {
	newID := uuid.New().String()
	carToBeAdd := model.Car{
		ID:        newID,
		OwnerID:   carInput.OwnerID,
		PlateNum:  carInput.PlateNum,
		PlateType: carInput.PlateType,
		IssuedAt:  carInput.IssuedAt,
		Color:     carInput.Color,
		Type:      carInput.Type,
		Brand:     carInput.Brand,
		Build:     carInput.Build,
	}
	_, err := op.db.NewInsert().Model(&carToBeAdd).Exec(ctx)
	return &carToBeAdd, err
}

func (op *SQLop) CarUpdateMulti(ctx context.Context, updateInput model.Car) (*model.Car, error) {
	err := ValidateID(updateInput.ID)
	if err != nil {
		return nil, err
	}
	_, err = op.db.NewUpdate().Model(&updateInput).Where("id = ?", updateInput.ID).Exec(ctx)
	if util.CheckErr(err) {
		return nil, err
	}
	resultCustomer, err := op.CarFindByID(ctx, updateInput.ID)
	return resultCustomer, err
}

func (op *SQLop) CarDelete(ctx context.Context, ID string) (*model.Car, error) {
	err := ValidateID(ID)
	if err != nil {
		return nil, err
	}
	resultCustomer, err := op.CarFindByID(ctx, ID)
	if util.CheckErr(err) {
		return nil, err
	}
	_, err = op.db.NewDelete().Model(op.carModel).Where("id = ?", ID).Exec(ctx)
	return resultCustomer, err
}

func (op *SQLop) CarDeleteAll(ctx context.Context) ([]*model.Car, error) {
	carArr, err := op.CarList(ctx)
	PrintIfErrorExist(err)
	for _, v := range carArr {
		_, err := op.CustomerDelete(ctx, v.ID)
		PrintIfErrorExist(err)
	}
	return carArr, err
}

func (op *SQLop) CarFindByID(ctx context.Context, ID string) (*model.Car, error) {
	err := ValidateID(ID)
	if err != nil {
		return nil, err
	}
	arrModel := new(model.Car)
	err = op.db.NewSelect().Model(op.carModel).Where("id = ?", ID).Scan(ctx, arrModel)
	return arrModel, err
}

func (op *SQLop) CarFindByOwner(ctx context.Context, ID string) ([]*model.Car, error) {
	err := ValidateID(ID)
	if err != nil {
		return nil, err
	}
	car := new([]*model.Car)
	err = op.db.NewSelect().Model(op.carModel).Where("owner_id = ?", ID).Scan(ctx, car)
	return *car, err
}

func (op *SQLop) CarList(ctx context.Context) ([]*model.Car, error) {
	car := new([]*model.Car)
	err := op.db.NewSelect().Model(op.carModel).Scan(ctx, car)
	return *car, err
}
