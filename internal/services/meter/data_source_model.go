// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package meter

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/m3ter-sdk-go"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

type MeterDataListDataSourceEnvelope struct {
	Data customfield.NestedObjectList[MeterDataSourceModel] `json:"data,computed"`
}

type MeterDataSourceModel struct {
	OrgID          types.String                                                    `tfsdk:"org_id" path:"orgId,optional"`
	ID             types.String                                                    `tfsdk:"id" path:"id,computed_optional"`
	Code           types.String                                                    `tfsdk:"code" json:"code,computed"`
	CreatedBy      types.String                                                    `tfsdk:"created_by" json:"createdBy,computed"`
	DtCreated      timetypes.RFC3339                                               `tfsdk:"dt_created" json:"dtCreated,computed" format:"date-time"`
	DtLastModified timetypes.RFC3339                                               `tfsdk:"dt_last_modified" json:"dtLastModified,computed" format:"date-time"`
	GroupID        types.String                                                    `tfsdk:"group_id" json:"groupId,computed"`
	LastModifiedBy types.String                                                    `tfsdk:"last_modified_by" json:"lastModifiedBy,computed"`
	Name           types.String                                                    `tfsdk:"name" json:"name,computed"`
	ProductID      types.String                                                    `tfsdk:"product_id" json:"productId,computed"`
	Version        types.Int64                                                     `tfsdk:"version" json:"version,computed"`
	CustomFields   customfield.Map[jsontypes.Normalized]                           `tfsdk:"custom_fields" json:"customFields,computed"`
	DataFields     customfield.NestedObjectList[MeterDataFieldsDataSourceModel]    `tfsdk:"data_fields" json:"dataFields,computed"`
	DerivedFields  customfield.NestedObjectList[MeterDerivedFieldsDataSourceModel] `tfsdk:"derived_fields" json:"derivedFields,computed"`
	FindOneBy      *MeterFindOneByDataSourceModel                                  `tfsdk:"find_one_by"`
}

func (m *MeterDataSourceModel) toListParams(_ context.Context) (params m3ter.MeterListParams, diags diag.Diagnostics) {
	mFindOneByCodes := []string{}
	for _, item := range *m.FindOneBy.Codes {
		mFindOneByCodes = append(mFindOneByCodes, item.ValueString())
	}
	mFindOneByIDs := []string{}
	for _, item := range *m.FindOneBy.IDs {
		mFindOneByIDs = append(mFindOneByIDs, item.ValueString())
	}
	mFindOneByProductID := []interface{}{}
	for _, item := range *m.FindOneBy.ProductID {
		mFindOneByProductID = append(mFindOneByProductID, item.ValueString())
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
	Category types.String `tfsdk:"category" json:"category,computed"`
	Code     types.String `tfsdk:"code" json:"code,computed"`
	Name     types.String `tfsdk:"name" json:"name,computed"`
	Unit     types.String `tfsdk:"unit" json:"unit,computed"`
}

type MeterFindOneByDataSourceModel struct {
	OrgID     types.String            `tfsdk:"org_id" path:"orgId,required"`
	Codes     *[]types.String         `tfsdk:"codes" query:"codes,optional"`
	IDs       *[]types.String         `tfsdk:"ids" query:"ids,optional"`
	ProductID *[]jsontypes.Normalized `tfsdk:"product_id" query:"productId,optional"`
}
