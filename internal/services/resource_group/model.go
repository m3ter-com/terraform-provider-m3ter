// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package resource_group

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/apijson"
)

type ResourceGroupModel struct {
	ID      types.String `tfsdk:"id" json:"id,computed"`
	Type    types.String `tfsdk:"type" path:"type,required"`
	OrgID   types.String `tfsdk:"org_id" path:"orgId,optional"`
	Name    types.String `tfsdk:"name" json:"name,required"`
	Version types.Int64  `tfsdk:"version" json:"version,computed,force_encode,encode_state_for_unknown"`
}

func (m ResourceGroupModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ResourceGroupModel) MarshalJSONForUpdate(state ResourceGroupModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
