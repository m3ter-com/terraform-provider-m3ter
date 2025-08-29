// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package product

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/apijson"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

type ProductModel struct {
	ID           types.String                       `tfsdk:"id" json:"id,computed"`
	OrgID        types.String                       `tfsdk:"org_id" path:"orgId,optional"`
	Code         types.String                       `tfsdk:"code" json:"code,required"`
	Name         types.String                       `tfsdk:"name" json:"name,required"`
	CustomFields customfield.NormalizedDynamicValue `tfsdk:"custom_fields" json:"customFields,optional"`
	Version      types.Int64                        `tfsdk:"version" json:"version,computed,force_encode,encode_state_for_unknown"`
}

func (m ProductModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ProductModel) MarshalJSONForUpdate(state ProductModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
