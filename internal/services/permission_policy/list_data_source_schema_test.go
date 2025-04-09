// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package permission_policy_test

import (
	"context"
	"testing"

	"github.com/m3ter-com/terraform-provider-m3ter/internal/services/permission_policy"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/test_helpers"
)

func TestPermissionPoliciesDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*permission_policy.PermissionPoliciesDataSourceModel)(nil)
	schema := permission_policy.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
