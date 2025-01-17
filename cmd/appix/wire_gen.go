// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"appix/internal/biz"
	"appix/internal/conf"
	"appix/internal/data"
	"appix/internal/data/sqldb"
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
func wireApp(confServer *conf.Server, confData *conf.Data, admin *conf.Admin, authz *conf.Authz, logger log.Logger) (*kratos.App, func(), error) {
	dataGorm, cleanup, err := sqldb.NewDataGorm(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	tagsRepo, err := sqldb.NewTagsRepoGorm(dataGorm, logger)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	appTagsRepo, err := sqldb.NewAppTagsRepoGorm(dataGorm, logger)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	hostgroupTagsRepo, err := sqldb.NewHostgroupTagsRepoGorm(dataGorm, logger)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	txManager := sqldb.NewTxManagerGorm(dataGorm, logger)
	tagsUsecase := biz.NewTagsUsecase(tagsRepo, logger, appTagsRepo, hostgroupTagsRepo, txManager)
	tagsService := service.NewTagsService(tagsUsecase, logger)
	featuresRepo, err := sqldb.NewFeaturesRepoGorm(dataGorm, logger)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	authzRepo, err := sqldb.NewAuthzRepoGorm(authz, dataGorm, logger)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	hostgroupFeaturesRepo, err := sqldb.NewHostgroupFeaturesRepoGorm(dataGorm, logger)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	appFeaturesRepo, err := sqldb.NewAppFeaturesRepoGorm(dataGorm, logger)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	featuresUsecase := biz.NewFeaturesUsecase(featuresRepo, authzRepo, hostgroupFeaturesRepo, appFeaturesRepo, logger, txManager)
	featuresService := service.NewFeaturesService(featuresUsecase, logger)
	teamsRepo, err := sqldb.NewTeamsRepoGorm(dataGorm, logger)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	hostgroupsRepo, err := sqldb.NewHostgroupsRepoGorm(dataGorm, logger)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	hostgroupTeamsRepo, err := sqldb.NewHostgroupTeamsRepoGorm(dataGorm, logger)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	applicationsRepo, err := sqldb.NewApplicationsRepoGorm(dataGorm, logger)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	teamsUsecase := biz.NewTeamsUsecase(teamsRepo, authzRepo, hostgroupsRepo, hostgroupTeamsRepo, applicationsRepo, logger, txManager)
	teamsService := service.NewTeamsService(teamsUsecase, logger)
	productsRepo, err := sqldb.NewProductsRepoGorm(dataGorm, logger)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	hostgroupProductsRepo, err := sqldb.NewHostgroupProductsRepoGorm(dataGorm, logger)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	productsUsecase := biz.NewProductsUsecase(productsRepo, hostgroupsRepo, applicationsRepo, hostgroupProductsRepo, logger, txManager)
	productsService := service.NewProductsService(productsUsecase, logger)
	envsRepo, err := sqldb.NewEnvsRepoGorm(dataGorm, logger)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	envsUsecase := biz.NewEnvsUsecase(envsRepo, authzRepo, hostgroupsRepo, logger, txManager)
	envsService := service.NewEnvsService(envsUsecase, logger)
	clustersRepo, err := sqldb.NewClustersRepoGorm(dataGorm, logger)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	clustersUsecase := biz.NewClustersUsecase(clustersRepo, authzRepo, hostgroupsRepo, logger, txManager)
	clustersService := service.NewClustersService(clustersUsecase, logger)
	datacentersRepo, err := sqldb.NewDatacentersRepoGorm(dataGorm, logger)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	datacentersUsecase := biz.NewDatacentersUsecase(datacentersRepo, authzRepo, hostgroupsRepo, logger, txManager)
	datacentersService := service.NewDatacentersService(datacentersUsecase, logger)
	appHostgroupsRepo, err := sqldb.NewAppHostgroupsRepoGorm(dataGorm, logger)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	hostgroupsUsecase := biz.NewHostgroupsUsecase(hostgroupsRepo, hostgroupTeamsRepo, hostgroupProductsRepo, hostgroupTagsRepo, hostgroupFeaturesRepo, clustersRepo, datacentersRepo, envsRepo, featuresRepo, tagsRepo, teamsRepo, productsRepo, appHostgroupsRepo, logger, txManager)
	hostgroupsService := service.NewHostgroupsService(hostgroupsUsecase, logger)
	adminRepo, err := sqldb.NewAdminRepoGorm(dataGorm, logger)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	applicationsUsecase := biz.NewApplicationsUsecase(applicationsRepo, appTagsRepo, appFeaturesRepo, appHostgroupsRepo, productsRepo, teamsRepo, featuresRepo, tagsRepo, hostgroupsRepo, hostgroupFeaturesRepo, authzRepo, adminRepo, logger, txManager)
	applicationsService := service.NewApplicationsService(applicationsUsecase, logger)
	tokenRepo := data.NewJwtMemRepo(admin)
	adminUsecase := biz.NewAdminUsecase(admin, adminRepo, tokenRepo, authzRepo, teamsRepo, txManager, logger)
	adminService := service.NewAdminService(adminUsecase, logger)
	grpcServer := server.NewGRPCServer(confServer, admin, tagsService, featuresService, teamsService, productsService, envsService, clustersService, datacentersService, hostgroupsService, applicationsService, adminService, logger)
	httpServer := server.NewHTTPServer(confServer, admin, tagsService, featuresService, teamsService, productsService, envsService, clustersService, datacentersService, hostgroupsService, applicationsService, adminService, logger)
	app := newApp(logger, grpcServer, httpServer)
	return app, func() {
		cleanup()
	}, nil
}
