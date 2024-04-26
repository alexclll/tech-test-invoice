package provider

import (
	"test-tech-invoice/src/framework/adapter/http"
	"test-tech-invoice/src/user/adapter/http/handler"
	"test-tech-invoice/src/user/adapter/mainStorage"
	"test-tech-invoice/src/user/useCase/getUsers"

	"go.uber.org/fx"
)

func GetProviders() []interface{} {
	return []interface{}{
		fx.Annotate(
			mainStorage.NewGetUsersRepository,
			fx.As(new(getUsers.UserRepository)),
		),
		getUsers.NewService,
		AsRoute(handler.NewGetUsersHandler),
	}
}

func AsRoute(handler any) any {
	return fx.Annotate(
		handler,
		fx.As(new(http.Route)),
		fx.ResultTags(`group:"routes"`),
	)
}
