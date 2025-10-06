// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package debit_reason

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*DebitReasonDataSource)(nil)

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
			"archived": schema.BoolAttribute{
				Description: "TRUE / FALSE flag indicating whether the data entity is archived. An entity can be archived if it is obsolete.",
				Computed:    true,
			},
			"code": schema.StringAttribute{
				Description: "The short code of the data entity.",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "The name of the data entity.",
				Computed:    true,
			},
			"version": schema.Int64Attribute{
				Description: "The version number:\n- **Create:** On initial Create to insert a new entity, the version is set at 1 in the response.\n- **Update:** On successful Update, the version is incremented by 1 in the response.",
				Computed:    true,
			},
			"find_one_by": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"archived": schema.BoolAttribute{
						Description: "Filter using the boolean archived flag. DebitReasons can be archived if they are obsolete.\n* TRUE includes DebitReasons that have been archived.\n* FALSE excludes archived DebitReasons.",
						Optional:    true,
					},
					"codes": schema.ListAttribute{
						Description: "List of Debit Reason short codes to retrieve.",
						Optional:    true,
						ElementType: types.StringType,
					},
					"ids": schema.ListAttribute{
						Description: "List of Debit Reason IDs to retrieve.",
						Optional:    true,
						ElementType: types.StringType,
					},
				},
			},
		},
	}
}

func (d *DebitReasonDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *DebitReasonDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("id"), path.MatchRoot("find_one_by")),
	}
}
