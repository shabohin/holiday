package usecases

import (
	"go.uber.org/fx"
)

var FXModule = fx.Options(
	fx.Provide(NewEventUseCase),
)
