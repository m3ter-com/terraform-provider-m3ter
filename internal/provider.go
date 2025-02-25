// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package internal

import (
	"context"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/m3ter-sdk-go"
	"github.com/m3ter-com/m3ter-sdk-go/option"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/services/account"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/services/account_plan"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/services/aggregation"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/services/balance"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/services/balance_transaction"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/services/bill"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/services/bill_config"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/services/bill_credit_line_item"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/services/bill_debit_line_item"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/services/bill_job"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/services/bill_line_item"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/services/commitment"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/services/compound_aggregation"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/services/contract"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/services/counter"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/services/counter_adjustment"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/services/counter_pricing"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/services/credit_reason"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/services/currency"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/services/custom_field"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/services/data_export_destination"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/services/data_export_job"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/services/data_export_schedule"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/services/debit_reason"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/services/event"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/services/external_mapping"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/services/integration_configuration"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/services/meter"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/services/notification_configuration"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/services/organization_config"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/services/permission_policy"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/services/plan"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/services/plan_group"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/services/plan_group_link"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/services/plan_template"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/services/pricing"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/services/product"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/services/resource_group"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/services/scheduled_event_configuration"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/services/transaction_type"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/services/usage_file_upload_job"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/services/user"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/services/user_invitation"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/services/webhook"
)

var _ provider.ProviderWithConfigValidators = (*M3terProvider)(nil)

// M3terProvider defines the provider implementation.
type M3terProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// M3terProviderModel describes the provider data model.
type M3terProviderModel struct {
	BaseURL   types.String `tfsdk:"base_url" json:"base_url,optional"`
	APIKey    types.String `tfsdk:"api_key" json:"api_key,required"`
	APISecret types.String `tfsdk:"api_secret" json:"api_secret,required"`
	Token     types.String `tfsdk:"token" json:"token,optional"`
	OrgID     types.String `tfsdk:"org_id" json:"org_id,required"`
}

func (p *M3terProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "m3ter"
	resp.Version = p.version
}

func ProviderSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"base_url": schema.StringAttribute{
				Description: "Set the base url that the provider connects to. This can be used for testing in other environments.",
				Optional:    true,
			},
			"api_key": schema.StringAttribute{
				Required: true,
			},
			"api_secret": schema.StringAttribute{
				Required: true,
			},
			"token": schema.StringAttribute{
				Optional: true,
			},
			"org_id": schema.StringAttribute{
				Required: true,
			},
		},
	}
}

func (p *M3terProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = ProviderSchema(ctx)
}

func (p *M3terProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {

	var data M3terProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	opts := []option.RequestOption{}

	if !data.BaseURL.IsNull() {
		opts = append(opts, option.WithBaseURL(data.BaseURL.ValueString()))
	}
	if o, ok := os.LookupEnv("M3TER_API_KEY"); ok {
		opts = append(opts, option.WithAPIKey(o))
	}
	if o, ok := os.LookupEnv("M3TER_API_SECRET"); ok {
		opts = append(opts, option.WithAPISecret(o))
	}
	if o, ok := os.LookupEnv("M3TER_API_TOKEN"); ok {
		opts = append(opts, option.WithToken(o))
	}
	if !data.APIKey.IsNull() {
		opts = append(opts, option.WithAPIKey(data.APIKey.ValueString()))
	}
	if !data.APISecret.IsNull() {
		opts = append(opts, option.WithAPISecret(data.APISecret.ValueString()))
	}
	if !data.Token.IsNull() {
		opts = append(opts, option.WithToken(data.Token.ValueString()))
	}
	if !data.OrgID.IsNull() {
		opts = append(opts, option.WithOrgID(data.OrgID.ValueString()))
	}

	client := m3ter.NewClient(
		opts...,
	)

	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *M3terProvider) ConfigValidators(_ context.Context) []provider.ConfigValidator {
	return []provider.ConfigValidator{}
}

func (p *M3terProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		account.NewResource,
		account_plan.NewResource,
		aggregation.NewResource,
		balance.NewResource,
		balance_transaction.NewResource,
		bill_credit_line_item.NewResource,
		bill_debit_line_item.NewResource,
		bill_config.NewResource,
		commitment.NewResource,
		bill_job.NewResource,
		compound_aggregation.NewResource,
		contract.NewResource,
		counter.NewResource,
		counter_adjustment.NewResource,
		counter_pricing.NewResource,
		credit_reason.NewResource,
		currency.NewResource,
		custom_field.NewResource,
		data_export_destination.NewResource,
		data_export_schedule.NewResource,
		debit_reason.NewResource,
		external_mapping.NewResource,
		integration_configuration.NewResource,
		meter.NewResource,
		notification_configuration.NewResource,
		organization_config.NewResource,
		permission_policy.NewResource,
		plan.NewResource,
		plan_group.NewResource,
		plan_group_link.NewResource,
		plan_template.NewResource,
		pricing.NewResource,
		product.NewResource,
		resource_group.NewResource,
		scheduled_event_configuration.NewResource,
		transaction_type.NewResource,
		user.NewResource,
		user_invitation.NewResource,
		webhook.NewResource,
	}
}

func (p *M3terProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		account.NewAccountDataSource,
		account.NewAccountsDataSource,
		account_plan.NewAccountPlanDataSource,
		account_plan.NewAccountPlansDataSource,
		aggregation.NewAggregationDataSource,
		aggregation.NewAggregationsDataSource,
		balance.NewBalanceDataSource,
		balance.NewBalancesDataSource,
		balance_transaction.NewBalanceTransactionsDataSource,
		bill.NewBillDataSource,
		bill.NewBillsDataSource,
		bill_credit_line_item.NewBillCreditLineItemDataSource,
		bill_credit_line_item.NewBillCreditLineItemsDataSource,
		bill_debit_line_item.NewBillDebitLineItemDataSource,
		bill_debit_line_item.NewBillDebitLineItemsDataSource,
		bill_line_item.NewBillLineItemDataSource,
		bill_line_item.NewBillLineItemsDataSource,
		bill_config.NewBillConfigDataSource,
		commitment.NewCommitmentDataSource,
		commitment.NewCommitmentsDataSource,
		bill_job.NewBillJobDataSource,
		bill_job.NewBillJobsDataSource,
		compound_aggregation.NewCompoundAggregationDataSource,
		compound_aggregation.NewCompoundAggregationsDataSource,
		contract.NewContractDataSource,
		contract.NewContractsDataSource,
		counter.NewCounterDataSource,
		counter.NewCountersDataSource,
		counter_adjustment.NewCounterAdjustmentDataSource,
		counter_adjustment.NewCounterAdjustmentsDataSource,
		counter_pricing.NewCounterPricingDataSource,
		counter_pricing.NewCounterPricingsDataSource,
		credit_reason.NewCreditReasonDataSource,
		credit_reason.NewCreditReasonsDataSource,
		currency.NewCurrencyDataSource,
		currency.NewCurrenciesDataSource,
		custom_field.NewCustomFieldDataSource,
		data_export_destination.NewDataExportDestinationDataSource,
		data_export_destination.NewDataExportDestinationsDataSource,
		data_export_job.NewDataExportJobDataSource,
		data_export_job.NewDataExportJobsDataSource,
		data_export_schedule.NewDataExportScheduleDataSource,
		data_export_schedule.NewDataExportSchedulesDataSource,
		debit_reason.NewDebitReasonDataSource,
		debit_reason.NewDebitReasonsDataSource,
		event.NewEventDataSource,
		event.NewEventsDataSource,
		external_mapping.NewExternalMappingDataSource,
		external_mapping.NewExternalMappingsDataSource,
		integration_configuration.NewIntegrationConfigurationDataSource,
		integration_configuration.NewIntegrationConfigurationsDataSource,
		meter.NewMeterDataSource,
		meter.NewMetersDataSource,
		notification_configuration.NewNotificationConfigurationDataSource,
		notification_configuration.NewNotificationConfigurationsDataSource,
		organization_config.NewOrganizationConfigDataSource,
		permission_policy.NewPermissionPolicyDataSource,
		permission_policy.NewPermissionPoliciesDataSource,
		plan.NewPlanDataSource,
		plan.NewPlansDataSource,
		plan_group.NewPlanGroupDataSource,
		plan_group.NewPlanGroupsDataSource,
		plan_group_link.NewPlanGroupLinkDataSource,
		plan_group_link.NewPlanGroupLinksDataSource,
		plan_template.NewPlanTemplateDataSource,
		plan_template.NewPlanTemplatesDataSource,
		pricing.NewPricingDataSource,
		pricing.NewPricingsDataSource,
		product.NewProductDataSource,
		product.NewProductsDataSource,
		resource_group.NewResourceGroupDataSource,
		resource_group.NewResourceGroupsDataSource,
		scheduled_event_configuration.NewScheduledEventConfigurationDataSource,
		scheduled_event_configuration.NewScheduledEventConfigurationsDataSource,
		transaction_type.NewTransactionTypeDataSource,
		transaction_type.NewTransactionTypesDataSource,
		usage_file_upload_job.NewUsageFileUploadJobDataSource,
		usage_file_upload_job.NewUsageFileUploadJobsDataSource,
		user.NewUserDataSource,
		user.NewUsersDataSource,
		user_invitation.NewUserInvitationDataSource,
		user_invitation.NewUserInvitationsDataSource,
		webhook.NewWebhookDataSource,
		webhook.NewWebhooksDataSource,
	}
}

func NewProvider(version string) func() provider.Provider {
	return func() provider.Provider {
		return &M3terProvider{
			version: version,
		}
	}
}
