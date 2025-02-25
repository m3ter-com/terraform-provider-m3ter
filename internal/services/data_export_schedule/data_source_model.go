// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package data_export_schedule

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/m3ter-sdk-go"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

type DataExportScheduleDataSourceModel struct {
	ID                   types.String                   `tfsdk:"id" path:"id,required"`
	OrgID                types.String                   `tfsdk:"org_id" path:"orgId,required"`
	Aggregation          types.String                   `tfsdk:"aggregation" json:"aggregation,computed"`
	AggregationFrequency types.String                   `tfsdk:"aggregation_frequency" json:"aggregationFrequency,computed"`
	TimePeriod           types.String                   `tfsdk:"time_period" json:"timePeriod,computed"`
	Version              types.Int64                    `tfsdk:"version" json:"version,computed"`
	AccountIDs           customfield.List[types.String] `tfsdk:"account_ids" json:"accountIds,computed"`
	MeterIDs             customfield.List[types.String] `tfsdk:"meter_ids" json:"meterIds,computed"`
	OperationalDataTypes customfield.List[types.String] `tfsdk:"operational_data_types" json:"operationalDataTypes,computed"`
}

func (m *DataExportScheduleDataSourceModel) toReadParams(_ context.Context) (params m3ter.DataExportScheduleGetParams, diags diag.Diagnostics) {
	params = m3ter.DataExportScheduleGetParams{
		OrgID: m3ter.F(m.OrgID.ValueString()),
	}

	return
}
