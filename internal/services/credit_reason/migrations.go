// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package credit_reason

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithUpgradeState = (*CreditReasonResource)(nil)

func (r *CreditReasonResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{}
}
