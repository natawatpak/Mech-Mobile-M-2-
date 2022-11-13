package resource

import (
	"context"

	"github.com/google/uuid"
	"github.com/natawatpak/Mech-Mobile-M-2-/backend/graph/model"
	"github.com/natawatpak/Mech-Mobile-M-2-/backend/util"
)

// ActiveTicket
func (op *SQLop) ActiveTicketCreate(ctx context.Context, activeTicketInput *model.ActiveTicketCreateInput) (*model.ActiveTicket, error) {
	newID := uuid.New().String()
	activeTicketToBeAdd := model.ActiveTicket{
		ID:         newID,
		CarID:      activeTicketInput.CarID,
		CustomerID: activeTicketInput.CustomerID,
		Problem:    activeTicketInput.Problem,
		ShopID:     activeTicketInput.ShopID,
		Status:     activeTicketInput.Status,
	}
	_, err := op.db.NewInsert().Model(&activeTicketToBeAdd).Exec(ctx)
	return &activeTicketToBeAdd, err
}

func (op *SQLop) ActiveTicketUpdateMulti(ctx context.Context, updateInput model.ActiveTicket) (*model.ActiveTicket, error) {
	_, err := op.db.NewUpdate().Model(&updateInput).Where("id = ?", updateInput.ID).Exec(ctx)
	if util.CheckErr(err) {
		return nil, err
	}
	resultActiveTicket, err := op.ActiveTicketFindByID(ctx, updateInput.ID)
	return resultActiveTicket, err
}

func (op *SQLop) ActiveTicketDelete(ctx context.Context, ID string) (*model.ActiveTicket, error) {
	resultActiveTicket, err := op.ActiveTicketFindByID(ctx, ID)
	if util.CheckErr(err) {
		return nil, err
	}
	_, err = op.db.NewDelete().Model(op.activeTicket).Where("id = ?", ID).Exec(ctx)
	return resultActiveTicket, err
}

func (op *SQLop) ActiveTicketDeleteActive(ctx context.Context) ([]*model.ActiveTicket, error) {
	activeTicketArr, err := op.ActiveTicketList(ctx)
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
	arrModel := new(model.ActiveTicket)
	err := op.db.NewSelect().Model(op.activeTicket).Where("id = ?", ID).Scan(ctx, arrModel)
	return arrModel, err
}

func (op *SQLop) ActiveTicketListStatus(ctx context.Context, input model.Status) ([]*model.ActiveTicket, error) {
	ActiveTicket := new([]*model.ActiveTicket)
	err := op.db.NewSelect().Model(op.activeTicket).Where("status = ?", input).Scan(ctx, ActiveTicket)
	return *ActiveTicket, err
}

func (op *SQLop) ActiveTicketList(ctx context.Context) ([]*model.ActiveTicket, error) {
	ActiveTicket := new([]*model.ActiveTicket)
	err := op.db.NewSelect().Model(op.activeTicket).Scan(ctx, ActiveTicket)
	return *ActiveTicket, err
}
