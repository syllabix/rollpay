package membership

import (
	"github.com/go-openapi/strfmt"
	api "github.com/syllabix/rollpay/backend/api/model"
	"github.com/syllabix/rollpay/backend/api/rest/operation/organization"
	"github.com/syllabix/rollpay/backend/common/id"
	db "github.com/syllabix/rollpay/backend/datastore/model"
)

func asModel(mem db.OrganizationMember) api.OrganizationMember {
	return api.OrganizationMember{
		ID:       id.AsPublic(mem.UserID),
		Username: mem.R.User.Username,
		Email:    strfmt.Email(mem.R.User.Email),
		Avatar:   mem.R.User.Avatar,
		Role:     mem.Role,
		Joined:   strfmt.DateTime(mem.CreatedAt),
	}
}

func asModels(mems db.OrganizationMemberSlice) []*api.OrganizationMember {
	results := make([]*api.OrganizationMember, len(mems))
	for i, mem := range mems {
		m := asModel(*mem)
		results[i] = &m
	}
	return results
}

func asNewMember(params organization.AddOrgMembersV1Params) (m db.OrganizationMember, err error) {
	if params.Member.UserID == nil || params.Member.Role == nil {
		return m, id.ErrInvalid
	}

	orgID, err := id.ToInternal(params.ID)
	if err != nil {
		return m, err
	}

	uID, err := id.ToInternal(*params.Member.UserID)
	if err != nil {
		return m, err
	}

	return db.OrganizationMember{
		UserID:         uID,
		OrganizationID: orgID,
		Role:           *params.Member.Role,
	}, nil
}
