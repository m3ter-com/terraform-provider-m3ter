// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package internal

import (
	"context"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
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
	"github.com/m3ter-com/terraform-provider-m3ter/internal/services/bill_config"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/services/commitment"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/services/compound_aggregation"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/services/contract"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/services/counter"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/services/counter_adjustment"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/services/counter_pricing"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/services/credit_reason"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/services/currency"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/services/custom_field"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/services/debit_reason"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/services/external_mapping"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/services/integration_configuration"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/services/meter"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/services/notification_configuration"
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
	APIKey    types.String `tfsdk:"api_key" json:"api_key,optional"`
	APISecret types.String `tfsdk:"api_secret" json:"api_secret,optional"`
	Token     types.String `tfsdk:"token" json:"token,optional"`
	OrgID     types.String `tfsdk:"org_id" json:"org_id,optional"`
}

func (p *M3terProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "m3ter"
	resp.Version = p.version
}

func ProviderSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"base_url": schema.StringAttribute{
				Description: "Set the base url that the provider connects to.",
				Optional:    true,
			},
			"api_key": schema.StringAttribute{
				Optional: true,
			},
			"api_secret": schema.StringAttribute{
				Optional: true,
			},
			"token": schema.StringAttribute{
				Optional: true,
			},
			"org_id": schema.StringAttribute{
				Optional: true,
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

	if !data.BaseURL.IsNull() && !data.BaseURL.IsUnknown() {
		opts = append(opts, option.WithBaseURL(data.BaseURL.ValueString()))
	} else if o, ok := os.LookupEnv("M3TER_BASE_URL"); ok {
		opts = append(opts, option.WithBaseURL(o))
	}

	if !data.APIKey.IsNull() && !data.APIKey.IsUnknown() {
		opts = append(opts, option.WithAPIKey(data.APIKey.ValueString()))
	} else if o, ok := os.LookupEnv("M3TER_API_KEY"); ok {
		opts = append(opts, option.WithAPIKey(o))
	} else {
		resp.Diagnostics.AddAttributeError(
			path.Root("api_key"),
			"Missing api_key value",
			"The api_key field is required. Set it in provider configuration or via the \"M3TER_API_KEY\" environment variable.",
		)
		return
	}

	if !data.APISecret.IsNull() && !data.APISecret.IsUnknown() {
		opts = append(opts, option.WithAPISecret(data.APISecret.ValueString()))
	} else if o, ok := os.LookupEnv("M3TER_API_SECRET"); ok {
		opts = append(opts, option.WithAPISecret(o))
	} else {
		resp.Diagnostics.AddAttributeError(
			path.Root("api_secret"),
			"Missing api_secret value",
			"The api_secret field is required. Set it in provider configuration or via the \"M3TER_API_SECRET\" environment variable.",
		)
		return
	}

	if !data.Token.IsNull() && !data.Token.IsUnknown() {
		opts = append(opts, option.WithToken(data.Token.ValueString()))
	} else if o, ok := os.LookupEnv("M3TER_API_TOKEN"); ok {
		opts = append(opts, option.WithToken(o))
	}

	if !data.OrgID.IsNull() && !data.OrgID.IsUnknown() {
		opts = append(opts, option.WithOrgID(data.OrgID.ValueString()))
	} else if o, ok := os.LookupEnv("M3TER_ORG_ID"); ok {
		opts = append(opts, option.WithOrgID(o))
	} else {
		resp.Diagnostics.AddAttributeError(
			path.Root("org_id"),
			"Missing org_id value",
			"The org_id field is required. Set it in provider configuration or via the \"M3TER_ORG_ID\" environment variable.",
		)
		return
	}

	client := m3ter.NewClientWithServiceUserAuth(
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
		bill_config.NewResource,
		commitment.NewResource,
		compound_aggregation.NewResource,
		contract.NewResource,
		counter.NewResource,
		counter_adjustment.NewResource,
		counter_pricing.NewResource,
		credit_reason.NewResource,
		currency.NewResource,
		custom_field.NewResource,
		debit_reason.NewResource,
		external_mapping.NewResource,
		integration_configuration.NewResource,
		meter.NewResource,
		notification_configuration.NewResource,
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
		bill_config.NewBillConfigDataSource,
		commitment.NewCommitmentDataSource,
		commitment.NewCommitmentsDataSource,
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
		debit_reason.NewDebitReasonDataSource,
		debit_reason.NewDebitReasonsDataSource,
		external_mapping.NewExternalMappingDataSource,
		external_mapping.NewExternalMappingsDataSource,
		integration_configuration.NewIntegrationConfigurationDataSource,
		integration_configuration.NewIntegrationConfigurationsDataSource,
		meter.NewMeterDataSource,
		meter.NewMetersDataSource,
		notification_configuration.NewNotificationConfigurationDataSource,
		notification_configuration.NewNotificationConfigurationsDataSource,
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
