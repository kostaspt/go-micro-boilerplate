// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"context"
	"github.com/go-kratos/kratos/v2"
	"github.com/kostaspt/go-tiny/config"
	"github.com/kostaspt/go-tiny/internal/scheduler"
	"github.com/kostaspt/go-tiny/internal/scheduler/job"
	"github.com/kostaspt/go-tiny/internal/scheduler/registrar"
)

// Injectors from wire.go:

func initApp(contextContext context.Context, configConfig *config.Config) (*kratos.App, func(), error) {
	dummy := job.NewDummy(contextContext)
	entries := registrar.NewEntries(dummy)
	registrarRegistrar := registrar.New(entries)
	schedulerScheduler, err := scheduler.New(registrarRegistrar)
	if err != nil {
		return nil, nil, err
	}
	app, cleanup, err := newApp(schedulerScheduler)
	if err != nil {
		return nil, nil, err
	}
	return app, func() {
		cleanup()
	}, nil
}
