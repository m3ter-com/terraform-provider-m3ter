// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package plan_group_link_test

import (
  "context"
  "testing"

  "github.com/m3ter-com/terraform-provider-m3ter/internal/services/plan_group_link"
  "github.com/m3ter-com/terraform-provider-m3ter/internal/test_helpers"
)

func TestPlanGroupLinksDataSourceModelSchemaParity(t *testing.T) {
  t.Parallel()
  model := (*plan_group_link.PlanGroupLinksDataSourceModel)(nil)
  schema := plan_group_link.ListDataSourceSchema(context.TODO())
  errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
