package containers

import (
	"context"

	"github.com/shabohin/holiday.git/pkg/log"
	"go.uber.org/fx/fxevent"

	"github.com/shabohin/holiday.git/internal/interceptors"
	"github.com/shabohin/holiday.git/internal/repositories"
	"github.com/shabohin/holiday.git/internal/usecases"

	"github.com/shabohin/holiday.git/pkg/clock"

	"github.com/shabohin/holiday.git/internal/configs"
	"go.uber.org/fx"
)

var appModule = fx.Options(
	fx.Provide(
		context.Background,
		clock.NewRealClock,
		func(config *configs.Config) (log.Logger, error) {
			return log.NewLog(config.LogLevel)
		},
	),
	configs.FXModule,
	repositories.FXModule,
	usecases.FXModule,
	interceptors.FXModule,
)

func NewHoliday(config string) *fx.App {
	app := fx.New(
		fx.Provide(func() string { return config }),
		appModule,
		fx.WithLogger(
			func(logger log.Logger) fxevent.Logger {
				return logger
			},
		),
	)
	return app
}
