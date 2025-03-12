// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package bill_test

import (
  "context"
  "testing"

  "github.com/m3ter-com/terraform-provider-m3ter/internal/services/bill"
  "github.com/m3ter-com/terraform-provider-m3ter/internal/test_helpers"
)

func TestBillDataSourceModelSchemaParity(t *testing.T) {
  t.Parallel()
  model := (*bill.BillDataSourceModel)(nil)
  schema := bill.DataSourceSchema(context.TODO())
  errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
