package membership

import (
	"context"

	"github.com/syllabix/rollpay/backend/api/model"
	"github.com/syllabix/rollpay/backend/api/rest/operation/organization"
	"github.com/syllabix/rollpay/backend/datastore/membership"
	"go.uber.org/zap"
)

// a monitor wraps and provides observability for a Service
type monitor struct {
	log *zap.Logger
	srv Service
}

func (s monitor) GetAllByOrgID(ctx context.Context, orgID string) (list model.MemberList, err error) {
	result, err := s.srv.GetAllByOrgID(ctx, orgID)
	if err != nil {
		s.log.Error("failed to get an organization member list",
			zap.Error(err),
			zap.String("requested-id", orgID))
	}
	return result, err
}

func (s monitor) AddMember(params organization.AddOrgMembersV1Params) (model.OrganizationMember, error) {
	result, err := s.srv.AddMember(params)
	if err != nil {
		s.log.Error("failed to add member to an organization",
			zap.String("org-id", params.ID),
			zap.Error(err))
	}

	return result, err
}

func NewService(store membership.Store, log *zap.Logger) Service {
	return monitor{
		log: log,
		srv: service{store},
	}
}
