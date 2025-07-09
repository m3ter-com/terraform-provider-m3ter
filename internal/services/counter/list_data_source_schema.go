// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package counter

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

var _ datasource.DataSourceWithConfigValidators = (*CountersDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"org_id": schema.StringAttribute{
				Optional:           true,
				DeprecationMessage: "the org id should be set at the client level instead",
			},
			"codes": schema.ListAttribute{
				Description: "List of Counter codes to retrieve. These are unique short codes to identify each Counter.",
				Optional:    true,
				ElementType: types.StringType,
			},
			"ids": schema.ListAttribute{
				Description: "List of Counter IDs to retrieve.",
				Optional:    true,
				ElementType: types.StringType,
			},
			"product_id": schema.ListAttribute{
				Description: "List of Products UUIDs to retrieve Counters for.",
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
				CustomType:  customfield.NewNestedObjectListType[CountersItemsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The UUID of the entity.",
							Computed:    true,
						},
						"code": schema.StringAttribute{
							Description: "Code of the Counter. A unique short code to identify the Counter.",
							Computed:    true,
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
						"name": schema.StringAttribute{
							Description: "Name of the Counter.",
							Computed:    true,
						},
						"product_id": schema.StringAttribute{
							Description: "UUID of the product the Counter belongs to. *(Optional)* - if no `productId` is returned, the Counter is Global. A Global Counter can be used to price Plans or Plan Templates belonging to any Product.",
							Computed:    true,
						},
						"unit": schema.StringAttribute{
							Description: "Label for units shown on Bill line items, and indicating to customers what they are being charged for.",
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

func (d *CountersDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *CountersDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
