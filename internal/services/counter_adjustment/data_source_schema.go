// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package counter_adjustment

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
)

var _ datasource.DataSourceWithConfigValidators = (*CounterAdjustmentDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
				Optional: true,
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
			"find_one_by": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"account_id": schema.StringAttribute{
						Description: "List CounterAdjustment items for the Account UUID.",
						Optional:    true,
					},
					"counter_id": schema.StringAttribute{
						Description: "List CounterAdjustment items for the Counter UUID.",
						Optional:    true,
					},
					"date": schema.StringAttribute{
						Description: "List CounterAdjustment items for the given date.",
						Optional:    true,
					},
					"date_end": schema.StringAttribute{
						Optional: true,
					},
					"date_start": schema.StringAttribute{
						Optional: true,
					},
					"end_date_end": schema.StringAttribute{
						Description: "Only include CounterAdjustments with end dates earlier than this date.",
						Optional:    true,
					},
					"end_date_start": schema.StringAttribute{
						Description: "Only include CounterAdjustments with end dates equal to or later than this date.",
						Optional:    true,
					},
					"sort_order": schema.StringAttribute{
						Description: "Sort order for the results",
						Optional:    true,
					},
				},
			},
		},
	}
}

func (d *CounterAdjustmentDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *CounterAdjustmentDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("id"), path.MatchRoot("find_one_by")),
	}
}
