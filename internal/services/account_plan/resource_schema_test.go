// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account_plan_test

import (
	"context"
	"testing"

	"github.com/m3ter-com/terraform-provider-m3ter/internal/services/account_plan"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/test_helpers"
)

func TestAccountPlanModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*account_plan.AccountPlanModel)(nil)
	schema := account_plan.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
