package main

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/roshith/uber-fx/internals/config"
	"github.com/roshith/uber-fx/internals/handlers"
	"github.com/roshith/uber-fx/internals/module"
	"github.com/roshith/uber-fx/internals/repository"
	"github.com/roshith/uber-fx/internals/service"
	"go.uber.org/fx"
)

func Newfiber() *fiber.App {
	return fiber.New()
}

func RegisterRoutes(app *fiber.App, h *handlers.UserHandlers) {
	h.HandlerRoutes(app)
}
func StartServer(lc fx.Lifecycle, cfg *config.Config, app *fiber.App) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go app.Listen(":" + cfg.APP_PORT)
			log.Println("Server running on port :", cfg.APP_PORT)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return app.Shutdown()
		},
	})
}

func main() {

	fx.New(
		fx.Provide(
			config.LoadConfig,
			Newfiber,
			repository.NewRepository,
			service.NewUserService,
			handlers.NewUserHandlers,
		),
		module.MySQLModule,
		fx.Invoke(
			RegisterRoutes,
			StartServer,
		),
	).Run()
}
