// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package bill_line_item

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

var _ datasource.DataSourceWithConfigValidators = (*BillLineItemDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"bill_id": schema.StringAttribute{
				Required: true,
			},
			"id": schema.StringAttribute{
				Required: true,
			},
			"org_id": schema.StringAttribute{
				Required: true,
			},
			"aggregation_id": schema.StringAttribute{
				Description: "A unique identifier (UUID) for the Aggregation that contributes to this Bill line item.",
				Computed:    true,
			},
			"average_unit_price": schema.Float64Attribute{
				Description: "Represents the average unit price calculated across all pricing bands or tiers for this line item.",
				Computed:    true,
			},
			"balance_id": schema.StringAttribute{
				Computed: true,
			},
			"charge_id": schema.StringAttribute{
				Computed: true,
			},
			"commitment_id": schema.StringAttribute{
				Description: "The unique identifier (UUID) of the Commitment *(if this is used)*.",
				Computed:    true,
			},
			"compound_aggregation_id": schema.StringAttribute{
				Description: "A unique identifier (UUID) for the Compound Aggregation, if applicable.",
				Computed:    true,
			},
			"contract_id": schema.StringAttribute{
				Description: "The unique identifier (UUID) for the contract associated with this line item.",
				Computed:    true,
			},
			"conversion_rate": schema.Float64Attribute{
				Description: "The currency conversion rate *(if used)* for the line item.",
				Computed:    true,
			},
			"converted_subtotal": schema.Float64Attribute{
				Description: "The subtotal amount for this line item after currency conversion, if applicable.",
				Computed:    true,
			},
			"counter_id": schema.StringAttribute{
				Computed: true,
			},
			"created_by": schema.StringAttribute{
				Description: "The unique identifier (UUID) for the user who created the Bill line item.",
				Computed:    true,
			},
			"credit_type_id": schema.StringAttribute{
				Description: "The unique identifier (UUID) for the type of credit applied to this line item.",
				Computed:    true,
			},
			"currency": schema.StringAttribute{
				Description: "The currency in which the line item is billed, represented as a currency code. For example, USD, GBP, or EUR.",
				Computed:    true,
			},
			"description": schema.StringAttribute{
				Description: "A detailed description providing context for the line item within the Bill.",
				Computed:    true,
			},
			"dt_created": schema.StringAttribute{
				Description: "The date and time *(in ISO 8601 format)* when the Bill line item was first created.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"dt_last_modified": schema.StringAttribute{
				Description: "The date and time *(in ISO 8601 format)* when the Bill line item was last modified.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"json_usage_generated": schema.BoolAttribute{
				Description: "Boolean flag indicating whether the Bill line item has associated statement usage in JSON format. When a Bill statement is generated, usage line items have their usage stored in JSON format.\n\nSee [Working with Bill Statements](https://www.m3ter.com/docs/guides/running-viewing-and-managing-bills/working-with-bill-statements) for more information.",
				Computed:    true,
			},
			"last_modified_by": schema.StringAttribute{
				Description: "The unique identifier (UUID) for the user who last modified this Bill line item.",
				Computed:    true,
			},
			"line_item_type": schema.StringAttribute{
				Description: `Available values: "STANDING_CHARGE", "USAGE", "COUNTER_RUNNING_TOTAL_CHARGE", "COUNTER_ADJUSTMENT_DEBIT", "COUNTER_ADJUSTMENT_CREDIT", "USAGE_CREDIT", "MINIMUM_SPEND", "MINIMUM_SPEND_REFUND", "CREDIT_DEDUCTION", "MANUAL_ADJUSTMENT", "CREDIT_MEMO", "DEBIT_MEMO", "COMMITMENT_CONSUMED", "COMMITMENT_FEE", "OVERAGE_SURCHARGE", "OVERAGE_USAGE", "BALANCE_CONSUMED", "BALANCE_FEE".`,
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"STANDING_CHARGE",
						"USAGE",
						"COUNTER_RUNNING_TOTAL_CHARGE",
						"COUNTER_ADJUSTMENT_DEBIT",
						"COUNTER_ADJUSTMENT_CREDIT",
						"USAGE_CREDIT",
						"MINIMUM_SPEND",
						"MINIMUM_SPEND_REFUND",
						"CREDIT_DEDUCTION",
						"MANUAL_ADJUSTMENT",
						"CREDIT_MEMO",
						"DEBIT_MEMO",
						"COMMITMENT_CONSUMED",
						"COMMITMENT_FEE",
						"OVERAGE_SURCHARGE",
						"OVERAGE_USAGE",
						"BALANCE_CONSUMED",
						"BALANCE_FEE",
					),
				},
			},
			"meter_id": schema.StringAttribute{
				Description: "The unique identifier (UUID) of the Meter responsible for tracking usage.",
				Computed:    true,
			},
			"plan_group_id": schema.StringAttribute{
				Description: "The UUID of the PlanGroup. \n\nThe unique identifier (UUID) for the PlanGroup, if applicable.",
				Computed:    true,
			},
			"plan_id": schema.StringAttribute{
				Description: "A unique identifier (UUID) for the billing Plan associated with this line item,",
				Computed:    true,
			},
			"pricing_id": schema.StringAttribute{
				Description: "The unique identifier (UUID) of the Pricing used for this line item,",
				Computed:    true,
			},
			"product_code": schema.StringAttribute{
				Computed: true,
			},
			"product_id": schema.StringAttribute{
				Description: "The unique identifier (UUID) for the associated Product.",
				Computed:    true,
			},
			"product_name": schema.StringAttribute{
				Description: "The name of the Product associated with this line item.",
				Computed:    true,
			},
			"quantity": schema.Float64Attribute{
				Description: "The amount of the product or service used in this line item.",
				Computed:    true,
			},
			"reason_id": schema.StringAttribute{
				Description: "The UUID of the reason used for the line item. \n\nA unique identifier (UUID) for the reason or justification for this line item, if applicable.",
				Computed:    true,
			},
			"referenced_bill_id": schema.StringAttribute{
				Description: "A unique identifier (UUID) for a Bill that this line item may be related to or derived from.",
				Computed:    true,
			},
			"referenced_line_item_id": schema.StringAttribute{
				Description: "A unique identifier (UUID) for another line item that this line item may be related to or derived from.",
				Computed:    true,
			},
			"sequence_number": schema.Int64Attribute{
				Description: "The number used for sequential invoices.",
				Computed:    true,
			},
			"service_period_end_date": schema.StringAttribute{
				Description: "The *(exclusive)* end date for the service period in ISO 68601 format.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"service_period_start_date": schema.StringAttribute{
				Description: "The *(inclusive)* start date for the service period in ISO 8601 format.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"subtotal": schema.Float64Attribute{
				Description: "The subtotal amount when not currency converted *(in the cases where currency conversion is required)*.",
				Computed:    true,
			},
			"unit": schema.StringAttribute{
				Description: "Specifies the unit type. For example: **MB**, **GB**, **api_calls**, and so on.",
				Computed:    true,
			},
			"units": schema.Float64Attribute{
				Description: "The number of units rated in the line item, each of which is of the type specified in the `unit` field. For example: 400 api_calls. \n\nIn this example, the unit type of **api_calls** is read from the `unit` field.",
				Computed:    true,
			},
			"version": schema.Int64Attribute{
				Description: "The version number:\n- **Create:** On initial Create to insert a new entity, the version is set at 1 in the response.\n- **Update:** On successful Update, the version is incremented by 1 in the response.",
				Computed:    true,
			},
			"group": schema.MapAttribute{
				Computed:    true,
				CustomType:  customfield.NewMapType[types.String](ctx),
				ElementType: types.StringType,
			},
			"segment": schema.MapAttribute{
				Description: "Specifies the segment name or identifier when segmented Aggregation is used. This is relevant for more complex billing structures.",
				Computed:    true,
				CustomType:  customfield.NewMapType[types.String](ctx),
				ElementType: types.StringType,
			},
			"band_usage": schema.ListNestedAttribute{
				Description: "Array containing the pricing band information, which shows the details for each pricing band or tier.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[BillLineItemBandUsageDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"band_quantity": schema.Float64Attribute{
							Description: "Usage amount within the band.",
							Computed:    true,
						},
						"band_subtotal": schema.Float64Attribute{
							Description: "Subtotal amount for the band.",
							Computed:    true,
						},
						"band_units": schema.Float64Attribute{
							Description: "The number of units used within the band.",
							Computed:    true,
						},
						"credit_type_id": schema.StringAttribute{
							Description: "The UUID of the credit type.",
							Computed:    true,
						},
						"fixed_price": schema.Float64Attribute{
							Description: "Fixed price is a charge entered for certain pricing types such as Stairstep, Custom Tiered, and Custom Volume. It is a set price and not dependent on usage.",
							Computed:    true,
						},
						"lower_limit": schema.Float64Attribute{
							Description: "The lower limit *(start)* of the pricing band.",
							Computed:    true,
						},
						"pricing_band_id": schema.StringAttribute{
							Description: "The UUID for the pricing band.",
							Computed:    true,
						},
						"unit_price": schema.Float64Attribute{
							Description: "The price per unit in the band.",
							Computed:    true,
						},
						"unit_subtotal": schema.Float64Attribute{
							Description: "The subtotal of the unit usage.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func (d *BillLineItemDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *BillLineItemDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
