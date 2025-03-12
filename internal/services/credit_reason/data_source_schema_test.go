// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package credit_reason_test

import (
  "context"
  "testing"

  "github.com/m3ter-com/terraform-provider-m3ter/internal/services/credit_reason"
  "github.com/m3ter-com/terraform-provider-m3ter/internal/test_helpers"
)

func TestCreditReasonDataSourceModelSchemaParity(t *testing.T) {
  t.Parallel()
  model := (*credit_reason.CreditReasonDataSourceModel)(nil)
  schema := credit_reason.DataSourceSchema(context.TODO())
  errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
