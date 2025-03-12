// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package transaction_type_test

import (
  "context"
  "testing"

  "github.com/m3ter-com/terraform-provider-m3ter/internal/services/transaction_type"
  "github.com/m3ter-com/terraform-provider-m3ter/internal/test_helpers"
)

func TestTransactionTypeDataSourceModelSchemaParity(t *testing.T) {
  t.Parallel()
  model := (*transaction_type.TransactionTypeDataSourceModel)(nil)
  schema := transaction_type.DataSourceSchema(context.TODO())
  errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
