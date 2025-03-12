// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package webhook

import (
  "context"
  "fmt"

  "github.com/hashicorp/terraform-plugin-framework/attr"
  "github.com/hashicorp/terraform-plugin-framework/datasource"
  "github.com/m3ter-com/m3ter-sdk-go"
  "github.com/m3ter-com/terraform-provider-m3ter/internal/apijson"
  "github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

type WebhooksDataSource struct {
  client *m3ter.Client
}

var _ datasource.DataSourceWithConfigure = (*WebhooksDataSource)(nil)

func NewWebhooksDataSource() datasource.DataSource {
  return &WebhooksDataSource{}
}

func (d *WebhooksDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
  resp.TypeName = req.ProviderTypeName + "_webhooks"
}

func (d *WebhooksDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *WebhooksDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
  var data *WebhooksDataSourceModel

  resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

  if resp.Diagnostics.HasError() {
    return
  }

  params, diags := data.toListParams(ctx)
  resp.Diagnostics.Append(diags...)
  if resp.Diagnostics.HasError() {
    return
  }

  env := WebhooksDataListDataSourceEnvelope{}
  maxItems := int(data.MaxItems.ValueInt64())
  acc := []attr.Value{}
  if maxItems <= 0 {
    maxItems = 1000
  }
  page, err := d.client.Webhooks.List(ctx, params)
  if err != nil {
    resp.Diagnostics.AddError("failed to make http request", err.Error())
    return
  }

  for page != nil && len(page.Data) > 0 {
    bytes := []byte(page.JSON.RawJSON())
    err = apijson.UnmarshalComputed(bytes, &env)
    if err != nil {
      resp.Diagnostics.AddError("failed to unmarshal http request", err.Error())
      return
    }
    acc = append(acc, env.Data.Elements()...)
    if len(acc) >= maxItems {
      break
    }
    page, err = page.GetNextPage()
    if err != nil {
      resp.Diagnostics.AddError("failed to fetch next page", err.Error())
      return
    }
  }

  acc = acc[:min(len(acc), maxItems)]
  result, diags := customfield.NewObjectListFromAttributes[WebhooksItemsDataSourceModel](ctx, acc)
  resp.Diagnostics.Append(diags...)
  data.Items = result

  resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
