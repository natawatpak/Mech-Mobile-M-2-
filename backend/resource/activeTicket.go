package resource

import (
	"context"
	"errors"

	"github.com/natawatpak/Mech-Mobile-M-2-/backend/graph/model"
	"github.com/natawatpak/Mech-Mobile-M-2-/backend/util"
)

// ActiveTicket
func (op *SQLop) ActiveTicketCreate(ctx context.Context, activeTicketInput *model.ActiveTicketCreateInput) (*model.ActiveTicket, error) {
	err := ValidateID(activeTicketInput.ID)
	if err != nil {
		return nil, err
	}
	activeTicketToBeAdd := model.ActiveTicket{
		ID:          activeTicketInput.ID,
		CarID:       activeTicketInput.CarID,
		CustomerID:  activeTicketInput.CustomerID,
		Problem:     activeTicketInput.Problem,
		Description: activeTicketInput.Description,
		ShopID:      activeTicketInput.ShopID,
		Status:      activeTicketInput.Status,
		Latitude:    activeTicketInput.Latitude,
		Longitude:   activeTicketInput.Longitude,
	}
	_, err = op.db.NewInsert().Model(&activeTicketToBeAdd).Exec(ctx)
	return &activeTicketToBeAdd, err
}

func (op *SQLop) ActiveTicketUpdateMulti(ctx context.Context, updateInput model.ActiveTicket) (*model.ActiveTicket, error) {
	err := ValidateID(updateInput.ID)
	if err != nil {
		return nil, err
	}
	_, err = op.db.NewUpdate().Model(&updateInput).Where("id = ?", updateInput.ID).Exec(ctx)
	if util.CheckErr(err) {
		return nil, err
	}
	resultActiveTicket, err := op.ActiveTicketFindByID(ctx, updateInput.ID)
	return resultActiveTicket, err
}

func (op *SQLop) ActiveTicketDelete(ctx context.Context, ID string) (*model.ActiveTicket, error) {
	err := ValidateID(ID)
	if err != nil {
		return nil, err
	}
	resultActiveTicket, err := op.ActiveTicketFindByID(ctx, ID)
	if util.CheckErr(err) {
		return nil, err
	}
	_, err = op.db.NewDelete().Model(op.activeTicket).Where("id = ?", ID).Exec(ctx)
	return resultActiveTicket, err
}

func (op *SQLop) ActiveTicketDeleteByStatus(ctx context.Context, input string) ([]*model.ActiveTicket, error) {
	err := ValidateID(input)
	if err != nil {
		return nil, err
	}
	activeTicketArr, err := op.ActiveTicketFindByStatus(ctx, input)
	PrintIfErrorExist(err)
	for _, v := range activeTicketArr {
		_, err := op.ActiveTicketDelete(ctx, v.ID)
		PrintIfErrorExist(err)
	}
	return activeTicketArr, err
}

func (op *SQLop) ActiveTicketDeleteAll(ctx context.Context) ([]*model.ActiveTicket, error) {
	activeTicketArr, err := op.ActiveTicketList(ctx)
	PrintIfErrorExist(err)
	for _, v := range activeTicketArr {
		_, err := op.ActiveTicketDelete(ctx, v.ID)
		PrintIfErrorExist(err)
	}
	return activeTicketArr, err
}

func (op *SQLop) ActiveTicketFindByID(ctx context.Context, ID string) (*model.ActiveTicket, error) {
	err := ValidateID(ID)
	if err != nil {
		return nil, err
	}
	arrModel := new(model.ActiveTicket)
	err = op.db.NewSelect().Model(op.activeTicket).Where("id = ?", ID).Scan(ctx, arrModel)
	return arrModel, err
}

func (op *SQLop) ActiveTicketFindByStatus(ctx context.Context, input string) ([]*model.ActiveTicket, error) {
	if len(input) < 1 {
		return nil, errors.New("status string is empty")
	}
	ActiveTicket := new([]*model.ActiveTicket)
	err := op.db.NewSelect().Model(op.activeTicket).Where("status = ?", input).Scan(ctx, ActiveTicket)
	return *ActiveTicket, err
}

func (op *SQLop) ActiveTicketFindByCustomer(ctx context.Context, ID string) ([]*model.ActiveTicket, error) {
	err := ValidateID(ID)
	if err != nil {
		return nil, err
	}
	ActiveTicket := new([]*model.ActiveTicket)
	err = op.db.NewSelect().Model(op.activeTicket).Where("customer_id = ?", ID).Scan(ctx, ActiveTicket)
	return *ActiveTicket, err
}

func (op *SQLop) ActiveTicketFindByShop(ctx context.Context, ID string) ([]*model.ActiveTicket, error) {
	err := ValidateID(ID)
	if err != nil {
		return nil, err
	}
	ActiveTicket := new([]*model.ActiveTicket)
	err = op.db.NewSelect().Model(op.activeTicket).Where("shop_id = ?", ID).Scan(ctx, ActiveTicket)
	return *ActiveTicket, err
}

func (op *SQLop) ActiveTicketList(ctx context.Context) ([]*model.ActiveTicket, error) {
	ActiveTicket := new([]*model.ActiveTicket)
	err := op.db.NewSelect().Model(op.activeTicket).Scan(ctx, ActiveTicket)
	return *ActiveTicket, err
}
