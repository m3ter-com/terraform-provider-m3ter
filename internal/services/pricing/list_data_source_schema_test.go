// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pricing_test

import (
  "context"
  "testing"

  "github.com/m3ter-com/terraform-provider-m3ter/internal/services/pricing"
  "github.com/m3ter-com/terraform-provider-m3ter/internal/test_helpers"
)

func TestPricingsDataSourceModelSchemaParity(t *testing.T) {
  t.Parallel()
  model := (*pricing.PricingsDataSourceModel)(nil)
  schema := pricing.ListDataSourceSchema(context.TODO())
  errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
