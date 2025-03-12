// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package balance

import (
  "context"

  "github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithUpgradeState = (*BalanceResource)(nil)

func (r *BalanceResource) UpgradeState(ctx context.Context) (map[int64]resource.StateUpgrader) {
  return map[int64]resource.StateUpgrader{}
}
