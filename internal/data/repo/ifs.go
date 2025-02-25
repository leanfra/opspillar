package repo

import (
	"context"
	"errors"
)

type TX interface {
	GetDB() interface{}
	Error(err error) bool
}

type TxManager interface {
	RunInTX(fn func(tx TX) error) error
}

type CountFilter interface {
	GetIds() []uint32
}

type RequireType string

const (
	RequireTeam       RequireType = "team"
	RequireTag        RequireType = "tag"
	RequireProduct    RequireType = "product"
	RequireHostgroup  RequireType = "hostgroup"
	RequireFeature    RequireType = "feature"
	RequireEnv        RequireType = "env"
	RequireDatacenter RequireType = "datacenter"
	RequireCluster    RequireType = "cluster"
	RequireApp        RequireType = "app"
	RequireUser       RequireType = "user"
)

var ErrorRequireIds = errors.New("invalid require ids or types")

type RequireCounter interface {
	CountRequire(ctx context.Context, tx TX, need RequireType, ids []uint32) (int64, error)
}

type AppTagsRepo interface {
	RequireCounter
	CreateAppTags(ctx context.Context, tx TX, apps []*AppTag) error
	UpdateAppTags(ctx context.Context, tx TX, apps []*AppTag) error
	DeleteAppTags(ctx context.Context, tx TX, ids []uint32) error
	DeleteAppTagsByAppId(ctx context.Context, tx TX, appids []uint32) error
	ListAppTags(ctx context.Context, tx TX, filter *AppTagsFilter) ([]*AppTag, error)
	//CountAppTags(ctx context.Context, tx TX, filter *AppTagsFilter) (int64, error)
}

type AppFeaturesRepo interface {
	RequireCounter
	CreateAppFeatures(ctx context.Context, tx TX, apps []*AppFeature) error
	UpdateAppFeatures(ctx context.Context, tx TX, apps []*AppFeature) error
	DeleteAppFeatures(ctx context.Context, tx TX, ids []uint32) error
	DeleteAppFeaturesByAppId(ctx context.Context, tx TX, appids []uint32) error
	ListAppFeatures(ctx context.Context, tx TX, filter *AppFeaturesFilter) ([]*AppFeature, error)
}

type AppHostgroupsRepo interface {
	RequireCounter
	CreateAppHostgroups(ctx context.Context, tx TX, apps []*AppHostgroup) error
	UpdateAppHostgroups(ctx context.Context, tx TX, apps []*AppHostgroup) error
	DeleteAppHostgroups(ctx context.Context, tx TX, ids []uint32) error
	DeleteAppHostgroupsByAppId(ctx context.Context, tx TX, appids []uint32) error
	ListAppHostgroups(ctx context.Context, tx TX, filter *AppHostgroupsFilter) ([]*AppHostgroup, error)
}

type HostgroupTeamsRepo interface {
	RequireCounter
	CreateHostgroupTeams(ctx context.Context, tx TX, hfs []*HostgroupTeam) error
	UpdateHostgroupTeams(ctx context.Context, tx TX, hfs []*HostgroupTeam) error
	DeleteHostgroupTeams(ctx context.Context, tx TX, ids []uint32) error
	ListHostgroupTeams(ctx context.Context, tx TX,
		filter *HostgroupTeamsFilter) ([]*HostgroupTeam, error)
}
type HostgroupProductsRepo interface {
	RequireCounter
	CreateHostgroupProducts(ctx context.Context, tx TX, hfs []*HostgroupProduct) error
	UpdateHostgroupProducts(ctx context.Context, tx TX, hfs []*HostgroupProduct) error
	DeleteHostgroupProducts(ctx context.Context, tx TX, ids []uint32) error
	ListHostgroupProducts(ctx context.Context, tx TX,
		filter *HostgroupProductsFilter) ([]*HostgroupProduct, error)
}

type HostgroupTagsRepo interface {
	RequireCounter
	CreateHostgroupTags(ctx context.Context, tx TX, hfs []*HostgroupTag) error
	UpdateHostgroupTags(ctx context.Context, tx TX, hfs []*HostgroupTag) error
	DeleteHostgroupTags(ctx context.Context, tx TX, ids []uint32) error
	ListHostgroupTags(ctx context.Context, tx TX,
		filter *HostgroupTagsFilter) ([]*HostgroupTag, error)
}

type HostgroupFeaturesRepo interface {
	RequireCounter
	CreateHostgroupFeatures(ctx context.Context, tx TX, hfs []*HostgroupFeature) error
	UpdateHostgroupFeatures(ctx context.Context, tx TX, hfs []*HostgroupFeature) error
	DeleteHostgroupFeatures(ctx context.Context, tx TX, ids []uint32) error
	ListHostgroupFeatures(ctx context.Context, tx TX,
		filter *HostgroupFeaturesFilter) ([]*HostgroupFeature, error)
	ListHostgroupMatchFeatures(ctx context.Context, tx TX,
		filter *HostgroupMatchFeaturesFilter) ([]uint32, error)
}
