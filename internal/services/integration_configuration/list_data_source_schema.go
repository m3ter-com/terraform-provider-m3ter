// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package integration_configuration

import (
  "context"

  "github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
  "github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
  "github.com/hashicorp/terraform-plugin-framework/datasource"
  "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
  "github.com/hashicorp/terraform-plugin-framework/schema/validator"
  "github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

var _ datasource.DataSourceWithConfigValidators = (*IntegrationConfigurationsDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) (schema.Schema) {
  return schema.Schema{
    Attributes: map[string]schema.Attribute{
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
        CustomType: customfield.NewNestedObjectListType[IntegrationConfigurationsItemsDataSourceModel](ctx),
        NestedObject: schema.NestedAttributeObject{
          Attributes: map[string]schema.Attribute{
            "id": schema.StringAttribute{
              Description: "The UUID of the entity.",
              Computed: true,
            },
            "destination": schema.StringAttribute{
              Description: "The type of destination *(e.g. Netsuite, webhooks)*.",
              Computed: true,
            },
            "entity_type": schema.StringAttribute{
              Description: "The type of entity the integration is for *(e.g. Bill)*.",
              Computed: true,
            },
            "version": schema.Int64Attribute{
              Description: "The version number:\n- **Create:** On initial Create to insert a new entity, the version is set at 1 in the response.\n- **Update:** On successful Update, the version is incremented by 1 in the response.",
              Computed: true,
            },
            "authorized": schema.BoolAttribute{
              Description: "A flag indicating whether the integration configuration is authorized. \n\n* TRUE - authorized.\n* FALSE - not authorized.",
              Computed: true,
            },
            "config_data": schema.MapAttribute{
              Description: "Configuration data for the integration",
              Computed: true,
              CustomType: customfield.NewMapType[jsontypes.Normalized](ctx),
              ElementType: jsontypes.NormalizedType{

              },
            },
            "created_by": schema.StringAttribute{
              Description: "The ID of the user who created this item.",
              Computed: true,
            },
            "destination_id": schema.StringAttribute{
              Description: "The unique identifier (UUID) of the entity the integration is for.",
              Computed: true,
            },
            "dt_created": schema.StringAttribute{
              Description: "The DateTime when this item was created *(in ISO-8601 format)*.",
              Computed: true,
              CustomType: timetypes.RFC3339Type{

              },
            },
            "dt_last_modified": schema.StringAttribute{
              Description: "The DateTime when this item was last modified *(in ISO-8601 format)*.",
              Computed: true,
              CustomType: timetypes.RFC3339Type{

              },
            },
            "enabled": schema.BoolAttribute{
              Description: "A flag indicating whether the integration configuration is currently enabled or disabled.\n\n* TRUE - enabled.\n* FALSE - disabled.",
              Computed: true,
            },
            "entity_id": schema.StringAttribute{
              Description: "The unique identifier (UUID) of the entity this integration is for *(e.g. the ID of a notification configuration. Optional.)*",
              Computed: true,
            },
            "integration_credentials_id": schema.StringAttribute{
              Description: "UUID of the credentials to use for this integration",
              Computed: true,
            },
            "last_modified_by": schema.StringAttribute{
              Description: "The ID of the user who last modified this item.",
              Computed: true,
            },
            "name": schema.StringAttribute{
              Description: "The name of the configuration",
              Computed: true,
            },
            "trigger_type": schema.StringAttribute{
              Description: "Specifies the type of trigger for the integration.\nAvailable values: \"EVENT\", \"SCHEDULE\".",
              Computed: true,
              Validators: []validator.String{
              stringvalidator.OneOfCaseInsensitive("EVENT", "SCHEDULE"),
              },
            },
          },
        },
      },
    },
  }
}

func (d *IntegrationConfigurationsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
  resp.Schema = ListDataSourceSchema(ctx)
}

func (d *IntegrationConfigurationsDataSource) ConfigValidators(_ context.Context) ([]datasource.ConfigValidator) {
  return []datasource.ConfigValidator{
  }
}
