// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package currency

import (
  "context"

  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
  "github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
  "github.com/hashicorp/terraform-plugin-framework/resource"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
  "github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = (*CurrencyResource)(nil)

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
      "name": schema.StringAttribute{
        Description: "The name of the entity.",
        Required: true,
      },
      "archived": schema.BoolAttribute{
        Description: "A Boolean TRUE / FALSE flag indicating whether the entity is archived. An entity can be archived if it is obsolete.\n\n* TRUE - the entity is in the archived state.\n* FALSE - the entity is not in the archived state.",
        Optional: true,
      },
      "code": schema.StringAttribute{
        Description: "The short code for the entity.",
        Optional: true,
      },
      "max_decimal_places": schema.Int64Attribute{
        Description: "Indicates the maximum number of decimal places to use for this Currency.",
        Optional: true,
        Validators: []validator.Int64{
        int64validator.Between(0, 5),
        },
      },
      "rounding_mode": schema.StringAttribute{
        Description: `Available values: "UP", "DOWN", "CEILING", "FLOOR", "HALF_UP", "HALF_DOWN", "HALF_EVEN", "UNNECESSARY".`,
        Optional: true,
        Validators: []validator.String{
        stringvalidator.OneOfCaseInsensitive(
          "UP",
          "DOWN",
          "CEILING",
          "FLOOR",
          "HALF_UP",
          "HALF_DOWN",
          "HALF_EVEN",
          "UNNECESSARY",
        ),
        },
      },
      "version": schema.Int64Attribute{
        Description: "The version number of the entity:\n- **Create entity:** Not valid for initial insertion of new entity - *do not use for Create*. On initial Create, version is set at 1 and listed in the response.\n- **Update Entity:**  On Update, version is required and must match the existing version because a check is performed to ensure sequential versioning is preserved. Version is incremented by 1 and listed in the response.",
        Optional: true,
      },
      "created_by": schema.StringAttribute{
        Description: "The unique identifier (UUID) of the user who created this Currency.",
        Computed: true,
      },
      "dt_created": schema.StringAttribute{
        Description: "The date and time *(in ISO-8601 format)* when the Currency was created.",
        Computed: true,
        CustomType: timetypes.RFC3339Type{

        },
      },
      "dt_last_modified": schema.StringAttribute{
        Description: "The date and time *(in ISO-8601 format)* when the Currency was last modified.",
        Computed: true,
        CustomType: timetypes.RFC3339Type{

        },
      },
      "last_modified_by": schema.StringAttribute{
        Description: "The unique identifier (UUID) of the user who last modified this Currency.",
        Computed: true,
      },
    },
  }
}

func (r *CurrencyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
  resp.Schema = ResourceSchema(ctx)
}

func (r *CurrencyResource) ConfigValidators(_ context.Context) ([]resource.ConfigValidator) {
  return []resource.ConfigValidator{
  }
}
