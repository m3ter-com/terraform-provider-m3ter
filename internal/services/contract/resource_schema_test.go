// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package contract_test

import (
  "context"
  "testing"

  "github.com/m3ter-com/terraform-provider-m3ter/internal/services/contract"
  "github.com/m3ter-com/terraform-provider-m3ter/internal/test_helpers"
)

func TestContractModelSchemaParity(t *testing.T) {
  t.Parallel()
  model := (*contract.ContractModel)(nil)
  schema := contract.ResourceSchema(context.TODO())
  errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
