package main

import (
	"test-tech-invoice/src/framework/adapter/dotenv"
	"test-tech-invoice/src/framework/adapter/http"
	frameworkProvider "test-tech-invoice/src/framework/adapter/provider"
	invoiceProvider "test-tech-invoice/src/invoice/adapter/provider"
	userProvider "test-tech-invoice/src/user/adapter/provider"

	"go.uber.org/fx"
)

func main() {
	dotenv.Load()

	fx.New(
		fx.Provide(flattenProviders()...),
		fx.Invoke(http.Invoke),
	).Run()
}

func flattenProviders() []interface{} {
	providers := getProviders()

	var allProviders []interface{}
	for _, domainProviders := range providers {
		allProviders = append(allProviders, domainProviders...)
	}

	return allProviders
}

func getProviders() [][]interface{} {
	return [][]interface{}{
		frameworkProvider.GetProviders(),
		invoiceProvider.GetProviders(),
		userProvider.GetProviders(),
	}
}
