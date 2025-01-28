// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package product

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/m3ter-sdk-go"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

type ProductDataListDataSourceEnvelope struct {
	Data customfield.NestedObjectList[ProductDataSourceModel] `json:"data,computed"`
}

type ProductDataSourceModel struct {
	ID             types.String                          `tfsdk:"id" path:"id,computed_optional"`
	OrgID          types.String                          `tfsdk:"org_id" path:"orgId,required"`
	Code           types.String                          `tfsdk:"code" json:"code,computed"`
	CreatedBy      types.String                          `tfsdk:"created_by" json:"createdBy,computed"`
	DtCreated      timetypes.RFC3339                     `tfsdk:"dt_created" json:"dtCreated,computed" format:"date-time"`
	DtLastModified timetypes.RFC3339                     `tfsdk:"dt_last_modified" json:"dtLastModified,computed" format:"date-time"`
	LastModifiedBy types.String                          `tfsdk:"last_modified_by" json:"lastModifiedBy,computed"`
	Name           types.String                          `tfsdk:"name" json:"name,computed"`
	Version        types.Int64                           `tfsdk:"version" json:"version,computed"`
	CustomFields   customfield.Map[jsontypes.Normalized] `tfsdk:"custom_fields" json:"customFields,computed"`
	FindOneBy      *ProductFindOneByDataSourceModel      `tfsdk:"find_one_by"`
}

func (m *ProductDataSourceModel) toListParams(_ context.Context) (params m3ter.ProductListParams, diags diag.Diagnostics) {
	mFindOneByIDs := []ProductListParamsID{}
	for _, item := range *m.FindOneBy.IDs {
		mFindOneByIDs = append(mFindOneByIDs, item.ValueString())
	}

	params = m3ter.ProductListParams{
		IDs: m3ter.F(mFindOneByIDs),
	}

	return
}

type ProductFindOneByDataSourceModel struct {
	IDs *[]types.String `tfsdk:"ids" query:"ids,optional"`
}
