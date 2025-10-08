// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package contract

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

var _ datasource.DataSourceWithConfigValidators = (*ContractDataSource)(nil)

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
			"account_id": schema.StringAttribute{
				Description: "The unique identifier (UUID) of the Account associated with this Contract.",
				Computed:    true,
			},
			"bill_grouping_key": schema.StringAttribute{
				Computed: true,
			},
			"code": schema.StringAttribute{
				Description: "The short code of the Contract.",
				Computed:    true,
			},
			"description": schema.StringAttribute{
				Description: "The description of the Contract, which provides context and information.",
				Computed:    true,
			},
			"end_date": schema.StringAttribute{
				Description: "The exclusive end date of the Contract *(in ISO-8601 format)*. This means the Contract is active until midnight on the day ***before*** this date.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"name": schema.StringAttribute{
				Description: "The name of the Contract.",
				Computed:    true,
			},
			"purchase_order_number": schema.StringAttribute{
				Description: "The Purchase Order Number associated with the Contract.",
				Computed:    true,
			},
			"start_date": schema.StringAttribute{
				Description: "The start date for the Contract *(in ISO-8601 format)*. This date is inclusive, meaning the Contract is active from this date onward.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
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
					"account_id": schema.StringAttribute{
						Optional: true,
					},
					"codes": schema.ListAttribute{
						Description: "An optional parameter to retrieve specific Contracts based on their short codes.",
						Optional:    true,
						ElementType: types.StringType,
					},
					"ids": schema.ListAttribute{
						Description: "An optional parameter to filter the list based on specific Contract unique identifiers (UUIDs).",
						Optional:    true,
						ElementType: types.StringType,
					},
				},
			},
		},
	}
}

func (d *ContractDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *ContractDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("id"), path.MatchRoot("find_one_by")),
	}
}
