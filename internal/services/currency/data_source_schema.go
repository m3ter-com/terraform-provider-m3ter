// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package currency

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*CurrencyDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Required: true,
			},
			"org_id": schema.StringAttribute{
				Required: true,
			},
			"archived": schema.BoolAttribute{
				Description: "TRUE / FALSE flag indicating whether the data entity is archived. An entity can be archived if it is obsolete.",
				Computed:    true,
			},
			"code": schema.StringAttribute{
				Description: "The short code of the data entity.",
				Computed:    true,
			},
			"created_by": schema.StringAttribute{
				Description: "The unique identifier (UUID) of the user who created this Currency.",
				Computed:    true,
			},
			"dt_created": schema.StringAttribute{
				Description: "The date and time *(in ISO-8601 format)* when the Currency was created.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"dt_last_modified": schema.StringAttribute{
				Description: "The date and time *(in ISO-8601 format)* when the Currency was last modified.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"last_modified_by": schema.StringAttribute{
				Description: "The unique identifier (UUID) of the user who last modified this Currency.",
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
				Description: "Available values: \"UP\", \"DOWN\", \"CEILING\", \"FLOOR\", \"HALF_UP\", \"HALF_DOWN\", \"HALF_EVEN\", \"UNNECESSARY\".",
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
		},
	}
}

func (d *CurrencyDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *CurrencyDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
