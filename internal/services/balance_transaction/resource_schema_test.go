// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package balance_transaction_test

import (
  "context"
  "testing"

  "github.com/m3ter-com/terraform-provider-m3ter/internal/services/balance_transaction"
  "github.com/m3ter-com/terraform-provider-m3ter/internal/test_helpers"
)

func TestBalanceTransactionModelSchemaParity(t *testing.T) {
  t.Parallel()
  model := (*balance_transaction.BalanceTransactionModel)(nil)
  schema := balance_transaction.ResourceSchema(context.TODO())
  errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
