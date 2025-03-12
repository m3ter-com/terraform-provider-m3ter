// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package debit_reason_test

import (
  "context"
  "testing"

  "github.com/m3ter-com/terraform-provider-m3ter/internal/services/debit_reason"
  "github.com/m3ter-com/terraform-provider-m3ter/internal/test_helpers"
)

func TestDebitReasonsDataSourceModelSchemaParity(t *testing.T) {
  t.Parallel()
  model := (*debit_reason.DebitReasonsDataSourceModel)(nil)
  schema := debit_reason.ListDataSourceSchema(context.TODO())
  errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
