// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package bill_job

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

var _ datasource.DataSourceWithConfigValidators = (*BillJobDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Required: true,
			},
			"org_id": schema.StringAttribute{
				Required: true,
			},
			"bill_date": schema.StringAttribute{
				Description: "The specific billing date *(in ISO 8601 format)*, determining when the Bill was generated.\n\nFor example: `\"2023-01-24\"`.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"bill_frequency_interval": schema.Int64Attribute{
				Description: "How often Bills are issued - used in conjunction with `billingFrequency`.\n\nFor example, if `billingFrequency` is set to Monthly and `billFrequencyInterval` is set to 3, Bills are issued every three months.",
				Computed:    true,
			},
			"billing_frequency": schema.StringAttribute{
				Description: "Defines how often Bills are generated.\n\n* **Daily**. Starting at midnight each day, covering a twenty-four hour period following.\n\n* **Weekly**. Starting at midnight on a Monday morning covering the seven-day period following.\n\n* **Monthly**. Starting at midnight on the morning of the first day of each month covering the entire calendar month following.\n\n* **Annually**. Starting at midnight on the morning of the first day of each year covering the entire calendar year following.\n\n* **Ad_Hoc**. Use this setting when a custom billing schedule is used for billing an Account, such as for billing of Prepayment/Commitment fees using a custom billing schedule.\nAvailable values: \"DAILY\", \"WEEKLY\", \"MONTHLY\", \"ANNUALLY\", \"AD_HOC\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"DAILY",
						"WEEKLY",
						"MONTHLY",
						"ANNUALLY",
						"AD_HOC",
					),
				},
			},
			"created_by": schema.StringAttribute{
				Description: "The unique identifier (UUID) for the user who created the BillJob.",
				Computed:    true,
			},
			"day_epoch": schema.StringAttribute{
				Description: "The starting date *(epoch)* for Daily billing frequency *(in ISO 8601 format)*, determining the first Bill date for daily Bills.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
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
			"due_date": schema.StringAttribute{
				Description: "The due date *(in ISO 8601 format)* for payment of the Bill.\n\nFor example: `\"2023-02-24\"`.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"external_invoice_date": schema.StringAttribute{
				Description: "For accounting purposes, the date set at Organization level to use for external invoicing with respect to billing periods - two options:\n* `FIRST_DAY_OF_NEXT_PERIOD` *(Default)*. Used when you want to recognize usage revenue in the following period.\n* `LAST_DAY_OF_ARREARS`. Used when you want to recognize usage revenue in the same period that it's consumed, instead of in the following period.\n\nFor example, if the retrieved Bill was on a monthly billing frequency and the billing period for the Bill is September 2023 and the *External invoice date* is set at `FIRST_DAY_OF_NEXT_PERIOD`, then the `externalInvoiceDate` will be `\"2023-10-01\"`.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"last_date_in_billing_period": schema.StringAttribute{
				Description: "Specifies the date *(in ISO 8601 format)* of the last day in the billing period, defining the time range for the associated Bills.\n\nFor example: `\"2023-03-24\"`.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"last_modified_by": schema.StringAttribute{
				Description: "The unique identifier (UUID) for the user who last modified this BillJob.",
				Computed:    true,
			},
			"month_epoch": schema.StringAttribute{
				Description: "The starting date *(epoch)* for Monthly billing frequency *(in ISO 8601 format)*, determining the first Bill date for monthly Bills.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
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
			"target_currency": schema.StringAttribute{
				Description: "The currency code used for the Bill, such as USD, GBP, or EUR.",
				Computed:    true,
			},
			"timezone": schema.StringAttribute{
				Description: "Specifies the time zone used for the generated Bills, ensuring alignment with the local time zone.",
				Computed:    true,
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
			"version": schema.Int64Attribute{
				Description: "The version number:\n- **Create:** On initial Create to insert a new entity, the version is set at 1 in the response.\n- **Update:** On successful Update, the version is incremented by 1 in the response.",
				Computed:    true,
			},
			"week_epoch": schema.StringAttribute{
				Description: "The starting date *(epoch)* for Weekly billing frequency *(in ISO 8601 format)*, determining the first Bill date for weekly Bills.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"year_epoch": schema.StringAttribute{
				Description: "The starting date *(epoch)* for Yearly billing frequency *(in ISO 8601 format)*, determining the first Bill date for yearly Bills.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"account_ids": schema.ListAttribute{
				Description: "An array of UUIDs representing the end customer Accounts associated with the BillJob.",
				Computed:    true,
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
			},
			"bill_ids": schema.ListAttribute{
				Description: "An array of Bill IDs related to the BillJob, providing references to the specific Bills generated.",
				Computed:    true,
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
			},
			"currency_conversions": schema.ListNestedAttribute{
				Description: "An array of currency conversion rates from Bill currency to Organization currency. For example, if Account is billed in GBP and Organization is set to USD, Bill line items are calculated in GBP and then converted to USD using the defined rate.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[BillJobCurrencyConversionsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"from": schema.StringAttribute{
							Description: "Currency to convert from. For example: GBP.",
							Computed:    true,
						},
						"to": schema.StringAttribute{
							Description: "Currency to convert to. For example: USD.",
							Computed:    true,
						},
						"multiplier": schema.Float64Attribute{
							Description: "Conversion rate between currencies.",
							Computed:    true,
							Validators: []validator.Float64{
								float64validator.AtLeast(0),
							},
						},
					},
				},
			},
		},
	}
}

func (d *BillJobDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *BillJobDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
