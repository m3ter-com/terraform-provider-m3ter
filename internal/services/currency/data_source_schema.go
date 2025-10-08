// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package currency

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*CurrencyDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"org_id": schema.StringAttribute{
				Optional:           true,
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
			"max_decimal_places": schema.Int64Attribute{
				Description: "This indicates the maximum number of decimal places to use for this Currency.",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "The name of the data entity.",
				Computed:    true,
			},
			"rounding_mode": schema.StringAttribute{
				Description: `Available values: "UP", "DOWN", "CEILING", "FLOOR", "HALF_UP", "HALF_DOWN", "HALF_EVEN", "UNNECESSARY".`,
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"UP",
						"DOWN",
						"CEILING",
						"FLOOR",
						"HALF_UP",
						"HALF_DOWN",
						"HALF_EVEN",
						"UNNECESSARY",
					),
				},
			},
			"version": schema.Int64Attribute{
				Description: "The version number:\n- **Create:** On initial Create to insert a new entity, the version is set at 1 in the response.\n- **Update:** On successful Update, the version is incremented by 1 in the response.",
				Computed:    true,
			},
			"find_one_by": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"archived": schema.BoolAttribute{
						Description: "Filter by archived flag. A True / False flag indicating whether to return Currencies that are archived *(obsolete)*.\n\n* TRUE - return archived Currencies.\n* FALSE - archived Currencies are not returned.",
						Optional:    true,
					},
					"codes": schema.ListAttribute{
						Description: "An optional parameter to retrieve specific Currencies based on their short codes.",
						Optional:    true,
						ElementType: types.StringType,
					},
					"ids": schema.ListAttribute{
						Description: "An optional parameter to filter the list based on specific Currency unique identifiers (UUIDs).",
						Optional:    true,
						ElementType: types.StringType,
					},
				},
			},
		},
	}
}

func (d *CurrencyDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *CurrencyDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("id"), path.MatchRoot("find_one_by")),
	}
}
