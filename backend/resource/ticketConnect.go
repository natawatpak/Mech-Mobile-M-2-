package resource

import (
	"context"

	"github.com/natawatpak/Mech-Mobile-M-2-/backend/graph/model"
	"github.com/natawatpak/Mech-Mobile-M-2-/backend/util"
)

func (op *SQLop) TicketConnectCreate(ctx context.Context, input *model.TicketConnectCreateInput) (*model.TicketConnect, error) {
	err := ValidateID(input.TicketID)
	if err != nil {
		return nil, err
	}
	ticketToBeAdd := model.TicketConnect{
		TicketID:             input.TicketID,
		CustomerConnectionID: input.CustomerConnectionID,
		ShopConnectionID:     input.ShopConnectionID,
	}
	_, err = op.db.NewInsert().Model(&ticketToBeAdd).Exec(ctx)
	return &ticketToBeAdd, err
}

func (op *SQLop) TicketConnectUpdateMulti(ctx context.Context, updateInput model.TicketConnectCreateInput) (*model.TicketConnect, error) {
	err := ValidateID(updateInput.TicketID)
	if err != nil {
		return nil, err
	}
	_, err = op.db.NewUpdate().Model(&updateInput).Where("ticket_id = ?", updateInput.TicketID).Exec(ctx)
	if util.CheckErr(err) {
		return nil, err
	}
	resultTicket, err := op.TicketConnectFindByID(ctx, updateInput.TicketID)
	return resultTicket, err
}

func (op *SQLop) TicketConnectDelete(ctx context.Context, ID string) (*model.TicketConnect, error) {
	err := ValidateID(ID)
	if err != nil {
		return nil, err
	}
	resultTicketConnect, err := op.TicketConnectFindByID(ctx, ID)
	if util.CheckErr(err) {
		return nil, err
	}
	_, err = op.db.NewDelete().Model(new(model.ShopConnect)).Where("ticket_id = ?", ID).Exec(ctx)
	return resultTicketConnect, err
}

func (op *SQLop) TicketConnectDeleteAll(ctx context.Context) ([]*model.TicketConnect, error) {
	ticketArr, err := op.TicketConnectList(ctx)
	PrintIfErrorExist(err)
	for _, v := range ticketArr {
		_, err := op.TicketConnectDelete(ctx, v.TicketID)
		PrintIfErrorExist(err)
	}
	return ticketArr, err
}

func (op *SQLop) TicketConnectFindByID(ctx context.Context, ID string) (*model.TicketConnect, error) {
	err := ValidateID(ID)
	if err != nil {
		return nil, err
	}
	arrModel := new(model.TicketConnect)
	err = op.db.NewSelect().Model(arrModel).Where("ticket_id = ?", ID).Scan(ctx, arrModel)
	return arrModel, err
}

func (op *SQLop) TicketConnectList(ctx context.Context) ([]*model.TicketConnect, error) {
	ticketConnect := new([]*model.TicketConnect)
	err := op.db.NewSelect().Model(new(model.TicketConnect)).Scan(ctx, ticketConnect)
	return *ticketConnect, err
}
