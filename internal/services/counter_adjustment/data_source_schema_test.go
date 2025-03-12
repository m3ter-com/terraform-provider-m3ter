// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package counter_adjustment_test

import (
  "context"
  "testing"

  "github.com/m3ter-com/terraform-provider-m3ter/internal/services/counter_adjustment"
  "github.com/m3ter-com/terraform-provider-m3ter/internal/test_helpers"
)

func TestCounterAdjustmentDataSourceModelSchemaParity(t *testing.T) {
  t.Parallel()
  model := (*counter_adjustment.CounterAdjustmentDataSourceModel)(nil)
  schema := counter_adjustment.DataSourceSchema(context.TODO())
  errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
