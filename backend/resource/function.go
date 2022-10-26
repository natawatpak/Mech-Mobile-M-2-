package resource

import (
	"database/sql"
	"fmt"

	"github.com/natawatpak/Mech-Mobile-M-2-/backend/graph/model"
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
		userModel: new(model.User),
		db:        db,
	}, err
}
