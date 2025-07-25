// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package plan_test

import (
	"context"
	"testing"

	"github.com/m3ter-com/terraform-provider-m3ter/internal/services/plan"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/test_helpers"
)

func TestPlanModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*plan.PlanModel)(nil)
	schema := plan.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
