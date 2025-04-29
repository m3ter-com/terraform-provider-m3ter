// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package bill_job

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

var _ resource.ResourceWithConfigValidators = (*BillJobResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "The UUID of the entity.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"org_id": schema.StringAttribute{
				Optional:           true,
				DeprecationMessage: "the org id should be set at the client level instead",
				PlanModifiers:      []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"bill_date": schema.StringAttribute{
				Description:   "The specific billing date *(in ISO 8601 format)*, determining when the Bill was generated.\n\nFor example: `\"2023-01-24\"`.",
				Optional:      true,
				CustomType:    timetypes.RFC3339Type{},
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"bill_frequency_interval": schema.Int64Attribute{
				Description:   "How often Bills are issued - used in conjunction with `billingFrequency`.\n\nFor example, if `billingFrequency` is set to Monthly and `billFrequencyInterval` is set to 3, Bills are issued every three months.",
				Optional:      true,
				PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
			},
			"billing_frequency": schema.StringAttribute{
				Description: "How often Bills are generated.\n\n* **Daily**. Starting at midnight each day, covering a twenty-four hour period following.\n\n* **Weekly**. Starting at midnight on a Monday morning covering the seven-day period following.\n\n* **Monthly**. Starting at midnight on the morning of the first day of each month covering the entire calendar month following.\n\n* **Annually**. Starting at midnight on the morning of the first day of each year covering the entire calendar year following.\n\n* **Ad_Hoc**. Use this setting when a custom billing schedule is used for billing an Account, such as for billing of Prepayment/Commitment fees using a custom billing schedule.\nAvailable values: \"DAILY\", \"WEEKLY\", \"MONTHLY\", \"ANNUALLY\", \"AD_HOC\".",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"DAILY",
						"WEEKLY",
						"MONTHLY",
						"ANNUALLY",
						"AD_HOC",
					),
				},
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"day_epoch": schema.StringAttribute{
				Description:   "The starting date *(epoch)* for Daily billing frequency *(in ISO 8601 format)*, determining the first Bill date for daily Bills.",
				Optional:      true,
				CustomType:    timetypes.RFC3339Type{},
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"due_date": schema.StringAttribute{
				Description:   "The due date *(in ISO 8601 format)* for payment of the Bill.\n\nFor example: `\"2023-02-24\"`.",
				Optional:      true,
				CustomType:    timetypes.RFC3339Type{},
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"external_invoice_date": schema.StringAttribute{
				Description:   "For accounting purposes, the date set at Organization level to use for external invoicing with respect to billing periods - two options:\n* `FIRST_DAY_OF_NEXT_PERIOD` *(Default)*. Used when you want to recognize usage revenue in the following period.\n* `LAST_DAY_OF_ARREARS`. Used when you want to recognize usage revenue in the same period that it's consumed, instead of in the following period.\n\nFor example, if the retrieved Bill was on a monthly billing frequency and the billing period for the Bill is September 2023 and the *External invoice date* is set at `FIRST_DAY_OF_NEXT_PERIOD`, then the `externalInvoiceDate` will be `\"2023-10-01\"`.\n\n**NOTE:** To change the `externalInvoiceDate` setting for your Organization, you can use the [Update OrganizationConfig](https://www.m3ter.com/docs/api#tag/OrganizationConfig/operation/GetOrganizationConfig) call.",
				Optional:      true,
				CustomType:    timetypes.RFC3339Type{},
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"last_date_in_billing_period": schema.StringAttribute{
				Description:   "Specifies the date *(in ISO 8601 format)* of the last day in the billing period, defining the time range for the associated Bills.\n\nFor example: `\"2023-03-24\"`.",
				Optional:      true,
				CustomType:    timetypes.RFC3339Type{},
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"month_epoch": schema.StringAttribute{
				Description:   "The starting date *(epoch)* for Monthly billing frequency *(in ISO 8601 format)*, determining the first Bill date for monthly Bills.",
				Optional:      true,
				CustomType:    timetypes.RFC3339Type{},
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"target_currency": schema.StringAttribute{
				Description:   "The currency code used for the Bill, such as USD, GBP, or EUR.",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"version": schema.Int64Attribute{
				Description:   "The version number of the entity:\n- **Create entity:** Not valid for initial insertion of new entity - *do not use for Create*. On initial Create, version is set at 1 and listed in the response.\n- **Update Entity:**  On Update, version is required and must match the existing version because a check is performed to ensure sequential versioning is preserved. Version is incremented by 1 and listed in the response.",
				Optional:      true,
				PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
			},
			"week_epoch": schema.StringAttribute{
				Description:   "The starting date *(epoch)* for Weekly billing frequency *(in ISO 8601 format)*, determining the first Bill date for weekly Bills.",
				Optional:      true,
				CustomType:    timetypes.RFC3339Type{},
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"year_epoch": schema.StringAttribute{
				Description:   "The starting date *(epoch)* for Yearly billing frequency *(in ISO 8601 format)*, determining the first Bill date for yearly Bills.",
				Optional:      true,
				CustomType:    timetypes.RFC3339Type{},
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"account_ids": schema.ListAttribute{
				Description:   "An array of UUIDs representing the end customer Accounts associated with the BillJob.",
				Optional:      true,
				ElementType:   types.StringType,
				PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
			},
			"timezone": schema.StringAttribute{
				Description:   "Specifies the time zone used for the generated Bills, ensuring alignment with the local time zone.",
				Computed:      true,
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
				Default:       stringdefault.StaticString("UTC"),
			},
			"currency_conversions": schema.ListNestedAttribute{
				Description: "An array of currency conversion rates from Bill currency to Organization currency. For example, if Account is billed in GBP and Organization is set to USD, Bill line items are calculated in GBP and then converted to USD using the defined rate.",
				Computed:    true,
				Optional:    true,
				CustomType:  customfield.NewNestedObjectListType[BillJobCurrencyConversionsModel](ctx),
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
				PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
			},
			"created_by": schema.StringAttribute{
				Description: "The unique identifier (UUID) for the user who created the BillJob.",
				Computed:    true,
			},
			"dt_created": schema.StringAttribute{
				Description: "The date and time *(in ISO 8601 format)* when the BillJob was first created.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"dt_last_modified": schema.StringAttribute{
				Description: "The date and time *(in ISO 8601 format)* when the BillJob was last modified.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"last_modified_by": schema.StringAttribute{
				Description: "The unique identifier (UUID) for the user who last modified this BillJob.",
				Computed:    true,
			},
			"pending": schema.Int64Attribute{
				Description: "The number of pending actions or calculations within the BillJob.",
				Computed:    true,
			},
			"status": schema.StringAttribute{
				Description: "The current status of the BillJob, indicating its progress or completion state.\nAvailable values: \"PENDING\", \"INITIALIZING\", \"RUNNING\", \"COMPLETE\", \"CANCELLED\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"PENDING",
						"INITIALIZING",
						"RUNNING",
						"COMPLETE",
						"CANCELLED",
					),
				},
			},
			"total": schema.Int64Attribute{
				Description: "The total number of Bills or calculations related to the BillJob.",
				Computed:    true,
			},
			"type": schema.StringAttribute{
				Description: "Specifies the type of BillJob. \n\n* **CREATE** Returned for a *Create BillJob* call.\n* **RECALCULATE** Returned for a successful *Create Recalculation BillJob* call.\nAvailable values: \"CREATE\", \"RECALCULATE\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("CREATE", "RECALCULATE"),
				},
			},
			"bill_ids": schema.ListAttribute{
				Description: "An array of Bill IDs related to the BillJob, providing references to the specific Bills generated.",
				Computed:    true,
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
			},
		},
	}
}

func (r *BillJobResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *BillJobResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
