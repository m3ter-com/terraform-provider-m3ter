// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package plan_group

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

var _ datasource.DataSourceWithConfigValidators = (*PlanGroupDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Required: true,
			},
			"org_id": schema.StringAttribute{
				Required: true,
			},
			"account_id": schema.StringAttribute{
				Description: "Optional. This PlanGroup was created as bespoke for the associated Account with this Account ID.",
				Computed:    true,
			},
			"code": schema.StringAttribute{
				Description: "The short code representing the PlanGroup.",
				Computed:    true,
			},
			"created_by": schema.StringAttribute{
				Description: "The unique identifier (UUID) for the user who created the PlanGroup.",
				Computed:    true,
			},
			"currency": schema.StringAttribute{
				Description: "Currency code for the PlanGroup (For example, USD).",
				Computed:    true,
			},
			"dt_created": schema.StringAttribute{
				Description: "The date and time *(in ISO 8601 format)* when the PlanGroup was first created.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"dt_last_modified": schema.StringAttribute{
				Description: "The date and time *(in ISO 8601 format)* when the PlanGroup was last modified.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"last_modified_by": schema.StringAttribute{
				Description: "The unique identifier (UUID) for the user who last modified the PlanGroup.",
				Computed:    true,
			},
			"minimum_spend": schema.Float64Attribute{
				Description: "The minimum spend amount for the PlanGroup.",
				Computed:    true,
			},
			"minimum_spend_accounting_product_id": schema.StringAttribute{
				Description: "Optional. Product ID to attribute the PlanGroup's minimum spend for accounting purposes.",
				Computed:    true,
			},
			"minimum_spend_bill_in_advance": schema.BoolAttribute{
				Description: "A boolean flag that determines when the minimum spend is billed. This flag overrides the setting at Organizational level for minimum spend billing in arrears/in advance.\n\n* **TRUE** - minimum spend is billed at the start of each billing period. \n* **FALSE** - minimum spend is billed at the end of each billing period.",
				Computed:    true,
			},
			"minimum_spend_description": schema.StringAttribute{
				Description: "Description of the minimum spend, displayed on the bill line item.",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "The name of the PlanGroup.",
				Computed:    true,
			},
			"standing_charge": schema.Float64Attribute{
				Description: "Standing charge amount for the PlanGroup.",
				Computed:    true,
			},
			"standing_charge_accounting_product_id": schema.StringAttribute{
				Description: "Optional. Product ID to attribute the PlanGroup's standing charge for accounting purposes.",
				Computed:    true,
			},
			"standing_charge_bill_in_advance": schema.BoolAttribute{
				Description: "A boolean flag that determines when the standing charge is billed. This flag overrides the setting at Organizational level for standing charge billing in arrears/in advance.\n\n* **TRUE** - standing charge is billed at the start of each billing period. \n* **FALSE** - standing charge is billed at the end of each billing period.",
				Computed:    true,
			},
			"standing_charge_description": schema.StringAttribute{
				Description: "Description of the standing charge, displayed on the bill line item.",
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

func (d *PlanGroupDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *PlanGroupDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
