// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package transaction_type

import (
  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework/types"
  "github.com/m3ter-com/terraform-provider-m3ter/internal/apijson"
)

type TransactionTypeModel struct {
ID types.String `tfsdk:"id" json:"id,computed"`
OrgID types.String `tfsdk:"org_id" path:"orgId,required"`
Name types.String `tfsdk:"name" json:"name,required"`
Archived types.Bool `tfsdk:"archived" json:"archived,optional"`
Code types.String `tfsdk:"code" json:"code,optional"`
Version types.Int64 `tfsdk:"version" json:"version,optional"`
CreatedBy types.String `tfsdk:"created_by" json:"createdBy,computed"`
DtCreated timetypes.RFC3339 `tfsdk:"dt_created" json:"dtCreated,computed" format:"date-time"`
DtLastModified timetypes.RFC3339 `tfsdk:"dt_last_modified" json:"dtLastModified,computed" format:"date-time"`
LastModifiedBy types.String `tfsdk:"last_modified_by" json:"lastModifiedBy,computed"`
}

func (m TransactionTypeModel) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(m)
}

func (m TransactionTypeModel) MarshalJSONForUpdate(state TransactionTypeModel) (data []byte, err error) {
  return apijson.MarshalForUpdate(m, state)
}
