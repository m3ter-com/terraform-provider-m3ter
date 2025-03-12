// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pricing

import (
  "context"
  "fmt"
  "io"
  "net/http"

  "github.com/hashicorp/terraform-plugin-framework/resource"
  "github.com/hashicorp/terraform-plugin-framework/types"
  "github.com/m3ter-com/m3ter-sdk-go"
  "github.com/m3ter-com/m3ter-sdk-go/option"
  "github.com/m3ter-com/terraform-provider-m3ter/internal/apijson"
  "github.com/m3ter-com/terraform-provider-m3ter/internal/importpath"
  "github.com/m3ter-com/terraform-provider-m3ter/internal/logging"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.ResourceWithConfigure = (*PricingResource)(nil)
var _ resource.ResourceWithModifyPlan = (*PricingResource)(nil)
var _ resource.ResourceWithImportState = (*PricingResource)(nil)

func NewResource() resource.Resource {
  return &PricingResource{}
}

// PricingResource defines the resource implementation.
type PricingResource struct {
  client *m3ter.Client
}

func (r *PricingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
  resp.TypeName = req.ProviderTypeName + "_pricing"
}

func (r *PricingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *PricingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
  var data *PricingModel

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
  _, err = r.client.Pricings.New(
    ctx,
    m3ter.PricingNewParams{
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

func (r *PricingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
  var data  *PricingModel

  resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

  if resp.Diagnostics.HasError() {
    return
  }

  var state  *PricingModel

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
  _, err = r.client.Pricings.Update(
    ctx,
    data.ID.ValueString(),
    m3ter.PricingUpdateParams{
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

func (r *PricingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
  var data  *PricingModel

  resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

  if resp.Diagnostics.HasError() {
    return
  }

  res := new(http.Response)
  _, err := r.client.Pricings.Get(
    ctx,
    data.ID.ValueString(),
    m3ter.PricingGetParams{
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

func (r *PricingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
  var data  *PricingModel

  resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

  if resp.Diagnostics.HasError() {
    return
  }

  _, err := r.client.Pricings.Delete(
    ctx,
    data.ID.ValueString(),
    m3ter.PricingDeleteParams{
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

func (r *PricingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
  var data *PricingModel = new(PricingModel)

  path_org_id := ""
  path_id := ""
  diags := importpath.ParseImportID(
    req.ID,
    "<org_id>/<id>",
    &path_org_id,
    &path_id,
  )
  resp.Diagnostics.Append(diags...)
  if resp.Diagnostics.HasError() {
    return
  }

  data.OrgID = types.StringValue(path_org_id)
  data.ID = types.StringValue(path_id)

  res := new(http.Response)
  _, err := r.client.Pricings.Get(
    ctx,
    path_id,
    m3ter.PricingGetParams{
      OrgID: m3ter.F(path_org_id),
    },
    option.WithResponseBodyInto(&res),
    option.WithMiddleware(logging.Middleware(ctx)),
  )
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

func (r *PricingResource) ModifyPlan(_ context.Context, _ resource.ModifyPlanRequest, _ *resource.ModifyPlanResponse) {

}
