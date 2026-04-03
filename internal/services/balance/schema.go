// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package balance

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

var _ resource.ResourceWithConfigValidators = (*BalanceResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Endpoints for creating/retrieving/updating/deleting Balances on Accounts.\n\nWhen you have created a Balance for an Account, you can create a positive or negative Transaction amounts for the Balance. To do this, you must first define Transaction Types for your Organization, and then use one of these Transaction Types when you add a specific Transaction to a Balance - see the [Create TransactionType](https://www.m3ter.com/docs/api#tag/TransactionType/operation/CreateTransactionType) call in the Transaction Type section in this API Reference for more details.\n\nBalances are typically used when a customer prepays an amount to add a credit to their Account, which can then be draw-down against charges due for product or service consumption. You can include options to top-up the original Balance.\n\nExamples of how Balances for end customer Accounts can be used:\n\n* Onboarding Balance/Free Trials. Offering an onboarding incentive to new customers as an initial free credit Balance on their Account. \n\n* Balance as initial commitment. Add a Balance amount to a new customer Account. This acts as an initial commitment, which allows them to use the service and gain an accurate insight into their usage level. \n\n* Managing Customer Satisfaction. Use Balance as credits that will be applied to subsequent Bills as compensation for acknowledged service delivery issues.\n\n* Facilitating Balance Adjustments:\n\t* Apply negative amounts to immediately write-off outstanding Balances.\n\n#### What is the difference between Balances and Commitments/Prepayments?\n\nTo manage credit amounts for your end-customer Accounts, you can use Balances or Commitments/Prepayments. However, these two kinds of credits for Accounts serve different purposes.\n\nCommitments - also referred to as Prepayments - are used for amounts end-customers have agreed to pay for consuming your product or services across a full contract term. A customer might pay the entire or only part of the agreed amount upfront, but ***the commitment or prepayment amount is payable regardless of the actual usage by the customer of your service or product.***\n\nIn contrast, a Balance - often referred to as a Top-Up or Prepaid draw-down - is used when a customer wants to add a credit amount to their Account at any time during the service period or when you as service provider want to add a credit to a customer Account. This Balance credit can then be drawn-down against for billing the Account for usage, minimum spend, standing charges, or recurring charges due. Balances therefore serve payment use cases in a more flexible way, for example to be used for a \"Free Credit\" sign-up scheme you offer to encourage sales or to enhance customer satisfaction by adding credit to an Account to compensate for service delivery issues.\n\nYou can use Commitments/Prepayments and Balances together on Account, and define at Organization or individual Account level the order in which any Balance/Commitment credit on an Account is drawn-down - Balance amounts first or Commitment/Prepayment amounts first.  \n",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "The UUID of the entity.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"org_id": schema.StringAttribute{
				Optional:           true,
				DeprecationMessage: "the org id should be set at the client level instead",
				PlanModifiers:      []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"account_id": schema.StringAttribute{
				Description: "The unique identifier (UUID) for the end customer Account.",
				Required:    true,
			},
			"code": schema.StringAttribute{
				Description: "Unique short code for the Balance.",
				Required:    true,
			},
			"currency": schema.StringAttribute{
				Description: "The currency code used for the Balance amount. For example: USD, GBP or EUR.",
				Required:    true,
			},
			"end_date": schema.StringAttribute{
				Description: "The date *(in ISO 8601 format)* after which the Balance will no longer be active for the Account.\n\n**Note:** You can use the `rolloverEndDate` request parameter to define an extended grace period for continued draw-down against the Balance if any amount remains when the specified `endDate` is reached.",
				Required:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"name": schema.StringAttribute{
				Description: "The official name for the Balance.",
				Required:    true,
			},
			"start_date": schema.StringAttribute{
				Description: "The date *(in ISO 8601 format)* when the Balance becomes active.",
				Required:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"allow_overdraft": schema.BoolAttribute{
				Description: "Allow balance amounts to fall below zero. This feature is enabled on request. Please get in touch with m3ter Support or your m3ter contact if you would like it enabling for your organization(s).",
				Optional:    true,
			},
			"balance_draw_down_description": schema.StringAttribute{
				Description: "A description for the bill line items for draw-down charges against the Balance. *(Optional).*",
				Optional:    true,
			},
			"consumptions_accounting_product_id": schema.StringAttribute{
				Description: "Product ID that any Balance Consumed line items will be attributed to for accounting purposes.(*Optional*)",
				Optional:    true,
			},
			"contract_id": schema.StringAttribute{
				Description: "The unique identifier (UUID) of a Contract on the Account that the Balance will be added to.",
				Optional:    true,
			},
			"description": schema.StringAttribute{
				Description: "A description of the Balance.",
				Optional:    true,
			},
			"fees_accounting_product_id": schema.StringAttribute{
				Description: "Product ID that any Balance Fees line items will be attributed to for accounting purposes.(*Optional*)",
				Optional:    true,
			},
			"overage_description": schema.StringAttribute{
				Description: "A description for Bill line items overage charges.",
				Optional:    true,
			},
			"overage_surcharge_percent": schema.Float64Attribute{
				Description: "Define a surcharge level, as a percentage of regular usage rating, applied to overages *(usage charges that exceed the Balance amount)*. For example, if the regular usage rate is $10 per unit of usage consumed and `overageSurchargePercent` is set at 10%, then any usage charged above the original Balance amount is charged at $11 per unit of usage.",
				Optional:    true,
			},
			"rollover_amount": schema.Float64Attribute{
				Description: "The maximum amount that can be carried over past the Balance end date for draw-down at billing if there is any unused Balance amount when the end date is reached. Works with `rolloverEndDate` to define the amount and duration of a Balance \"grace period\". *(Optional)*\n\n**Notes:**\n- If you leave `rolloverAmount` empty and only enter a `rolloverEndDate`, any amount left over after the Balance end date is reached will be drawn-down against up to the specified `rolloverEndDate`.\n- You must enter a `rolloverEndDate`. If you only enter a `rolloverAmount` without entering a `rolloverEndDate`, you'll receive an error when trying to create or update the Balance.\n- If you don't want to grant any grace period for outstanding Balance amounts, then do not use `rolloverAmount` and `rolloverEndDate`.",
				Optional:    true,
				Validators: []validator.Float64{
					float64validator.AtLeast(0),
				},
			},
			"rollover_end_date": schema.StringAttribute{
				Description: "The end date *(in ISO 8601 format)* for the grace period during which unused Balance amounts can be carried over and drawn-down against at billing.\n\n**Note:** Use `rolloverAmount` if you want to specify a maximum amount that can be carried over and made available for draw-down.",
				Optional:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"line_item_types": schema.ListAttribute{
				Description: "Specify the line item charge types that can draw-down at billing against the  Balance amount. Options are:\n- `\"MINIMUM_SPEND\"`\n- `\"STANDING_CHARGE\"`\n- `\"USAGE\"`\n- `\"COUNTER_RUNNING_TOTAL_CHARGE\"`\n- `\"COUNTER_ADJUSTMENT_DEBIT\"`\n- `AD_HOC`\n\n**NOTE:** If no charge types are specified, by default *all types* can draw-down against the Balance amount at billing.",
				Optional:    true,
				Validators: []validator.List{
					listvalidator.ValueStringsAre(
						stringvalidator.OneOfCaseInsensitive(
							"STANDING_CHARGE",
							"USAGE",
							"MINIMUM_SPEND",
							"COUNTER_RUNNING_TOTAL_CHARGE",
							"COUNTER_ADJUSTMENT_DEBIT",
							"AD_HOC",
						),
					),
				},
				ElementType: types.StringType,
			},
			"product_ids": schema.ListAttribute{
				Description: "Specify the Products whose consumption charges due at billing can be drawn-down against the Balance amount.\n\n**Note:** If you don't specify any Products for Balance draw-down, by default the consumption charges for any Product the Account consumes will be drawn-down against the Balance amount.",
				Optional:    true,
				ElementType: types.StringType,
			},
			"custom_fields": schema.DynamicAttribute{
				Description:   "User defined fields enabling you to attach custom data. The value for a custom field can be either a string or a number.\n\nIf `customFields` can also be defined for this entity at the Organizational level, `customField` values defined at individual level override values of `customFields` with the same name defined at Organization level.\n\nSee [Working with Custom Fields](https://www.m3ter.com/docs/guides/creating-and-managing-products/working-with-custom-fields) in the m3ter documentation for more information.",
				Optional:      true,
				CustomType:    customfield.NormalizedDynamicType{},
				PlanModifiers: []planmodifier.Dynamic{customfield.NormalizeDynamicPlanModifier()},
			},
			"amount": schema.Float64Attribute{
				Description: "The financial value that the Balance holds.",
				Computed:    true,
			},
			"version": schema.Int64Attribute{
				Description: "The version number:\n- **Create:** On initial Create to insert a new entity, the version is set at 1 in the response.\n- **Update:** On successful Update, the version is incremented by 1 in the response.",
				Computed:    true,
			},
		},
	}
}

func (r *BalanceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *BalanceResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
