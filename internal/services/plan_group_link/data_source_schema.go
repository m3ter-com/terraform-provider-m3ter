// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package plan_group_link

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*PlanGroupLinkDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"org_id": schema.StringAttribute{
				Required:           true,
				DeprecationMessage: "the org id should be set at the client level instead",
			},
			"plan_group_id": schema.StringAttribute{
				Description: "ID of the linked PlanGroup",
				Computed:    true,
			},
			"plan_id": schema.StringAttribute{
				Description: "ID of the linked Plan",
				Computed:    true,
			},
			"version": schema.Int64Attribute{
				Description: "The version number:\n- **Create:** On initial Create to insert a new entity, the version is set at 1 in the response.\n- **Update:** On successful Update, the version is incremented by 1 in the response.",
				Computed:    true,
			},
			"find_one_by": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"ids": schema.ListAttribute{
						Description: "list of IDs to retrieve",
						Optional:    true,
						ElementType: types.StringType,
					},
					"plan": schema.StringAttribute{
						Description: "UUID of the Plan to retrieve PlanGroupLinks for",
						Optional:    true,
					},
					"plan_group": schema.StringAttribute{
						Description: "UUID of the PlanGroup to retrieve PlanGroupLinks for",
						Optional:    true,
					},
				},
			},
		},
	}
}

func (d *PlanGroupLinkDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *PlanGroupLinkDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("id"), path.MatchRoot("find_one_by")),
	}
}
