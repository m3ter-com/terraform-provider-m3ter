// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package resource_group

import (
  "context"

  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
  "github.com/hashicorp/terraform-plugin-framework/datasource"
  "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
  "github.com/hashicorp/terraform-plugin-framework/schema/validator"
  "github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

var _ datasource.DataSourceWithConfigValidators = (*ResourceGroupsDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) (schema.Schema) {
  return schema.Schema{
    Attributes: map[string]schema.Attribute{
      "type": schema.StringAttribute{
        Required: true,
      },
      "org_id": schema.StringAttribute{
        Optional: true,
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
        CustomType: customfield.NewNestedObjectListType[ResourceGroupsItemsDataSourceModel](ctx),
        NestedObject: schema.NestedAttributeObject{
          Attributes: map[string]schema.Attribute{
            "id": schema.StringAttribute{
              Description: "The unique identifier (UUID) of the Resource Group.",
              Computed: true,
            },
            "created_by": schema.StringAttribute{
              Description: "The unique identifier (UUID) of the user who created this Resource Group.",
              Computed: true,
            },
            "dt_created": schema.StringAttribute{
              Description: "The date and time *(in ISO-8601 format)* when the Resource Group was created.",
              Computed: true,
              CustomType: timetypes.RFC3339Type{

              },
            },
            "dt_last_modified": schema.StringAttribute{
              Description: "The date and time *(in ISO-8601 format)* when the Resource Group was last modified.",
              Computed: true,
              CustomType: timetypes.RFC3339Type{

              },
            },
            "last_modified_by": schema.StringAttribute{
              Description: "The unique identifier (UUID) of the user who last modified this Resource Group.",
              Computed: true,
            },
            "name": schema.StringAttribute{
              Description: "The name of the Resource Group.",
              Computed: true,
            },
            "version": schema.Int64Attribute{
              Description: "The version number. Default value when newly created is one.",
              Computed: true,
            },
          },
        },
      },
    },
  }
}

func (d *ResourceGroupsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
  resp.Schema = ListDataSourceSchema(ctx)
}

func (d *ResourceGroupsDataSource) ConfigValidators(_ context.Context) ([]datasource.ConfigValidator) {
  return []datasource.ConfigValidator{
  }
}
