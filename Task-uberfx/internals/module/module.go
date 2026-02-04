package module

import (
	"github.com/roshith/uber-fx/internals/database"
	"go.uber.org/fx"
)

var Psqlmodule = fx.Module(
	"database",
	fx.Provide(database.PostgresDB),
)

var MySQLModule = fx.Module(
	"database",
	fx.Provide(database.MysqlConnection),
)
