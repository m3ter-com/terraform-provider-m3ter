// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package meter

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/m3ter-sdk-go"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

type MeterDataSourceModel struct {
	ID            types.String                                                    `tfsdk:"id" path:"id,computed_optional"`
	OrgID         types.String                                                    `tfsdk:"org_id" path:"orgId,required"`
	Code          types.String                                                    `tfsdk:"code" json:"code,computed"`
	GroupID       types.String                                                    `tfsdk:"group_id" json:"groupId,computed"`
	Name          types.String                                                    `tfsdk:"name" json:"name,computed"`
	ProductID     types.String                                                    `tfsdk:"product_id" json:"productId,computed"`
	Version       types.Int64                                                     `tfsdk:"version" json:"version,computed,force_encode,encode_state_for_unknown"`
	DataFields    customfield.NestedObjectList[MeterDataFieldsDataSourceModel]    `tfsdk:"data_fields" json:"dataFields,computed"`
	DerivedFields customfield.NestedObjectList[MeterDerivedFieldsDataSourceModel] `tfsdk:"derived_fields" json:"derivedFields,computed"`
	CustomFields  customfield.NormalizedDynamicValue                              `tfsdk:"custom_fields" json:"customFields,computed"`
	FindOneBy     *MeterFindOneByDataSourceModel                                  `tfsdk:"find_one_by"`
}

func (m *MeterDataSourceModel) toReadParams(_ context.Context) (params m3ter.MeterGetParams, diags diag.Diagnostics) {
	params = m3ter.MeterGetParams{}

	if !m.OrgID.IsNull() {
		params.OrgID = m3ter.F(m.OrgID.ValueString())
	}

	return
}

func (m *MeterDataSourceModel) toListParams(_ context.Context) (params m3ter.MeterListParams, diags diag.Diagnostics) {
	mFindOneByCodes := []string{}
	if m.FindOneBy.Codes != nil {
		for _, item := range *m.FindOneBy.Codes {
			mFindOneByCodes = append(mFindOneByCodes, item.ValueString())
		}
	}
	mFindOneByIDs := []string{}
	if m.FindOneBy.IDs != nil {
		for _, item := range *m.FindOneBy.IDs {
			mFindOneByIDs = append(mFindOneByIDs, item.ValueString())
		}
	}
	mFindOneByProductID := []string{}
	if m.FindOneBy.ProductID != nil {
		for _, item := range *m.FindOneBy.ProductID {
			mFindOneByProductID = append(mFindOneByProductID, item.ValueString())
		}
	}

	params = m3ter.MeterListParams{
		Codes:     m3ter.F(mFindOneByCodes),
		IDs:       m3ter.F(mFindOneByIDs),
		ProductID: m3ter.F(mFindOneByProductID),
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
	Category    types.String `tfsdk:"category" json:"category,computed"`
	Code        types.String `tfsdk:"code" json:"code,computed"`
	Name        types.String `tfsdk:"name" json:"name,computed"`
	Unit        types.String `tfsdk:"unit" json:"unit,computed"`
	Calculation types.String `tfsdk:"calculation" json:"calculation,computed"`
}

type MeterFindOneByDataSourceModel struct {
	Codes     *[]types.String `tfsdk:"codes" query:"codes,optional"`
	IDs       *[]types.String `tfsdk:"ids" query:"ids,optional"`
	ProductID *[]types.String `tfsdk:"product_id" query:"productId,optional"`
}
