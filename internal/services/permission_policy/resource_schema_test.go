// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package permission_policy_test

import (
	"context"
	"testing"

	"github.com/m3ter-com/terraform-provider-m3ter/internal/services/permission_policy"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/test_helpers"
)

func TestPermissionPolicyModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*permission_policy.PermissionPolicyModel)(nil)
	schema := permission_policy.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
