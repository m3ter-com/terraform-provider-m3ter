// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package external_mapping_test

import (
	"context"
	"testing"

	"github.com/m3ter-com/terraform-provider-m3ter/internal/services/external_mapping"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/test_helpers"
)

func TestExternalMappingsDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*external_mapping.ExternalMappingsDataSourceModel)(nil)
	schema := external_mapping.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
