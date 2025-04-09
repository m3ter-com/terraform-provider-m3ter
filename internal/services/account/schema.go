// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
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

var _ resource.ResourceWithConfigValidators = (*AccountResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "The UUID of the entity.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"org_id": schema.StringAttribute{
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"code": schema.StringAttribute{
				Description: "Code of the Account. \nThis is a unique short code used for the Account.",
				Required:    true,
			},
			"email_address": schema.StringAttribute{
				Description: "Contact email for the Account.",
				Required:    true,
			},
			"name": schema.StringAttribute{
				Description: "Name of the Account.",
				Required:    true,
			},
			"auto_generate_statement_mode": schema.StringAttribute{
				Description: "Specify whether to auto-generate statements once Bills are approved or locked.\n\n- **None**. Statements will not be auto-generated.\n- **JSON**. Statements are auto-generated in JSON format.\n- **JSON and CSV**. Statements are auto-generated in both JSON and CSV formats.\nAvailable values: \"NONE\", \"JSON\", \"JSON_AND_CSV\".",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"NONE",
						"JSON",
						"JSON_AND_CSV",
					),
				},
			},
			"bill_epoch": schema.StringAttribute{
				Description: "Optional setting to define a *billing cycle date*, which sets the date of the first Bill and acts as a reference for when in the applied billing frequency period subsequent bills are created:\n* For example, if you attach a Plan to an Account where the Plan is configured for monthly billing frequency and you've defined the period the Plan will apply to the Account to be from January 1st, 2022 until January 1st, 2023. You then set a `billEpoch` date of February 15th, 2022. The first Bill will be created for the Account on February 15th, and subsequent Bills created on the 15th of the months following for the remainder of the billing period - March 15th, April 15th, and so on.\n* If not defined, then the relevant Epoch date set for the billing frequency period at Organization level will be used instead.\n* The date is in ISO-8601 format.",
				Optional:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"currency": schema.StringAttribute{
				Description: "Account level billing currency, such as USD or GBP. Optional attribute:\n- If you define an Account currency, this will be used for bills.\n- If you do not define a currency, the billing currency defined at Organizational level will be used.\n\n**Note:** If you've attached a Plan to the Account that uses a different currency to the billing currency, then you must add the relevant currency conversion rate at Organization level to ensure the billing process can convert line items calculated using the Plan currency into the selected billing currency. If you don't add these conversion rates, then bills will fail for the Account.",
				Optional:    true,
			},
			"days_before_bill_due": schema.Int64Attribute{
				Description: "Enter the number of days after the Bill generation date that you want to show on Bills as the due date.\n\n**Note:** If you define `daysBeforeBillDue` at individual Account level, this will take precedence over any `daysBeforeBillDue` setting defined at Organization level.",
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
			},
			"parent_account_id": schema.StringAttribute{
				Description: "Parent Account ID, or null if this Account does not have a parent.",
				Optional:    true,
			},
			"purchase_order_number": schema.StringAttribute{
				Description: "Purchase Order Number of the Account.\n\nOptional attribute - allows you to set a purchase order number that comes through into invoicing. For example, your financial systems might require this as a reference for clearing payments.",
				Optional:    true,
			},
			"statement_definition_id": schema.StringAttribute{
				Description: "The UUID of the statement definition used when Bill statements are generated for the Account. If no statement definition is specified for the Account, the statement definition specified at Organizational level is used.\n\nBill statements can be used as informative backing sheets to invoices. Based on the usage breakdown defined in the statement definition, generated statements give a breakdown of usage charges on Account Bills, which helps customers better understand usage charges incurred over the billing period.\n\nSee [Working with Bill Statements](https://www.m3ter.com/docs/guides/running-viewing-and-managing-bills/working-with-bill-statements) in the m3ter documentation for more details.",
				Optional:    true,
			},
			"version": schema.Int64Attribute{
				Description: "The version number of the entity:\n- **Create entity:** Not valid for initial insertion of new entity - *do not use for Create*. On initial Create, version is set at 1 and listed in the response.\n- **Update Entity:**  On Update, version is required and must match the existing version because a check is performed to ensure sequential versioning is preserved. Version is incremented by 1 and listed in the response.",
				Optional:    true,
			},
			"config_data": schema.MapAttribute{
				Description: "Configuration data for the Account\nSupported settings:\n * SendBillsToThirdParties (\"true\"/\"false\")",
				Optional:    true,
				ElementType: jsontypes.NormalizedType{},
			},
			"credit_application_order": schema.ListAttribute{
				Description: "Define the order in which any Prepayment or Balance amounts on the Account are to be drawn-down against for billing. Four options:\n- `\"PREPAYMENT\",\"BALANCE\"`. Draw-down against Prepayment credit before Balance credit.\n- `\"BALANCE\",\"PREPAYMENT\"`. Draw-down against Balance credit before Prepayment credit.\n- `\"PREPAYMENT\"`. Only draw-down against Prepayment credit.\n- `\"BALANCE\"`. Only draw-down against Balance credit.\n\n**NOTES:**\n* Any setting you define here overrides the setting for credit application order at Organization level.\n* If the Account belongs to a Parent/Child Account hierarchy, then the `creditApplicationOrder` settings are not available, and the draw-down order defaults always to Prepayment then Balance order.",
				Optional:    true,
				Validators: []validator.List{
					listvalidator.ValueStringsAre(
						stringvalidator.OneOfCaseInsensitive("PREPAYMENT", "BALANCE"),
					),
				},
				ElementType: types.StringType,
			},
			"custom_fields": schema.MapAttribute{
				Description: "User defined fields enabling you to attach custom data. The value for a custom field can be either a string or a number.\n\nIf `customFields` can also be defined for this entity at the Organizational level, `customField` values defined at individual level override values of `customFields` with the same name defined at Organization level.\n\nSee [Working with Custom Fields](https://www.m3ter.com/docs/guides/creating-and-managing-products/working-with-custom-fields) in the m3ter documentation for more information.",
				Optional:    true,
				ElementType: types.DynamicType,
			},
			"address": schema.SingleNestedAttribute{
				Description: "Contact address.",
				Computed:    true,
				Optional:    true,
				CustomType:  customfield.NewNestedObjectType[AccountAddressModel](ctx),
				Attributes: map[string]schema.Attribute{
					"address_line1": schema.StringAttribute{
						Optional: true,
					},
					"address_line2": schema.StringAttribute{
						Optional: true,
					},
					"address_line3": schema.StringAttribute{
						Optional: true,
					},
					"address_line4": schema.StringAttribute{
						Optional: true,
					},
					"country": schema.StringAttribute{
						Optional: true,
					},
					"locality": schema.StringAttribute{
						Optional: true,
					},
					"post_code": schema.StringAttribute{
						Optional: true,
					},
					"region": schema.StringAttribute{
						Optional: true,
					},
				},
			},
			"created_by": schema.StringAttribute{
				Description: "The ID of the user who created the account.",
				Computed:    true,
			},
			"dt_created": schema.StringAttribute{
				Description: "The DateTime when the Account was created *(in ISO 8601 format)*.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"dt_last_modified": schema.StringAttribute{
				Description: "The DateTime when the Account was last modified *(in ISO 8601 format)*.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"last_modified_by": schema.StringAttribute{
				Description: "The ID of the user who last modified the Account.",
				Computed:    true,
			},
		},
	}
}

func (r *AccountResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *AccountResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
