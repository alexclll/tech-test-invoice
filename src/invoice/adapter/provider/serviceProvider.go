package provider

import (
	"test-tech-invoice/src/framework/adapter/http"
	"test-tech-invoice/src/invoice/adapter/http/handler"
	"test-tech-invoice/src/invoice/adapter/mainStorage"
	"test-tech-invoice/src/invoice/business"
	"test-tech-invoice/src/invoice/useCase/createInvoice"
	"test-tech-invoice/src/invoice/useCase/validateInvoice"

	"go.uber.org/fx"
)

func GetProviders() []interface{} {
	return []interface{}{
		fx.Annotate(
			mainStorage.NewInvoiceRepository,
			fx.As(new(business.InvoiceRepository)),
		),
		fx.Annotate(
			mainStorage.NewUserRepository,
			fx.As(new(business.UserRepository)),
		),

		createInvoice.NewService,
		asRoute(handler.NewCreateInvoiceHandler),

		validateInvoice.NewService,
		asRoute(handler.NewValidateInvoiceHandler),
	}
}

func asRoute(handler any) any {
	return fx.Annotate(
		handler,
		fx.As(new(http.Route)),
		fx.ResultTags(`group:"routes"`),
	)
}
