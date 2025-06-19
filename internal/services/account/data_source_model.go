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

type AccountDataSourceModel struct {
	ID                        types.String                                            `tfsdk:"id" path:"id,required"`
	OrgID                     types.String                                            `tfsdk:"org_id" path:"orgId,required"`
	AutoGenerateStatementMode types.String                                            `tfsdk:"auto_generate_statement_mode" json:"autoGenerateStatementMode,computed"`
	BillEpoch                 timetypes.RFC3339                                       `tfsdk:"bill_epoch" json:"billEpoch,computed" format:"date"`
	Code                      types.String                                            `tfsdk:"code" json:"code,computed"`
	CreatedBy                 types.String                                            `tfsdk:"created_by" json:"createdBy,computed"`
	Currency                  types.String                                            `tfsdk:"currency" json:"currency,computed"`
	DaysBeforeBillDue         types.Int64                                             `tfsdk:"days_before_bill_due" json:"daysBeforeBillDue,computed"`
	DtCreated                 timetypes.RFC3339                                       `tfsdk:"dt_created" json:"dtCreated,computed" format:"date-time"`
	DtLastModified            timetypes.RFC3339                                       `tfsdk:"dt_last_modified" json:"dtLastModified,computed" format:"date-time"`
	EmailAddress              types.String                                            `tfsdk:"email_address" json:"emailAddress,computed"`
	LastModifiedBy            types.String                                            `tfsdk:"last_modified_by" json:"lastModifiedBy,computed"`
	Name                      types.String                                            `tfsdk:"name" json:"name,computed"`
	ParentAccountID           types.String                                            `tfsdk:"parent_account_id" json:"parentAccountId,computed"`
	PurchaseOrderNumber       types.String                                            `tfsdk:"purchase_order_number" json:"purchaseOrderNumber,computed"`
	StatementDefinitionID     types.String                                            `tfsdk:"statement_definition_id" json:"statementDefinitionId,computed"`
	Version                   types.Int64                                             `tfsdk:"version" json:"version,computed"`
	ConfigData                customfield.Map[jsontypes.Normalized]                   `tfsdk:"config_data" json:"configData,computed"`
	CreditApplicationOrder    customfield.List[types.String]                          `tfsdk:"credit_application_order" json:"creditApplicationOrder,computed"`
	CustomFields              customfield.Map[types.Dynamic]                          `tfsdk:"custom_fields" json:"customFields,computed"`
	Address                   customfield.NestedObject[AccountAddressDataSourceModel] `tfsdk:"address" json:"address,computed"`
}

func (m *AccountDataSourceModel) toReadParams(_ context.Context) (params m3ter.AccountGetParams, diags diag.Diagnostics) {
	params = m3ter.AccountGetParams{}

	if !m.OrgID.IsNull() {
		params.OrgID = m3ter.F(m.OrgID.ValueString())
	}

	return
}

type AccountAddressDataSourceModel struct {
	AddressLine1 types.String `tfsdk:"address_line1" json:"addressLine1,computed"`
	AddressLine2 types.String `tfsdk:"address_line2" json:"addressLine2,computed"`
	AddressLine3 types.String `tfsdk:"address_line3" json:"addressLine3,computed"`
	AddressLine4 types.String `tfsdk:"address_line4" json:"addressLine4,computed"`
	Country      types.String `tfsdk:"country" json:"country,computed"`
	Locality     types.String `tfsdk:"locality" json:"locality,computed"`
	PostCode     types.String `tfsdk:"post_code" json:"postCode,computed"`
	Region       types.String `tfsdk:"region" json:"region,computed"`
}
