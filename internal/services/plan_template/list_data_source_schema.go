// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package plan_template

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

var _ datasource.DataSourceWithConfigValidators = (*PlanTemplatesDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"org_id": schema.StringAttribute{
				Optional:           true,
				DeprecationMessage: "the org id should be set at the client level instead",
			},
			"product_id": schema.StringAttribute{
				Description: "The unique identifiers (UUIDs) of the Products to retrieve associated PlanTemplates.",
				Optional:    true,
			},
			"ids": schema.ListAttribute{
				Description: "List of specific PlanTemplate UUIDs to retrieve.",
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
				CustomType:  customfield.NewNestedObjectListType[PlanTemplatesItemsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The UUID of the entity.",
							Computed:    true,
						},
						"bill_frequency": schema.StringAttribute{
							Description: "Determines the frequency at which bills are generated.\n\n* **Daily**. Starting at midnight each day, covering the twenty-four hour period following.\n\n* **Weekly**. Starting at midnight on a Monday, covering the seven-day period following.\n\n* **Monthly**. Starting at midnight on the first day of each month, covering the entire calendar month following.\n\n* **Annually**. Starting at midnight on first day of each year covering the entire calendar year following.\nAvailable values: \"DAILY\", \"WEEKLY\", \"MONTHLY\", \"ANNUALLY\", \"AD_HOC\", \"MIXED\".",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"DAILY",
									"WEEKLY",
									"MONTHLY",
									"ANNUALLY",
									"AD_HOC",
									"MIXED",
								),
							},
						},
						"bill_frequency_interval": schema.Int64Attribute{
							Description: "How often bills are issued. \nFor example, if `billFrequency` is Monthly and `billFrequencyInterval` is 3, bills are issued every three months.",
							Computed:    true,
						},
						"code": schema.StringAttribute{
							Description: "A unique, short code reference for the PlanTemplate. This code should not contain control characters or spaces.",
							Computed:    true,
						},
						"currency": schema.StringAttribute{
							Description: "The ISO currency code for the pricing currency used by Plans based on the Plan Template to define charge rates for Product consumption - for example USD, GBP, EUR.",
							Computed:    true,
						},
						"custom_fields": schema.DynamicAttribute{
							Description: "User defined fields enabling you to attach custom data. The value for a custom field can be either a string or a number.\n\nIf `customFields` can also be defined for this entity at the Organizational level,`customField` values defined at individual level override values of `customFields` with the same name defined at Organization level.\n\nSee [Working with Custom Fields](https://www.m3ter.com/docs/guides/creating-and-managing-products/working-with-custom-fields) in the m3ter documentation for more information.",
							Computed:    true,
							CustomType:  customfield.NormalizedDynamicType{},
						},
						"minimum_spend": schema.Float64Attribute{
							Description: "The Product minimum spend amount per billing cycle for end customer Accounts on a pricing Plan based on the PlanTemplate. This must be a non-negative number.",
							Computed:    true,
						},
						"minimum_spend_bill_in_advance": schema.BoolAttribute{
							Description: "A boolean that determines when the minimum spend is billed.\n\n* TRUE - minimum spend is billed at the start of each billing period.\n* FALSE - minimum spend is billed at the end of each billing period.\n\nOverrides the setting at Organizational level for minimum spend billing in arrears/in advance.",
							Computed:    true,
						},
						"minimum_spend_description": schema.StringAttribute{
							Description: "Minimum spend description *(displayed on the bill line item)*.",
							Computed:    true,
						},
						"name": schema.StringAttribute{
							Description: "Descriptive name for the PlanTemplate.",
							Computed:    true,
						},
						"ordinal": schema.Int64Attribute{
							Description: "The ranking of the PlanTemplate among your pricing plans. Lower numbers represent more basic plans, while higher numbers represent premium plans. This must be a non-negative integer.\n\n**NOTE:** **DEPRECATED** - no longer used.",
							Computed:    true,
						},
						"product_id": schema.StringAttribute{
							Description: "The unique identifier (UUID) of the Product associated with this PlanTemplate.",
							Computed:    true,
						},
						"standing_charge": schema.Float64Attribute{
							Description: "The fixed charge *(standing charge)* applied to customer bills. This charge is prorated and must be a non-negative number.",
							Computed:    true,
						},
						"standing_charge_bill_in_advance": schema.BoolAttribute{
							Description: "A boolean that determines when the standing charge is billed.\n\n* TRUE - standing charge is billed at the start of each billing period.\n* FALSE - standing charge is billed at the end of each billing period.\n\nOverrides the setting at Organizational level for standing charge billing in arrears/in advance.",
							Computed:    true,
						},
						"standing_charge_description": schema.StringAttribute{
							Description: "Standing charge description *(displayed on the bill line item)*.",
							Computed:    true,
						},
						"standing_charge_interval": schema.Int64Attribute{
							Description: "How often the standing charge is applied. \nFor example, if the bill is issued every three months and `standingChargeInterval` is 2, then the standing charge is applied every six months.",
							Computed:    true,
						},
						"standing_charge_offset": schema.Int64Attribute{
							Description: "Defines an offset for when the standing charge is first applied. \nFor example, if the bill is issued every three months and the `standingChargeOfset` is 0, then the charge is applied to the first bill *(at three months)*; if 1, it would be applied to the second bill *(at six months)*, and so on.",
							Computed:    true,
						},
						"version": schema.Int64Attribute{
							Description: "The version number:\n- **Create:** On initial Create to insert a new entity, the version is set at 1 in the response.\n- **Update:** On successful Update, the version is incremented by 1 in the response.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func (d *PlanTemplatesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *PlanTemplatesDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
