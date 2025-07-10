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

type MetersDataListDataSourceEnvelope struct {
	Data customfield.NestedObjectList[MetersItemsDataSourceModel] `json:"data,computed"`
}

type MetersDataSourceModel struct {
	OrgID     types.String                                             `tfsdk:"org_id" path:"orgId,optional"`
	Codes     *[]types.String                                          `tfsdk:"codes" query:"codes,optional"`
	IDs       *[]types.String                                          `tfsdk:"ids" query:"ids,optional"`
	ProductID *[]types.String                                          `tfsdk:"product_id" query:"productId,optional"`
	MaxItems  types.Int64                                              `tfsdk:"max_items"`
	Items     customfield.NestedObjectList[MetersItemsDataSourceModel] `tfsdk:"items"`
}

func (m *MetersDataSourceModel) toListParams(_ context.Context) (params m3ter.MeterListParams, diags diag.Diagnostics) {
	mCodes := []string{}
	for _, item := range *m.Codes {
		mCodes = append(mCodes, item.ValueString())
	}
	mIDs := []string{}
	for _, item := range *m.IDs {
		mIDs = append(mIDs, item.ValueString())
	}
	mProductID := []string{}
	for _, item := range *m.ProductID {
		mProductID = append(mProductID, item.ValueString())
	}

	params = m3ter.MeterListParams{
		Codes:     m3ter.F(mCodes),
		IDs:       m3ter.F(mIDs),
		ProductID: m3ter.F(mProductID),
	}

	if !m.OrgID.IsNull() {
		params.OrgID = m3ter.F(m.OrgID.ValueString())
	}

	return
}

type MetersItemsDataSourceModel struct {
	ID             types.String                                                     `tfsdk:"id" json:"id,computed"`
	Code           types.String                                                     `tfsdk:"code" json:"code,computed"`
	CreatedBy      types.String                                                     `tfsdk:"created_by" json:"createdBy,computed"`
	CustomFields   types.Dynamic                                                    `tfsdk:"custom_fields" json:"customFields,computed"`
	DataFields     customfield.NestedObjectList[MetersDataFieldsDataSourceModel]    `tfsdk:"data_fields" json:"dataFields,computed"`
	DerivedFields  customfield.NestedObjectList[MetersDerivedFieldsDataSourceModel] `tfsdk:"derived_fields" json:"derivedFields,computed"`
	DtCreated      timetypes.RFC3339                                                `tfsdk:"dt_created" json:"dtCreated,computed" format:"date-time"`
	DtLastModified timetypes.RFC3339                                                `tfsdk:"dt_last_modified" json:"dtLastModified,computed" format:"date-time"`
	GroupID        types.String                                                     `tfsdk:"group_id" json:"groupId,computed"`
	LastModifiedBy types.String                                                     `tfsdk:"last_modified_by" json:"lastModifiedBy,computed"`
	Name           types.String                                                     `tfsdk:"name" json:"name,computed"`
	ProductID      types.String                                                     `tfsdk:"product_id" json:"productId,computed"`
	Version        types.Int64                                                      `tfsdk:"version" json:"version,computed,force_encode,encode_state_for_unknown"`
}

type MetersDataFieldsDataSourceModel struct {
	Category types.String `tfsdk:"category" json:"category,computed"`
	Code     types.String `tfsdk:"code" json:"code,computed"`
	Name     types.String `tfsdk:"name" json:"name,computed"`
	Unit     types.String `tfsdk:"unit" json:"unit,computed"`
}

type MetersDerivedFieldsDataSourceModel struct {
	Category    types.String `tfsdk:"category" json:"category,computed"`
	Code        types.String `tfsdk:"code" json:"code,computed"`
	Name        types.String `tfsdk:"name" json:"name,computed"`
	Unit        types.String `tfsdk:"unit" json:"unit,computed"`
	Calculation types.String `tfsdk:"calculation" json:"calculation,computed"`
}
