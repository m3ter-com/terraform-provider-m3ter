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
		Description: "Endpoints for updating and retreiving the Bill Configuration for an Organization. \nThe Organization represents your company as a direct customer of the m3ter service.\n\nYou can use the **Update BillConfig** endpoint to set a global lock date for **all** Bills - any Bill with a service period end date on or before the set date will be locked and cannot be updated.\n\n**Warning: Ensure all Bills are Approved!** If you try to set a global lock date when there remains Bills in a *Pending* state whose service period end date is on or before the specified lock date, then you'll receive an error.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:           true,
				DeprecationMessage: "the org id should be set at the client level instead",
			},
			"org_id": schema.StringAttribute{
				Optional:           true,
				DeprecationMessage: "the org id should be set at the client level instead",
			},
			"bill_lock_date": schema.StringAttribute{
				Description: "The global lock date *(in ISO 8601 format)* when all Bills will be locked.\n\nFor example: `\"2024-03-01\"`.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"version": schema.Int64Attribute{
				Description: "The version number:\n- **Create:** On initial Create to insert a new entity, the version is set at 1 in the response.\n- **Update:** On successful Update, the version is incremented by 1 in the response.",
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
