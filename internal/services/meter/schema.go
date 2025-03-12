// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package meter

import (
  "context"

  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
  "github.com/hashicorp/terraform-plugin-framework/resource"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
  "github.com/hashicorp/terraform-plugin-framework/schema/validator"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*MeterResource)(nil)

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
      "code": schema.StringAttribute{
        Description: "Code of the Meter - unique short code used to identify the Meter.\n\n**NOTE:** Code has a maximum length of 80 characters and must not contain non-printable or whitespace characters (except space), and cannot start/end with whitespace.",
        Required: true,
      },
      "name": schema.StringAttribute{
        Description: "Descriptive name for the Meter.",
        Required: true,
      },
      "data_fields": schema.ListNestedAttribute{
        Description: "Used to submit categorized raw usage data values for ingest into the platform - either numeric quantitative values or non-numeric data values. At least one required per Meter; maximum 15 per Meter.",
        Required: true,
        NestedObject: schema.NestedAttributeObject{
          Attributes: map[string]schema.Attribute{
            "category": schema.StringAttribute{
              Description: "The type of field (WHO, WHAT, WHERE, MEASURE, METADATA, INCOME, COST, OTHER).\nAvailable values: \"WHO\", \"WHERE\", \"WHAT\", \"OTHER\", \"METADATA\", \"MEASURE\", \"INCOME\", \"COST\".",
              Required: true,
              Validators: []validator.String{
              stringvalidator.OneOfCaseInsensitive(
                "WHO",
                "WHERE",
                "WHAT",
                "OTHER",
                "METADATA",
                "MEASURE",
                "INCOME",
                "COST",
              ),
              },
            },
            "code": schema.StringAttribute{
              Description: "Short code to identify the field\n\n**NOTE:** Code has a maximum length of 80 characters and can only contain letters, numbers, underscore, and the dollar character, and must not start with a number.",
              Required: true,
            },
            "name": schema.StringAttribute{
              Description: "Descriptive name of the field.",
              Required: true,
            },
            "unit": schema.StringAttribute{
              Description: "The units to measure the data with. Should conform to *Unified Code for Units of Measure* (UCUM). Required only for numeric field categories.",
              Optional: true,
            },
          },
        },
      },
      "derived_fields": schema.ListNestedAttribute{
        Description: "Used to submit usage data values for ingest into the platform that are the result of a calculation performed on `dataFields`, `customFields`, or system `Timestamp` fields. Raw usage data is not submitted using `derivedFields`. Maximum 15 per Meter. *(Optional)*.\n\n**Note:** Required parameter. If you want to create a Meter without Derived Fields, use an empty array `[]`. If you use a `null`, you'll receive an error.",
        Required: true,
        NestedObject: schema.NestedAttributeObject{
          Attributes: map[string]schema.Attribute{
            "category": schema.StringAttribute{
              Description: "The type of field (WHO, WHAT, WHERE, MEASURE, METADATA, INCOME, COST, OTHER).\nAvailable values: \"WHO\", \"WHERE\", \"WHAT\", \"OTHER\", \"METADATA\", \"MEASURE\", \"INCOME\", \"COST\".",
              Required: true,
              Validators: []validator.String{
              stringvalidator.OneOfCaseInsensitive(
                "WHO",
                "WHERE",
                "WHAT",
                "OTHER",
                "METADATA",
                "MEASURE",
                "INCOME",
                "COST",
              ),
              },
            },
            "code": schema.StringAttribute{
              Description: "Short code to identify the field\n\n**NOTE:** Code has a maximum length of 80 characters and can only contain letters, numbers, underscore, and the dollar character, and must not start with a number.",
              Required: true,
            },
            "name": schema.StringAttribute{
              Description: "Descriptive name of the field.",
              Required: true,
            },
            "unit": schema.StringAttribute{
              Description: "The units to measure the data with. Should conform to *Unified Code for Units of Measure* (UCUM). Required only for numeric field categories.",
              Optional: true,
            },
            "calculation": schema.StringAttribute{
              Description: "The calculation used to transform the value of submitted `dataFields` in usage data. Calculation can reference `dataFields`, `customFields`, or system `Timestamp` fields. \n*(Example: datafieldms  datafieldgb)*",
              Required: true,
            },
          },
        },
      },
      "group_id": schema.StringAttribute{
        Description: "UUID of the group the Meter belongs to. *(Optional)*.",
        Optional: true,
      },
      "product_id": schema.StringAttribute{
        Description: "UUID of the product the Meter belongs to. *(Optional)* - if left blank, the Meter is global.",
        Optional: true,
      },
      "version": schema.Int64Attribute{
        Description: "The version number of the entity:\n- **Create entity:** Not valid for initial insertion of new entity - *do not use for Create*. On initial Create, version is set at 1 and listed in the response.\n- **Update Entity:**  On Update, version is required and must match the existing version because a check is performed to ensure sequential versioning is preserved. Version is incremented by 1 and listed in the response.",
        Optional: true,
      },
      "custom_fields": schema.MapAttribute{
        Description: "User defined fields enabling you to attach custom data. The value for a custom field can be either a string or a number.\n\nIf `customFields` can also be defined for this entity at the Organizational level, `customField` values defined at individual level override values of `customFields` with the same name defined at Organization level.\n\nSee [Working with Custom Fields](https://www.m3ter.com/docs/guides/creating-and-managing-products/working-with-custom-fields) in the m3ter documentation for more information.",
        Optional: true,
        ElementType: types.DynamicType,
      },
      "created_by": schema.StringAttribute{
        Description: "The id of the user who created this meter.",
        Computed: true,
      },
      "dt_created": schema.StringAttribute{
        Description: "The DateTime when the meter was created *(in ISO-8601 format)*.",
        Computed: true,
        CustomType: timetypes.RFC3339Type{

        },
      },
      "dt_last_modified": schema.StringAttribute{
        Description: "The DateTime when the meter was last modified *(in ISO-8601 format)*.",
        Computed: true,
        CustomType: timetypes.RFC3339Type{

        },
      },
      "last_modified_by": schema.StringAttribute{
        Description: "The id of the user who last modified this meter.",
        Computed: true,
      },
    },
  }
}

func (r *MeterResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
  resp.Schema = ResourceSchema(ctx)
}

func (r *MeterResource) ConfigValidators(_ context.Context) ([]resource.ConfigValidator) {
  return []resource.ConfigValidator{
  }
}
