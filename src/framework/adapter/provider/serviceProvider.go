package provider

import (
	"test-tech-invoice/src/framework/adapter/http"
	"test-tech-invoice/src/framework/adapter/mainStorage"

	"go.uber.org/fx"
)

func GetProviders() []interface{} {
	return []interface{}{
		http.NewServer,
		fx.Annotate(
			http.NewServeMux,
			fx.ParamTags(`group:"routes"`),
		),
		mainStorage.NewConnection,
	}
}
