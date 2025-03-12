// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package data_export_destination

import (
  "context"
  "fmt"
  "io"
  "net/http"

  "github.com/hashicorp/terraform-plugin-framework/resource"
  "github.com/m3ter-com/m3ter-sdk-go"
  "github.com/m3ter-com/m3ter-sdk-go/option"
  "github.com/m3ter-com/terraform-provider-m3ter/internal/apijson"
  "github.com/m3ter-com/terraform-provider-m3ter/internal/logging"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.ResourceWithConfigure = (*DataExportDestinationResource)(nil)
var _ resource.ResourceWithModifyPlan = (*DataExportDestinationResource)(nil)

func NewResource() resource.Resource {
  return &DataExportDestinationResource{}
}

// DataExportDestinationResource defines the resource implementation.
type DataExportDestinationResource struct {
  client *m3ter.Client
}

func (r *DataExportDestinationResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
  resp.TypeName = req.ProviderTypeName + "_data_export_destination"
}

func (r *DataExportDestinationResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

  r.client = client
}

func (r *DataExportDestinationResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
  var data *DataExportDestinationModel

  resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

  if resp.Diagnostics.HasError() {
    return
  }

  dataBytes, err := data.MarshalJSON()
  if err != nil {
    resp.Diagnostics.AddError("failed to serialize http request", err.Error())
    return
  }
  res := new(http.Response)
  _, err = r.client.DataExports.Destinations.New(
    ctx,
    m3ter.DataExportDestinationNewParams{
      OrgID: m3ter.F(data.OrgID.ValueString()),
    },
    option.WithRequestBody("application/json", dataBytes),
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

func (r *DataExportDestinationResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
  var data  *DataExportDestinationModel

  resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

  if resp.Diagnostics.HasError() {
    return
  }

  var state  *DataExportDestinationModel

  resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

  if resp.Diagnostics.HasError() {
    return
  }

  dataBytes, err := data.MarshalJSONForUpdate(*state)
  if err != nil {
    resp.Diagnostics.AddError("failed to serialize http request", err.Error())
    return
  }
  res := new(http.Response)
  _, err = r.client.DataExports.Destinations.Update(
    ctx,
    data.ID.ValueString(),
    m3ter.DataExportDestinationUpdateParams{
      OrgID: m3ter.F(data.OrgID.ValueString()),
    },
    option.WithRequestBody("application/json", dataBytes),
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

func (r *DataExportDestinationResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
  var data  *DataExportDestinationModel

  resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

  if resp.Diagnostics.HasError() {
    return
  }

  res := new(http.Response)
  _, err := r.client.DataExports.Destinations.Get(
    ctx,
    data.ID.ValueString(),
    m3ter.DataExportDestinationGetParams{
      OrgID: m3ter.F(data.OrgID.ValueString()),
    },
    option.WithResponseBodyInto(&res),
    option.WithMiddleware(logging.Middleware(ctx)),
  )
  if res != nil && res.StatusCode == 404 {
  resp.Diagnostics.AddWarning("Resource not found", "The resource was not found on the server and will be removed from state.")
    resp.State.RemoveResource(ctx)
    return
  }
  if err != nil {
    resp.Diagnostics.AddError("failed to make http request", err.Error())
    return
  }
  bytes, _ := io.ReadAll(res.Body)
  err = apijson.Unmarshal(bytes, &data)
  if err != nil {
    resp.Diagnostics.AddError("failed to deserialize http request", err.Error())
    return
  }

  resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DataExportDestinationResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
  var data  *DataExportDestinationModel

  resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

  if resp.Diagnostics.HasError() {
    return
  }

  _, err := r.client.DataExports.Destinations.Delete(
    ctx,
    data.ID.ValueString(),
    m3ter.DataExportDestinationDeleteParams{
      OrgID: m3ter.F(data.OrgID.ValueString()),
    },
    option.WithMiddleware(logging.Middleware(ctx)),
  )
  if err != nil {
    resp.Diagnostics.AddError("failed to make http request", err.Error())
    return
  }

  resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DataExportDestinationResource) ModifyPlan(_ context.Context, _ resource.ModifyPlanRequest, _ *resource.ModifyPlanResponse) {

}
