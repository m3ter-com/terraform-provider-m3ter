// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account_plan

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/apijson"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

type AccountPlanModel struct {
	ID               types.String                       `tfsdk:"id" json:"id,computed"`
	OrgID            types.String                       `tfsdk:"org_id" path:"orgId,optional"`
	AccountID        types.String                       `tfsdk:"account_id" json:"accountId,required"`
	StartDate        timetypes.RFC3339                  `tfsdk:"start_date" json:"startDate,required" format:"date-time"`
	BillEpoch        timetypes.RFC3339                  `tfsdk:"bill_epoch" json:"billEpoch,optional" format:"date"`
	ChildBillingMode types.String                       `tfsdk:"child_billing_mode" json:"childBillingMode,optional"`
	Code             types.String                       `tfsdk:"code" json:"code,optional"`
	ContractID       types.String                       `tfsdk:"contract_id" json:"contractId,optional"`
	EndDate          timetypes.RFC3339                  `tfsdk:"end_date" json:"endDate,optional" format:"date-time"`
	PlanGroupID      types.String                       `tfsdk:"plan_group_id" json:"planGroupId,optional"`
	PlanID           types.String                       `tfsdk:"plan_id" json:"planId,optional"`
	CustomFields     customfield.NormalizedDynamicValue `tfsdk:"custom_fields" json:"customFields,optional"`
	ProductID        types.String                       `tfsdk:"product_id" json:"productId,computed"`
	Version          types.Int64                        `tfsdk:"version" json:"version,computed,force_encode,encode_state_for_unknown"`
}

func (m AccountPlanModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m AccountPlanModel) MarshalJSONForUpdate(state AccountPlanModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
