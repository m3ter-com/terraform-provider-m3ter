// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/m3ter-sdk-go"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

type AccountsDataListDataSourceEnvelope struct {
	Data customfield.NestedObjectList[AccountsItemsDataSourceModel] `json:"data,computed"`
}

type AccountsDataSourceModel struct {
	OrgID    types.String                                               `tfsdk:"org_id" path:"orgId,optional"`
	Codes    *[]types.String                                            `tfsdk:"codes" query:"codes,optional"`
	IDs      *[]types.String                                            `tfsdk:"ids" query:"ids,optional"`
	MaxItems types.Int64                                                `tfsdk:"max_items"`
	Items    customfield.NestedObjectList[AccountsItemsDataSourceModel] `tfsdk:"items"`
}

func (m *AccountsDataSourceModel) toListParams(_ context.Context) (params m3ter.AccountListParams, diags diag.Diagnostics) {
	mCodes := []string{}
	for _, item := range *m.Codes {
		mCodes = append(mCodes, item.ValueString())
	}
	mIDs := []string{}
	for _, item := range *m.IDs {
		mIDs = append(mIDs, item.ValueString())
	}

	params = m3ter.AccountListParams{
		Codes: m3ter.F(mCodes),
		IDs:   m3ter.F(mIDs),
	}

	if !m.OrgID.IsNull() {
		params.OrgID = m3ter.F(m.OrgID.ValueString())
	}

	return
}

type AccountsItemsDataSourceModel struct {
	ID                        types.String                                             `tfsdk:"id" json:"id,computed"`
	Address                   customfield.NestedObject[AccountsAddressDataSourceModel] `tfsdk:"address" json:"address,computed"`
	AutoGenerateStatementMode types.String                                             `tfsdk:"auto_generate_statement_mode" json:"autoGenerateStatementMode,computed"`
	BillEpoch                 timetypes.RFC3339                                        `tfsdk:"bill_epoch" json:"billEpoch,computed" format:"date"`
	Code                      types.String                                             `tfsdk:"code" json:"code,computed"`
	ConfigData                customfield.Map[jsontypes.Normalized]                    `tfsdk:"config_data" json:"configData,computed"`
	CreditApplicationOrder    customfield.List[types.String]                           `tfsdk:"credit_application_order" json:"creditApplicationOrder,computed"`
	Currency                  types.String                                             `tfsdk:"currency" json:"currency,computed"`
	CustomFields              customfield.NormalizedDynamicValue                       `tfsdk:"custom_fields" json:"customFields,computed"`
	DaysBeforeBillDue         types.Int64                                              `tfsdk:"days_before_bill_due" json:"daysBeforeBillDue,computed"`
	EmailAddress              types.String                                             `tfsdk:"email_address" json:"emailAddress,computed"`
	Name                      types.String                                             `tfsdk:"name" json:"name,computed"`
	ParentAccountID           types.String                                             `tfsdk:"parent_account_id" json:"parentAccountId,computed"`
	PurchaseOrderNumber       types.String                                             `tfsdk:"purchase_order_number" json:"purchaseOrderNumber,computed"`
	StatementDefinitionID     types.String                                             `tfsdk:"statement_definition_id" json:"statementDefinitionId,computed"`
	Version                   types.Int64                                              `tfsdk:"version" json:"version,computed,force_encode,encode_state_for_unknown"`
}

type AccountsAddressDataSourceModel struct {
	AddressLine1 types.String `tfsdk:"address_line1" json:"addressLine1,computed"`
	AddressLine2 types.String `tfsdk:"address_line2" json:"addressLine2,computed"`
	AddressLine3 types.String `tfsdk:"address_line3" json:"addressLine3,computed"`
	AddressLine4 types.String `tfsdk:"address_line4" json:"addressLine4,computed"`
	Country      types.String `tfsdk:"country" json:"country,computed"`
	Locality     types.String `tfsdk:"locality" json:"locality,computed"`
	PostCode     types.String `tfsdk:"post_code" json:"postCode,computed"`
	Region       types.String `tfsdk:"region" json:"region,computed"`
}
