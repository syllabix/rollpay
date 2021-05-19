package organization

import (
	"context"

	"github.com/syllabix/rollpay/backend/api/model"
	api "github.com/syllabix/rollpay/backend/api/rest/operation/organization"
	"github.com/syllabix/rollpay/backend/common/id"
	"github.com/syllabix/rollpay/backend/common/media"
	db "github.com/syllabix/rollpay/backend/datastore/model"
	"github.com/syllabix/rollpay/backend/datastore/organization"
)

type Service interface {
	Get(ctx context.Context, id string) (model.Organization, error)
	GetAll(ctx context.Context) ([]*model.Organization, error)
	Create(api.CreateOrganizationV1Params) (model.Organization, error)
	Update(api.UpdateOrganizationByIDV1Params) (model.Organization, error)
	Delete(ctx context.Context, id string) error
}

type service struct {
	store organization.Store
}

func (s service) Get(ctx context.Context, userID string) (model.Organization, error) {
	uID, err := id.ToInternal(userID)
	if err != nil {
		return failure(err)
	}

	user, err := s.store.GetByID(ctx, uID)
	if err != nil {
		return failure(err)
	}

	return asModel(user), nil
}

func (s service) GetAll(ctx context.Context) ([]*model.Organization, error) {
	orgs, err := s.store.GetAll(ctx)
	if err != nil {
		return nil, mapErr(err)
	}

	return asModels(orgs), nil
}

func (s service) Create(params api.CreateOrganizationV1Params) (model.Organization, error) {
	newOrg, err := asNewOrg(params)
	if err != nil {
		return failure(err)
	}

	newOrg, err = s.store.Create(params.HTTPRequest.Context(), newOrg)
	if err != nil {
		return failure(err)
	}

	return asModel(newOrg), nil
}

func (s service) Update(params api.UpdateOrganizationByIDV1Params) (model.Organization, error) {
	orgID, err := id.ToInternal(params.ID)
	if err != nil {
		return failure(err)
	}

	update := db.Organization{ID: orgID}

	if params.Logo != nil {
		img, err := media.Process(params.HTTPRequest.Context(), params.Logo)
		if err != nil {
			return failure(err)
		}
		update.Logo = img
	}

	if params.Name != nil {
		update.Name = *params.Name
	}

	update, err = s.store.Update(params.HTTPRequest.Context(), update)
	if err != nil {
		return failure(err)
	}

	return asModel(update), nil
}

func (s service) Delete(ctx context.Context, orgID string) error {
	oID, err := id.ToInternal(orgID)
	if err != nil {
		return mapErr(err)
	}

	err = s.store.Delete(ctx, oID)
	if err != nil {
		return mapErr(err)
	}

	return nil
}
