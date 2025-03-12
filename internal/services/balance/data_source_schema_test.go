// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package balance_test

import (
  "context"
  "testing"

  "github.com/m3ter-com/terraform-provider-m3ter/internal/services/balance"
  "github.com/m3ter-com/terraform-provider-m3ter/internal/test_helpers"
)

func TestBalanceDataSourceModelSchemaParity(t *testing.T) {
  t.Parallel()
  model := (*balance.BalanceDataSourceModel)(nil)
  schema := balance.DataSourceSchema(context.TODO())
  errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
