// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package integration_configuration

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*IntegrationConfigurationDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"org_id": schema.StringAttribute{
				Required:           true,
				DeprecationMessage: "the org id should be set at the client level instead",
			},
			"destination": schema.StringAttribute{
				Description: "The destination system for the integration run.",
				Computed:    true,
			},
			"dt_completed": schema.StringAttribute{
				Description: "The date and time the integration was completed. *(in ISO-8601 format)*.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"dt_started": schema.StringAttribute{
				Description: "The date and time the integration run was started *(in ISO-8601 format)*.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"entity_id": schema.StringAttribute{
				Description: "The unique identifier (UUID) of the entity the integration run is for.",
				Computed:    true,
			},
			"entity_type": schema.StringAttribute{
				Description: "The type of entity the integration run is for. Two options:\n* Bill\n* Notification",
				Computed:    true,
			},
			"error": schema.StringAttribute{
				Description: "Describes any errors encountered during the integration run.",
				Computed:    true,
			},
			"external_id": schema.StringAttribute{
				Description: "The external ID in the destination system if available.",
				Computed:    true,
			},
			"status": schema.StringAttribute{
				Description: `Available values: "WAITING", "STARTED", "COMPLETE", "ERROR", "AWAITING_RETRY", "AUTH_FAILED", "ACCOUNTING_PERIOD_CLOSED", "INVOICE_ALREADY_PAID", "DISABLED", "TIMEOUT_LIMIT_EXCEEDED", "RATE_LIMIT_RETRY".`,
				Computed:    true,
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
						"TIMEOUT_LIMIT_EXCEEDED",
						"RATE_LIMIT_RETRY",
					),
				},
			},
			"url": schema.StringAttribute{
				Description: "The URL of the entity in the destination system if available.",
				Computed:    true,
			},
			"version": schema.Int64Attribute{
				Description: "The version number:\n- **Create:** On initial Create to insert a new entity, the version is set at 1 in the response.\n- **Update:** On successful Update, the version is incremented by 1 in the response.",
				Computed:    true,
			},
			"find_one_by": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"destination_id": schema.StringAttribute{
						Description: "optional filter for a specific destination",
						Optional:    true,
					},
				},
			},
		},
	}
}

func (d *IntegrationConfigurationDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *IntegrationConfigurationDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("id"), path.MatchRoot("find_one_by")),
	}
}
