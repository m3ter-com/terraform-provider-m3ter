// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account_plan

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/m3ter-sdk-go"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

type AccountPlanDataSourceModel struct {
	ID               types.String                       `tfsdk:"id" path:"id,required"`
	OrgID            types.String                       `tfsdk:"org_id" path:"orgId,required"`
	AccountID        types.String                       `tfsdk:"account_id" json:"accountId,computed"`
	BillEpoch        timetypes.RFC3339                  `tfsdk:"bill_epoch" json:"billEpoch,computed" format:"date"`
	ChildBillingMode types.String                       `tfsdk:"child_billing_mode" json:"childBillingMode,computed"`
	Code             types.String                       `tfsdk:"code" json:"code,computed"`
	ContractID       types.String                       `tfsdk:"contract_id" json:"contractId,computed"`
	EndDate          timetypes.RFC3339                  `tfsdk:"end_date" json:"endDate,computed" format:"date-time"`
	PlanGroupID      types.String                       `tfsdk:"plan_group_id" json:"planGroupId,computed"`
	PlanID           types.String                       `tfsdk:"plan_id" json:"planId,computed"`
	ProductID        types.String                       `tfsdk:"product_id" json:"productId,computed"`
	StartDate        timetypes.RFC3339                  `tfsdk:"start_date" json:"startDate,computed" format:"date-time"`
	Version          types.Int64                        `tfsdk:"version" json:"version,computed,force_encode,encode_state_for_unknown"`
	CustomFields     customfield.NormalizedDynamicValue `tfsdk:"custom_fields" json:"customFields,computed"`
}

func (m *AccountPlanDataSourceModel) toReadParams(_ context.Context) (params m3ter.AccountPlanGetParams, diags diag.Diagnostics) {
	params = m3ter.AccountPlanGetParams{}

	if !m.OrgID.IsNull() {
		params.OrgID = m3ter.F(m.OrgID.ValueString())
	}

	return
}
