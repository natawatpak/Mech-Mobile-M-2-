package resource

import (
	"github.com/natawatpak/Mech-Mobile-M-2-/backend/graph/model"
	"github.com/uptrace/bun"
)

type SQLop struct {
	db        *bun.DB
	userModel *model.User
}
