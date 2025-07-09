// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package notification_configuration

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/m3ter-sdk-go"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

type NotificationConfigurationsDataListDataSourceEnvelope struct {
	Data customfield.NestedObjectList[NotificationConfigurationsItemsDataSourceModel] `json:"data,computed"`
}

type NotificationConfigurationsDataSourceModel struct {
	OrgID     types.String                                                                 `tfsdk:"org_id" path:"orgId,optional"`
	Active    types.Bool                                                                   `tfsdk:"active" query:"active,optional"`
	EventName types.String                                                                 `tfsdk:"event_name" query:"eventName,optional"`
	IDs       *[]types.String                                                              `tfsdk:"ids" query:"ids,optional"`
	MaxItems  types.Int64                                                                  `tfsdk:"max_items"`
	Items     customfield.NestedObjectList[NotificationConfigurationsItemsDataSourceModel] `tfsdk:"items"`
}

func (m *NotificationConfigurationsDataSourceModel) toListParams(_ context.Context) (params m3ter.NotificationConfigurationListParams, diags diag.Diagnostics) {
	mIDs := []string{}
	for _, item := range *m.IDs {
		mIDs = append(mIDs, item.ValueString())
	}

	params = m3ter.NotificationConfigurationListParams{
		IDs: m3ter.F(mIDs),
	}

	if !m.Active.IsNull() {
		params.Active = m3ter.F(m.Active.ValueBool())
	}
	if !m.EventName.IsNull() {
		params.EventName = m3ter.F(m.EventName.ValueString())
	}
	if !m.OrgID.IsNull() {
		params.OrgID = m3ter.F(m.OrgID.ValueString())
	}

	return
}

type NotificationConfigurationsItemsDataSourceModel struct {
	ID              types.String      `tfsdk:"id" json:"id,computed"`
	Code            types.String      `tfsdk:"code" json:"code,computed"`
	Description     types.String      `tfsdk:"description" json:"description,computed"`
	Name            types.String      `tfsdk:"name" json:"name,computed"`
	Active          types.Bool        `tfsdk:"active" json:"active,computed"`
	AlwaysFireEvent types.Bool        `tfsdk:"always_fire_event" json:"alwaysFireEvent,computed"`
	Calculation     types.String      `tfsdk:"calculation" json:"calculation,computed"`
	CreatedBy       types.String      `tfsdk:"created_by" json:"createdBy,computed"`
	DtCreated       timetypes.RFC3339 `tfsdk:"dt_created" json:"dtCreated,computed" format:"date-time"`
	DtLastModified  timetypes.RFC3339 `tfsdk:"dt_last_modified" json:"dtLastModified,computed" format:"date-time"`
	EventName       types.String      `tfsdk:"event_name" json:"eventName,computed"`
	LastModifiedBy  types.String      `tfsdk:"last_modified_by" json:"lastModifiedBy,computed"`
	Version         types.Int64       `tfsdk:"version" json:"version,computed"`
}
