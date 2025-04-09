// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package credit_reason

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

var _ datasource.DataSourceWithConfigValidators = (*CreditReasonsDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"org_id": schema.StringAttribute{
				Optional: true,
			},
			"archived": schema.BoolAttribute{
				Description: "TRUE / FALSE archived flag to filter the list. CreditReasons can be archived once they are obsolete.\n\n* TRUE includes archived CreditReasons.\n* FALSE excludes CreditReasons that are archived.",
				Optional:    true,
			},
			"codes": schema.ListAttribute{
				Description: "List of Credit Reason short codes to retrieve.",
				Optional:    true,
				ElementType: types.StringType,
			},
			"ids": schema.ListAttribute{
				Description: "List of Credit Reason IDs to retrieve.",
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
				CustomType:  customfield.NewNestedObjectListType[CreditReasonsItemsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The UUID of the entity.",
							Computed:    true,
						},
						"version": schema.Int64Attribute{
							Description: "The version number:\n- **Create:** On initial Create to insert a new entity, the version is set at 1 in the response.\n- **Update:** On successful Update, the version is incremented by 1 in the response.",
							Computed:    true,
						},
						"archived": schema.BoolAttribute{
							Description: "TRUE / FALSE flag indicating whether the data entity is archived. An entity can be archived if it is obsolete.",
							Computed:    true,
						},
						"code": schema.StringAttribute{
							Description: "The short code of the data entity.",
							Computed:    true,
						},
						"created_by": schema.StringAttribute{
							Description: "The id of the user who created this credit reason.",
							Computed:    true,
						},
						"dt_created": schema.StringAttribute{
							Description: "The DateTime when the credit reason was created *(in ISO-8601 format)*.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"dt_last_modified": schema.StringAttribute{
							Description: "The DateTime when the credit reason was last modified *(in ISO-8601 format)*.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"last_modified_by": schema.StringAttribute{
							Description: "The id of the user who last modified this credit reason.",
							Computed:    true,
						},
						"name": schema.StringAttribute{
							Description: "The name of the data entity.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func (d *CreditReasonsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *CreditReasonsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
