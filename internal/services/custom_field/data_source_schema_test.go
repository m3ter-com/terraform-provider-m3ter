// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_field_test

import (
	"context"
	"testing"

	"github.com/m3ter-com/terraform-provider-m3ter/internal/services/custom_field"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/test_helpers"
)

func TestCustomFieldDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*custom_field.CustomFieldDataSourceModel)(nil)
	schema := custom_field.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
