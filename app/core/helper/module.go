package helper

import (
	"ubereats/app/core/helper/database"

	"go.uber.org/fx"
)

var Module = fx.Module(
	"helper",
	fx.Provide(
		database.NewMysql,
	),
)
