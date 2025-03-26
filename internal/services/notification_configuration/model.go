// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package notification_configuration

import (
  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework/types"
  "github.com/m3ter-com/terraform-provider-m3ter/internal/apijson"
)

type NotificationConfigurationModel struct {
ID types.String `tfsdk:"id" json:"id,computed"`
OrgID types.String `tfsdk:"org_id" path:"orgId,optional"`
Code types.String `tfsdk:"code" json:"code,required"`
Description types.String `tfsdk:"description" json:"description,required"`
EventName types.String `tfsdk:"event_name" json:"eventName,required"`
Name types.String `tfsdk:"name" json:"name,required"`
Active types.Bool `tfsdk:"active" json:"active,optional"`
AlwaysFireEvent types.Bool `tfsdk:"always_fire_event" json:"alwaysFireEvent,optional"`
Calculation types.String `tfsdk:"calculation" json:"calculation,optional"`
Version types.Int64 `tfsdk:"version" json:"version,optional"`
CreatedBy types.String `tfsdk:"created_by" json:"createdBy,computed"`
DtCreated timetypes.RFC3339 `tfsdk:"dt_created" json:"dtCreated,computed" format:"date-time"`
DtLastModified timetypes.RFC3339 `tfsdk:"dt_last_modified" json:"dtLastModified,computed" format:"date-time"`
LastModifiedBy types.String `tfsdk:"last_modified_by" json:"lastModifiedBy,computed"`
}

func (m NotificationConfigurationModel) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(m)
}

func (m NotificationConfigurationModel) MarshalJSONForUpdate(state NotificationConfigurationModel) (data []byte, err error) {
  return apijson.MarshalForUpdate(m, state)
}
