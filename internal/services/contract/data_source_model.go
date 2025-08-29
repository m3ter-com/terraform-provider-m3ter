// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package contract

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/m3ter-sdk-go"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

type ContractDataSourceModel struct {
	ID                  types.String                       `tfsdk:"id" path:"id,required"`
	OrgID               types.String                       `tfsdk:"org_id" path:"orgId,required"`
	AccountID           types.String                       `tfsdk:"account_id" json:"accountId,computed"`
	BillGroupingKey     types.String                       `tfsdk:"bill_grouping_key" json:"billGroupingKey,computed"`
	Code                types.String                       `tfsdk:"code" json:"code,computed"`
	Description         types.String                       `tfsdk:"description" json:"description,computed"`
	EndDate             timetypes.RFC3339                  `tfsdk:"end_date" json:"endDate,computed" format:"date"`
	Name                types.String                       `tfsdk:"name" json:"name,computed"`
	PurchaseOrderNumber types.String                       `tfsdk:"purchase_order_number" json:"purchaseOrderNumber,computed"`
	StartDate           timetypes.RFC3339                  `tfsdk:"start_date" json:"startDate,computed" format:"date"`
	Version             types.Int64                        `tfsdk:"version" json:"version,computed,force_encode,encode_state_for_unknown"`
	CustomFields        customfield.NormalizedDynamicValue `tfsdk:"custom_fields" json:"customFields,computed"`
}

func (m *ContractDataSourceModel) toReadParams(_ context.Context) (params m3ter.ContractGetParams, diags diag.Diagnostics) {
	params = m3ter.ContractGetParams{}

	if !m.OrgID.IsNull() {
		params.OrgID = m3ter.F(m.OrgID.ValueString())
	}

	return
}
