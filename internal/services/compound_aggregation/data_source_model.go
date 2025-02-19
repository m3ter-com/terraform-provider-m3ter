// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package compound_aggregation

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

type CompoundAggregationDataSourceModel struct {
	OrgID                    types.String                                    `tfsdk:"org_id" path:"orgId,required"`
	ID                       types.String                                    `tfsdk:"id" path:"id,computed"`
	AccountingProductID      types.String                                    `tfsdk:"accounting_product_id" json:"accountingProductId,computed"`
	Calculation              types.String                                    `tfsdk:"calculation" json:"calculation,computed"`
	Code                     types.String                                    `tfsdk:"code" json:"code,computed"`
	CreatedBy                types.String                                    `tfsdk:"created_by" json:"createdBy,computed"`
	DtCreated                timetypes.RFC3339                               `tfsdk:"dt_created" json:"dtCreated,computed" format:"date-time"`
	DtLastModified           timetypes.RFC3339                               `tfsdk:"dt_last_modified" json:"dtLastModified,computed" format:"date-time"`
	EvaluateNullAggregations types.Bool                                      `tfsdk:"evaluate_null_aggregations" json:"evaluateNullAggregations,computed"`
	LastModifiedBy           types.String                                    `tfsdk:"last_modified_by" json:"lastModifiedBy,computed"`
	Name                     types.String                                    `tfsdk:"name" json:"name,computed"`
	ProductID                types.String                                    `tfsdk:"product_id" json:"productId,computed"`
	QuantityPerUnit          types.Float64                                   `tfsdk:"quantity_per_unit" json:"quantityPerUnit,computed"`
	Rounding                 types.String                                    `tfsdk:"rounding" json:"rounding,computed"`
	Unit                     types.String                                    `tfsdk:"unit" json:"unit,computed"`
	Version                  types.Int64                                     `tfsdk:"version" json:"version,computed"`
	CustomFields             customfield.Map[types.Dynamic]                  `tfsdk:"custom_fields" json:"customFields,computed"`
	Segments                 customfield.List[customfield.Map[types.String]] `tfsdk:"segments" json:"segments,computed"`
}
