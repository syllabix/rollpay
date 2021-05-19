package membership

import (
	"context"

	"github.com/syllabix/rollpay/backend/api/model"
	"github.com/syllabix/rollpay/backend/api/rest/operation/organization"
	"github.com/syllabix/rollpay/backend/common/id"
	"github.com/syllabix/rollpay/backend/datastore/membership"
)

type Service interface {
	GetAllByOrgID(context.Context, string) (model.MemberList, error)
	AddMember(organization.AddOrgMembersV1Params) (model.OrganizationMember, error)
}

type service struct {
	store membership.Store
}

func (s service) GetAllByOrgID(ctx context.Context, orgID string) (list model.MemberList, err error) {
	oID, err := id.ToInternal(orgID)
	if err != nil {
		return list, mapErr(err)
	}

	members, err := s.store.GetAllByOrgID(ctx, oID)
	if err != nil {
		return list, mapErr(err)
	}

	return model.MemberList{
		Results: asModels(members),
	}, nil
}

func (s service) AddMember(params organization.AddOrgMembersV1Params) (model.OrganizationMember, error) {
	newMember, err := asNewMember(params)
	if err != nil {
		return failure(err)
	}

	newMember, err = s.store.AddMember(params.HTTPRequest.Context(), newMember)
	if err != nil {
		return failure(err)
	}

	return asModel(newMember), nil
}
