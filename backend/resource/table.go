package resource

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"

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

func ValidateID(sIn string) error {
	if len(sIn) < 1 {
		return errors.New("ID must not be empty")
	}
	return nil
}

func NewDBOperator(inDSN string) (*SQLop, error) {
	db, err := sql.Open("postgres", inDSN)
	bunDB := bun.NewDB(db, pgdialect.New())
	bunDB.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
		bundebug.FromEnv("BUNDEBUG"),
	))
	util.CheckErr(err)
	err = db.Ping()
	util.CheckErr(err)
	return &SQLop{
		cusModel: new(model.Customer),
		db:       bunDB,
	}, err
}

// TABLE
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
