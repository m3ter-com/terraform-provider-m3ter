// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package commitment

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

var _ datasource.DataSourceWithConfigValidators = (*CommitmentsDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"org_id": schema.StringAttribute{
				Required: true,
			},
			"account_id": schema.StringAttribute{
				Description: "The unique identifier (UUID) for the Account. This parameter helps filter the Commitments related to a specific end-customer Account.",
				Optional:    true,
			},
			"contract_id": schema.StringAttribute{
				Optional: true,
			},
			"date": schema.StringAttribute{
				Description: "A date *(in ISO-8601 format)* to filter Commitments which are active on this specific date.",
				Optional:    true,
			},
			"end_date_end": schema.StringAttribute{
				Description: "A date *(in ISO-8601 format)* used to filter Commitments. Only Commitments with end dates before this date will be included.",
				Optional:    true,
			},
			"end_date_start": schema.StringAttribute{
				Description: "A date *(in ISO-8601 format)* used to filter Commitments. Only Commitments with end dates on or after this date will be included.",
				Optional:    true,
			},
			"product_id": schema.StringAttribute{
				Description: "The unique identifier (UUID) for the Product. This parameter helps filter the Commitments related to a specific Product.",
				Optional:    true,
			},
			"ids": schema.ListAttribute{
				Description: "A list of unique identifiers (UUIDs) for the Commitments to retrieve. Use this to fetch specific Commitments in a single request.",
				Optional:    true,
				ElementType: types.StringType,
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
				CustomType:  customfield.NewNestedObjectListType[CommitmentsItemsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The UUID of the entity.",
							Computed:    true,
						},
						"version": schema.Int64Attribute{
							Description: "The version number:\n- **Create:** On initial Create to insert a new entity, the version is set at 1 in the response.\n- **Update:** On successful Update, the version is incremented by 1 in the response.",
							Computed:    true,
						},
						"account_id": schema.StringAttribute{
							Description: "The unique identifier (UUID) for the end customer Account the Commitment is added to.",
							Computed:    true,
						},
						"accounting_product_id": schema.StringAttribute{
							Description: "The unique identifier (UUID) for the Product linked to the Commitment for accounting purposes.",
							Computed:    true,
						},
						"amount": schema.Float64Attribute{
							Description: "The total amount that the customer has committed to pay.",
							Computed:    true,
						},
						"amount_first_bill": schema.Float64Attribute{
							Description: "The amount to be billed in the first invoice.",
							Computed:    true,
						},
						"amount_pre_paid": schema.Float64Attribute{
							Description: "The amount that the customer has already paid upfront at the start of the Commitment service period.",
							Computed:    true,
						},
						"amount_spent": schema.Float64Attribute{
							Description: "The total amount of the Commitment that the customer has spent so far.",
							Computed:    true,
						},
						"bill_epoch": schema.StringAttribute{
							Description: "The starting date *(in ISO-8601 date format)* from which the billing cycles are calculated.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"billing_interval": schema.Int64Attribute{
							Description: "How often the Commitment fees are applied to bills. For example, if the plan being used to bill for Commitment fees is set to issue bills every three months and the `billingInterval` is set to 2, then the Commitment fees are applied every six months.",
							Computed:    true,
						},
						"billing_offset": schema.Int64Attribute{
							Description: "The offset for when the Commitment fees are first applied to bills on the Account. For example, if bills are issued every three months and the `billingOffset` is 0, then the charge is applied to the first bill (at three months); if set to 1, it's applied to the next bill (at six months), and so on.",
							Computed:    true,
						},
						"billing_plan_id": schema.StringAttribute{
							Description: "The unique identifier (UUID) for the Product Plan used for billing Commitment fees due.",
							Computed:    true,
						},
						"child_billing_mode": schema.StringAttribute{
							Description: "If the Account is either a Parent or a Child Account, this specifies the Account hierarchy billing mode. The mode determines how billing will be handled and shown on bills for charges due on the Parent Account, and charges due on Child Accounts:\n\n* **Parent Breakdown** - a separate bill line item per Account. Default setting.\n\n* **Parent Summary** - single bill line item for all Accounts.\n\n* **Child** - the Child Account is billed.\nAvailable values: \"PARENT_SUMMARY\", \"PARENT_BREAKDOWN\", \"CHILD\".",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"PARENT_SUMMARY",
									"PARENT_BREAKDOWN",
									"CHILD",
								),
							},
						},
						"commitment_fee_bill_in_advance": schema.BoolAttribute{
							Description: "A boolean value indicating whether the Commitment fee is billed in advance *(start of each billing period)* or arrears *(end of each billing period)*.\n\n* **TRUE** - bill in advance *(start of each billing period)*.\n* **FALSE** - bill in arrears *(end of each billing period)*.",
							Computed:    true,
						},
						"commitment_fee_description": schema.StringAttribute{
							Description: "A textual description of the Commitment fee.",
							Computed:    true,
						},
						"commitment_usage_description": schema.StringAttribute{
							Description: "A textual description of the Commitment usage.",
							Computed:    true,
						},
						"contract_id": schema.StringAttribute{
							Description: "The unique identifier (UUID) for a Contract you've created for the Account and to which the Commitment has been added.",
							Computed:    true,
						},
						"created_by": schema.StringAttribute{
							Description: "The unique identifier (UUID) of the user who created this Commitment.",
							Computed:    true,
						},
						"currency": schema.StringAttribute{
							Description: "The currency used for the Commitment. For example, 'USD'.",
							Computed:    true,
						},
						"drawdowns_accounting_product_id": schema.StringAttribute{
							Computed: true,
						},
						"dt_created": schema.StringAttribute{
							Description: "The date and time *(in ISO-8601 format)* when the Commitment was created.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"dt_last_modified": schema.StringAttribute{
							Description: "The date and time *(in ISO-8601 format)* when the Commitment was last modified.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"end_date": schema.StringAttribute{
							Description: "The end date of the Commitment period in ISO-8601 format.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"fee_dates": schema.ListNestedAttribute{
							Description: "Used for billing any outstanding Commitment fees *on a schedule*.\n\nAn array defining a series of bill dates and amounts covering specified service periods:\n- `date` - the billing date *(in ISO-8601 format)*.\n- `amount` - the billed amount.\n- `servicePeriodStartDate` and `servicePeriodEndDate` - defines the service period the bill covers *(in ISO-8601 format)*.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectListType[CommitmentsFeeDatesDataSourceModel](ctx),
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"amount": schema.Float64Attribute{
										Computed: true,
										Validators: []validator.Float64{
											float64validator.AtLeast(0),
										},
									},
									"date": schema.StringAttribute{
										Computed:   true,
										CustomType: timetypes.RFC3339Type{},
									},
									"service_period_end_date": schema.StringAttribute{
										Computed:   true,
										CustomType: timetypes.RFC3339Type{},
									},
									"service_period_start_date": schema.StringAttribute{
										Computed:   true,
										CustomType: timetypes.RFC3339Type{},
									},
								},
							},
						},
						"fees_accounting_product_id": schema.StringAttribute{
							Computed: true,
						},
						"last_modified_by": schema.StringAttribute{
							Description: "The unique identifier (UUID) of the user who last modified this Commitment.",
							Computed:    true,
						},
						"line_item_types": schema.ListAttribute{
							Description: "Specifies the line item charge types that can draw-down at billing against the Commitment amount. Options are:\n- `MINIMUM_SPEND`\n- `STANDING_CHARGE`\n- `USAGE`\n- `\"COUNTER_RUNNING_TOTAL_CHARGE\"`\n- `\"COUNTER_ADJUSTMENT_DEBIT\"`",
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
						"overage_description": schema.StringAttribute{
							Description: "A textual description of the overage charges.",
							Computed:    true,
						},
						"overage_surcharge_percent": schema.Float64Attribute{
							Description: "The percentage surcharge applied to the usage charges that exceed the Commitment amount.",
							Computed:    true,
						},
						"product_ids": schema.ListAttribute{
							Description: "A list of unique identifiers (UUIDs) for Products the Account consumes. Charges due for these Products will be made available for draw-down against the Commitment.\n\n**Note:** If not used, then charges due for all Products the Account consumes will be made available for draw-down against the Commitment.",
							Computed:    true,
							CustomType:  customfield.NewListType[types.String](ctx),
							ElementType: types.StringType,
						},
						"separate_overage_usage": schema.BoolAttribute{
							Description: "A boolean value indicating whether the overage usage is billed separately or together. If overage usage is separated and a Commitment amount has been consumed by an Account, any subsequent line items on Bills against the Account for usage will show as separate \"overage usage\" charges, not simply as \"usage\" charges:\n\n* **TRUE** - billed separately.\n* **FALSE** - billed together.",
							Computed:    true,
						},
						"start_date": schema.StringAttribute{
							Description: "The start date of the Commitment period in ISO-8601 format.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
					},
				},
			},
		},
	}
}

func (d *CommitmentsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *CommitmentsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
