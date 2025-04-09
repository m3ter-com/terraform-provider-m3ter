// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package permission_policy

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*PermissionPolicyResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "The unique identifier (UUID) for this Permission Policy.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"org_id": schema.StringAttribute{
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"name": schema.StringAttribute{
				Required: true,
			},
			"permission_policy": schema.ListNestedAttribute{
				Required: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"action": schema.ListAttribute{
							Description: "The actions available to users who are assigned the Permission Policy - what they can do or cannot do with respect to the specified resource.\n\n**NOTE:** Use lower case and a colon-separated format, for example, if you want to confer full CRUD, use:\n```\n\"config:create\",\n\"config:delete\",\n\"config:retrieve\",\n\"config:update\"\n```",
							Required:    true,
							Validators: []validator.List{
								listvalidator.ValueStringsAre(
									stringvalidator.OneOfCaseInsensitive(
										"ALL",
										"CONFIG_CREATE",
										"CONFIG_RETRIEVE",
										"CONFIG_UPDATE",
										"CONFIG_DELETE",
										"CONFIG_EXPORT",
										"ANALYTICS_QUERY",
										"MEASUREMENTS_UPLOAD",
										"MEASUREMENTS_FILEUPLOAD",
										"MEASUREMENTS_RETRIEVE",
										"MEASUREMENTS_EXPORT",
										"FORECAST_RETRIEVE",
										"HEALTHSCORES_RETRIEVE",
										"ANOMALIES_RETRIEVE",
										"EXPORTS_DOWNLOAD",
									),
								),
							},
							ElementType: types.StringType,
						},
						"effect": schema.StringAttribute{
							Description: "Specifies whether or not the user is allowed to perform the action on the resource.\n\n**NOTE:** Use lower case, for example: `\"allow\"`. If you use upper case, you'll receive an error.\nAvailable values: \"ALLOW\", \"DENY\".",
							Required:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("ALLOW", "DENY"),
							},
						},
						"resource": schema.ListAttribute{
							Description: "See [Statements - Available Resources](https://www.m3ter.com/docs/guides/managing-organization-and-users/creating-and-managing-permissions#statements---available-resources) for a listing of available resources for Permission Policy statements.",
							Required:    true,
							ElementType: types.StringType,
						},
					},
				},
			},
			"version": schema.Int64Attribute{
				Description: "The version number of the entity:\n* **Create entity:** Not valid for initial insertion of new entity - do not use for Create. On initial Create, version is set at 1 and listed in the response.\n* **Update Entity:** On Update, version is required and must match the existing version because a check is performed to ensure sequential versioning is preserved. Version is incremented by 1 and listed in the response.",
				Optional:    true,
			},
			"created_by": schema.StringAttribute{
				Description: "The unique identifier (UUID) of the user who created this Permission Policy.",
				Computed:    true,
			},
			"dt_created": schema.StringAttribute{
				Description: "The date and time *(in ISO-8601 format)* when the Permission Policy was created.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"dt_last_modified": schema.StringAttribute{
				Description: "The date and time *(in ISO-8601 format)* when the Permission Policy was last modified.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"last_modified_by": schema.StringAttribute{
				Description: "The unique identifier (UUID) of the user who last modified this Permission Policy.",
				Computed:    true,
			},
			"managed_policy": schema.BoolAttribute{
				Description: "Indicates whether this is a system generated Managed Permission Policy.",
				Computed:    true,
			},
		},
	}
}

func (r *PermissionPolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *PermissionPolicyResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
