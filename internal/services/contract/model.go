// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package contract

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/apijson"
)

type ContractModel struct {
	ID                  types.String      `tfsdk:"id" json:"id,computed"`
	OrgID               types.String      `tfsdk:"org_id" path:"orgId,optional"`
	AccountID           types.String      `tfsdk:"account_id" json:"accountId,required"`
	EndDate             timetypes.RFC3339 `tfsdk:"end_date" json:"endDate,required" format:"date"`
	Name                types.String      `tfsdk:"name" json:"name,required"`
	StartDate           timetypes.RFC3339 `tfsdk:"start_date" json:"startDate,required" format:"date"`
	Code                types.String      `tfsdk:"code" json:"code,optional"`
	Description         types.String      `tfsdk:"description" json:"description,optional"`
	PurchaseOrderNumber types.String      `tfsdk:"purchase_order_number" json:"purchaseOrderNumber,optional"`
	CustomFields        types.Dynamic     `tfsdk:"custom_fields" json:"customFields,optional"`
	Version             types.Int64       `tfsdk:"version" json:"version,computed,force_encode,encode_state_for_unknown"`
}

func (m ContractModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ContractModel) MarshalJSONForUpdate(state ContractModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
