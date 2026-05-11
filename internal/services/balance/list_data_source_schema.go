// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package balance

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

var _ datasource.DataSourceWithConfigValidators = (*BalancesDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Endpoints for creating/retrieving/updating/deleting Balances on Accounts.\n\nWhen you have created a Balance for an Account, you can create a positive or negative Transaction amounts for the Balance. To do this, you must first define Transaction Types for your Organization, and then use one of these Transaction Types when you add a specific Transaction to a Balance - see the [Create TransactionType](https://www.m3ter.com/docs/api#tag/TransactionType/operation/CreateTransactionType) call in the Transaction Type section in this API Reference for more details.\n\nBalances are typically used when a customer prepays an amount to add a credit to their Account, which can then be draw-down against charges due for product or service consumption. You can include options to top-up the original Balance.\n\nExamples of how Balances for end customer Accounts can be used:\n\n* Onboarding Balance/Free Trials. Offering an onboarding incentive to new customers as an initial free credit Balance on their Account. \n\n* Balance as initial commitment. Add a Balance amount to a new customer Account. This acts as an initial commitment, which allows them to use the service and gain an accurate insight into their usage level. \n\n* Managing Customer Satisfaction. Use Balance as credits that will be applied to subsequent Bills as compensation for acknowledged service delivery issues.\n\n* Facilitating Balance Adjustments:\n\t* Apply negative amounts to immediately write-off outstanding Balances.\n\n#### What is the difference between Balances and Commitments/Prepayments?\n\nTo manage credit amounts for your end-customer Accounts, you can use Balances or Commitments/Prepayments. However, these two kinds of credits for Accounts serve different purposes.\n\nCommitments - also referred to as Prepayments - are used for amounts end-customers have agreed to pay for consuming your product or services across a full contract term. A customer might pay the entire or only part of the agreed amount upfront, but ***the commitment or prepayment amount is payable regardless of the actual usage by the customer of your service or product.***\n\nIn contrast, a Balance - often referred to as a Top-Up or Prepaid draw-down - is used when a customer wants to add a credit amount to their Account at any time during the service period or when you as service provider want to add a credit to a customer Account. This Balance credit can then be drawn-down against for billing the Account for usage, minimum spend, standing charges, or recurring charges due. Balances therefore serve payment use cases in a more flexible way, for example to be used for a \"Free Credit\" sign-up scheme you offer to encourage sales or to enhance customer satisfaction by adding credit to an Account to compensate for service delivery issues.\n\nYou can use Commitments/Prepayments and Balances together on Account, and define at Organization or individual Account level the order in which any Balance/Commitment credit on an Account is drawn-down - Balance amounts first or Commitment/Prepayment amounts first.  \n",
		Attributes: map[string]schema.Attribute{
			"org_id": schema.StringAttribute{
				Optional:           true,
				DeprecationMessage: "the org id should be set at the client level instead",
			},
			"account_id": schema.StringAttribute{
				Description: "The unique identifier (UUID) for the end customer's account.",
				Optional:    true,
			},
			"contract": schema.StringAttribute{
				Optional: true,
			},
			"contract_id": schema.StringAttribute{
				Description: "Filter Balances by contract id. Use '' with accountId to fetch unlinked balances.",
				Optional:    true,
			},
			"end_date_end": schema.StringAttribute{
				Description: "Only include Balances with end dates earlier than this date. If a Balance has a rollover amount configured, then the `rolloverEndDate` will be used as the end date.",
				Optional:    true,
			},
			"end_date_start": schema.StringAttribute{
				Description: "Only include Balances with end dates equal to or later than this date. If a Balance has a rollover amount configured, then the `rolloverEndDate` will be used as the end date.",
				Optional:    true,
			},
			"ids": schema.ListAttribute{
				Description: "A list of unique identifiers (UUIDs) for specific Balances to retrieve.",
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

func (d *BalancesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *BalancesDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
