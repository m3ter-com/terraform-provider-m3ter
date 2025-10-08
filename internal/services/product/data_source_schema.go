// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package product

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

var _ datasource.DataSourceWithConfigValidators = (*ProductDataSource)(nil)

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
			"code": schema.StringAttribute{
				Description: "A unique short code to identify the Product. It should not contain control chracters or spaces.",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "Descriptive name for the Product providing context and information.",
				Computed:    true,
			},
			"version": schema.Int64Attribute{
				Description: "The version number:\n- **Create:** On initial Create to insert a new entity, the version is set at 1 in the response.\n- **Update:** On successful Update, the version is incremented by 1 in the response.",
				Computed:    true,
			},
			"custom_fields": schema.DynamicAttribute{
				Description: "User defined fields enabling you to attach custom data. The value for a custom field can be either a string or a number.\n\nIf `customFields` can also be defined for this entity at the Organizational level,`customField` values defined at individual level override values of `customFields` with the same name defined at Organization level.\n\nSee [Working with Custom Fields](https://www.m3ter.com/docs/guides/creating-and-managing-products/working-with-custom-fields) in the m3ter documentation for more information.",
				Computed:    true,
				CustomType:  customfield.NormalizedDynamicType{},
			},
			"find_one_by": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"ids": schema.ListAttribute{
						Description: "List of specific Product UUIDs to retrieve.",
						Optional:    true,
						ElementType: types.StringType,
					},
				},
			},
		},
	}
}

func (d *ProductDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *ProductDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("id"), path.MatchRoot("find_one_by")),
	}
}
