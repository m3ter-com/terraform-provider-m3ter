// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package scheduled_event_configuration

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/apijson"
)

type ScheduledEventConfigurationModel struct {
	ID      types.String `tfsdk:"id" json:"id,computed"`
	OrgID   types.String `tfsdk:"org_id" path:"orgId,optional"`
	Entity  types.String `tfsdk:"entity" json:"entity,required"`
	Field   types.String `tfsdk:"field" json:"field,required"`
	Name    types.String `tfsdk:"name" json:"name,required"`
	Offset  types.Int64  `tfsdk:"offset" json:"offset,required"`
	Version types.Int64  `tfsdk:"version" json:"version,computed,force_encode,encode_state_for_unknown"`
}

func (m ScheduledEventConfigurationModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ScheduledEventConfigurationModel) MarshalJSONForUpdate(state ScheduledEventConfigurationModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
