package resource

import (
	"context"

	"github.com/natawatpak/Mech-Mobile-M-2-/backend/graph/model"
	"github.com/natawatpak/Mech-Mobile-M-2-/backend/util"
)

// TicketService
func (op *SQLop) TicketServiceCreate(ctx context.Context, TicketServiceInput *model.TicketServiceCreateInput) (*model.TicketService, error) {
	TicketServiceToBeAdd := model.TicketService{
		TicketID:  TicketServiceInput.TicketID,
		ServiceID: TicketServiceInput.ServiceID,
	}
	_, err := op.db.NewInsert().Model(&TicketServiceToBeAdd).Exec(ctx)
	return &TicketServiceToBeAdd, err
}

func (op *SQLop) TicketServiceDelete(ctx context.Context, input *model.TicketServiceCreateInput) (*model.TicketService, error) {
	resultTicketService, err := op.TicketServiceFindByID(ctx, model.TicketServiceCreateInput{TicketID: input.TicketID, ServiceID: input.ServiceID})
	if util.CheckErr(err) {
		return nil, err
	}
	_, err = op.db.NewDelete().Model(op.ticketService).
		Where("ticket_id = ? AND service_id = ?", input.TicketID, input.ServiceID).Exec(ctx)
	return resultTicketService, err
}

func (op *SQLop) TicketServiceDeleteAll(ctx context.Context) ([]*model.TicketService, error) {
	TicketServiceArr, err := op.TicketServiceList(ctx)
	PrintIfErrorExist(err)
	for _, v := range TicketServiceArr {
		_, err := op.TicketServiceDelete(ctx, &model.TicketServiceCreateInput{TicketID: v.TicketID, ServiceID: v.ServiceID})
		PrintIfErrorExist(err)
	}
	return TicketServiceArr, err
}

func (op *SQLop) TicketServiceFindByID(ctx context.Context, input model.TicketServiceCreateInput) (*model.TicketService, error) {
	arrModel := new(model.TicketService)
	err := op.db.NewSelect().Model(op.ticketService).Where("ticket_id = ? AND service_id = ?", input.TicketID, input.ServiceID).Scan(ctx, arrModel)
	return arrModel, err
}

func (op *SQLop) TicketServiceList(ctx context.Context) ([]*model.TicketService, error) {
	TicketService := new([]*model.TicketService)
	err := op.db.NewSelect().Model(op.ticketService).Scan(ctx, TicketService)
	return *TicketService, err
}
