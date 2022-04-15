// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"context"
	"github.com/go-kratos/kratos/v2"
	"github.com/kostaspt/go-tiny/config"
	"github.com/kostaspt/go-tiny/internal/http/handler"
	"github.com/kostaspt/go-tiny/internal/http/server"
)

// Injectors from wire.go:

func initApp(contextContext context.Context, configConfig *config.Config) (*kratos.App, func(), error) {
	handlerHandler := handler.NewHandler(configConfig)
	middleware := server.NewMiddleware(configConfig)
	httpServer := server.NewServer(configConfig, handlerHandler, middleware)
	app := newApp(httpServer)
	return app, func() {
	}, nil
}