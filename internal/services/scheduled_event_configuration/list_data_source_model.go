// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package scheduled_event_configuration

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/m3ter-sdk-go"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

type ScheduledEventConfigurationsDataListDataSourceEnvelope struct {
	Data customfield.NestedObjectList[ScheduledEventConfigurationsItemsDataSourceModel] `json:"data,computed"`
}

type ScheduledEventConfigurationsDataSourceModel struct {
	OrgID    types.String                                                                   `tfsdk:"org_id" path:"orgId,optional"`
	IDs      *[]types.String                                                                `tfsdk:"ids" query:"ids,optional"`
	MaxItems types.Int64                                                                    `tfsdk:"max_items"`
	Items    customfield.NestedObjectList[ScheduledEventConfigurationsItemsDataSourceModel] `tfsdk:"items"`
}

func (m *ScheduledEventConfigurationsDataSourceModel) toListParams(_ context.Context) (params m3ter.ScheduledEventConfigurationListParams, diags diag.Diagnostics) {
	mIDs := []string{}
	for _, item := range *m.IDs {
		mIDs = append(mIDs, item.ValueString())
	}

	params = m3ter.ScheduledEventConfigurationListParams{
		IDs: m3ter.F(mIDs),
	}

	if !m.OrgID.IsNull() {
		params.OrgID = m3ter.F(m.OrgID.ValueString())
	}

	return
}

type ScheduledEventConfigurationsItemsDataSourceModel struct {
	ID      types.String `tfsdk:"id" json:"id,computed"`
	Entity  types.String `tfsdk:"entity" json:"entity,computed"`
	Field   types.String `tfsdk:"field" json:"field,computed"`
	Name    types.String `tfsdk:"name" json:"name,computed"`
	Offset  types.Int64  `tfsdk:"offset" json:"offset,computed"`
	Version types.Int64  `tfsdk:"version" json:"version,computed,force_encode,encode_state_for_unknown"`
}
