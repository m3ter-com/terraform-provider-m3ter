// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package data_export_schedule

import (
  "context"

  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
  "github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
  "github.com/hashicorp/terraform-plugin-framework/datasource"
  "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
  "github.com/hashicorp/terraform-plugin-framework/schema/validator"
  "github.com/hashicorp/terraform-plugin-framework/types"
  "github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

var _ datasource.DataSourceWithConfigValidators = (*DataExportSchedulesDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) (schema.Schema) {
  return schema.Schema{
    Attributes: map[string]schema.Attribute{
      "org_id": schema.StringAttribute{
        Required: true,
      },
      "ids": schema.ListAttribute{
        Description: "Data Export Schedule IDs to filter the returned list by.",
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
        CustomType: customfield.NewNestedObjectListType[DataExportSchedulesItemsDataSourceModel](ctx),
        NestedObject: schema.NestedAttributeObject{
          Attributes: map[string]schema.Attribute{
            "id": schema.StringAttribute{
              Description: "The id of the Data Export Schedule.",
              Computed: true,
            },
            "version": schema.Int64Attribute{
              Description: "The version number:\n- **Create:** On initial Create to insert a new entity, the version is set at 1 in the response.\n- **Update:** On successful Update, the version is incremented by 1 in the response.",
              Computed: true,
            },
            "code": schema.StringAttribute{
              Description: "Unique short code of the Data Export Schedule.",
              Computed: true,
            },
            "created_by": schema.StringAttribute{
              Description: "The id of the user who created this Schedule.",
              Computed: true,
            },
            "destination_ids": schema.ListAttribute{
              Description: "The Export Destination ids.",
              Computed: true,
              CustomType: customfield.NewListType[types.String](ctx),
              ElementType: types.StringType,
            },
            "dt_created": schema.StringAttribute{
              Description: "The DateTime when the Data Export Schedule was created.",
              Computed: true,
              CustomType: timetypes.RFC3339Type{

              },
            },
            "dt_last_modified": schema.StringAttribute{
              Description: "The DateTime when the Schedule was last modified.",
              Computed: true,
              CustomType: timetypes.RFC3339Type{

              },
            },
            "export_file_format": schema.StringAttribute{
              Description: `Available values: "CSV", "JSON".`,
              Computed: true,
              Validators: []validator.String{
              stringvalidator.OneOfCaseInsensitive("CSV", "JSON"),
              },
            },
            "last_modified_by": schema.StringAttribute{
              Description: "The id of the user who last modified this Data Export Schedule.",
              Computed: true,
            },
            "name": schema.StringAttribute{
              Description: "The name of the Data Export Schedule.",
              Computed: true,
            },
            "period": schema.Int64Attribute{
              Description: "Defines the Schedule frequency for the Data Export to run in Hours or Days. Used in conjunction with the `scheduleType` parameter.",
              Computed: true,
            },
            "schedule_type": schema.StringAttribute{
              Description: `Available values: "HOURLY", "DAILY", "AD_HOC".`,
              Computed: true,
              Validators: []validator.String{
              stringvalidator.OneOfCaseInsensitive(
                "HOURLY",
                "DAILY",
                "AD_HOC",
              ),
              },
            },
            "source_type": schema.StringAttribute{
              Description: `Available values: "USAGE", "OPERATIONAL".`,
              Computed: true,
              Validators: []validator.String{
              stringvalidator.OneOfCaseInsensitive("USAGE", "OPERATIONAL"),
              },
            },
          },
        },
      },
    },
  }
}

func (d *DataExportSchedulesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
  resp.Schema = ListDataSourceSchema(ctx)
}

func (d *DataExportSchedulesDataSource) ConfigValidators(_ context.Context) ([]datasource.ConfigValidator) {
  return []datasource.ConfigValidator{
  }
}
