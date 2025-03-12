// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package meter_test

import (
  "context"
  "testing"

  "github.com/m3ter-com/terraform-provider-m3ter/internal/services/meter"
  "github.com/m3ter-com/terraform-provider-m3ter/internal/test_helpers"
)

func TestMeterModelSchemaParity(t *testing.T) {
  t.Parallel()
  model := (*meter.MeterModel)(nil)
  schema := meter.ResourceSchema(context.TODO())
  errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
