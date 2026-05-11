// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package plan_group

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

var _ datasource.DataSourceWithConfigValidators = (*PlanGroupsDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Endpoints for PlanGroup related operations such as creation, update, retrieve, list and delete.\n\nPlanGroups are constructs that group multiple plans together. This enables a unified approach to efficiently handle various uses cases across different plans. For example applying a minimum spend amount at billing, across several of your products or features that are each priced separately. ",
		Attributes: map[string]schema.Attribute{
			"org_id": schema.StringAttribute{
				Optional:           true,
				DeprecationMessage: "the org id should be set at the client level instead",
			},
			"account_id": schema.ListAttribute{
				Description: "Optional filter. The list of Account IDs to which the PlanGroups belong.",
				Optional:    true,
				ElementType: types.StringType,
			},
			"ids": schema.ListAttribute{
				Description: "Optional filter. The list of PlanGroup IDs to retrieve.",
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

func (d *PlanGroupsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *PlanGroupsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
