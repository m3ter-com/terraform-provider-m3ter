// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user

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
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

var _ resource.ResourceWithConfigValidators = (*UserResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"org_id": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"dt_end_access": schema.StringAttribute{
				Description: "The date and time *(in ISO 8601 format)* when the user's access will end. Use this to set or update the expiration of the user's access.",
				Optional:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"version": schema.Int64Attribute{
				Description: "The version number of the entity:\n\n- **Newly created entity:** On initial Create, version is set at 1 and listed in the response.\n- **Update Entity:** On Update, version is required and must match the existing version because a check is performed to ensure sequential versioning is preserved. Version is incremented by 1 and listed in the response.",
				Optional:    true,
			},
			"permission_policy": schema.ListNestedAttribute{
				Description: "An array of permission statements for the user. Each permission statement defines a specific permission for the user.\n\nSee [Understanding, Creating, and Managing Permission Policies](https://www.m3ter.com/docs/guides/organization-and-access-management/creating-and-managing-permissions) for more information.",
				Computed:    true,
				Optional:    true,
				CustomType:  customfield.NewNestedObjectListType[UserPermissionPolicyModel](ctx),
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
							Description: "Specifies whether or not the user is allowed to perform the action on the resource.\n\n**NOTE:** Use lower case, for example: `\"allow\"`. If you use upper case, you'll receive an error.\navailable values: \"ALLOW\", \"DENY\"",
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
			"contact_number": schema.StringAttribute{
				Description: "The user's contact telephone number.",
				Computed:    true,
			},
			"created_by": schema.StringAttribute{
				Description: "The user who created this user.",
				Computed:    true,
			},
			"dt_created": schema.StringAttribute{
				Description: "The date and time *(in ISO-8601 format)* when the user was created.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"dt_last_modified": schema.StringAttribute{
				Description: "The date and time *(in ISO-8601 format)* when the user was last modified.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"email": schema.StringAttribute{
				Description: "The email address for this user.",
				Computed:    true,
			},
			"first_accepted_terms_and_conditions": schema.StringAttribute{
				Description: "The date and time *(in ISO 8601 format)* when this user first accepted the the m3ter terms and conditions.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"first_name": schema.StringAttribute{
				Description: "The first name of the user.",
				Computed:    true,
			},
			"last_accepted_terms_and_conditions": schema.StringAttribute{
				Description: "The date and time *(in ISO 8601 format)* when this user last accepted the the m3ter terms and conditions.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"last_modified_by": schema.StringAttribute{
				Description: "The unique identifier (UUID) of the user who last modified this user record.",
				Computed:    true,
			},
			"last_name": schema.StringAttribute{
				Description: "The surname of the user.",
				Computed:    true,
			},
			"support_user": schema.BoolAttribute{
				Description: "Indicates whether this is a m3ter Support user.",
				Computed:    true,
			},
			"organizations": schema.ListAttribute{
				Description: "An array listing the Organizations where this user has access.",
				Computed:    true,
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
			},
		},
	}
}

func (r *UserResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *UserResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
