// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package aggregation

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

var _ datasource.DataSourceWithConfigValidators = (*AggregationsDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Endpoints for listing, creating, updating, retrieving, or deleting Aggregations.\n\nAn Aggregation links to a Meter and targets a Data Field or Derived Field on the Meter. You define the method of aggregation used to convert the usage data collected by the targeted Meter field into a numerical unit of measurement. \n\nYou can then use the unit of measurement an Aggregation yields as a metric for pricing Product Plans and apply usage-based pricing to your products and services. You might also want to aggregate raw data measures for other purposes, such as to feed into analytical or business performance tools.\n\n**Notes:**\n* **Contrast with Compound Aggregations**. Standard or simple Aggregations of this type, which apply an aggregation method directly to Meter usage data fields, are contrasted with [Compound Aggregations](https://www.m3ter.com/docs/api#tag/CompoundAggregation). A Compound Aggregation typically references one or more simple Aggregations and applies a calculation to them to derive pricing metrics needed to serve more complex usage-based pricing scenarios.\n* **Segmented Aggregations**. Segmented Aggregations allow you to segment the usage data collected by a single Meter. This capability is very useful for implementing some pricing and billing use cases. See [Segmented Aggregations](https://www.m3ter.com/docs/guides/usage-data-aggregations/segmented-aggregations) in our main documentation for more details.\n",
		Attributes: map[string]schema.Attribute{
			"org_id": schema.StringAttribute{
				Optional:           true,
				DeprecationMessage: "the org id should be set at the client level instead",
			},
			"codes": schema.ListAttribute{
				Description: "List of Aggregation codes to retrieve. These are unique short codes to identify each Aggregation.",
				Optional:    true,
				ElementType: types.StringType,
			},
			"ids": schema.ListAttribute{
				Description: "List of Aggregation IDs to retrieve.",
				Optional:    true,
				ElementType: types.StringType,
			},
			"product_id": schema.ListAttribute{
				Description: "The UUIDs of the Products to retrieve Aggregations for.",
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

func (d *AggregationsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *AggregationsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
