// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package plan

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

var _ datasource.DataSourceWithConfigValidators = (*PlanDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"org_id": schema.StringAttribute{
				Required: true,
			},
			"id": schema.StringAttribute{
				Computed: true,
			},
			"account_id": schema.StringAttribute{
				Description: "*(Optional)*. The Account ID for which this plan was created as custom/bespoke. A custom/bespoke Plan can only be attached to the specified Account.",
				Computed:    true,
			},
			"bespoke": schema.BoolAttribute{
				Description: "TRUE/FALSE flag indicating whether the plan is custom/bespoke for a particular Account.",
				Computed:    true,
			},
			"code": schema.StringAttribute{
				Description: "Unique short code reference for the Plan.",
				Computed:    true,
			},
			"created_by": schema.StringAttribute{
				Description: "The id of the user who created this plan.",
				Computed:    true,
			},
			"dt_created": schema.StringAttribute{
				Description: "The DateTime *(in ISO-8601 format)* when the plan was created.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"dt_last_modified": schema.StringAttribute{
				Description: "The DateTime *(in ISO-8601 format)* when the plan was last modified.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"last_modified_by": schema.StringAttribute{
				Description: "The id of the user who last modified this plan.",
				Computed:    true,
			},
			"minimum_spend": schema.Float64Attribute{
				Description: "The product minimum spend amount per billing cycle for end customer Accounts on a priced Plan.\n\n*(Optional)*. Overrides PlanTemplate value.",
				Computed:    true,
			},
			"minimum_spend_bill_in_advance": schema.BoolAttribute{
				Description: "When TRUE, minimum spend is billed at the start of each billing period.\n\nWhen FALSE, minimum spend is billed at the end of each billing period.\n\n*(Optional)*. Overrides the setting at PlanTemplate level for minimum spend billing in arrears/in advance.",
				Computed:    true,
			},
			"minimum_spend_description": schema.StringAttribute{
				Description: "Minimum spend description *(displayed on the bill line item)*.",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "Descriptive name for the Plan.",
				Computed:    true,
			},
			"ordinal": schema.Int64Attribute{
				Description: "Assigns a rank or position to the Plan in your order of pricing plans - lower numbers represent more basic pricing plans; higher numbers represent more premium pricing plans.\n\n*(Optional)*. Overrides PlanTemplate value.\n\n**NOTE:** **DEPRECATED** - no longer used.",
				Computed:    true,
			},
			"plan_template_id": schema.StringAttribute{
				Description: "UUID of the PlanTemplate the Plan belongs to. ",
				Computed:    true,
			},
			"product_id": schema.StringAttribute{
				Description: "UUID of the Product the Plan belongs to. ",
				Computed:    true,
			},
			"standing_charge": schema.Float64Attribute{
				Description: "The standing charge applied to bills for end customers. This is prorated.\n\n*(Optional)*. Overrides PlanTemplate value.",
				Computed:    true,
			},
			"standing_charge_bill_in_advance": schema.BoolAttribute{
				Description: "When TRUE, standing charge is billed at the start of each billing period.\n\nWhen FALSE, standing charge is billed at the end of each billing period.\n\n*(Optional)*. Overrides the setting at PlanTemplate level for standing charge billing in arrears/in advance.",
				Computed:    true,
			},
			"standing_charge_description": schema.StringAttribute{
				Description: "Standing charge description *(displayed on the bill line item)*.",
				Computed:    true,
			},
			"version": schema.Int64Attribute{
				Description: "The version number:\n- **Create:** On initial Create to insert a new entity, the version is set at 1 in the response.\n- **Update:** On successful Update, the version is incremented by 1 in the response.",
				Computed:    true,
			},
			"custom_fields": schema.MapAttribute{
				Description: "User defined fields enabling you to attach custom data. The value for a custom field can be either a string or a number.\n\nIf `customFields` can also be defined for this entity at the Organizational level,`customField` values defined at individual level override values of `customFields` with the same name defined at Organization level.\n\nSee [Working with Custom Fields](https://www.m3ter.com/docs/guides/creating-and-managing-products/working-with-custom-fields) in the m3ter documentation for more information.",
				Computed:    true,
				CustomType:  customfield.NewMapType[types.Dynamic](ctx),
				ElementType: types.DynamicType,
			},
		},
	}
}

func (d *PlanDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *PlanDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
