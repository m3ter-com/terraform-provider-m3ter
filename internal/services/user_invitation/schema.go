// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user_invitation

import (
  "context"

  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework/resource"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*UserInvitationResource)(nil)

func ResourceSchema(ctx context.Context) (schema.Schema) {
  return schema.Schema{
    Attributes: map[string]schema.Attribute{
      "id": schema.StringAttribute{
        Description: "The UUID of the invitation.",
        Computed: true,
        PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
      },
      "org_id": schema.StringAttribute{
        Required: true,
        PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
      },
      "email": schema.StringAttribute{
        Required: true,
        PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
      },
      "first_name": schema.StringAttribute{
        Required: true,
        PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
      },
      "last_name": schema.StringAttribute{
        Required: true,
        PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
      },
      "contact_number": schema.StringAttribute{
        Optional: true,
        PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
      },
      "dt_end_access": schema.StringAttribute{
        Description: "The date when access will end for the user *(in ISO-8601 format)*. Leave blank for no end date, which gives the user permanent access.",
        Optional: true,
        CustomType: timetypes.RFC3339Type{

        },
        PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
      },
      "dt_expiry": schema.StringAttribute{
        Description: "The date when the invite expires *(in ISO-8601 format)*. After this date the invited user can no longer accept the invite. By default, any invite is valid for 30 days from the date the invite is sent.",
        Optional: true,
        CustomType: timetypes.RFC3339Type{

        },
        PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
      },
      "m3ter_user": schema.BoolAttribute{
        Optional: true,
        PlanModifiers: []planmodifier.Bool{boolplanmodifier.RequiresReplace()},
      },
      "version": schema.Int64Attribute{
        Optional: true,
        PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
      },
      "permission_policy_ids": schema.ListAttribute{
        Description: "The IDs of the permission policies the invited user has been assigned. This controls the access rights and privileges that this user will have when working in the m3ter Organization.",
        Optional: true,
        ElementType: types.StringType,
        PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
      },
      "accepted": schema.BoolAttribute{
        Description: "Boolean indicating whether the user has accepted the invitation.\n\n* TRUE - the invite has been accepted.\n* FALSE - the invite has not yet been accepted.",
        Computed: true,
      },
      "created_by": schema.StringAttribute{
        Description: "The UUID of the user who created the invitation.",
        Computed: true,
      },
      "dt_created": schema.StringAttribute{
        Description: "The DateTime when the invitation was created *(in ISO-8601 format)*.",
        Computed: true,
        CustomType: timetypes.RFC3339Type{

        },
      },
      "dt_last_modified": schema.StringAttribute{
        Description: "The DateTime when the invitation was last modified *(in ISO-8601 format)*.",
        Computed: true,
        CustomType: timetypes.RFC3339Type{

        },
      },
      "inviting_principal_id": schema.StringAttribute{
        Description: "The UUID of the user who sent the invite.",
        Computed: true,
      },
      "last_modified_by": schema.StringAttribute{
        Description: "The UUID of the user who last modified the invitation.",
        Computed: true,
      },
    },
  }
}

func (r *UserInvitationResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
  resp.Schema = ResourceSchema(ctx)
}

func (r *UserInvitationResource) ConfigValidators(_ context.Context) ([]resource.ConfigValidator) {
  return []resource.ConfigValidator{
  }
}
