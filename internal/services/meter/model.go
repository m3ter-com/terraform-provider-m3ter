// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package meter

import (
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/apijson"
)

type MeterModel struct {
	ID             types.String                     `tfsdk:"id" json:"id,computed"`
	OrgID          types.String                     `tfsdk:"org_id" path:"orgId,required"`
	Code           types.String                     `tfsdk:"code" json:"code,required"`
	Name           types.String                     `tfsdk:"name" json:"name,required"`
	DataFields     *[]*MeterDataFieldsModel         `tfsdk:"data_fields" json:"dataFields,required"`
	DerivedFields  *[]*MeterDerivedFieldsModel      `tfsdk:"derived_fields" json:"derivedFields,required"`
	GroupID        types.String                     `tfsdk:"group_id" json:"groupId,optional"`
	ProductID      types.String                     `tfsdk:"product_id" json:"productId,optional"`
	Version        types.Int64                      `tfsdk:"version" json:"version,optional"`
	CustomFields   *map[string]jsontypes.Normalized `tfsdk:"custom_fields" json:"customFields,optional"`
	CreatedBy      types.String                     `tfsdk:"created_by" json:"createdBy,computed"`
	DtCreated      timetypes.RFC3339                `tfsdk:"dt_created" json:"dtCreated,computed" format:"date-time"`
	DtLastModified timetypes.RFC3339                `tfsdk:"dt_last_modified" json:"dtLastModified,computed" format:"date-time"`
	LastModifiedBy types.String                     `tfsdk:"last_modified_by" json:"lastModifiedBy,computed"`
}

func (m MeterModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m MeterModel) MarshalJSONForUpdate(state MeterModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type MeterDataFieldsModel struct {
	Category types.String `tfsdk:"category" json:"category,required"`
	Code     types.String `tfsdk:"code" json:"code,required"`
	Name     types.String `tfsdk:"name" json:"name,required"`
	Unit     types.String `tfsdk:"unit" json:"unit,optional"`
}

type MeterDerivedFieldsModel struct {
	Calculation types.String `tfsdk:"calculation" json:"calculation,required"`
	Category    types.String `tfsdk:"category" json:"category,required"`
	Code        types.String `tfsdk:"code" json:"code,required"`
	Name        types.String `tfsdk:"name" json:"name,required"`
	Unit        types.String `tfsdk:"unit" json:"unit,optional"`
}
