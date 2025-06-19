// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package plan_template

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/m3ter-sdk-go"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

type PlanTemplatesDataListDataSourceEnvelope struct {
	Data customfield.NestedObjectList[PlanTemplatesItemsDataSourceModel] `json:"data,computed"`
}

type PlanTemplatesDataSourceModel struct {
	OrgID     types.String                                                    `tfsdk:"org_id" path:"orgId,optional"`
	ProductID types.String                                                    `tfsdk:"product_id" query:"productId,optional"`
	IDs       *[]types.String                                                 `tfsdk:"ids" query:"ids,optional"`
	MaxItems  types.Int64                                                     `tfsdk:"max_items"`
	Items     customfield.NestedObjectList[PlanTemplatesItemsDataSourceModel] `tfsdk:"items"`
}

func (m *PlanTemplatesDataSourceModel) toListParams(_ context.Context) (params m3ter.PlanTemplateListParams, diags diag.Diagnostics) {
	mIDs := []string{}
	for _, item := range *m.IDs {
		mIDs = append(mIDs, item.ValueString())
	}

	params = m3ter.PlanTemplateListParams{
		IDs: m3ter.F(mIDs),
	}

	if !m.ProductID.IsNull() {
		params.ProductID = m3ter.F(m.ProductID.ValueString())
	}
	if !m.OrgID.IsNull() {
		params.OrgID = m3ter.F(m.OrgID.ValueString())
	}

	return
}

type PlanTemplatesItemsDataSourceModel struct {
	ID                          types.String                   `tfsdk:"id" json:"id,computed"`
	Version                     types.Int64                    `tfsdk:"version" json:"version,computed"`
	BillFrequency               types.String                   `tfsdk:"bill_frequency" json:"billFrequency,computed"`
	BillFrequencyInterval       types.Int64                    `tfsdk:"bill_frequency_interval" json:"billFrequencyInterval,computed"`
	Code                        types.String                   `tfsdk:"code" json:"code,computed"`
	CreatedBy                   types.String                   `tfsdk:"created_by" json:"createdBy,computed"`
	Currency                    types.String                   `tfsdk:"currency" json:"currency,computed"`
	CustomFields                customfield.Map[types.Dynamic] `tfsdk:"custom_fields" json:"customFields,computed"`
	DtCreated                   timetypes.RFC3339              `tfsdk:"dt_created" json:"dtCreated,computed" format:"date-time"`
	DtLastModified              timetypes.RFC3339              `tfsdk:"dt_last_modified" json:"dtLastModified,computed" format:"date-time"`
	LastModifiedBy              types.String                   `tfsdk:"last_modified_by" json:"lastModifiedBy,computed"`
	MinimumSpend                types.Float64                  `tfsdk:"minimum_spend" json:"minimumSpend,computed"`
	MinimumSpendBillInAdvance   types.Bool                     `tfsdk:"minimum_spend_bill_in_advance" json:"minimumSpendBillInAdvance,computed"`
	MinimumSpendDescription     types.String                   `tfsdk:"minimum_spend_description" json:"minimumSpendDescription,computed"`
	Name                        types.String                   `tfsdk:"name" json:"name,computed"`
	Ordinal                     types.Int64                    `tfsdk:"ordinal" json:"ordinal,computed"`
	ProductID                   types.String                   `tfsdk:"product_id" json:"productId,computed"`
	StandingCharge              types.Float64                  `tfsdk:"standing_charge" json:"standingCharge,computed"`
	StandingChargeBillInAdvance types.Bool                     `tfsdk:"standing_charge_bill_in_advance" json:"standingChargeBillInAdvance,computed"`
	StandingChargeDescription   types.String                   `tfsdk:"standing_charge_description" json:"standingChargeDescription,computed"`
	StandingChargeInterval      types.Int64                    `tfsdk:"standing_charge_interval" json:"standingChargeInterval,computed"`
	StandingChargeOffset        types.Int64                    `tfsdk:"standing_charge_offset" json:"standingChargeOffset,computed"`
}
