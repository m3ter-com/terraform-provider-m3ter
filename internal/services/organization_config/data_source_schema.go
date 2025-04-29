// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package organization_config

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

var _ datasource.DataSourceWithConfigValidators = (*OrganizationConfigDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"org_id": schema.StringAttribute{
				Required:           true,
				DeprecationMessage: "the org id should be set at the client level instead",
			},
			"auto_approve_bills_grace_period": schema.Int64Attribute{
				Description: "Grace period before bills are auto-approved. Used in combination with the field `autoApproveBillsGracePeriodUnit`.",
				Computed:    true,
			},
			"auto_approve_bills_grace_period_unit": schema.StringAttribute{
				Description: `Available values: "MINUTES", "HOURS", "DAYS".`,
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"MINUTES",
						"HOURS",
						"DAYS",
					),
				},
			},
			"auto_generate_statement_mode": schema.StringAttribute{
				Description: "Specifies whether to auto-generate statements once Bills are *approved* or *locked*. It will not auto-generate if a bill is in *pending* state.\n\nThe default value is **None**.\n\n- **None**. Statements will not be auto-generated.\n- **JSON**. Statements are auto-generated in JSON format.\n- **JSON and CSV**. Statements are auto-generated in both JSON and CSV formats.\nAvailable values: \"NONE\", \"JSON\", \"JSON_AND_CSV\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"NONE",
						"JSON",
						"JSON_AND_CSV",
					),
				},
			},
			"bill_prefix": schema.StringAttribute{
				Description: "Prefix to be used for sequential invoice numbers. This will be combined with the `sequenceStartNumber`.",
				Computed:    true,
			},
			"commitment_fee_bill_in_advance": schema.BoolAttribute{
				Description: "Specifies whether commitments *(prepayments)* are billed in advance at the start of each billing period, or billed in arrears at the end of each billing period. \n\n* **TRUE** - bill in advance *(start of each billing period)*.\n* **FALSE** - bill in arrears *(end of each billing period)*.",
				Computed:    true,
			},
			"consolidate_bills": schema.BoolAttribute{
				Description: "Specifies whether to consolidate different billing frequencies onto the same bill.\n\n* **TRUE** - consolidate different billing frequencies onto the same bill.\n* **FALSE** - bills are not consolidated.",
				Computed:    true,
			},
			"created_by": schema.StringAttribute{
				Description: "The id of the user who created this organization config.",
				Computed:    true,
			},
			"currency": schema.StringAttribute{
				Description: "The currency code for the currency used in this Organization. For example: USD, GBP, or EUR.",
				Computed:    true,
			},
			"day_epoch": schema.StringAttribute{
				Description: "The first bill date *(in ISO-8601 format)* for daily billing periods.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"days_before_bill_due": schema.Int64Attribute{
				Description: "The number of days after the Bill generation date shown on Bills as the due date.",
				Computed:    true,
			},
			"default_statement_definition_id": schema.StringAttribute{
				Description: "Organization level default `statementDefinitionId` to be used when there is no statement definition linked to the account.\n\nStatement definitions are used to generate bill statements, which are informative backing sheets to invoices.",
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
			"external_invoice_date": schema.StringAttribute{
				Description: `Available values: "LAST_DAY_OF_ARREARS", "FIRST_DAY_OF_NEXT_PERIOD".`,
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("LAST_DAY_OF_ARREARS", "FIRST_DAY_OF_NEXT_PERIOD"),
				},
			},
			"id": schema.StringAttribute{
				Description: "The UUID of the entity.",
				Computed:    true,
			},
			"last_modified_by": schema.StringAttribute{
				Description: "The id of the user who last modified this organization config.",
				Computed:    true,
			},
			"minimum_spend_bill_in_advance": schema.BoolAttribute{
				Description: "Specifies whether minimum spend amounts are billed in advance at the start of each billing period, or billed in arrears at the end of each billing period. \n\n* **TRUE** - bill in advance *(start of each billing period)*.\n* **FALSE** - bill in arrears *(end of each billing period)*.",
				Computed:    true,
			},
			"month_epoch": schema.StringAttribute{
				Description: "The first bill date *(in ISO-8601 format)* for monthly billing periods.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"scheduled_bill_interval": schema.Float64Attribute{
				Description: "Specifies the required interval for updating bills. \n\n* **For portions of an hour (minutes)**. Two options: **0.25** (15 minutes) and **0.5** (30 minutes).\n* **For full hours.** Eight possible values: **1**, **2**, **3**, **4**, **6**, **8**, **12**, or **24**.\n* **Default.** The default is **0**, which disables scheduling.",
				Computed:    true,
			},
			"sequence_start_number": schema.Int64Attribute{
				Description: "The starting number to be used for sequential invoice numbers. This will be combined with the `billPrefix`.",
				Computed:    true,
			},
			"standing_charge_bill_in_advance": schema.BoolAttribute{
				Description: "Specifies whether the standing charge is billed in advance at the start of each billing period, or billed in arrears at the end of each billing period. \n\n* **TRUE** - bill in advance *(start of each billing period)*.\n* **FALSE** - bill in arrears *(end of each billing period)*.",
				Computed:    true,
			},
			"suppressed_empty_bills": schema.BoolAttribute{
				Description: "Specifies whether to supress generating bills that have no line items. \n\n* **TRUE** - prevents generating bills with no line items.\n* **FALSE** - bills are still generated even when they have no line items.",
				Computed:    true,
			},
			"timezone": schema.StringAttribute{
				Description: "The timezone for the Organization.",
				Computed:    true,
			},
			"version": schema.Int64Attribute{
				Description: "The version number:\n- **Create:** On initial Create to insert a new entity, the version is set at 1 in the response.\n- **Update:** On successful Update, the version is incremented by 1 in the response.",
				Computed:    true,
			},
			"week_epoch": schema.StringAttribute{
				Description: "The first bill date *(in ISO-8601 format)* for weekly billing periods.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"year_epoch": schema.StringAttribute{
				Description: "The first bill date *(in ISO-8601 format)* for yearly billing periods.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"credit_application_order": schema.ListAttribute{
				Description: "The order in which any Prepayment or Balance credit amounts on Accounts are to be drawn-down against for billing. Four options:\n- `\"PREPAYMENT\",\"BALANCE\"`. Draw-down against Prepayment credit before Balance credit.\n- `\"BALANCE\",\"PREPAYMENT\"`. Draw-down against Balance credit before Prepayment credit.\n- `\"PREPAYMENT\"`. Only draw-down against Prepayment credit.\n- `\"BALANCE\"`. Only draw-down against Balance credit.",
				Computed:    true,
				Validators: []validator.List{
					listvalidator.ValueStringsAre(
						stringvalidator.OneOfCaseInsensitive("PREPAYMENT", "BALANCE"),
					),
				},
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
			},
			"currency_conversions": schema.ListNestedAttribute{
				Description: "Currency conversion rates from Bill currency to Organization currency. \n\nFor example, if Account is billed in GBP and Organization is set to USD, Bill line items are calculated in GBP and then converted to USD using the defined rate.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[OrganizationConfigCurrencyConversionsDataSourceModel](ctx),
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

func (d *OrganizationConfigDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *OrganizationConfigDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
