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
	"github.com/m3ter-com/terraform-provider-m3ter/internal/services/aggregation"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/services/compound_aggregation"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/services/counter"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/services/meter"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/services/product"
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
		aggregation.NewResource,
		compound_aggregation.NewResource,
		counter.NewResource,
		meter.NewResource,
		product.NewResource,
	}
}

func (p *M3terProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		aggregation.NewAggregationDataSource,
		aggregation.NewAggregationsDataSource,
		compound_aggregation.NewCompoundAggregationDataSource,
		compound_aggregation.NewCompoundAggregationsDataSource,
		counter.NewCounterDataSource,
		counter.NewCountersDataSource,
		meter.NewMeterDataSource,
		meter.NewMetersDataSource,
		product.NewProductDataSource,
		product.NewProductsDataSource,
	}
}

func NewProvider(version string) func() provider.Provider {
	return func() provider.Provider {
		return &M3terProvider{
			version: version,
		}
	}
}
