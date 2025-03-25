// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_field

import (
  "context"

  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework/resource"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*CustomFieldResource)(nil)

func ResourceSchema(ctx context.Context) (schema.Schema) {
  return schema.Schema{
    Attributes: map[string]schema.Attribute{
      "id": schema.StringAttribute{
        Description: "The UUID of the entity.",
        Computed: true,
        PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
      },
      "org_id": schema.StringAttribute{
        Required: true,
        PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
      },
      "version": schema.Int64Attribute{
        Description: "The version number of the entity:\n- **Create entity:** Not valid for initial insertion of new entity - *do not use for Create*. On initial Create, version is set at 1 and listed in the response.\n- **Update Entity:**  On Update, version is required and must match the existing version because a check is performed to ensure sequential versioning is preserved. Version is incremented by 1 and listed in the response.",
        Optional: true,
      },
      "account": schema.MapAttribute{
        Description: "Updates to Account entity CustomFields.",
        Optional: true,
        ElementType: types.DynamicType,
      },
      "account_plan": schema.MapAttribute{
        Description: "Updates to accountPlan entity CustomFields.",
        Optional: true,
        ElementType: types.DynamicType,
      },
      "aggregation": schema.MapAttribute{
        Description: "Updates to simple Aggregation entity CustomFields.",
        Optional: true,
        ElementType: types.DynamicType,
      },
      "compound_aggregation": schema.MapAttribute{
        Description: "Updates to Compound Aggregation entity CustomFields.",
        Optional: true,
        ElementType: types.DynamicType,
      },
      "meter": schema.MapAttribute{
        Description: "Updates to Meter entity CustomFields.",
        Optional: true,
        ElementType: types.DynamicType,
      },
      "organization": schema.MapAttribute{
        Description: "Updates to Organization CustomFields.",
        Optional: true,
        ElementType: types.DynamicType,
      },
      "plan": schema.MapAttribute{
        Description: "Updates to Plan entity CustomFields.",
        Optional: true,
        ElementType: types.DynamicType,
      },
      "plan_template": schema.MapAttribute{
        Description: "Updates to planTemplate entity CustomFields.",
        Optional: true,
        ElementType: types.DynamicType,
      },
      "product": schema.MapAttribute{
        Description: "Updates to Product entity CustomFields.",
        Optional: true,
        ElementType: types.DynamicType,
      },
      "created_by": schema.StringAttribute{
        Description: "The id of the user who created this custom field.",
        Computed: true,
      },
      "dt_created": schema.StringAttribute{
        Description: "The DateTime when the Organization was created *(in ISO-8601 format)*.",
        Computed: true,
        CustomType: timetypes.RFC3339Type{

        },
      },
      "dt_last_modified": schema.StringAttribute{
        Description: "The DateTime when a custom field was last modified - created, modified, or deleted - for the Organization *(in ISO-8601 format)*.",
        Computed: true,
        CustomType: timetypes.RFC3339Type{

        },
      },
      "last_modified_by": schema.StringAttribute{
        Description: "The id of the user who last modified this custom field.",
        Computed: true,
      },
    },
  }
}

func (r *CustomFieldResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
  resp.Schema = ResourceSchema(ctx)
}

func (r *CustomFieldResource) ConfigValidators(_ context.Context) ([]resource.ConfigValidator) {
  return []resource.ConfigValidator{
  }
}
