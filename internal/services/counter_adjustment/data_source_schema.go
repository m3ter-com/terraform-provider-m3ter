// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package counter_adjustment

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = (*CounterAdjustmentDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Required: true,
			},
			"org_id": schema.StringAttribute{
				Required:           true,
				DeprecationMessage: "the org id should be set at the client level instead",
			},
			"account_id": schema.StringAttribute{
				Description: "The Account ID the CounterAdjustment was created for.",
				Computed:    true,
			},
			"counter_id": schema.StringAttribute{
				Description: "The ID of the Counter that was used to make the CounterAdjustment on the Account.",
				Computed:    true,
			},
			"date": schema.StringAttribute{
				Description: "The date the CounterAdjustment was created for the Account *(in ISO-8601 date format)*.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"purchase_order_number": schema.StringAttribute{
				Description: "Purchase Order Number for the Counter Adjustment. *(Optional)*",
				Computed:    true,
			},
			"value": schema.Int64Attribute{
				Description: "Integer Value of the Counter that was used to make the CounterAdjustment.",
				Computed:    true,
			},
			"version": schema.Int64Attribute{
				Description: "The version number:\n- **Create:** On initial Create to insert a new entity, the version is set at 1 in the response.\n- **Update:** On successful Update, the version is incremented by 1 in the response.",
				Computed:    true,
			},
		},
	}
}

func (d *CounterAdjustmentDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *CounterAdjustmentDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
