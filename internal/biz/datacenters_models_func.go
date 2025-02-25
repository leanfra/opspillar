package biz

import (
	"opspillar/internal/data/repo"
	"fmt"
)

func (f *Datacenter) Validate(isNew bool) error {
	if len(f.Name) == 0 {
		return fmt.Errorf("InvalidNameValue")
	}
	if !isNew {
		if f.Id <= 0 {
			return fmt.Errorf("InvalidId")
		}
	}
	if e := ValidateName(f.Name); e != nil {
		return e
	}
	return nil
}

func (lf *ListDatacentersFilter) Validate() error {
	if lf == nil {
		return nil
	}
	if len(lf.Ids) > MaxFilterValues || len(lf.Names) > MaxFilterValues {
		return ErrFilterValuesExceedMax
	}
	if lf.PageSize == 0 || lf.PageSize > MaxPageSize {
		return ErrFilterInvalidPagesize
	}
	if lf.Page == 0 {
		return ErrFilterInvalidPage
	}
	return nil
}

func DefaultDatacentersFilter() *ListDatacentersFilter {
	return &ListDatacentersFilter{
		Page:     1,
		PageSize: DefaultPageSize,
	}
}

func ToDBDatacenter(t *Datacenter) (*repo.Datacenter, error) {
	return &repo.Datacenter{
		ID:          t.Id,
		Name:        t.Name,
		Description: t.Description,
	}, nil
}

func ToDBDatacenters(es []*Datacenter) ([]*repo.Datacenter, error) {
	var clusters = make([]*repo.Datacenter, len(es))
	for i, f := range es {
		nf, err := ToDBDatacenter(f)
		if err != nil {
			return nil, err
		}
		clusters[i] = nf
	}
	return clusters, nil
}

func ToBizDatacenter(t *repo.Datacenter) (*Datacenter, error) {
	return &Datacenter{
		Id:          t.ID,
		Name:        t.Name,
		Description: t.Description,
	}, nil
}

func ToBizDatacenters(es []*repo.Datacenter) ([]*Datacenter, error) {
	var biz_clusters = make([]*Datacenter, len(es))
	for i, f := range es {
		biz_clusters[i] = &Datacenter{
			Id:          f.ID,
			Name:        f.Name,
			Description: f.Description,
		}
	}
	return biz_clusters, nil
}

func ToDBDatacentersFilter(filter *ListDatacentersFilter) *repo.DatacentersFilter {
	return &repo.DatacentersFilter{
		Ids:      filter.Ids,
		Names:    filter.Names,
		Page:     filter.Page,
		PageSize: filter.PageSize,
	}
}
