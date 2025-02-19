// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package balance

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

var _ datasource.DataSourceWithConfigValidators = (*BalancesDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"org_id": schema.StringAttribute{
				Required: true,
			},
			"account_id": schema.StringAttribute{
				Description: "The unique identifier (UUID) for the end customer's account.",
				Optional:    true,
			},
			"end_date_end": schema.StringAttribute{
				Description: "Only include Balances with end dates earlier than this date.",
				Optional:    true,
			},
			"end_date_start": schema.StringAttribute{
				Description: "Only include Balances with end dates equal to or later than this date.",
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
				CustomType:  customfield.NewNestedObjectListType[BalancesItemsDataSourceModel](ctx),
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
							Description: "The unique identifier (UUID) for the end customer Account the Balance belongs to.",
							Computed:    true,
						},
						"amount": schema.Float64Attribute{
							Description: "The financial value that the Balance holds.",
							Computed:    true,
						},
						"balance_draw_down_description": schema.StringAttribute{
							Description: "A description for the bill line items for charges drawn-down against the Balance.",
							Computed:    true,
						},
						"code": schema.StringAttribute{
							Description: "A unique short code assigned to the Balance.",
							Computed:    true,
						},
						"consumptions_accounting_product_id": schema.StringAttribute{
							Computed: true,
						},
						"created_by": schema.StringAttribute{
							Description: "The unique identifier (UUID) for the user who created the Balance.",
							Computed:    true,
						},
						"currency": schema.StringAttribute{
							Description: "The currency code used for the Balance amount. For example: USD, GBP or EUR.",
							Computed:    true,
						},
						"description": schema.StringAttribute{
							Description: "A description of the Balance.",
							Computed:    true,
						},
						"dt_created": schema.StringAttribute{
							Description: "The date and time *(in ISO 8601 format)* when the Balance was first created.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"dt_last_modified": schema.StringAttribute{
							Description: "The date and time *(in ISO 8601 format)* when the Balance was last modified.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"end_date": schema.StringAttribute{
							Description: "The date *(in ISO 8601 format)* after which the Balance will no longer be active.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"fees_accounting_product_id": schema.StringAttribute{
							Computed: true,
						},
						"last_modified_by": schema.StringAttribute{
							Description: "The unique identifier (UUID) for the user who last modified the Balance.",
							Computed:    true,
						},
						"line_item_types": schema.ListAttribute{
							Description: "A list of line item charge types that can draw-down against the Balance amount at billing. ",
							Computed:    true,
							Validators: []validator.List{
								listvalidator.ValueStringsAre(
									stringvalidator.OneOfCaseInsensitive(
										"STANDING_CHARGE",
										"USAGE",
										"MINIMUM_SPEND",
										"COUNTER_RUNNING_TOTAL_CHARGE",
										"COUNTER_ADJUSTMENT_DEBIT",
									),
								),
							},
							CustomType:  customfield.NewListType[types.String](ctx),
							ElementType: types.StringType,
						},
						"name": schema.StringAttribute{
							Description: "The official name of the Balance.",
							Computed:    true,
						},
						"overage_description": schema.StringAttribute{
							Description: "A description for overage charges.",
							Computed:    true,
						},
						"overage_surcharge_percent": schema.Float64Attribute{
							Description: "The percentage surcharge applied to overage charges *(usage above the Balance)*.",
							Computed:    true,
						},
						"product_ids": schema.ListAttribute{
							Description: "A list of Product IDs whose consumption charges due at billing can be drawn-down against the Balance amount.",
							Computed:    true,
							CustomType:  customfield.NewListType[types.String](ctx),
							ElementType: types.StringType,
						},
						"rollover_amount": schema.Float64Attribute{
							Description: "The maximum amount that can be carried over past the Balance end date and draw-down against for billing if there is an unused Balance amount remaining when the Balance end date is reached.\n",
							Computed:    true,
						},
						"rollover_end_date": schema.StringAttribute{
							Description: "The end date *(in ISO 8601 format)* for the rollover grace period, which is the period that unused Balance amounts can be carried over beyond the specified Balance `endDate` and continue to be drawn-down against for billing.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"start_date": schema.StringAttribute{
							Description: "The date *(in ISO 8601 format)* when the Balance becomes active.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
					},
				},
			},
		},
	}
}

func (d *BalancesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *BalancesDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
