// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package integration_configuration

import (
  "context"

  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
  "github.com/hashicorp/terraform-plugin-framework/datasource"
  "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
  "github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*IntegrationConfigurationDataSource)(nil)

func DataSourceSchema(ctx context.Context) (schema.Schema) {
  return schema.Schema{
    Attributes: map[string]schema.Attribute{
      "id": schema.StringAttribute{
        Required: true,
      },
      "org_id": schema.StringAttribute{
        Required: true,
      },
      "created_by": schema.StringAttribute{
        Description: "The ID of the user who created this item.",
        Computed: true,
      },
      "destination": schema.StringAttribute{
        Description: "The destination system for the integration.",
        Computed: true,
      },
      "dt_completed": schema.StringAttribute{
        Description: "The date and time the integration was completed *(in ISO-8601 format)*.",
        Computed: true,
        CustomType: timetypes.RFC3339Type{

        },
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
      "dt_started": schema.StringAttribute{
        Description: "The date and time the integration was started *(in ISO-8601 format)*.",
        Computed: true,
        CustomType: timetypes.RFC3339Type{

        },
      },
      "entity_id": schema.StringAttribute{
        Description: "The unique identifier (UUID) of the entity the integration is for.",
        Computed: true,
      },
      "entity_type": schema.StringAttribute{
        Description: "The type of entity the integration is for *(e.g. Bill)*.",
        Computed: true,
      },
      "error": schema.StringAttribute{
        Description: "Describes any errors encountered during the integration run.",
        Computed: true,
      },
      "external_id": schema.StringAttribute{
        Description: "The external ID in the destination system if available.",
        Computed: true,
      },
      "last_modified_by": schema.StringAttribute{
        Description: "The ID of the user who last modified this item.",
        Computed: true,
      },
      "status": schema.StringAttribute{
        Description: `Available values: "WAITING", "STARTED", "COMPLETE", "ERROR", "AWAITING_RETRY", "AUTH_FAILED", "ACCOUNTING_PERIOD_CLOSED", "INVOICE_ALREADY_PAID", "DISABLED".`,
        Computed: true,
        Validators: []validator.String{
        stringvalidator.OneOfCaseInsensitive(
          "WAITING",
          "STARTED",
          "COMPLETE",
          "ERROR",
          "AWAITING_RETRY",
          "AUTH_FAILED",
          "ACCOUNTING_PERIOD_CLOSED",
          "INVOICE_ALREADY_PAID",
          "DISABLED",
        ),
        },
      },
      "url": schema.StringAttribute{
        Description: "The URL of the entity in the destination system if available.",
        Computed: true,
      },
      "version": schema.Int64Attribute{
        Description: "The version number:\n- **Create:** On initial Create to insert a new entity, the version is set at 1 in the response.\n- **Update:** On successful Update, the version is incremented by 1 in the response.",
        Computed: true,
      },
    },
  }
}

func (d *IntegrationConfigurationDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
  resp.Schema = DataSourceSchema(ctx)
}

func (d *IntegrationConfigurationDataSource) ConfigValidators(_ context.Context) ([]datasource.ConfigValidator) {
  return []datasource.ConfigValidator{
  }
}
