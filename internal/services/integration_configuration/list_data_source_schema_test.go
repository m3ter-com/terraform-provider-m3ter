// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package integration_configuration_test

import (
	"context"
	"testing"

	"github.com/m3ter-com/terraform-provider-m3ter/internal/services/integration_configuration"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/test_helpers"
)

func TestIntegrationConfigurationsDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*integration_configuration.IntegrationConfigurationsDataSourceModel)(nil)
	schema := integration_configuration.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
