package repo

import "context"

const FeatureTable = "features"

type Feature struct {
	Id          uint32 `gorm:"primaryKey;autoIncrement"`
	Name        string `gorm:"type:varchar(255);index:idx_feature_name_value,unique"`
	Value       string `gorm:"type:varchar(255);index:idx_feature_name_value"`
	Description string `gorm:"type:text"`
}

type FeaturesFilter struct {
	Page     uint32
	PageSize uint32
	Ids      []uint32
	Names    []string
	Kvs      []string
}

func (f *FeaturesFilter) GetIds() []uint32 {
	return f.Ids
}

type FeaturesRepo interface {
	RequireCounter
	CreateFeatures(ctx context.Context, tx TX, features []*Feature) error
	UpdateFeatures(ctx context.Context, tx TX, features []*Feature) error
	DeleteFeatures(ctx context.Context, tx TX, ids []uint32) error
	GetFeatures(ctx context.Context, id uint32) (*Feature, error)
	ListFeatures(ctx context.Context, tx TX, filter *FeaturesFilter) ([]*Feature, error)
	CountFeatures(ctx context.Context, tx TX, filter CountFilter) (int64, error)
}
