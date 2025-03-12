// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package bill_debit_line_item_test

import (
  "context"
  "testing"

  "github.com/m3ter-com/terraform-provider-m3ter/internal/services/bill_debit_line_item"
  "github.com/m3ter-com/terraform-provider-m3ter/internal/test_helpers"
)

func TestBillDebitLineItemDataSourceModelSchemaParity(t *testing.T) {
  t.Parallel()
  model := (*bill_debit_line_item.BillDebitLineItemDataSourceModel)(nil)
  schema := bill_debit_line_item.DataSourceSchema(context.TODO())
  errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
