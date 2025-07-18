// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_field

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = (*CustomFieldDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"org_id": schema.StringAttribute{
				Required:           true,
				DeprecationMessage: "the org id should be set at the client level instead",
			},
			"created_by": schema.StringAttribute{
				Description: "The id of the user who created this custom field.",
				Computed:    true,
			},
			"dt_created": schema.StringAttribute{
				Description: "The DateTime when the Organization was created *(in ISO-8601 format)*.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"dt_last_modified": schema.StringAttribute{
				Description: "The DateTime when a custom field was last modified - created, modified, or deleted - for the Organization *(in ISO-8601 format)*.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"id": schema.StringAttribute{
				Description: "The UUID of the entity.",
				Computed:    true,
			},
			"last_modified_by": schema.StringAttribute{
				Description: "The id of the user who last modified this custom field.",
				Computed:    true,
			},
			"version": schema.Int64Attribute{
				Description: "The version number:\n- **Create:** On initial Create to insert a new entity, the version is set at 1 in the response.\n- **Update:** On successful Update, the version is incremented by 1 in the response.",
				Computed:    true,
			},
			"account": schema.DynamicAttribute{
				Description: "CustomFields added to Account entities.",
				Computed:    true,
			},
			"account_plan": schema.DynamicAttribute{
				Description: "CustomFields added to accountPlan entities.",
				Computed:    true,
			},
			"aggregation": schema.DynamicAttribute{
				Description: "CustomFields added to simple Aggregation entities.",
				Computed:    true,
			},
			"compound_aggregation": schema.DynamicAttribute{
				Description: "CustomFields added to Compound Aggregation entities.",
				Computed:    true,
			},
			"contract": schema.DynamicAttribute{
				Description: "CustomFields added to Contract entities.",
				Computed:    true,
			},
			"meter": schema.DynamicAttribute{
				Description: "CustomFields added to Meter entities.",
				Computed:    true,
			},
			"organization": schema.DynamicAttribute{
				Description: "CustomFields added to the Organization.",
				Computed:    true,
			},
			"plan": schema.DynamicAttribute{
				Description: "CustomFields added to Plan entities.",
				Computed:    true,
			},
			"plan_template": schema.DynamicAttribute{
				Description: "CustomFields added to planTemplate entities.",
				Computed:    true,
			},
			"product": schema.DynamicAttribute{
				Description: "CustomFields added to Product entities.",
				Computed:    true,
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
