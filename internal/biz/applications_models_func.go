package biz

import (
	"fmt"
	"opspillar/internal/data/repo"
)

func (m *Application) Validate(isNew bool) error {
	if len(m.Name) == 0 {
		return fmt.Errorf("InvalidNameValue")
	}
	if e := ValidateName(m.Name); e != nil {
		return e
	}
	if !isNew {
		if m.Id == 0 {
			return fmt.Errorf("InvalidId")
		}
	}
	if len(m.Name) == 0 {
		return fmt.Errorf("InvalidNameValue")
	}
	if m.OwnerId == 0 {
		return fmt.Errorf("InvalidOwnerIdValue")
	}
	if m.ProductId <= 0 {
		return fmt.Errorf("InvalidProductId")
	}
	if m.TeamId <= 0 {
		return fmt.Errorf("InvalidTeamId")
	}
	return nil
}

func (m *ListApplicationsFilter) Validate() error {
	if m == nil {
		return nil
	}
	if len(m.Ids) > MaxFilterValues ||
		len(m.Names) > MaxFilterValues ||
		len(m.ProductsId) > MaxFilterValues ||
		len(m.TeamsId) > MaxFilterValues ||
		len(m.FeaturesId) > MaxFilterValues ||
		len(m.HostgroupsId) > MaxFilterValues ||
		len(m.TagsId) > MaxFilterValues {

		return ErrFilterValuesExceedMax
	}

	if m.PageSize == 0 || m.PageSize > MaxPageSize {
		return ErrFilterInvalidPagesize
	}
	if m.Page == 0 {
		return ErrFilterInvalidPage
	}

	if m.IsStateful != IsStatefulFalse && m.IsStateful != IsStatefulTrue && m.IsStateful != IsStatefulNone {
		return fmt.Errorf("InvalidIsStateful")
	}

	return nil
}

func DefaultApplicationFilter() *ListApplicationsFilter {
	return &ListApplicationsFilter{
		Page:       1,
		PageSize:   DefaultPageSize,
		IsStateful: IsStatefulNone,
	}
}

func ToDBApplication(app *Application) (*repo.Application, error) {
	if app == nil {
		return nil, nil
	}
	return &repo.Application{
		ChangeInfo:  repo.ChangeInfo{},
		Id:          app.Id,
		Name:        app.Name,
		Description: app.Description,
		OwnerId:     app.OwnerId,
		IsStateful:  app.IsStateful,
		ProductId:   app.ProductId,
		TeamId:      app.TeamId,
	}, nil
}

func ToDBApplications(apps []*Application) ([]*repo.Application, error) {
	db_apps := make([]*repo.Application, len(apps))
	for i, a := range apps {
		db_app, e := ToDBApplication(a)
		if e != nil {
			return nil, e
		}
		db_apps[i] = db_app
	}
	return db_apps, nil
}

func ToBizApplication(t *repo.Application) (*Application, error) {
	if t == nil {
		return nil, nil
	}
	return &Application{
		Id:          t.Id,
		Name:        t.Name,
		Description: t.Description,
		OwnerId:     t.OwnerId,
		IsStateful:  t.IsStateful,
		ProductId:   t.ProductId,
		TeamId:      t.TeamId,
		ChangeInfo: ChangeInfo{
			CreatedAt: t.CreatedAt,
			UpdatedAt: t.UpdatedAt,
			CreatedBy: t.CreatedBy,
			UpdatedBy: t.UpdatedBy,
		},
	}, nil
}

func ToBizApplications(es []*repo.Application) ([]*Application, error) {
	apps := make([]*Application, len(es))
	for i, e := range es {
		app, err := ToBizApplication(e)
		if err != nil {
			return nil, err
		}
		apps[i] = app
	}
	return apps, nil
}

func ToDBApplicationsFilter(filter *ListApplicationsFilter) *repo.ApplicationsFilter {
	if filter == nil {
		return nil
	}
	return &repo.ApplicationsFilter{
		Ids:        filter.Ids,
		Names:      filter.Names,
		ProductsId: filter.ProductsId,
		TeamsId:    filter.TeamsId,
		IsStateful: filter.IsStateful,
		Page:       filter.Page,
		PageSize:   filter.PageSize,
	}
}
