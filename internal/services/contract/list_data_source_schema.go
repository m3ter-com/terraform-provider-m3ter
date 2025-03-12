// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package contract

import (
  "context"

  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
  "github.com/hashicorp/terraform-plugin-framework/datasource"
  "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
  "github.com/hashicorp/terraform-plugin-framework/schema/validator"
  "github.com/hashicorp/terraform-plugin-framework/types"
  "github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

var _ datasource.DataSourceWithConfigValidators = (*ContractsDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) (schema.Schema) {
  return schema.Schema{
    Attributes: map[string]schema.Attribute{
      "org_id": schema.StringAttribute{
        Required: true,
      },
      "account_id": schema.StringAttribute{
        Optional: true,
      },
      "codes": schema.ListAttribute{
        Description: "An optional parameter to retrieve specific Contracts based on their short codes.",
        Optional: true,
        ElementType: types.StringType,
      },
      "ids": schema.ListAttribute{
        Description: "An optional parameter to filter the list based on specific Contract unique identifiers (UUIDs).",
        Optional: true,
        ElementType: types.StringType,
      },
      "max_items": schema.Int64Attribute{
        Description: "Max items to fetch, default: 1000",
        Optional: true,
        Validators: []validator.Int64{
        int64validator.AtLeast(0),
        },
      },
      "items": schema.ListNestedAttribute{
        Description: "The items returned by the data source",
        Computed: true,
        CustomType: customfield.NewNestedObjectListType[ContractsItemsDataSourceModel](ctx),
        NestedObject: schema.NestedAttributeObject{
          Attributes: map[string]schema.Attribute{
            "id": schema.StringAttribute{
              Description: "The UUID of the entity.",
              Computed: true,
            },
            "version": schema.Int64Attribute{
              Description: "The version number:\n- **Create:** On initial Create to insert a new entity, the version is set at 1 in the response.\n- **Update:** On successful Update, the version is incremented by 1 in the response.",
              Computed: true,
            },
            "account_id": schema.StringAttribute{
              Description: "The unique identifier (UUID) of the Account associated with this Contract.",
              Computed: true,
            },
            "code": schema.StringAttribute{
              Description: "The short code of the Contract.",
              Computed: true,
            },
            "created_by": schema.StringAttribute{
              Description: "The unique identifier (UUID) of the user who created this Contract.",
              Computed: true,
            },
            "custom_fields": schema.MapAttribute{
              Description: "User defined fields enabling you to attach custom data. The value for a custom field can be either a string or a number.\n\nIf `customFields` can also be defined for this entity at the Organizational level,`customField` values defined at individual level override values of `customFields` with the same name defined at Organization level.\n\nSee [Working with Custom Fields](https://www.m3ter.com/docs/guides/creating-and-managing-products/working-with-custom-fields) in the m3ter documentation for more information.",
              Computed: true,
              CustomType: customfield.NewMapType[types.Dynamic](ctx),
              ElementType: types.DynamicType,
            },
            "description": schema.StringAttribute{
              Description: "The description of the Contract, which provides context and information.",
              Computed: true,
            },
            "dt_created": schema.StringAttribute{
              Description: "The date and time *(in ISO-8601 format)* when the Contract was created.",
              Computed: true,
              CustomType: timetypes.RFC3339Type{

              },
            },
            "dt_last_modified": schema.StringAttribute{
              Description: "The date and time *(in ISO-8601 format)* when the Contract was last modified.",
              Computed: true,
              CustomType: timetypes.RFC3339Type{

              },
            },
            "end_date": schema.StringAttribute{
              Description: "The exclusive end date of the Contract *(in ISO-8601 format)*. This means the Contract is active until midnight on the day ***before*** this date.",
              Computed: true,
              CustomType: timetypes.RFC3339Type{

              },
            },
            "last_modified_by": schema.StringAttribute{
              Description: "The unique identifier (UUID) of the user who last modified this Contract.",
              Computed: true,
            },
            "name": schema.StringAttribute{
              Description: "The name of the Contract.",
              Computed: true,
            },
            "purchase_order_number": schema.StringAttribute{
              Description: "The Purchase Order Number associated with the Contract.",
              Computed: true,
            },
            "start_date": schema.StringAttribute{
              Description: "The start date for the Contract *(in ISO-8601 format)*. This date is inclusive, meaning the Contract is active from this date onward.",
              Computed: true,
              CustomType: timetypes.RFC3339Type{

              },
            },
          },
        },
      },
    },
  }
}

func (d *ContractsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
  resp.Schema = ListDataSourceSchema(ctx)
}

func (d *ContractsDataSource) ConfigValidators(_ context.Context) ([]datasource.ConfigValidator) {
  return []datasource.ConfigValidator{
  }
}
