// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package bill_config

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = (*BillConfigDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"org_id": schema.StringAttribute{
				Required:           true,
				DeprecationMessage: "the org id should be set at the client level instead",
			},
			"bill_lock_date": schema.StringAttribute{
				Description: "The global lock date *(in ISO 8601 format)* when all Bills will be locked.\n\nFor example: `\"2024-03-01\"`.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"id": schema.StringAttribute{
				Description: "The Organization UUID. The Organization represents your company as a direct customer of the m3ter service.",
				Computed:    true,
			},
			"version": schema.Int64Attribute{
				Description: "The version number:\n* Default value when newly created is one.\n* Incremented by 1 each time it is updated.",
				Computed:    true,
			},
		},
	}
}

func (d *BillConfigDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *BillConfigDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
