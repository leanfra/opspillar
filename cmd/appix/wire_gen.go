// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"appix/internal/biz"
	"appix/internal/conf"
	"appix/internal/data"
	"appix/internal/server"
	"appix/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	greeterRepo := data.NewGreeterRepo(dataData, logger)
	greeterUsecase := biz.NewGreeterUsecase(greeterRepo, logger)
	greeterService := service.NewGreeterService(greeterUsecase)
	tagsRepo, err := data.NewTagsRepoImpl(dataData, logger)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	tagsUsecase := biz.NewTagsUsecase(tagsRepo, logger)
	tagsService := service.NewTagsService(tagsUsecase, logger)
	featuresRepo, err := data.NewFeaturesRepoImpl(dataData, logger)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	featuresUsecase := biz.NewFeaturesUsecase(featuresRepo, logger)
	featuresService := service.NewFeaturesService(featuresUsecase, logger)
	teamsRepo, err := data.NewTeamsRepoImpl(dataData, logger)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	teamsUsecase := biz.NewTeamsUsecase(teamsRepo, logger)
	teamsService := service.NewTeamsService(teamsUsecase, logger)
	productsRepo, err := data.NewProductsRepoImpl(dataData, logger)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	productsUsecase := biz.NewProductsUsecase(productsRepo, logger)
	productsService := service.NewProductsService(productsUsecase, logger)
	envsRepo, err := data.NewEnvsRepoImpl(dataData, logger)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	envsUsecase := biz.NewEnvsUsecase(envsRepo, logger)
	envsService := service.NewEnvsService(envsUsecase, logger)
	grpcServer := server.NewGRPCServer(confServer, greeterService, tagsService, featuresService, teamsService, productsService, envsService, logger)
	httpServer := server.NewHTTPServer(confServer, greeterService, tagsService, featuresService, teamsService, productsService, envsService, logger)
	app := newApp(logger, grpcServer, httpServer)
	return app, func() {
		cleanup()
	}, nil
}
