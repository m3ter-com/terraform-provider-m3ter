// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package meter

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/apijson"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

type MeterModel struct {
	ID            types.String                       `tfsdk:"id" json:"id,computed"`
	OrgID         types.String                       `tfsdk:"org_id" path:"orgId,optional"`
	Code          types.String                       `tfsdk:"code" json:"code,required"`
	Name          types.String                       `tfsdk:"name" json:"name,required"`
	DataFields    *[]*MeterDataFieldsModel           `tfsdk:"data_fields" json:"dataFields,required"`
	DerivedFields *[]*MeterDerivedFieldsModel        `tfsdk:"derived_fields" json:"derivedFields,required"`
	GroupID       types.String                       `tfsdk:"group_id" json:"groupId,optional"`
	ProductID     types.String                       `tfsdk:"product_id" json:"productId,optional"`
	CustomFields  customfield.NormalizedDynamicValue `tfsdk:"custom_fields" json:"customFields,optional"`
	Version       types.Int64                        `tfsdk:"version" json:"version,computed,force_encode,encode_state_for_unknown"`
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
	Category    types.String `tfsdk:"category" json:"category,required"`
	Code        types.String `tfsdk:"code" json:"code,required"`
	Name        types.String `tfsdk:"name" json:"name,required"`
	Unit        types.String `tfsdk:"unit" json:"unit,optional"`
	Calculation types.String `tfsdk:"calculation" json:"calculation,required"`
}
