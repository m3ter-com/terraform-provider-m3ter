// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package permission_policy

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithUpgradeState = (*PermissionPolicyResource)(nil)

func (r *PermissionPolicyResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{}
}
