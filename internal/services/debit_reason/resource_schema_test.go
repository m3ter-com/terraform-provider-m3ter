// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package debit_reason_test

import (
  "context"
  "testing"

  "github.com/m3ter-com/terraform-provider-m3ter/internal/services/debit_reason"
  "github.com/m3ter-com/terraform-provider-m3ter/internal/test_helpers"
)

func TestDebitReasonModelSchemaParity(t *testing.T) {
  t.Parallel()
  model := (*debit_reason.DebitReasonModel)(nil)
  schema := debit_reason.ResourceSchema(context.TODO())
  errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
