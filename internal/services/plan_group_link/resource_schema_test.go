// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package plan_group_link_test

import (
  "context"
  "testing"

  "github.com/m3ter-com/terraform-provider-m3ter/internal/services/plan_group_link"
  "github.com/m3ter-com/terraform-provider-m3ter/internal/test_helpers"
)

func TestPlanGroupLinkModelSchemaParity(t *testing.T) {
  t.Parallel()
  model := (*plan_group_link.PlanGroupLinkModel)(nil)
  schema := plan_group_link.ResourceSchema(context.TODO())
  errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
