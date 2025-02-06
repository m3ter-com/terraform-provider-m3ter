// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package organization_config_test

import (
	"context"
	"testing"

	"github.com/m3ter-com/terraform-provider-m3ter/internal/services/organization_config"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/test_helpers"
)

func TestOrganizationConfigModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*organization_config.OrganizationConfigModel)(nil)
	schema := organization_config.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
