// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_field

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

var _ datasource.DataSourceWithConfigValidators = (*CustomFieldDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:           true,
				DeprecationMessage: "the org id should be set at the client level instead",
			},
			"org_id": schema.StringAttribute{
				Optional:           true,
				DeprecationMessage: "the org id should be set at the client level instead",
			},
			"version": schema.Int64Attribute{
				Description: "The version number:\n- **Create:** On initial Create to insert a new entity, the version is set at 1 in the response.\n- **Update:** On successful Update, the version is incremented by 1 in the response.",
				Computed:    true,
			},
			"account": schema.DynamicAttribute{
				Description: "CustomFields added to Account entities.",
				Computed:    true,
				CustomType:  customfield.NormalizedDynamicType{},
			},
			"account_plan": schema.DynamicAttribute{
				Description: "CustomFields added to accountPlan entities.",
				Computed:    true,
				CustomType:  customfield.NormalizedDynamicType{},
			},
			"aggregation": schema.DynamicAttribute{
				Description: "CustomFields added to simple Aggregation entities.",
				Computed:    true,
				CustomType:  customfield.NormalizedDynamicType{},
			},
			"compound_aggregation": schema.DynamicAttribute{
				Description: "CustomFields added to Compound Aggregation entities.",
				Computed:    true,
				CustomType:  customfield.NormalizedDynamicType{},
			},
			"contract": schema.DynamicAttribute{
				Description: "CustomFields added to Contract entities.",
				Computed:    true,
				CustomType:  customfield.NormalizedDynamicType{},
			},
			"meter": schema.DynamicAttribute{
				Description: "CustomFields added to Meter entities.",
				Computed:    true,
				CustomType:  customfield.NormalizedDynamicType{},
			},
			"organization": schema.DynamicAttribute{
				Description: "CustomFields added to the Organization.",
				Computed:    true,
				CustomType:  customfield.NormalizedDynamicType{},
			},
			"plan": schema.DynamicAttribute{
				Description: "CustomFields added to Plan entities.",
				Computed:    true,
				CustomType:  customfield.NormalizedDynamicType{},
			},
			"plan_template": schema.DynamicAttribute{
				Description: "CustomFields added to planTemplate entities.",
				Computed:    true,
				CustomType:  customfield.NormalizedDynamicType{},
			},
			"product": schema.DynamicAttribute{
				Description: "CustomFields added to Product entities.",
				Computed:    true,
				CustomType:  customfield.NormalizedDynamicType{},
			},
		},
	}
}

func (d *CustomFieldDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *CustomFieldDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
