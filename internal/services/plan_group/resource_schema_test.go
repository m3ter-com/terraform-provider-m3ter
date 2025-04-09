// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package plan_group_test

import (
	"context"
	"testing"

	"github.com/m3ter-com/terraform-provider-m3ter/internal/services/plan_group"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/test_helpers"
)

func TestPlanGroupModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*plan_group.PlanGroupModel)(nil)
	schema := plan_group.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
