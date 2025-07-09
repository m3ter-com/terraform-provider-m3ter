// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package balance

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

var _ datasource.DataSourceWithConfigValidators = (*BalanceDataSource)(nil)

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
			"contract_id": schema.StringAttribute{
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
			"rollover_amount": schema.Float64Attribute{
				Description: "The maximum amount that can be carried over past the Balance end date and draw-down against for billing if there is an unused Balance amount remaining when the Balance end date is reached.",
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
			"version": schema.Int64Attribute{
				Description: "The version number:\n- **Create:** On initial Create to insert a new entity, the version is set at 1 in the response.\n- **Update:** On successful Update, the version is incremented by 1 in the response.",
				Computed:    true,
			},
			"line_item_types": schema.ListAttribute{
				Description: "A list of line item charge types that can draw-down against the Balance amount at billing.",
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
			"product_ids": schema.ListAttribute{
				Description: "A list of Product IDs whose consumption charges due at billing can be drawn-down against the Balance amount.",
				Computed:    true,
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
			},
			"custom_fields": schema.DynamicAttribute{
				Description: "User defined fields enabling you to attach custom data. The value for a custom field can be either a string or a number.\n\nIf `customFields` can also be defined for this entity at the Organizational level,`customField` values defined at individual level override values of `customFields` with the same name defined at Organization level.\n\nSee [Working with Custom Fields](https://www.m3ter.com/docs/guides/creating-and-managing-products/working-with-custom-fields) in the m3ter documentation for more information.",
				Computed:    true,
			},
		},
	}
}

func (d *BalanceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *BalanceDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
