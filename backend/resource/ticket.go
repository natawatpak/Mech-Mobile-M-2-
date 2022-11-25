package resource

import (
	"context"

	"github.com/google/uuid"
	"github.com/natawatpak/Mech-Mobile-M-2-/backend/graph/model"
	"github.com/natawatpak/Mech-Mobile-M-2-/backend/util"
)

// Ticket
func (op *SQLop) TicketCreate(ctx context.Context, ticketInput *model.TicketCreateInput) (*model.Ticket, error) {
	newID := uuid.New().String()
	ticketToBeAdd := model.Ticket{
		ID:           newID,
		CustomerID:   ticketInput.CustomerID,
		CarID:        ticketInput.CarID,
		Problem:      ticketInput.Problem,
		CreateTime:   ticketInput.CreateTime,
		ShopID:       ticketInput.ShopID,
		AcceptedTime: ticketInput.AcceptedTime,
		Status:       ticketInput.Status,
		Longitude:    ticketInput.Longitude,
		Latitude:     ticketInput.Latitude,
	}
	_, err := op.db.NewInsert().Model(&ticketToBeAdd).Exec(ctx)
	return &ticketToBeAdd, err
}

func (op *SQLop) TicketUpdateMulti(ctx context.Context, updateInput model.Ticket) (*model.Ticket, error) {
	err := ValidateID(updateInput.ID)
	if err != nil {
		return nil, err
	}
	_, err = op.db.NewUpdate().Model(&updateInput).Where("id = ?", updateInput.ID).Exec(ctx)
	if util.CheckErr(err) {
		return nil, err
	}
	resultTicket, err := op.TicketFindByID(ctx, updateInput.ID)
	return resultTicket, err
}

func (op *SQLop) TicketDelete(ctx context.Context, ID string) (*model.Ticket, error) {
	err := ValidateID(ID)
	if err != nil {
		return nil, err
	}
	resultTicket, err := op.TicketFindByID(ctx, ID)
	if util.CheckErr(err) {
		return nil, err
	}
	_, err = op.db.NewDelete().Model(op.ticketModel).Where("id = ?", ID).Exec(ctx)
	return resultTicket, err
}

func (op *SQLop) TicketDeleteAll(ctx context.Context) ([]*model.Ticket, error) {
	ticketArr, err := op.TicketList(ctx)
	PrintIfErrorExist(err)
	for _, v := range ticketArr {
		_, err := op.TicketDelete(ctx, v.ID)
		PrintIfErrorExist(err)
	}
	return ticketArr, err
}

func (op *SQLop) TicketFindByID(ctx context.Context, ID string) (*model.Ticket, error) {
	err := ValidateID(ID)
	if err != nil {
		return nil, err
	}
	arrModel := new(model.Ticket)
	err = op.db.NewSelect().Model(op.ticketModel).Where("id = ?", ID).Scan(ctx, arrModel)
	return arrModel, err
}

func (op *SQLop) TicketFindByCustomer(ctx context.Context, input model.TicketByCustomerInput) ([]*model.Ticket, error) {
	err := ValidateID(input.CustomerID)
	if err != nil {
		return nil, err
	}
	Ticket := new([]*model.Ticket)
	err = op.db.NewSelect().Model(op.ticketModel).
		Where("customer_id = ? AND create_time >= ? AND create_time <= ?", input.CustomerID, input.FromTime, input.ToTime).Scan(ctx, Ticket)
	return *Ticket, err
}

func (op *SQLop) TicketFindByShop(ctx context.Context, input model.TicketByShopInput) ([]*model.Ticket, error) {
	err := ValidateID(input.ShopID)
	if err != nil {
		return nil, err
	}
	Ticket := new([]*model.Ticket)
	err = op.db.NewSelect().Model(op.ticketModel).
		Where("shop_id = ? AND create_time >= ? AND create_time <= ?", input.ShopID, input.FromTime, input.ToTime).Scan(ctx, Ticket)
	return *Ticket, err
}

func (op *SQLop) TicketList(ctx context.Context) ([]*model.Ticket, error) {
	Ticket := new([]*model.Ticket)
	err := op.db.NewSelect().Model(op.ticketModel).Scan(ctx, Ticket)
	return *Ticket, err
}
