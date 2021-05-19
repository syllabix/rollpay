package organization

import (
	"github.com/syllabix/rollpay/backend/datastore/model"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

// updatable returns the columns that should be updated
// based upon the state of the organization model
func updatable(org model.Organization) boil.Columns {
	var cols []string

	if len(org.Logo) > 1 {
		cols = append(cols, model.OrganizationColumns.Logo)
	}

	if len(org.Name) > 1 {
		cols = append(cols, model.OrganizationColumns.Name)
	}

	return boil.Whitelist(cols...)
}
