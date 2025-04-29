// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package organization_config

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

var _ resource.ResourceWithConfigValidators = (*OrganizationConfigResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
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
			"days_before_bill_due": schema.Int64Attribute{
				Description: "Enter the number of days after the Bill generation date that you want to show on Bills as the due date.\n\n**Note:** If you define `daysBeforeBillDue` at individual Account level, this will take precedence over any `daysBeforeBillDue` setting defined at Organization level.",
				Required:    true,
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
			},
			"auto_approve_bills_grace_period": schema.Int64Attribute{
				Description: "Grace period before bills are auto-approved. Used in combination with `autoApproveBillsGracePeriodUnit` parameter.\n\n**Note:** When used in combination with `autoApproveBillsGracePeriodUnit` enables auto-approval of Bills for Organization, which occurs when the specified time period has elapsed after Bill generation.",
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
			},
			"auto_approve_bills_grace_period_unit": schema.StringAttribute{
				Description: "Time unit of grace period before bills are auto-approved. Used in combination with `autoApproveBillsGracePeriod` parameter. Allowed options are MINUTES, HOURS, or DAYS.\n\n**Note:** When used in combination with `autoApproveBillsGracePeriod` enables auto-approval of Bills for Organization, which occurs when the specified time period has elapsed after Bill generation.",
				Optional:    true,
			},
			"auto_generate_statement_mode": schema.StringAttribute{
				Description: "Specify whether to auto-generate statements once Bills are *approved* or *locked*. It will not auto-generate if a bill is in *pending* state.\n\nThe default value is **None**.\n\n- **None**. Statements will not be auto-generated.\n- **JSON**. Statements are auto-generated in JSON format.\n- **JSON and CSV**. Statements are auto-generated in both JSON and CSV formats.\nAvailable values: \"NONE\", \"JSON\", \"JSON_AND_CSV\".",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"NONE",
						"JSON",
						"JSON_AND_CSV",
					),
				},
			},
			"bill_prefix": schema.StringAttribute{
				Description: "Prefix to be used for sequential invoice numbers. This will be combined with the `sequenceStartNumber`.\n\n**NOTES:**\n* If you do not define a `billPrefix`, a default will be used in the Console for the Bill **REFERENCE** number. This default will concatenate **INV-** with the last four characters of the `billId`.\n* If you do not define a `billPrefix`, the Bill response schema for API calls that retrieve Bill data will not contain a `sequentialInvoiceNumber`.",
				Optional:    true,
			},
			"commitment_fee_bill_in_advance": schema.BoolAttribute{
				Description: "Boolean setting to specify whether commitments *(prepayments)* are billed in advance at the start of each billing period, or billed in arrears at the end of each billing period. \n\n* **TRUE** - bill in advance *(start of each billing period)*.\n* **FALSE** - bill in arrears *(end of each billing period)*.",
				Optional:    true,
			},
			"consolidate_bills": schema.BoolAttribute{
				Description: "Boolean setting to consolidate different billing frequencies onto the same bill.\n\n* **TRUE** - consolidate different billing frequencies onto the same bill.\n* **FALSE** - bills are not consolidated.",
				Optional:    true,
			},
			"default_statement_definition_id": schema.StringAttribute{
				Description: "Organization level default `statementDefinitionId` to be used when there is no statement definition linked to the account.\n\nStatement definitions are used to generate bill statements, which are informative backing sheets to invoices.",
				Optional:    true,
			},
			"external_invoice_date": schema.StringAttribute{
				Description: "Date to use for the invoice date. Allowed values are `FIRST_DAY_OF_NEXT_PERIOD` or `LAST_DAY_OF_ARREARS`.",
				Optional:    true,
			},
			"minimum_spend_bill_in_advance": schema.BoolAttribute{
				Description: "Boolean setting to specify whether minimum spend amounts are billed in advance at the start of each billing period, or billed in arrears at the end of each billing period. \n\n* **TRUE** - bill in advance *(start of each billing period)*.\n* **FALSE** - bill in arrears *(end of each billing period)*.",
				Optional:    true,
			},
			"scheduled_bill_interval": schema.Float64Attribute{
				Description: "Sets the required interval for updating bills. It is an optional parameter that can be set as:\n\n* **For portions of an hour (minutes)**. Two options: **0.25** (15 minutes) and **0.5** (30 minutes).\n* **For full hours.** Enter **1** for every hour, **2** for every two hours, and so on. Eight options: **1**, **2**, **3**, **4**, **6**, **8**, **12**, or **24**.\n* **Default.** The default is **0**, which disables scheduling.",
				Optional:    true,
			},
			"sequence_start_number": schema.Int64Attribute{
				Description: "The starting number to be used for sequential invoice numbers. This will be combined with the `billPrefix`.\n\nFor example, if you define `billPrefix` to be **INVOICE-** and you set the `seqenceStartNumber` as **100**, the first Bill created after updating your Organization Configuration will have a `sequentialInvoiceNumber` assigned of **INVOICE-101**. Subsequent Bills created will be numbered in time sequence for their initial creation date/time.",
				Optional:    true,
			},
			"standing_charge_bill_in_advance": schema.BoolAttribute{
				Description: "Boolean setting to specify whether the standing charge is billed in advance at the start of each billing period, or billed in arrears at the end of each billing period. \n\n* **TRUE** - bill in advance *(start of each billing period)*.\n* **FALSE** - bill in arrears *(end of each billing period)*.",
				Optional:    true,
			},
			"suppressed_empty_bills": schema.BoolAttribute{
				Description: "Boolean setting that supresses generating bills that have no line items. \n\n* **TRUE** - prevents generating bills with no line items.\n* **FALSE** - bills are still generated even when they have no line items.",
				Optional:    true,
			},
			"version": schema.Int64Attribute{
				Description: "The version number of the entity:\n- **Create entity:** Not valid for initial insertion of new entity - *do not use for Create*. On initial Create, version is set at 1 and listed in the response.\n- **Update Entity:**  On Update, version is required and must match the existing version because a check is performed to ensure sequential versioning is preserved. Version is incremented by 1 and listed in the response.",
				Optional:    true,
			},
			"credit_application_order": schema.ListAttribute{
				Description: "Define the order in which any Prepayment or Balance amounts on Accounts are to be drawn-down against for billing. Four options:\n- `\"PREPAYMENT\",\"BALANCE\"`. Draw-down against Prepayment credit before Balance credit.\n- `\"BALANCE\",\"PREPAYMENT\"`. Draw-down against Balance credit before Prepayment credit.\n- `\"PREPAYMENT\"`. Only draw-down against Prepayment credit.\n- `\"BALANCE\"`. Only draw-down against Balance credit.\n\n**NOTES:**\n* You can override this Organization-level setting for `creditApplicationOrder` at the level of an individual Account.\n* If the Account belongs to a Parent/Child Account hierarchy, then the `creditApplicationOrder` settings are not available, and the draw-down order defaults always to Prepayment then Balance order.",
				Optional:    true,
				Validators: []validator.List{
					listvalidator.ValueStringsAre(
						stringvalidator.OneOfCaseInsensitive("PREPAYMENT", "BALANCE"),
					),
				},
				ElementType: types.StringType,
			},
			"currency": schema.StringAttribute{
				Description: "The currency code for the Organization. For example: USD, GBP, or EUR:\n* This defines the *billing currency* for the Organization. You can override this by selecting a different billing currency at individual Account level.\n* You must first define the currencies you want to use in your Organization. See the [Currency](https://www.m3ter.com/docs/api#tag/Currency) section in this API Reference.\n\n**Note:** If you use a different currency as the *pricing currency* for Plans to set charge rates for Product consumption by an Account, you must define a currency conversion rate from the pricing currency to the billing currency before you run billing for the Account, otherwise billing will fail. See below for the `currencyConversions` request parameter.",
				Computed:    true,
				Optional:    true,
				Default:     stringdefault.StaticString("USD"),
			},
			"day_epoch": schema.StringAttribute{
				Description: "Optional setting that defines the billing cycle date for Accounts that are billed daily. Defines the date of the first Bill:\n* For example, suppose the Plan you attach to an Account is configured for daily billing frequency and will apply to the Account from January 1st, 2022 until June 30th, 2022. If you set a `dayEpoch` date of January 2nd, 2022, then the first Bill is created for the Account on that date and subsequent Bills are created for the Account each day following through to the end of the billing service period.\n* The date is in ISO-8601 format.",
				Computed:    true,
				Optional:    true,
				Default:     stringdefault.StaticString("2022-01-01"),
			},
			"month_epoch": schema.StringAttribute{
				Description: "Optional setting that defines the billing cycle date for Accounts that are billed monthly. Defines the date of the first Bill and then acts as reference for when subsequent Bills are created for the Account:\n* For example, suppose the Plan you attach to an Account is configured for monthly billing frequency and will apply to the Account from January 1st, 2022 until June 30th, 2022. If you set a `monthEpoch` date of January 15th, 2022, then the first Bill is created for the Account on that date and subsequent Bills are created for the Account on the 15th of each month following through to the end of the billing service period - February 15th, March 15th, and so on.\n* The date is in ISO-8601 format.",
				Computed:    true,
				Optional:    true,
				Default:     stringdefault.StaticString("2022-01-01"),
			},
			"timezone": schema.StringAttribute{
				Description: "Sets the timezone for the Organization.",
				Computed:    true,
				Optional:    true,
				Default:     stringdefault.StaticString("UTC"),
			},
			"week_epoch": schema.StringAttribute{
				Description: "Optional setting that defines the billing cycle date for Accounts that are billed weekly. Defines the date of the first Bill and then acts as reference for when subsequent Bills are created for the Account:\n* For example, suppose the Plan you attach to an Account is configured for weekly billing frequency and will apply to the Account from January 1st, 2022 until June 30th, 2022. If you set a `weekEpoch` date of January 15th, 2022, which falls on a Saturday, then the first Bill is created for the Account on that date and subsequent Bills are created for the Account on Saturday of each week following through to the end of the billing service period.\n* The date is in ISO-8601 format.",
				Computed:    true,
				Optional:    true,
				Default:     stringdefault.StaticString("2022-01-04"),
			},
			"year_epoch": schema.StringAttribute{
				Description: "Optional setting that defines the billing cycle date for Accounts that are billed yearly. Defines the date of the first Bill and then acts as reference for when subsequent Bills are created for the Account:\n* For example, suppose the Plan you attach to an Account is configured for yearly billing frequency and will apply to the Account from January 1st, 2022 until January 15th, 2028. If you set a `yearEpoch` date of January 1st, 2023, then the first Bill is created for the Account on that date and subsequent Bills are created for the Account on January 1st of each year following through to the end of the billing service period - January 1st, 2023, January 1st, 2024 and so on.\n* The date is in ISO-8601 format.",
				Computed:    true,
				Optional:    true,
				Default:     stringdefault.StaticString("2022-01-01"),
			},
			"currency_conversions": schema.ListNestedAttribute{
				Description: "Define currency conversion rates from *pricing currency* to *billing currency*:\n* You can use the `currency` request parameter with this call to define the billing currency for your Organization - see above.\n* You can also define a billing currency at the individual Account level and this will override the Organization billing currency.\n* A Plan used to set Product consumption charge rates on an Account might use a different pricing currency. At billing, charges are calculated in the pricing currency and then converted into billing currency amounts to appear on Bills. If you haven't defined a currency conversion rate from pricing to billing currency, billing will fail for the Account.",
				Computed:    true,
				Optional:    true,
				CustomType:  customfield.NewNestedObjectListType[OrganizationConfigCurrencyConversionsModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"from": schema.StringAttribute{
							Description: "Currency to convert from. For example: GBP.",
							Required:    true,
						},
						"to": schema.StringAttribute{
							Description: "Currency to convert to. For example: USD.",
							Required:    true,
						},
						"multiplier": schema.Float64Attribute{
							Description: "Conversion rate between currencies.",
							Optional:    true,
							Validators: []validator.Float64{
								float64validator.AtLeast(0),
							},
						},
					},
				},
			},
			"created_by": schema.StringAttribute{
				Description: "The id of the user who created this organization config.",
				Computed:    true,
			},
			"dt_created": schema.StringAttribute{
				Description: "The DateTime when the organization config was created *(in ISO-8601 format)*.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"dt_last_modified": schema.StringAttribute{
				Description: "The DateTime when the organization config was last modified *(in ISO-8601 format)*.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"last_modified_by": schema.StringAttribute{
				Description: "The id of the user who last modified this organization config.",
				Computed:    true,
			},
		},
	}
}

func (r *OrganizationConfigResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *OrganizationConfigResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
