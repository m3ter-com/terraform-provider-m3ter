// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package counter

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/m3ter-sdk-go"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

type CounterDataListDataSourceEnvelope struct {
	Data customfield.NestedObjectList[CounterDataSourceModel] `json:"data,computed"`
}

type CounterDataSourceModel struct {
	OrgID          types.String                     `tfsdk:"org_id" path:"orgId,optional"`
	ID             types.String                     `tfsdk:"id" path:"id,computed_optional"`
	Code           types.String                     `tfsdk:"code" json:"code,computed"`
	CreatedBy      types.String                     `tfsdk:"created_by" json:"createdBy,computed"`
	DtCreated      timetypes.RFC3339                `tfsdk:"dt_created" json:"dtCreated,computed" format:"date-time"`
	DtLastModified timetypes.RFC3339                `tfsdk:"dt_last_modified" json:"dtLastModified,computed" format:"date-time"`
	LastModifiedBy types.String                     `tfsdk:"last_modified_by" json:"lastModifiedBy,computed"`
	Name           types.String                     `tfsdk:"name" json:"name,computed"`
	ProductID      types.String                     `tfsdk:"product_id" json:"productId,computed"`
	Unit           types.String                     `tfsdk:"unit" json:"unit,computed"`
	Version        types.Int64                      `tfsdk:"version" json:"version,computed"`
	FindOneBy      *CounterFindOneByDataSourceModel `tfsdk:"find_one_by"`
}

func (m *CounterDataSourceModel) toListParams(_ context.Context) (params m3ter.CounterListParams, diags diag.Diagnostics) {
	mFindOneByCodes := []string{}
	for _, item := range *m.FindOneBy.Codes {
		mFindOneByCodes = append(mFindOneByCodes, item.ValueString())
	}
	mFindOneByIDs := []string{}
	for _, item := range *m.FindOneBy.IDs {
		mFindOneByIDs = append(mFindOneByIDs, item.ValueString())
	}
	mFindOneByProductID := []string{}
	for _, item := range *m.FindOneBy.ProductID {
		mFindOneByProductID = append(mFindOneByProductID, item.ValueString())
	}

	params = m3ter.CounterListParams{
		Codes:     m3ter.F(mFindOneByCodes),
		IDs:       m3ter.F(mFindOneByIDs),
		ProductID: m3ter.F(mFindOneByProductID),
	}

	return
}

type CounterFindOneByDataSourceModel struct {
	OrgID     types.String    `tfsdk:"org_id" path:"orgId,required"`
	Codes     *[]types.String `tfsdk:"codes" query:"codes,optional"`
	IDs       *[]types.String `tfsdk:"ids" query:"ids,optional"`
	ProductID *[]types.String `tfsdk:"product_id" query:"productId,optional"`
}
