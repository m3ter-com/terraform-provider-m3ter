// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package data_export_destination

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

var _ datasource.DataSourceWithConfigValidators = (*DataExportDestinationsDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) (schema.Schema) {
  return schema.Schema{
    Attributes: map[string]schema.Attribute{
      "org_id": schema.StringAttribute{
        Optional: true,
      },
      "ids": schema.ListAttribute{
        Description: "List of Export Destination UUIDs to retrieve.",
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
        CustomType: customfield.NewNestedObjectListType[DataExportDestinationsItemsDataSourceModel](ctx),
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
            "code": schema.StringAttribute{
              Description: "The code of the data Export Destination.",
              Computed: true,
            },
            "created_by": schema.StringAttribute{
              Description: "The id of the user who created the Export Destination.",
              Computed: true,
            },
            "dt_created": schema.StringAttribute{
              Description: "The DateTime when the Export Destination was created.",
              Computed: true,
              CustomType: timetypes.RFC3339Type{

              },
            },
            "dt_last_modified": schema.StringAttribute{
              Description: "The DateTime when the Export Destination was last modified.",
              Computed: true,
              CustomType: timetypes.RFC3339Type{

              },
            },
            "last_modified_by": schema.StringAttribute{
              Description: "The id of the user who last modified the Export Destination.",
              Computed: true,
            },
            "name": schema.StringAttribute{
              Description: "The name of the data Export Destination.",
              Computed: true,
            },
          },
        },
      },
    },
  }
}

func (d *DataExportDestinationsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
  resp.Schema = ListDataSourceSchema(ctx)
}

func (d *DataExportDestinationsDataSource) ConfigValidators(_ context.Context) ([]datasource.ConfigValidator) {
  return []datasource.ConfigValidator{
  }
}
