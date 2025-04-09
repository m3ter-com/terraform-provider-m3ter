// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account

import (
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/apijson"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

type AccountModel struct {
	ID                        types.String                                  `tfsdk:"id" json:"id,computed"`
	OrgID                     types.String                                  `tfsdk:"org_id" path:"orgId,optional"`
	Code                      types.String                                  `tfsdk:"code" json:"code,required"`
	EmailAddress              types.String                                  `tfsdk:"email_address" json:"emailAddress,required"`
	Name                      types.String                                  `tfsdk:"name" json:"name,required"`
	AutoGenerateStatementMode types.String                                  `tfsdk:"auto_generate_statement_mode" json:"autoGenerateStatementMode,optional"`
	BillEpoch                 timetypes.RFC3339                             `tfsdk:"bill_epoch" json:"billEpoch,optional" format:"date"`
	Currency                  types.String                                  `tfsdk:"currency" json:"currency,optional"`
	DaysBeforeBillDue         types.Int64                                   `tfsdk:"days_before_bill_due" json:"daysBeforeBillDue,optional"`
	ParentAccountID           types.String                                  `tfsdk:"parent_account_id" json:"parentAccountId,optional"`
	PurchaseOrderNumber       types.String                                  `tfsdk:"purchase_order_number" json:"purchaseOrderNumber,optional"`
	StatementDefinitionID     types.String                                  `tfsdk:"statement_definition_id" json:"statementDefinitionId,optional"`
	Version                   types.Int64                                   `tfsdk:"version" json:"version,optional"`
	ConfigData                *map[string]jsontypes.Normalized              `tfsdk:"config_data" json:"configData,optional"`
	CreditApplicationOrder    *[]types.String                               `tfsdk:"credit_application_order" json:"creditApplicationOrder,optional"`
	CustomFields              *map[string]types.Dynamic                     `tfsdk:"custom_fields" json:"customFields,optional"`
	Address                   customfield.NestedObject[AccountAddressModel] `tfsdk:"address" json:"address,computed_optional"`
	CreatedBy                 types.String                                  `tfsdk:"created_by" json:"createdBy,computed"`
	DtCreated                 timetypes.RFC3339                             `tfsdk:"dt_created" json:"dtCreated,computed" format:"date-time"`
	DtLastModified            timetypes.RFC3339                             `tfsdk:"dt_last_modified" json:"dtLastModified,computed" format:"date-time"`
	LastModifiedBy            types.String                                  `tfsdk:"last_modified_by" json:"lastModifiedBy,computed"`
}

func (m AccountModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m AccountModel) MarshalJSONForUpdate(state AccountModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type AccountAddressModel struct {
	AddressLine1 types.String `tfsdk:"address_line1" json:"addressLine1,optional"`
	AddressLine2 types.String `tfsdk:"address_line2" json:"addressLine2,optional"`
	AddressLine3 types.String `tfsdk:"address_line3" json:"addressLine3,optional"`
	AddressLine4 types.String `tfsdk:"address_line4" json:"addressLine4,optional"`
	Country      types.String `tfsdk:"country" json:"country,optional"`
	Locality     types.String `tfsdk:"locality" json:"locality,optional"`
	PostCode     types.String `tfsdk:"post_code" json:"postCode,optional"`
	Region       types.String `tfsdk:"region" json:"region,optional"`
}
