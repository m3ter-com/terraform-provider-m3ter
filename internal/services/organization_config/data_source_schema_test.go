// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package organization_config_test

import (
	"context"
	"testing"

	"github.com/m3ter-com/terraform-provider-m3ter/internal/services/organization_config"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/test_helpers"
)

func TestOrganizationConfigDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*organization_config.OrganizationConfigDataSourceModel)(nil)
	schema := organization_config.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
