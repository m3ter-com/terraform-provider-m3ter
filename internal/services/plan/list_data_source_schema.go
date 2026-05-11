// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package plan

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

var _ datasource.DataSourceWithConfigValidators = (*PlansDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Endpoints for listing, creating, updating, retrieving, or deleting Plans.\n\nA Plan is based on a PlanTemplate and represents a specific pricing plan for one of your products or services. Each Plan inherits general billing attributes or pricing structure from its parent Plan Template. Some attributes can be overriden for the specific Plan.\n\nWhen you've created the Plan Templates and Plans you need for your Products, you can configure the exact pricing structures for Plans to charge customers that consume one or more of your Products.\n\nYou can then attach the appropriately priced Plans to customer Accounts to create [Account Plans](https://www.m3ter.com/docs/api#tag/AccountPlan) and enable charges to be calculated correctly for billing against those Accounts.\n\nSee also:\n- [Reviewing Options for Plans and Plan Templates](https://www.m3ter.com/docs/guides/working-with-plan-templates-and-plans/reviewing-configuration-options-for-plans-and-plan-templates).",
		Attributes: map[string]schema.Attribute{
			"org_id": schema.StringAttribute{
				Optional:           true,
				DeprecationMessage: "the org id should be set at the client level instead",
			},
			"product_id": schema.StringAttribute{
				Description: "UUID of the Product to retrieve Plans for.",
				Optional:    true,
			},
			"account_id": schema.ListAttribute{
				Description: "List of Account IDs the Plan belongs to.",
				Optional:    true,
				ElementType: types.StringType,
			},
			"ids": schema.ListAttribute{
				Description: "List of Plan IDs to retrieve.",
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

func (d *PlansDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *PlansDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
