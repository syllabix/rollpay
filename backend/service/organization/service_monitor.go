package organization

import (
	"context"

	"github.com/syllabix/rollpay/backend/api/model"
	"github.com/syllabix/rollpay/backend/datastore/organization"

	api "github.com/syllabix/rollpay/backend/api/rest/operation/organization"
	"go.uber.org/zap"
)

// a monitor wraps and provides observability for a Service
type monitor struct {
	log *zap.Logger
	srv Service
}

func (s monitor) Get(ctx context.Context, orgID string) (model.Organization, error) {
	result, err := s.srv.Get(ctx, orgID)
	if err != nil {
		s.log.Error("failed to get an organization by id",
			zap.Error(err),
			zap.String("requested-id", orgID))
	}
	return result, err
}

func (s monitor) GetAll(ctx context.Context) ([]*model.Organization, error) {
	result, err := s.srv.GetAll(ctx)
	if err != nil {
		s.log.Error("failed to get all organizations",
			zap.Error(err))
	}

	return result, err
}

func (s monitor) Create(params api.CreateOrganizationV1Params) (model.Organization, error) {
	result, err := s.srv.Create(params)
	if err != nil {
		s.log.Error("failed to create a new organization",
			zap.Error(err))
	}
	return result, err
}

func (s monitor) Update(params api.UpdateOrganizationByIDV1Params) (model.Organization, error) {
	result, err := s.srv.Update(params)
	if err != nil {
		s.log.Error("failed to update an organization",
			zap.String("org-id", params.ID),
			zap.Error(err))
	}
	return result, err
}

func (s monitor) Delete(ctx context.Context, orgID string) error {
	err := s.srv.Delete(ctx, orgID)
	if err != nil {
		s.log.Error("failed to delete an organization",
			zap.Error(err))
	}
	return err
}

func NewService(store organization.Store, log *zap.Logger) Service {
	return monitor{
		log: log,
		srv: service{store},
	}
}
