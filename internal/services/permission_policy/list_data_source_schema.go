// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package permission_policy

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

var _ datasource.DataSourceWithConfigValidators = (*PermissionPoliciesDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"org_id": schema.StringAttribute{
				Optional: true,
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
				CustomType:  customfield.NewNestedObjectListType[PermissionPoliciesItemsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The unique identifier (UUID) for this Permission Policy.",
							Computed:    true,
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
						"name": schema.StringAttribute{
							Description: "The name of the Permission Policy.",
							Computed:    true,
						},
						"permission_policy": schema.ListNestedAttribute{
							Description: "Array containing the Permission Policies information.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectListType[PermissionPoliciesPermissionPolicyDataSourceModel](ctx),
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
						"version": schema.Int64Attribute{
							Description: "The version number. Default value when newly created is one.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func (d *PermissionPoliciesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *PermissionPoliciesDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
