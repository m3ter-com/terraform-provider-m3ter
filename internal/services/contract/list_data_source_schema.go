// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package contract

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

var _ datasource.DataSourceWithConfigValidators = (*ContractsDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Endpoints for Contract related operations such as creation, update, list and delete.\n\nContracts are created for Accounts, which are your end-user customers. Contracts can be used for:\n* **Accounts Reporting**. To serve your general accounting operations and processes, you can report on total Contract values for an Account.\n* **Contract Billing**. Various billing entities associated with an Account can  be linked to Contracts on the Account to meet your specific Contract billing use cases. ",
		Attributes: map[string]schema.Attribute{
			"org_id": schema.StringAttribute{
				Optional:           true,
				DeprecationMessage: "the org id should be set at the client level instead",
			},
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

func (d *ContractsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *ContractsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
