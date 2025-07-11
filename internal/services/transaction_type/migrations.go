// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package transaction_type

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithUpgradeState = (*TransactionTypeResource)(nil)

func (r *TransactionTypeResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{}
}
