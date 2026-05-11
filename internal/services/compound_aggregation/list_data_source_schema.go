// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package compound_aggregation

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

var _ datasource.DataSourceWithConfigValidators = (*CompoundAggregationsDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Endpoints for Compound Aggregation related operations such as creation, update, list and delete.\n\nUse Compound Aggregations to create numerical measures from usage data by applying a calculation to one or more simple Aggregations or Custom Fields. These numerical measures can then be used as pricing metrics to price your Product Plans, enabling you to implement a wide range of usage-based pricing use cases.\n\nYou can create two types of Compound Aggregation:\n\n**Global**\n- Pricing: Not tied to any specific product and can be used to price Plans belonging to any Product.\n- Calculation: can reference all simple Aggregations - both Global simple Aggregations and any product-specific simple Aggregations.\n\n**Product-specific**\n- Pricing: belong to a specific Product and can only be used to price Plans belonging to the same Product.\n- Calculation: can reference any simple Aggregations belonging to the same Product and any Global simple Aggregations.\n\n**IMPORTANT!** If a simple Aggregation referenced by a Compound Aggregation has a **Quantity per unit** defined or a **Rounding** defined, these will not be factored into the value used by the calculation. For example, if the simple Aggregation referenced has a base value of 100 and has **Quantity per unit** set at 10, the Compound Aggregation calculation *will use the base value of 100 not 10*. \n\nTo better understand and use Compound Aggregations, refer to the example [Compound Aggregation Use Case](https://www.m3ter.com/docs/guides/setting-up-usage-data-meters-and-aggregations/compound-aggregations#example-use-case) in the m3ter documentation.\n\n",
		Attributes: map[string]schema.Attribute{
			"org_id": schema.StringAttribute{
				Optional:           true,
				DeprecationMessage: "the org id should be set at the client level instead",
			},
			"codes": schema.ListAttribute{
				Description: "An optional parameter to retrieve specific CompoundAggregations based on their short codes.",
				Optional:    true,
				ElementType: types.StringType,
			},
			"ids": schema.ListAttribute{
				Description: "An optional parameter to retrieve specific CompoundAggregations based on their unique identifiers (UUIDs).",
				Optional:    true,
				ElementType: types.StringType,
			},
			"product_id": schema.ListAttribute{
				Description: "An optional parameter to filter the CompoundAggregations based on specific Product unique identifiers (UUIDs).",
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

func (d *CompoundAggregationsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *CompoundAggregationsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
