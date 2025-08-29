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

type ContractsDataListDataSourceEnvelope struct {
	Data customfield.NestedObjectList[ContractsItemsDataSourceModel] `json:"data,computed"`
}

type ContractsDataSourceModel struct {
	OrgID     types.String                                                `tfsdk:"org_id" path:"orgId,optional"`
	AccountID types.String                                                `tfsdk:"account_id" query:"accountId,optional"`
	Codes     *[]types.String                                             `tfsdk:"codes" query:"codes,optional"`
	IDs       *[]types.String                                             `tfsdk:"ids" query:"ids,optional"`
	MaxItems  types.Int64                                                 `tfsdk:"max_items"`
	Items     customfield.NestedObjectList[ContractsItemsDataSourceModel] `tfsdk:"items"`
}

func (m *ContractsDataSourceModel) toListParams(_ context.Context) (params m3ter.ContractListParams, diags diag.Diagnostics) {
	mCodes := []string{}
	if m.Codes != nil {
		for _, item := range *m.Codes {
			mCodes = append(mCodes, item.ValueString())
		}
	}
	mIDs := []string{}
	if m.IDs != nil {
		for _, item := range *m.IDs {
			mIDs = append(mIDs, item.ValueString())
		}
	}

	params = m3ter.ContractListParams{
		Codes: m3ter.F(mCodes),
		IDs:   m3ter.F(mIDs),
	}

	if !m.AccountID.IsNull() {
		params.AccountID = m3ter.F(m.AccountID.ValueString())
	}
	if !m.OrgID.IsNull() {
		params.OrgID = m3ter.F(m.OrgID.ValueString())
	}

	return
}

type ContractsItemsDataSourceModel struct {
	ID                  types.String                       `tfsdk:"id" json:"id,computed"`
	AccountID           types.String                       `tfsdk:"account_id" json:"accountId,computed"`
	BillGroupingKey     types.String                       `tfsdk:"bill_grouping_key" json:"billGroupingKey,computed"`
	Code                types.String                       `tfsdk:"code" json:"code,computed"`
	CustomFields        customfield.NormalizedDynamicValue `tfsdk:"custom_fields" json:"customFields,computed"`
	Description         types.String                       `tfsdk:"description" json:"description,computed"`
	EndDate             timetypes.RFC3339                  `tfsdk:"end_date" json:"endDate,computed" format:"date"`
	Name                types.String                       `tfsdk:"name" json:"name,computed"`
	PurchaseOrderNumber types.String                       `tfsdk:"purchase_order_number" json:"purchaseOrderNumber,computed"`
	StartDate           timetypes.RFC3339                  `tfsdk:"start_date" json:"startDate,computed" format:"date"`
	Version             types.Int64                        `tfsdk:"version" json:"version,computed,force_encode,encode_state_for_unknown"`
}
