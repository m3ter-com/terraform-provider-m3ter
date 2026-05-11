// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package plan_template

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

var _ datasource.DataSourceWithConfigValidators = (*PlanTemplatesDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Endpoints for listing, creating, updating, retrieving, or deleting PlanTemplates.\n\nUse PlanTemplates to define default values for Plans. These default values control the billing operations you want applied to your products. PlanTemplates avoid repetition in configuration work - many Plans will share settings for billing operations and differ only in the details of their pricing structures.\n\nA PlanTemplate is linked to a Product, and each Plan is a child of a PlanTemplate. ",
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
			"items": schema.DynamicAttribute{
				Description: "The items returned by the data source",
				Computed:    true,
				CustomType:  customfield.NormalizedDynamicType{},
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
