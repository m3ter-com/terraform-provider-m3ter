// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package bill_job

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
var _ resource.ResourceWithConfigure = (*BillJobResource)(nil)
var _ resource.ResourceWithModifyPlan = (*BillJobResource)(nil)
var _ resource.ResourceWithImportState = (*BillJobResource)(nil)

func NewResource() resource.Resource {
  return &BillJobResource{}
}

// BillJobResource defines the resource implementation.
type BillJobResource struct {
  client *m3ter.Client
}

func (r *BillJobResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
  resp.TypeName = req.ProviderTypeName + "_bill_job"
}

func (r *BillJobResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *BillJobResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
  var data *BillJobModel

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
  _, err = r.client.BillJobs.New(
    ctx,
    m3ter.BillJobNewParams{
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

func (r *BillJobResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
  var data  *BillJobModel

  resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

  if resp.Diagnostics.HasError() {
    return
  }

  var state  *BillJobModel

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
  _, err = r.client.BillJobs.New(
    ctx,
    m3ter.BillJobNewParams{
      OrgID: m3ter.F(data.ID.ValueString()),
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

func (r *BillJobResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
  var data  *BillJobModel

  resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

  if resp.Diagnostics.HasError() {
    return
  }

  res := new(http.Response)
  _, err := r.client.BillJobs.Get(
    ctx,
    data.ID.ValueString(),
    m3ter.BillJobGetParams{
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

func (r *BillJobResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

}

func (r *BillJobResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
  var data *BillJobModel = new(BillJobModel)

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
  _, err := r.client.BillJobs.Get(
    ctx,
    path_id,
    m3ter.BillJobGetParams{
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

func (r *BillJobResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
  if req.State.Raw.IsNull() {
      resp.Diagnostics.AddWarning(
          "Resource Destruction Considerations",
          "This resource cannot be destroyed from Terraform. If you create this resource, it will be "+
          "present in the API until manually deleted.",
      )
  }
  if req.Plan.Raw.IsNull() {
      resp.Diagnostics.AddWarning(
          "Resource Destruction Considerations",
          "Applying this resource destruction will remove the resource from the Terraform state "+
              "but will not change it in the API. If you would like to destroy or reset this resource "+
              "in the API, refer to the documentation for how to do it manually.",
      )
  }
}
