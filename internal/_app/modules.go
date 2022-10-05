package app

import (
	"context"

	"github.com/chaewonkong/go-clean-arch-template/internal/_app/controllers"

	"github.com/labstack/echo"
	"go.uber.org/fx"
)

var Modules = fx.Options(
	controllers.Modules,
	fx.Provide(NewServer),
	fx.Invoke(RegisterHooks),
)

func RegisterHooks(lifecycle fx.Lifecycle, e *echo.Echo) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go e.Start(":8080")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return e.Shutdown(ctx)
		},
	})
}
