// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package resource_group

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

var _ datasource.DataSourceWithConfigValidators = (*ResourceGroupsDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Description: "Endpoints for ResourceGroup related operations such as creation, update, list and delete.\n\nResourceGroups are used in the context of Permission Policies, which controls what a User who has been given access to your Organization can and cannot do. For example, you might want to create a Permissions Policy that denies Users the ability to retrieve Meters. \n\nResources are defined as m3ter Resource Identifiers *(MRIs)* in the format:\n\n```\nservice:resource-type/item-type/id\n```\n\nWhere:\n\n* service is a distinct part of the overall m3ter system, and which forms a natural functional grouping, such as \"config\" or \"billing\".\n\n* resource-type is the resource type item accessed - for example: \"Plan\", \"Meter\", \"Bill\"\n\n* item-type is one of:\n\n\t* \"item\" - to specify an individual item.\n\n\t* \"group\" - to specify a resource group.\n\n* id is the resource group id or the resource item id\n\nResources can be assigned to one or more ResourceGroups. For example, a Plan can be assigned to Plan ResourceGroups, a Meter can be assigned to Meter ResourceGroups, and so on. This is useful for cases where you want to create Permission Policies which allow or deny access to a specific subset of resources. For example, grant a user access to only some of the Plans in your Organization.\n\nThis concept of grouping resources applies to every resource in m3ter, including ResourceGroups themselves. This allows you to nest ResourceGroups to support hierarchies of groups.\n\nSee [Understanding, Creating, and Managing Permission Policies](https://www.m3ter.com/docs/guides/managing-organization-and-users/creating-and-managing-permissions) in the m3ter documentation for more information.\n\n**Note: User Resource Groups** You can create a User Resource Group to group resources of type = `user`. You can then retrieve a list of the User Resource Groups a user belongs to. For more details, see the [Retrieve OrgUser Groups](https://www.m3ter.com/docs/api#tag/OrgUsers/operation/GetOrgUserGroups) call in the OrgUsers section.",
		Attributes: map[string]schema.Attribute{
			"type": schema.StringAttribute{
				Required: true,
			},
			"org_id": schema.StringAttribute{
				Optional:           true,
				DeprecationMessage: "the org id should be set at the client level instead",
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
				CustomType:  customfield.NewNestedObjectListType[ResourceGroupsItemsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The UUID of the entity.",
							Computed:    true,
						},
						"name": schema.StringAttribute{
							Description: "The name of the Resource Group.",
							Computed:    true,
						},
						"version": schema.Int64Attribute{
							Description: "The version number:\n- **Create:** On initial Create to insert a new entity, the version is set at 1 in the response.\n- **Update:** On successful Update, the version is incremented by 1 in the response.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func (d *ResourceGroupsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *ResourceGroupsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
