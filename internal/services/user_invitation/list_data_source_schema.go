// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user_invitation

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

var _ datasource.DataSourceWithConfigValidators = (*UserInvitationsDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"org_id": schema.StringAttribute{
				Required: true,
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
				CustomType:  customfield.NewNestedObjectListType[UserInvitationsItemsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The UUID of the invitation. ",
							Computed:    true,
						},
						"accepted": schema.BoolAttribute{
							Description: "Boolean indicating whether the user has accepted the invitation.\n\n* TRUE - the invite has been accepted.\n* FALSE - the invite has not yet been accepted.",
							Computed:    true,
						},
						"dt_end_access": schema.StringAttribute{
							Description: "The date that access will end for the user *(in ISO-8601 format)*. If this is blank, there is no end date  meaning that the user has permanent access. ",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"dt_expiry": schema.StringAttribute{
							Description: "The date when the invite expires *(in ISO-8601 format)*. After this date the invited user can no longer accept the invite. By default, any invite is valid for 30 days from the date the invite is sent.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"email": schema.StringAttribute{
							Description: "The email address of the invitee. The invitation will be sent to this email address. ",
							Computed:    true,
						},
						"first_name": schema.StringAttribute{
							Description: "The first name of the invitee.",
							Computed:    true,
						},
						"inviting_principal_id": schema.StringAttribute{
							Description: "The UUID of the user who sent the invite. ",
							Computed:    true,
						},
						"last_name": schema.StringAttribute{
							Description: "The surname of the invitee.",
							Computed:    true,
						},
						"permission_policy_ids": schema.ListAttribute{
							Description: "The IDs of the permission policies the invited user has been assigned. This controls the access rights and privileges that this user will have when working in the m3ter Organization. ",
							Computed:    true,
							CustomType:  customfield.NewListType[types.String](ctx),
							ElementType: types.StringType,
						},
						"version": schema.Int64Attribute{
							Description: "The version number. Default value when newly created is one. ",
							Computed:    true,
						},
						"created_by": schema.StringAttribute{
							Description: "The UUID of the user who created the invitation.",
							Computed:    true,
						},
						"dt_created": schema.StringAttribute{
							Description: "The DateTime when the invitation was created *(in ISO-8601 format)*.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"dt_last_modified": schema.StringAttribute{
							Description: "The DateTime when the invitation was last modified *(in ISO-8601 format)*.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"last_modified_by": schema.StringAttribute{
							Description: "The UUID of the user who last modified the invitation.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func (d *UserInvitationsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *UserInvitationsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
