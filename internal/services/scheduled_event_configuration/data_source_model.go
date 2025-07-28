// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package scheduled_event_configuration

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/m3ter-sdk-go"
)

type ScheduledEventConfigurationDataSourceModel struct {
	ID      types.String `tfsdk:"id" path:"id,required"`
	OrgID   types.String `tfsdk:"org_id" path:"orgId,required"`
	Entity  types.String `tfsdk:"entity" json:"entity,computed"`
	Field   types.String `tfsdk:"field" json:"field,computed"`
	Name    types.String `tfsdk:"name" json:"name,computed"`
	Offset  types.Int64  `tfsdk:"offset" json:"offset,computed"`
	Version types.Int64  `tfsdk:"version" json:"version,computed,force_encode,encode_state_for_unknown"`
}

func (m *ScheduledEventConfigurationDataSourceModel) toReadParams(_ context.Context) (params m3ter.ScheduledEventConfigurationGetParams, diags diag.Diagnostics) {
	params = m3ter.ScheduledEventConfigurationGetParams{}

	if !m.OrgID.IsNull() {
		params.OrgID = m3ter.F(m.OrgID.ValueString())
	}

	return
}
