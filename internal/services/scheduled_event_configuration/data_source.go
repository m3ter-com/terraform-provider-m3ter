// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package scheduled_event_configuration

import (
  "context"
  "fmt"
  "io"
  "net/http"

  "github.com/hashicorp/terraform-plugin-framework/datasource"
  "github.com/m3ter-com/m3ter-sdk-go"
  "github.com/m3ter-com/m3ter-sdk-go/option"
  "github.com/m3ter-com/terraform-provider-m3ter/internal/apijson"
  "github.com/m3ter-com/terraform-provider-m3ter/internal/logging"
)

type ScheduledEventConfigurationDataSource struct {
  client *m3ter.Client
}

var _ datasource.DataSourceWithConfigure = (*ScheduledEventConfigurationDataSource)(nil)

func NewScheduledEventConfigurationDataSource() datasource.DataSource {
  return &ScheduledEventConfigurationDataSource{}
}

func (d *ScheduledEventConfigurationDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
  resp.TypeName = req.ProviderTypeName + "_scheduled_event_configuration"
}

func (d *ScheduledEventConfigurationDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
  if req.ProviderData == nil {
    return
  }

  client, ok := req.ProviderData.(*m3ter.Client)

  if !ok {
    resp.Diagnostics.AddError(
      "unexpected resource configure type",
      fmt.Sprintf("Expected *m3ter.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
    )

    return
  }

  d.client = client
}

func (d *ScheduledEventConfigurationDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
  var data *ScheduledEventConfigurationDataSourceModel

  resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

  if resp.Diagnostics.HasError() {
    return
  }

  params, diags := data.toReadParams(ctx)
  resp.Diagnostics.Append(diags...)
  if resp.Diagnostics.HasError() {
    return
  }

  res := new(http.Response)
  _, err := d.client.ScheduledEventConfigurations.Get(
    ctx,
    data.ID.ValueString(),
    params,
    option.WithResponseBodyInto(&res),
    option.WithMiddleware(logging.Middleware(ctx)),
  )
  if err != nil {
    resp.Diagnostics.AddError("failed to make http request", err.Error())
    return
  }
  bytes, _ := io.ReadAll(res.Body)
  err = apijson.UnmarshalComputed(bytes, &data)
  if err != nil {
    resp.Diagnostics.AddError("failed to deserialize http request", err.Error())
    return
  }

  resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
