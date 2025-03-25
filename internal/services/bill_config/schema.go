// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package bill_config

import (
  "context"

  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework/resource"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

var _ resource.ResourceWithConfigValidators = (*BillConfigResource)(nil)

func ResourceSchema(ctx context.Context) (schema.Schema) {
  return schema.Schema{
    Attributes: map[string]schema.Attribute{
      "id": schema.StringAttribute{
        Description: "The Organization UUID. The Organization represents your company as a direct customer of the m3ter service.",
        Computed: true,
        PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
      },
      "org_id": schema.StringAttribute{
        Required: true,
        PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
      },
      "bill_lock_date": schema.StringAttribute{
        Description: "The global lock date when all Bills will be locked *(in ISO 8601 format)*.\n\nFor example: `\"2024-03-01\"`.",
        Optional: true,
        CustomType: timetypes.RFC3339Type{

        },
      },
      "version": schema.Int64Attribute{
        Description: "The version number:\n* Default value when newly created is one.\n* On Update, version is required and must match the existing version because a check is performed to ensure sequential versioning is preserved. Version is incremented by 1 and listed in the response",
        Optional: true,
      },
      "created_by": schema.StringAttribute{
        Description: "The id of the user who created this bill config.",
        Computed: true,
      },
      "dt_created": schema.StringAttribute{
        Description: "The DateTime *(in ISO-8601 format)* when the bill config was first created.",
        Computed: true,
        CustomType: timetypes.RFC3339Type{

        },
      },
      "dt_last_modified": schema.StringAttribute{
        Description: "The DateTime *(in ISO-8601 format)* when the bill config was last modified.",
        Computed: true,
        CustomType: timetypes.RFC3339Type{

        },
      },
      "last_modified_by": schema.StringAttribute{
        Description: "The id of the user who last modified this bill config.",
        Computed: true,
      },
    },
  }
}

func (r *BillConfigResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
  resp.Schema = ResourceSchema(ctx)
}

func (r *BillConfigResource) ConfigValidators(_ context.Context) ([]resource.ConfigValidator) {
  return []resource.ConfigValidator{
  }
}
