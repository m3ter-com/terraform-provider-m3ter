// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package webhook

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = (*WebhookResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "The UUID of the entity.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"org_id": schema.StringAttribute{
				Optional:           true,
				DeprecationMessage: "the org id should be set at the client level instead",
				PlanModifiers:      []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"description": schema.StringAttribute{
				Required: true,
			},
			"name": schema.StringAttribute{
				Required: true,
			},
			"url": schema.StringAttribute{
				Description: "The URL to which the webhook requests will be sent.",
				Required:    true,
			},
			"credentials": schema.SingleNestedAttribute{
				Description: "This schema defines the credentials required for m3ter request signing.",
				Required:    true,
				Attributes: map[string]schema.Attribute{
					"api_key": schema.StringAttribute{
						Description: "The API key provided by m3ter. This key is part of the credential set required for signing requests and authenticating with m3ter services.",
						Required:    true,
					},
					"secret": schema.StringAttribute{
						Description: "The secret associated with the API key. This secret is used in conjunction with the API key to generate a signature for secure authentication.",
						Required:    true,
					},
					"type": schema.StringAttribute{
						Description: "Specifies the authorization type. For this schema, it is exclusively set to M3TER_SIGNED_REQUEST.\nAvailable values: \"M3TER_SIGNED_REQUEST\".",
						Required:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("M3TER_SIGNED_REQUEST"),
						},
					},
					"empty": schema.BoolAttribute{
						Description: "A flag to indicate whether the credentials are empty. \n\n* TRUE - empty credentials.\n* FALSE - credential details required.",
						Optional:    true,
					},
					"version": schema.Int64Attribute{
						Description: "The version number of the entity:\n- **Create entity:** Not valid for initial insertion of new entity - *do not use for Create*. On initial Create, version is set at 1 and listed in the response.\n- **Update Entity:**  On Update, version is required and must match the existing version because a check is performed to ensure sequential versioning is preserved. Version is incremented by 1 and listed in the response.",
						Computed:    true,
					},
				},
			},
			"active": schema.BoolAttribute{
				Optional: true,
			},
			"code": schema.StringAttribute{
				Optional: true,
			},
			"created_by": schema.StringAttribute{
				Description: "The ID of the user who created this item.",
				Computed:    true,
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
			"last_modified_by": schema.StringAttribute{
				Description: "The ID of the user who last modified this item.",
				Computed:    true,
			},
			"version": schema.Int64Attribute{
				Description: "The version number:\n- **Create:** On initial Create to insert a new entity, the version is set at 1 in the response.\n- **Update:** On successful Update, the version is incremented by 1 in the response.",
				Computed:    true,
			},
		},
	}
}

func (r *WebhookResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *WebhookResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
