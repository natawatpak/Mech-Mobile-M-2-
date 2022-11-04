package resource

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
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

func (op *SQLop) CreateCountryTable(ctx context.Context) (sql.Result, error) {
	sqlResult, err := op.db.NewCreateTable().
		Model((*model.Country)(nil)).
		Exec(ctx)
	return sqlResult, err
}

func (op *SQLop) CountryCreate(ctx context.Context, countryInput *model.CountryCreateInput) (*model.Country, error) {
	userToBeAdd := model.Country{
		Code:      countryInput.Code,
		Name:      countryInput.Name,
		Continent: countryInput.Continent,
	}
	_, err := op.db.NewInsert().Model(&userToBeAdd).Exec(ctx)
	return &userToBeAdd, err
}

func (op *SQLop) CreateUserTable(ctx context.Context) (sql.Result, error) {
	sqlResult, err := op.db.NewCreateTable().
		Model((*model.User)(nil)).
		ForeignKey(`("country_code") REFERENCES Country ("Code") ON DELETE CASCADE`).
		Exec(ctx)
	return sqlResult, err
}

func (op *SQLop) UserCreate(ctx context.Context, userInput *model.UserCreateInput) (*model.User, error) {
	newID := uuid.New().String()
	userToBeAdd := model.User{
		ID:          newID,
		FName:       userInput.FName,
		Email:       userInput.Email,
		Gender:      userInput.Gender,
		DoB:         userInput.DoB,
		CountryCode: userInput.CountryCode,
		CreateTime:  userInput.CreateTime,
	}
	_, err := op.db.NewInsert().Model(&userToBeAdd).Exec(ctx)
	return &userToBeAdd, err
}

func (op *SQLop) UserList(ctx context.Context) ([]*model.User, error) {
	user := new([]*model.User)
	err := op.db.NewSelect().Model(user).Scan(ctx, user)
	return *user, err
}
