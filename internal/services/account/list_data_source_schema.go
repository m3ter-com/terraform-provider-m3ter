// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

var _ datasource.DataSourceWithConfigValidators = (*AccountsDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Endpoints for Account related operations such as creation, update, list and delete. \nAn Account represents one of your end-customer accounts. \n\nAccounts do not belong to a Product to allow for cases where an end customer takes more than one of your Products, and the charges for these Products differ.\n\nYou typically attach a priced Plan or Plan Template to an Account before you can generate bills for the Account:\n- If a customer consumes several of your Products, you can attach a priced Plan or Plan Template to the Account for charging against each Product.\n- If an Account is charged solely on the basis of an agreed Prepayment/Commitment amount but not all of the Prepayment is prepaid, you can use a customized billing schedule for outstanding fees without having to attach a Plan to the Account to generate Bills.\n\nYou can create Child Accounts for end customers who hold multiple Accounts with you. You can then set up billing for the Parent/Child Account usage to have the end-customer billed once for the Parent Account, instead of having separate bills issued for usage against each of their multiple Accounts.\n\n**IMPORTANT! - use of PII:** The use of any of your end-customers' Personally Identifiable Information (PII) in m3ter is restricted to a few fields on the **Account** entity. Please ensure that only the ``name``, ``address``, or ``emailAddress`` fields contain any end-customer PII data on any Accounts you create. See the [Introduction section](https://www.m3ter.com/docs/api#section/Introduction) above for more details.",
		Attributes: map[string]schema.Attribute{
			"org_id": schema.StringAttribute{
				Optional:           true,
				DeprecationMessage: "the org id should be set at the client level instead",
			},
			"codes": schema.ListAttribute{
				Description: "List of Account Codes to retrieve. \nThese are unique short codes for each Account.",
				Optional:    true,
				ElementType: types.StringType,
			},
			"ids": schema.ListAttribute{
				Description: "List of Account IDs to retrieve.",
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

func (d *AccountsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *AccountsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
