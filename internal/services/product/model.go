// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package product

import (
  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework/types"
  "github.com/m3ter-com/terraform-provider-m3ter/internal/apijson"
)

type ProductModel struct {
ID types.String `tfsdk:"id" json:"id,computed"`
OrgID types.String `tfsdk:"org_id" path:"orgId,optional"`
Code types.String `tfsdk:"code" json:"code,required"`
Name types.String `tfsdk:"name" json:"name,required"`
Version types.Int64 `tfsdk:"version" json:"version,optional"`
CustomFields *map[string]types.Dynamic `tfsdk:"custom_fields" json:"customFields,optional"`
CreatedBy types.String `tfsdk:"created_by" json:"createdBy,computed"`
DtCreated timetypes.RFC3339 `tfsdk:"dt_created" json:"dtCreated,computed" format:"date-time"`
DtLastModified timetypes.RFC3339 `tfsdk:"dt_last_modified" json:"dtLastModified,computed" format:"date-time"`
LastModifiedBy types.String `tfsdk:"last_modified_by" json:"lastModifiedBy,computed"`
}

func (m ProductModel) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(m)
}

func (m ProductModel) MarshalJSONForUpdate(state ProductModel) (data []byte, err error) {
  return apijson.MarshalForUpdate(m, state)
}
