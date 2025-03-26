// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package bill_credit_line_item

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
var _ resource.ResourceWithConfigure = (*BillCreditLineItemResource)(nil)
var _ resource.ResourceWithModifyPlan = (*BillCreditLineItemResource)(nil)
var _ resource.ResourceWithImportState = (*BillCreditLineItemResource)(nil)

func NewResource() resource.Resource {
  return &BillCreditLineItemResource{}
}

// BillCreditLineItemResource defines the resource implementation.
type BillCreditLineItemResource struct {
  client *m3ter.Client
}

func (r *BillCreditLineItemResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
  resp.TypeName = req.ProviderTypeName + "_bill_credit_line_item"
}

func (r *BillCreditLineItemResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *BillCreditLineItemResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
  var data *BillCreditLineItemModel

  resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

  if resp.Diagnostics.HasError() {
    return
  }

  params := m3ter.BillCreditLineItemNewParams{

  }

  if !data.OrgID.IsNull() {
    params.OrgID = m3ter.F(data.OrgID.ValueString())
  }

  dataBytes, err := data.MarshalJSON()
  if err != nil {
    resp.Diagnostics.AddError("failed to serialize http request", err.Error())
    return
  }
  res := new(http.Response)
  _, err = r.client.Bills.CreditLineItems.New(
    ctx,
    data.BillID.ValueString(),
    params,
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

func (r *BillCreditLineItemResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
  var data  *BillCreditLineItemModel

  resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

  if resp.Diagnostics.HasError() {
    return
  }

  var state  *BillCreditLineItemModel

  resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

  if resp.Diagnostics.HasError() {
    return
  }

  params := m3ter.BillCreditLineItemUpdateParams{

  }

  if !data.OrgID.IsNull() {
    params.OrgID = m3ter.F(data.OrgID.ValueString())
  }

  dataBytes, err := data.MarshalJSONForUpdate(*state)
  if err != nil {
    resp.Diagnostics.AddError("failed to serialize http request", err.Error())
    return
  }
  res := new(http.Response)
  _, err = r.client.Bills.CreditLineItems.Update(
    ctx,
    data.BillID.ValueString(),
    data.ID.ValueString(),
    params,
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

func (r *BillCreditLineItemResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
  var data  *BillCreditLineItemModel

  resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

  if resp.Diagnostics.HasError() {
    return
  }

  params := m3ter.BillCreditLineItemGetParams{

  }

  if !data.OrgID.IsNull() {
    params.OrgID = m3ter.F(data.OrgID.ValueString())
  }

  res := new(http.Response)
  _, err := r.client.Bills.CreditLineItems.Get(
    ctx,
    data.BillID.ValueString(),
    data.ID.ValueString(),
    params,
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

func (r *BillCreditLineItemResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
  var data  *BillCreditLineItemModel

  resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

  if resp.Diagnostics.HasError() {
    return
  }

  params := m3ter.BillCreditLineItemDeleteParams{

  }

  if !data.OrgID.IsNull() {
    params.OrgID = m3ter.F(data.OrgID.ValueString())
  }

  _, err := r.client.Bills.CreditLineItems.Delete(
    ctx,
    data.BillID.ValueString(),
    data.ID.ValueString(),
    params,
    option.WithMiddleware(logging.Middleware(ctx)),
  )
  if err != nil {
    resp.Diagnostics.AddError("failed to make http request", err.Error())
    return
  }

  resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *BillCreditLineItemResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
  var data *BillCreditLineItemModel = new(BillCreditLineItemModel)

  path_org_id := ""
  path_bill_id := ""
  path_id := ""
  diags := importpath.ParseImportID(
    req.ID,
    "<org_id>/<bill_id>/<id>",
    &path_org_id,
    &path_bill_id,
    &path_id,
  )
  resp.Diagnostics.Append(diags...)
  if resp.Diagnostics.HasError() {
    return
  }

  data.OrgID = types.StringValue(path_org_id)
  data.BillID = types.StringValue(path_bill_id)
  data.ID = types.StringValue(path_id)

  res := new(http.Response)
  _, err := r.client.Bills.CreditLineItems.Get(
    ctx,
    path_bill_id,
    path_id,
    m3ter.BillCreditLineItemGetParams{
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

func (r *BillCreditLineItemResource) ModifyPlan(_ context.Context, _ resource.ModifyPlanRequest, _ *resource.ModifyPlanResponse) {

}
