// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package commitment_test

import (
  "context"
  "testing"

  "github.com/m3ter-com/terraform-provider-m3ter/internal/services/commitment"
  "github.com/m3ter-com/terraform-provider-m3ter/internal/test_helpers"
)

func TestCommitmentDataSourceModelSchemaParity(t *testing.T) {
  t.Parallel()
  model := (*commitment.CommitmentDataSourceModel)(nil)
  schema := commitment.DataSourceSchema(context.TODO())
  errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
