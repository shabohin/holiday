package repositories

import (
	"go.uber.org/fx"
)

var FXModule = fx.Options(
	fx.Provide(NewPostgresEventRepository),
)
