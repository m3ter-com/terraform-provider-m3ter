// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package plan_template_test

import (
  "context"
  "testing"

  "github.com/m3ter-com/terraform-provider-m3ter/internal/services/plan_template"
  "github.com/m3ter-com/terraform-provider-m3ter/internal/test_helpers"
)

func TestPlanTemplateDataSourceModelSchemaParity(t *testing.T) {
  t.Parallel()
  model := (*plan_template.PlanTemplateDataSourceModel)(nil)
  schema := plan_template.DataSourceSchema(context.TODO())
  errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
  errs.Report(t)
}
