// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package integration_configuration

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = (*IntegrationConfigurationResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "The UUID of the entity.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"org_id": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"destination": schema.StringAttribute{
				Description: "Denotes the integration destination. This field identifies the target platform or service for the integration.",
				Required:    true,
			},
			"destination_id": schema.StringAttribute{
				Description: "The unique identifier (UUID) for the integration destination.",
				Required:    true,
			},
			"entity_id": schema.StringAttribute{
				Description: "The unique identifier (UUID) of the entity. This field is used to specify which entity's integration configuration you're updating.",
				Required:    true,
			},
			"entity_type": schema.StringAttribute{
				Description: "Specifies the type of entity for which the integration configuration is being updated. Must be a valid alphanumeric string.",
				Required:    true,
			},
			"integration_credentials_id": schema.StringAttribute{
				Required: true,
			},
			"name": schema.StringAttribute{
				Required: true,
			},
			"config_data": schema.MapAttribute{
				Description: "A flexible object to include any additional configuration data specific to the integration.",
				Required:    true,
				ElementType: jsontypes.NormalizedType{},
			},
			"credentials": schema.SingleNestedAttribute{
				Description: "Base model for defining integration credentials across different types of integrations.",
				Required:    true,
				Attributes: map[string]schema.Attribute{
					"type": schema.StringAttribute{
						Description: "Specifies the type of authorization required for the integration.\nAvailable values: \"HTTP_BASIC\", \"OAUTH_CLIENT_CREDENTIALS\", \"M3TER_SIGNED_REQUEST\", \"AWS_INTEGRATION\", \"PADDLE_AUTH\", \"NETSUITE_AUTH\", \"CHARGEBEE_AUTH\", \"M3TER_SERVICE_USER\", \"STRIPE_SIGNED_REQUEST\".",
						Required:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"HTTP_BASIC",
								"OAUTH_CLIENT_CREDENTIALS",
								"M3TER_SIGNED_REQUEST",
								"AWS_INTEGRATION",
								"PADDLE_AUTH",
								"NETSUITE_AUTH",
								"CHARGEBEE_AUTH",
								"M3TER_SERVICE_USER",
								"STRIPE_SIGNED_REQUEST",
							),
						},
					},
					"destination": schema.StringAttribute{
						Description: `Available values: "WEBHOOK", "NETSUITE", "STRIPE", "STRIPE_TEST", "AWS", "PADDLE", "PADDLE_SANDBOX", "SALESFORCE", "XERO", "CHARGEBEE", "QUICKBOOKS", "QUICKBOOKS_SANDBOX", "M3TER".`,
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"WEBHOOK",
								"NETSUITE",
								"STRIPE",
								"STRIPE_TEST",
								"AWS",
								"PADDLE",
								"PADDLE_SANDBOX",
								"SALESFORCE",
								"XERO",
								"CHARGEBEE",
								"QUICKBOOKS",
								"QUICKBOOKS_SANDBOX",
								"M3TER",
							),
						},
					},
					"empty": schema.BoolAttribute{
						Description: "A flag to indicate whether the credentials are empty. \n\n* TRUE - empty credentials.\n* FALSE - credential details required.",
						Optional:    true,
					},
					"name": schema.StringAttribute{
						Description: "The name of the credentials",
						Optional:    true,
					},
					"version": schema.Int64Attribute{
						Description: "The version number of the entity:\n- **Create entity:** Not valid for initial insertion of new entity - *do not use for Create*. On initial Create, version is set at 1 and listed in the response.\n- **Update Entity:**  On Update, version is required and must match the existing version because a check is performed to ensure sequential versioning is preserved. Version is incremented by 1 and listed in the response.",
						Optional:    true,
					},
				},
			},
			"version": schema.Int64Attribute{
				Description: "The version number of the entity:\n- **Create entity:** Not valid for initial insertion of new entity - *do not use for Create*. On initial Create, version is set at 1 and listed in the response.\n- **Update Entity:**  On Update, version is required and must match the existing version because a check is performed to ensure sequential versioning is preserved. Version is incremented by 1 and listed in the response.",
				Optional:    true,
			},
			"authorized": schema.BoolAttribute{
				Description: "A flag indicating whether the integration configuration is authorized. \n\n* TRUE - authorized.\n* FALSE - not authorized.",
				Computed:    true,
			},
			"created_by": schema.StringAttribute{
				Description: "The ID of the user who created this item.",
				Computed:    true,
			},
			"dt_completed": schema.StringAttribute{
				Description: "The date and time the integration was completed *(in ISO-8601 format)*.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"dt_created": schema.StringAttribute{
				Description: "The DateTime when this item was created *(in ISO-8601 format)*.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"dt_last_modified": schema.StringAttribute{
				Description: "The DateTime when this item was last modified *(in ISO-8601 format)*.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"dt_started": schema.StringAttribute{
				Description: "The date and time the integration was started *(in ISO-8601 format)*.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"enabled": schema.BoolAttribute{
				Description: "A flag indicating whether the integration configuration is currently enabled or disabled.\n\n* TRUE - enabled.\n* FALSE - disabled.",
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
			"last_modified_by": schema.StringAttribute{
				Description: "The ID of the user who last modified this item.",
				Computed:    true,
			},
			"status": schema.StringAttribute{
				Description: `Available values: "WAITING", "STARTED", "COMPLETE", "ERROR", "AWAITING_RETRY", "AUTH_FAILED", "ACCOUNTING_PERIOD_CLOSED", "INVOICE_ALREADY_PAID", "DISABLED".`,
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
					),
				},
			},
			"trigger_type": schema.StringAttribute{
				Description: "Specifies the type of trigger for the integration.\nAvailable values: \"EVENT\", \"SCHEDULE\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("EVENT", "SCHEDULE"),
				},
			},
			"url": schema.StringAttribute{
				Description: "The URL of the entity in the destination system if available.",
				Computed:    true,
			},
		},
	}
}

func (r *IntegrationConfigurationResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *IntegrationConfigurationResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
