// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package contract

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

type ContractDataSourceModel struct {
	ID                  types.String                   `tfsdk:"id" path:"id,required"`
	OrgID               types.String                   `tfsdk:"org_id" path:"orgId,required"`
	AccountID           types.String                   `tfsdk:"account_id" json:"accountId,computed"`
	Code                types.String                   `tfsdk:"code" json:"code,computed"`
	CreatedBy           types.String                   `tfsdk:"created_by" json:"createdBy,computed"`
	Description         types.String                   `tfsdk:"description" json:"description,computed"`
	DtCreated           timetypes.RFC3339              `tfsdk:"dt_created" json:"dtCreated,computed" format:"date-time"`
	DtLastModified      timetypes.RFC3339              `tfsdk:"dt_last_modified" json:"dtLastModified,computed" format:"date-time"`
	EndDate             timetypes.RFC3339              `tfsdk:"end_date" json:"endDate,computed" format:"date"`
	LastModifiedBy      types.String                   `tfsdk:"last_modified_by" json:"lastModifiedBy,computed"`
	Name                types.String                   `tfsdk:"name" json:"name,computed"`
	PurchaseOrderNumber types.String                   `tfsdk:"purchase_order_number" json:"purchaseOrderNumber,computed"`
	StartDate           timetypes.RFC3339              `tfsdk:"start_date" json:"startDate,computed" format:"date"`
	Version             types.Int64                    `tfsdk:"version" json:"version,computed"`
	CustomFields        customfield.Map[types.Dynamic] `tfsdk:"custom_fields" json:"customFields,computed"`
}
