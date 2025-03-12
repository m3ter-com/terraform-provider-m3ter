// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package bill_debit_line_item_test

import (
  "context"
  "testing"

  "github.com/m3ter-com/terraform-provider-m3ter/internal/services/bill_debit_line_item"
  "github.com/m3ter-com/terraform-provider-m3ter/internal/test_helpers"
)

func TestBillDebitLineItemModelSchemaParity(t *testing.T) {
  t.Parallel()
  model := (*bill_debit_line_item.BillDebitLineItemModel)(nil)
  schema := bill_debit_line_item.ResourceSchema(context.TODO())
  errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
