package resource

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/google/uuid"
	"github.com/natawatpak/Mech-Mobile-M-2-/backend/graph/model"
	"github.com/natawatpak/Mech-Mobile-M-2-/backend/util"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/extra/bundebug"
)

func PrintIfErrorExist(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func NewDBOperator(inDSN string) (*SQLop, error) {
	db, err := sql.Open("postgres", inDSN)
	bunDB := bun.NewDB(db, pgdialect.New())
	bunDB.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
		bundebug.FromEnv("BUNDEBUG"),
	))
	err = db.Ping()
	util.CheckErr(err)
	return &SQLop{
		cusModel: new(model.Customer),
		db:       bunDB,
	}, err
}

func (op *SQLop) DropTable(ctx context.Context) (bool, error) {
	//ticket service
	_, err := op.db.NewDropTable().
		Model((*model.TicketService)(nil)).
		IfExists().
		Exec(ctx)
	if util.CheckErr(err) {
		return false, err
	}
	//active ticket
	_, err = op.db.NewDropTable().
		Model((*model.ActiveTicket)(nil)).
		IfExists().
		Exec(ctx)
	if util.CheckErr(err) {
		return false, err
	}
	//ticket
	_, err = op.db.NewDropTable().
		Model((*model.Ticket)(nil)).
		IfExists().
		Exec(ctx)
	if util.CheckErr(err) {
		return false, err
	}
	//shop Service
	_, err = op.db.NewDropTable().
		Model((*model.ShopService)(nil)).
		IfExists().
		Exec(ctx)
	if util.CheckErr(err) {
		return false, err
	}
	//service
	_, err = op.db.NewDropTable().
		Model((*model.Service)(nil)).
		IfExists().
		Exec(ctx)
	if util.CheckErr(err) {
		return false, err
	}
	//shop
	_, err = op.db.NewDropTable().
		Model((*model.Shop)(nil)).
		IfExists().
		Exec(ctx)
	if util.CheckErr(err) {
		return false, err
	}
	//car
	_, err = op.db.NewDropTable().
		Model((*model.Car)(nil)).
		IfExists().
		Exec(ctx)
	if util.CheckErr(err) {
		return false, err
	}
	//customer
	_, err = op.db.NewDropTable().
		Model((*model.Customer)(nil)).
		IfExists().
		Exec(ctx)
	if util.CheckErr(err) {
		return false, err
	}
	return true, nil
}

func (op *SQLop) CreateTables(ctx context.Context) (bool, error) {
	// customer
	_, err := op.db.NewCreateTable().
		Model((*model.Customer)(nil)).
		Exec(ctx)
	if util.CheckErr(err) {
		return false, err
	}
	//car
	_, err = op.db.NewCreateTable().
		Model((*model.Car)(nil)).
		ForeignKey(`("owner_id") REFERENCES "customers" ("id") ON DELETE CASCADE`).
		Exec(ctx)
	if util.CheckErr(err) {
		return false, err
	}
	//shop
	_, err = op.db.NewCreateTable().
		Model((*model.Shop)(nil)).
		Exec(ctx)
	if util.CheckErr(err) {
		return false, err
	}
	//Service
	_, err = op.db.NewCreateTable().
		Model((*model.Service)(nil)).
		Exec(ctx)
	if util.CheckErr(err) {
		return false, err
	}
	//shop service
	_, err = op.db.NewCreateTable().
		Model((*model.ShopService)(nil)).
		ForeignKey(`("shop_id") REFERENCES "shops" ("id") ON DELETE CASCADE`).
		ForeignKey(`("service_id") REFERENCES "services" ("id") ON DELETE CASCADE`).
		Exec(ctx)
	if util.CheckErr(err) {
		return false, err
	}
	//Ticket
	_, err = op.db.NewCreateTable().
		Model((*model.Ticket)(nil)).
		ForeignKey(`("customer_id") REFERENCES "customers" ("id") ON DELETE CASCADE`).
		ForeignKey(`("car_id") REFERENCES "cars" ("id") ON DELETE CASCADE`).
		ForeignKey(`("shop_id") REFERENCES "shops" ("id") ON DELETE CASCADE`).
		Exec(ctx)
	if util.CheckErr(err) {
		return false, err
	}
	//Active Ticket
	_, err = op.db.NewCreateTable().
		Model((*model.ActiveTicket)(nil)).
		ForeignKey(`("id") REFERENCES "tickets" ("id") ON DELETE CASCADE`).
		ForeignKey(`("car_id") REFERENCES "cars" ("id") ON DELETE CASCADE`).
		ForeignKey(`("customer_id") REFERENCES "customers" ("id") ON DELETE CASCADE`).
		ForeignKey(`("shop_id") REFERENCES "shops" ("id") ON DELETE CASCADE`).
		Exec(ctx)
	if util.CheckErr(err) {
		return false, err
	}
	//Ticket Service
	_, err = op.db.NewCreateTable().
		Model((*model.TicketService)(nil)).
		ForeignKey(`("ticket_id") REFERENCES "tickets" ("id") ON DELETE CASCADE`).
		ForeignKey(`("service_id") REFERENCES "services" ("id") ON DELETE CASCADE`).
		ForeignKey(`("ticket_id") REFERENCES "active_tickets" ("id") ON DELETE CASCADE`).
		Exec(ctx)
	if util.CheckErr(err) {
		return false, err
	}
	_, err = op.db.NewCreateIndex().
		Model((*model.Ticket)(nil)).
		Index("ticket_index_0").
		Column("customer_id", "car_id").
		Exec(ctx)
	if util.CheckErr(err) {
		return false, err
	}

	return true, err
}

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

func (op *SQLop) CustomerUpdateMulti(ctx context.Context, updateInput model.Customer) ([]*model.Customer, error) {
	_, err := op.db.NewUpdate().Model(&updateInput).Where("id = ?", updateInput.ID).Exec(ctx)
	resultCustomer, _ := op.CustomerFindByID(ctx, updateInput.ID)
	return resultCustomer, err
}

func (op *SQLop) CustomerDelete(ctx context.Context, ID string) (*model.Customer, error) {
	resultCustomer, err := op.CustomerFindByID(ctx, ID)
	PrintIfErrorExist(err)
	_, err = op.db.NewDelete().Model(op.cusModel).Where("id = ?", ID).Exec(ctx)
	targetID := -1
	for i, v := range resultCustomer {
		if ID == v.ID {
			targetID = i
		}
	}
	return resultCustomer[targetID], err
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

func (op *SQLop) CustomerFindByID(ctx context.Context, ID string) ([]*model.Customer, error) {
	arrModel := new([]*model.Customer)
	err := op.db.NewSelect().Model(op.cusModel).Where("id = ?", ID).Scan(ctx, arrModel)
	return *arrModel, err
}

func (op *SQLop) CustomerList(ctx context.Context) ([]*model.Customer, error) {
	customer := new([]*model.Customer)
	err := op.db.NewSelect().Model(op.cusModel).Scan(ctx, customer)
	return *customer, err
}
