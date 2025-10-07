// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package counter

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*CounterDataSource)(nil)

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
			"code": schema.StringAttribute{
				Description: "Code of the Counter. A unique short code to identify the Counter.",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "Name of the Counter.",
				Computed:    true,
			},
			"product_id": schema.StringAttribute{
				Description: "UUID of the product the Counter belongs to. *(Optional)* - if no `productId` is returned, the Counter is Global. A Global Counter can be used to price Plans or Plan Templates belonging to any Product.",
				Computed:    true,
			},
			"unit": schema.StringAttribute{
				Description: "Label for units shown on Bill line items, and indicating to customers what they are being charged for.",
				Computed:    true,
			},
			"version": schema.Int64Attribute{
				Description: "The version number:\n- **Create:** On initial Create to insert a new entity, the version is set at 1 in the response.\n- **Update:** On successful Update, the version is incremented by 1 in the response.",
				Computed:    true,
			},
			"find_one_by": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"codes": schema.ListAttribute{
						Description: "List of Counter codes to retrieve. These are unique short codes to identify each Counter.",
						Optional:    true,
						ElementType: types.StringType,
					},
					"ids": schema.ListAttribute{
						Description: "List of Counter IDs to retrieve.",
						Optional:    true,
						ElementType: types.StringType,
					},
					"product_id": schema.ListAttribute{
						Description: "List of Products UUIDs to retrieve Counters for.",
						Optional:    true,
						ElementType: types.StringType,
					},
				},
			},
		},
	}
}

func (d *CounterDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *CounterDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("id"), path.MatchRoot("find_one_by")),
	}
}
