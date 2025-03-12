// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package credit_reason_test

import (
  "context"
  "testing"

  "github.com/m3ter-com/terraform-provider-m3ter/internal/services/credit_reason"
  "github.com/m3ter-com/terraform-provider-m3ter/internal/test_helpers"
)

func TestCreditReasonModelSchemaParity(t *testing.T) {
  t.Parallel()
  model := (*credit_reason.CreditReasonModel)(nil)
  schema := credit_reason.ResourceSchema(context.TODO())
  errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
