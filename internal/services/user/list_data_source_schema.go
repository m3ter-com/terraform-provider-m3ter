// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

var _ datasource.DataSourceWithConfigValidators = (*UsersDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"org_id": schema.StringAttribute{
				Required: true,
			},
			"ids": schema.ListAttribute{
				Description: "list of ids to retrieve",
				Optional:    true,
				ElementType: types.StringType,
			},
			"max_items": schema.Int64Attribute{
				Description: "Max items to fetch, default: 1000",
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
			},
			"items": schema.ListNestedAttribute{
				Description: "The items returned by the data source",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[UsersItemsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The unique identifier (UUID) of this user.",
							Computed:    true,
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
						"dt_end_access": schema.StringAttribute{
							Description: "The date and time *(in ISO 8601 format)* when the user's access will end. Used to set or update the date and time a user's access expires.",
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
						"organizations": schema.ListAttribute{
							Description: "An array listing the Organizations where this user has access.",
							Computed:    true,
							CustomType:  customfield.NewListType[types.String](ctx),
							ElementType: types.StringType,
						},
						"permission_policy": schema.ListNestedAttribute{
							Description: "An array of permission statements for the user. Each permission statement defines a specific permission for the user.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectListType[UsersPermissionPolicyDataSourceModel](ctx),
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"action": schema.ListAttribute{
										Description: "The actions available to users who are assigned the Permission Policy - what they can do or cannot do with respect to the specified resource.\n\n**NOTE:** Use lower case and a colon-separated format, for example, if you want to confer full CRUD, use:\n```\n\"config:create\",\n\"config:delete\",\n\"config:retrieve\",\n\"config:update\"\n```",
										Computed:    true,
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
										CustomType:  customfield.NewListType[types.String](ctx),
										ElementType: types.StringType,
									},
									"effect": schema.StringAttribute{
										Description: "Specifies whether or not the user is allowed to perform the action on the resource.\n\n**NOTE:** Use lower case, for example: `\"allow\"`. If you use upper case, you'll receive an error.\nAvailable values: \"ALLOW\", \"DENY\".",
										Computed:    true,
										Validators: []validator.String{
											stringvalidator.OneOfCaseInsensitive("ALLOW", "DENY"),
										},
									},
									"resource": schema.ListAttribute{
										Description: "See [Statements - Available Resources](https://www.m3ter.com/docs/guides/managing-organization-and-users/creating-and-managing-permissions#statements---available-resources) for a listing of available resources for Permission Policy statements.",
										Computed:    true,
										CustomType:  customfield.NewListType[types.String](ctx),
										ElementType: types.StringType,
									},
								},
							},
						},
						"support_user": schema.BoolAttribute{
							Description: "Indicates whether this is a m3ter Support user.",
							Computed:    true,
						},
						"version": schema.Int64Attribute{
							Description: "The version number:\n\n- **Create:** On initial Create to insert a new entity, the version is set at 1 in the response.\n- **Update:** On successful Update, the version is incremented by 1 in the response.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func (d *UsersDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *UsersDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
