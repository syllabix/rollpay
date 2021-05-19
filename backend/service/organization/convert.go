package organization

import (
	"github.com/go-openapi/strfmt"
	api "github.com/syllabix/rollpay/backend/api/model"
	"github.com/syllabix/rollpay/backend/api/rest/operation/organization"
	"github.com/syllabix/rollpay/backend/common/id"
	"github.com/syllabix/rollpay/backend/common/media"
	db "github.com/syllabix/rollpay/backend/datastore/model"
)

// asModel converts a db model to an api model
func asModel(org db.Organization) api.Organization {
	return api.Organization{
		ID:        id.AsPublic(org.ID),
		Name:      org.Name,
		Logo:      org.Logo,
		Accounts:  mapAccounts(org),
		CreatedAt: strfmt.DateTime(org.CreatedAt),
		UpdatedAt: strfmt.DateTime(org.UpdatedAt),
	}
}

func mapAccounts(org db.Organization) []*api.LinkedAccount {
	if org.R == nil {
		return []*api.LinkedAccount{}
	}

	accounts := make([]*api.LinkedAccount, len(org.R.OrganizationAccounts))
	for i, account := range org.R.OrganizationAccounts {
		link := account.R.LinkedAccount
		accounts[i] = &api.LinkedAccount{
			ID:        id.AsPublic(link.ID),
			Alias:     link.Alias,
			CreatedAt: strfmt.DateTime(link.CreatedAt),
			UpdatedAt: strfmt.DateTime(link.UpdatedAt),
		}
	}
	return accounts
}

func asModels(orgs db.OrganizationSlice) []*api.Organization {
	result := make([]*api.Organization, len(orgs))
	for i, org := range orgs {
		m := asModel(*org)
		result[i] = &m
	}
	return result
}

func asNewOrg(params organization.CreateOrganizationV1Params) (db.Organization, error) {
	img, err := media.Process(params.HTTPRequest.Context(), params.Logo)
	if err != nil {
		return db.Organization{}, err
	}

	return db.Organization{
		Name: params.Name,
		Logo: img,
	}, nil
}
