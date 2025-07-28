// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package notification_configuration

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/m3ter-sdk-go"
)

type NotificationConfigurationDataSourceModel struct {
	ID              types.String `tfsdk:"id" path:"id,required"`
	OrgID           types.String `tfsdk:"org_id" path:"orgId,required"`
	Active          types.Bool   `tfsdk:"active" json:"active,computed"`
	AlwaysFireEvent types.Bool   `tfsdk:"always_fire_event" json:"alwaysFireEvent,computed"`
	Calculation     types.String `tfsdk:"calculation" json:"calculation,computed"`
	Code            types.String `tfsdk:"code" json:"code,computed"`
	Description     types.String `tfsdk:"description" json:"description,computed"`
	EventName       types.String `tfsdk:"event_name" json:"eventName,computed"`
	Name            types.String `tfsdk:"name" json:"name,computed"`
	Version         types.Int64  `tfsdk:"version" json:"version,computed,force_encode,encode_state_for_unknown"`
}

func (m *NotificationConfigurationDataSourceModel) toReadParams(_ context.Context) (params m3ter.NotificationConfigurationGetParams, diags diag.Diagnostics) {
	params = m3ter.NotificationConfigurationGetParams{}

	if !m.OrgID.IsNull() {
		params.OrgID = m3ter.F(m.OrgID.ValueString())
	}

	return
}
