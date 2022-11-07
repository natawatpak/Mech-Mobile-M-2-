package resource

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/natawatpak/Mech-Mobile-M-2-/backend/graph/model"
	"github.com/natawatpak/Mech-Mobile-M-2-/backend/util"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
	"github.com/uptrace/bun/extra/bundebug"
)

func PrintIfErrorExist(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func NewDBOperator(dbName string) (*SQLop, error) {
	sqldb, err := sql.Open(sqliteshim.ShimName, dbName)
	db := bun.NewDB(sqldb, sqlitedialect.New())
	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
		bundebug.FromEnv("BUNDEBUG"),
	))
	return &SQLop{
		cusModel: new(model.Customer),
		db:       db,
	}, err
}

func (op *SQLop) DropAllTable(ctx context.Context) (sql.Result, error) {
	sqlDeleteCustomer, err := op.db.NewDropTable().
		Model((*model.Customer)(nil)).
		IfExists().
		Exec(ctx)
	if util.CheckErr(err) {
		return sqlDeleteCustomer, err
	}
	sqlDeleteCar, err := op.db.NewDropTable().
		Model((*model.Car)(nil)).
		IfExists().
		Exec(ctx)
	if util.CheckErr(err) {
		return sqlDeleteCar, err
	}
	sqlDeleteShop, err := op.db.NewDropTable().
		Model((*model.Shop)(nil)).
		IfExists().
		Exec(ctx)
	if util.CheckErr(err) {
		return sqlDeleteShop, err
	}
	sqlDeleteShopService, err := op.db.NewDropTable().
		Model((*model.ShopService)(nil)).
		IfExists().
		Exec(ctx)
	if util.CheckErr(err) {
		return sqlDeleteShopService, err
	}
	sqlDeleteService, err := op.db.NewDropTable().
		Model((*model.Service)(nil)).
		IfExists().
		Exec(ctx)
	if util.CheckErr(err) {
		return sqlDeleteService, err
	}
	sqlDeleteTicketService, err := op.db.NewDropTable().
		Model((*model.TicketService)(nil)).
		IfExists().
		Exec(ctx)
	if util.CheckErr(err) {
		return sqlDeleteTicketService, err
	}
	sqlDeleteActiveTicket, err := op.db.NewDropTable().
		Model((*model.ActiveTicket)(nil)).
		IfExists().
		Exec(ctx)
	if util.CheckErr(err) {
		return sqlDeleteActiveTicket, err
	}
	sqlDeleteTicket, err := op.db.NewDropTable().
		Model((*model.Ticket)(nil)).
		IfExists().
		Exec(ctx)
	if util.CheckErr(err) {
		return sqlDeleteTicket, err
	}
	return sqlDeleteTicket, err
}

func (op *SQLop) CreateAllTables(ctx context.Context) (sql.Result, error) {
	sqlResultCus, err := op.db.NewCreateTable().
		Model((*model.Customer)(nil)).
		Exec(ctx)
	if util.CheckErr(err) {
		return sqlResultCus, err
	}
	sqlResultCar, err := op.db.NewCreateTable().
		Model((*model.Car)(nil)).
		Exec(ctx)
	if util.CheckErr(err) {
		return sqlResultCar, err
	}
	sqlResultShop, err := op.db.NewCreateTable().
		Model((*model.Shop)(nil)).
		Exec(ctx)
	if util.CheckErr(err) {
		return sqlResultShop, err
	}
	sqlResultShopService, err := op.db.NewCreateTable().
		Model((*model.ShopService)(nil)).
		Exec(ctx)
	if util.CheckErr(err) {
		return sqlResultShopService, err
	}
	sqlResultService, err := op.db.NewCreateTable().
		Model((*model.Service)(nil)).
		Exec(ctx)
	if util.CheckErr(err) {
		return sqlResultService, err
	}
	sqlResultTicketService, err := op.db.NewCreateTable().
		Model((*model.TicketService)(nil)).
		Exec(ctx)
	if util.CheckErr(err) {
		return sqlResultTicketService, err
	}
	sqlResultActiveTicket, err := op.db.NewCreateTable().
		Model((*model.ActiveTicket)(nil)).
		Exec(ctx)
	if util.CheckErr(err) {
		return sqlResultActiveTicket, err
	}
	sqlResultTicket, err := op.db.NewCreateTable().
		Model((*model.Ticket)(nil)).
		Exec(ctx)
	if util.CheckErr(err) {
		return sqlResultTicket, err
	}

	sqlIndexResult, err := op.db.NewCreateIndex().
		Model((*model.Ticket)(nil)).
		Index("ticket_index_0").
		Column("customerId", "carId").
		Exec(ctx)
	if util.CheckErr(err) {
		return sqlIndexResult, err
	}

	return sqlIndexResult, err
}

func (op *SQLop) CustomerCreate(ctx context.Context, cusInput *model.CustomerCreateInput) (*model.Customer, error) {
	newID := uuid.New().String()
	cusToBeAdd := model.Customer{
		ID:    newID,
		FName: cusInput.FName,
		Email: cusInput.Email,
	}
	_, err := op.db.NewInsert().Model(&cusToBeAdd).Exec(ctx)
	return &cusToBeAdd, err
}

func (op *SQLop) CustomerList(ctx context.Context) ([]*model.Customer, error) {
	customer := new([]*model.Customer)
	err := op.db.NewSelect().Model(customer).Scan(ctx, customer)
	return *customer, err
}
