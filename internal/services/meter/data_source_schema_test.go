// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package meter_test

import (
  "context"
  "testing"

  "github.com/m3ter-com/terraform-provider-m3ter/internal/services/meter"
  "github.com/m3ter-com/terraform-provider-m3ter/internal/test_helpers"
)

func TestMeterDataSourceModelSchemaParity(t *testing.T) {
  t.Parallel()
  model := (*meter.MeterDataSourceModel)(nil)
  schema := meter.DataSourceSchema(context.TODO())
  errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
