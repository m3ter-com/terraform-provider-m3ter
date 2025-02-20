// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package meter

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/m3ter-sdk-go"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

type MeterDataSourceModel struct {
	ID             types.String                                                    `tfsdk:"id" path:"id,required"`
	OrgID          types.String                                                    `tfsdk:"org_id" path:"orgId,required"`
	Code           types.String                                                    `tfsdk:"code" json:"code,computed"`
	CreatedBy      types.String                                                    `tfsdk:"created_by" json:"createdBy,computed"`
	DtCreated      timetypes.RFC3339                                               `tfsdk:"dt_created" json:"dtCreated,computed" format:"date-time"`
	DtLastModified timetypes.RFC3339                                               `tfsdk:"dt_last_modified" json:"dtLastModified,computed" format:"date-time"`
	GroupID        types.String                                                    `tfsdk:"group_id" json:"groupId,computed"`
	LastModifiedBy types.String                                                    `tfsdk:"last_modified_by" json:"lastModifiedBy,computed"`
	Name           types.String                                                    `tfsdk:"name" json:"name,computed"`
	ProductID      types.String                                                    `tfsdk:"product_id" json:"productId,computed"`
	Version        types.Int64                                                     `tfsdk:"version" json:"version,computed"`
	CustomFields   customfield.Map[types.Dynamic]                                  `tfsdk:"custom_fields" json:"customFields,computed"`
	DataFields     customfield.NestedObjectList[MeterDataFieldsDataSourceModel]    `tfsdk:"data_fields" json:"dataFields,computed"`
	DerivedFields  customfield.NestedObjectList[MeterDerivedFieldsDataSourceModel] `tfsdk:"derived_fields" json:"derivedFields,computed"`
}

func (m *MeterDataSourceModel) toReadParams(_ context.Context) (params m3ter.MeterGetParams, diags diag.Diagnostics) {
	params = m3ter.MeterGetParams{
		OrgID: m3ter.F(m.OrgID.ValueString()),
	}

	return
}

type MeterDataFieldsDataSourceModel struct {
	Category types.String `tfsdk:"category" json:"category,computed"`
	Code     types.String `tfsdk:"code" json:"code,computed"`
	Name     types.String `tfsdk:"name" json:"name,computed"`
	Unit     types.String `tfsdk:"unit" json:"unit,computed"`
}

type MeterDerivedFieldsDataSourceModel struct {
	Calculation types.String `tfsdk:"calculation" json:"calculation,computed"`
	Category    types.String `tfsdk:"category" json:"category,computed"`
	Code        types.String `tfsdk:"code" json:"code,computed"`
	Name        types.String `tfsdk:"name" json:"name,computed"`
	Unit        types.String `tfsdk:"unit" json:"unit,computed"`
}
