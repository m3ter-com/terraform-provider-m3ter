// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package counter_adjustment

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

var _ datasource.DataSourceWithConfigValidators = (*CounterAdjustmentsDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"org_id": schema.StringAttribute{
				Required: true,
			},
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
			"max_items": schema.Int64Attribute{
				Description: "Max items to fetch, default: 1000",
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
			},
			"items": schema.ListNestedAttribute{
				Description: "The items returned by the data source",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[CounterAdjustmentsItemsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The UUID of the entity. ",
							Computed:    true,
						},
						"version": schema.Int64Attribute{
							Description: "The version number:\n- **Create:** On initial Create to insert a new entity, the version is set at 1 in the response.\n- **Update:** On successful Update, the version is incremented by 1 in the response.",
							Computed:    true,
						},
						"account_id": schema.StringAttribute{
							Description: "The Account ID the CounterAdjustment was created for.",
							Computed:    true,
						},
						"counter_id": schema.StringAttribute{
							Description: "The ID of the Counter that was used to make the CounterAdjustment on the Account.",
							Computed:    true,
						},
						"created_by": schema.StringAttribute{
							Description: "The ID of the user who created this item.",
							Computed:    true,
						},
						"date": schema.StringAttribute{
							Description: "The date the CounterAdjustment was created for the Account *(in ISO-8601 date format)*.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"dt_created": schema.StringAttribute{
							Description: "The DateTime when this item was created *(in ISO-8601 format)*.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"dt_last_modified": schema.StringAttribute{
							Description: "The DateTime when this item was last modified *(in ISO-8601 format)*.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"last_modified_by": schema.StringAttribute{
							Description: "The ID of the user who last modified this item.",
							Computed:    true,
						},
						"purchase_order_number": schema.StringAttribute{
							Description: "Purchase Order Number for the Counter Adjustment. *(Optional)*",
							Computed:    true,
						},
						"value": schema.Int64Attribute{
							Description: "Integer Value of the Counter that was used to make the CounterAdjustment.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func (d *CounterAdjustmentsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *CounterAdjustmentsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
