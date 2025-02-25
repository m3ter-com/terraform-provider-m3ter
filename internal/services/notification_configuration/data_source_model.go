// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package notification_configuration

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/m3ter-sdk-go"
)

type NotificationConfigurationDataSourceModel struct {
	ID              types.String      `tfsdk:"id" path:"id,required"`
	OrgID           types.String      `tfsdk:"org_id" path:"orgId,required"`
	Active          types.Bool        `tfsdk:"active" json:"active,computed"`
	AlwaysFireEvent types.Bool        `tfsdk:"always_fire_event" json:"alwaysFireEvent,computed"`
	Calculation     types.String      `tfsdk:"calculation" json:"calculation,computed"`
	Code            types.String      `tfsdk:"code" json:"code,computed"`
	CreatedBy       types.String      `tfsdk:"created_by" json:"createdBy,computed"`
	Description     types.String      `tfsdk:"description" json:"description,computed"`
	DtCreated       timetypes.RFC3339 `tfsdk:"dt_created" json:"dtCreated,computed" format:"date-time"`
	DtLastModified  timetypes.RFC3339 `tfsdk:"dt_last_modified" json:"dtLastModified,computed" format:"date-time"`
	EventName       types.String      `tfsdk:"event_name" json:"eventName,computed"`
	LastModifiedBy  types.String      `tfsdk:"last_modified_by" json:"lastModifiedBy,computed"`
	Name            types.String      `tfsdk:"name" json:"name,computed"`
	Version         types.Int64       `tfsdk:"version" json:"version,computed"`
}

func (m *NotificationConfigurationDataSourceModel) toReadParams(_ context.Context) (params m3ter.NotificationConfigurationGetParams, diags diag.Diagnostics) {
	params = m3ter.NotificationConfigurationGetParams{
		OrgID: m3ter.F(m.OrgID.ValueString()),
	}

	return
}
