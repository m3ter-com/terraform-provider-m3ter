// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package scheduled_event_configuration

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

var _ datasource.DataSourceWithConfigValidators = (*ScheduledEventConfigurationsDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"org_id": schema.StringAttribute{
				Required: true,
			},
			"ids": schema.ListAttribute{
				Description: "list of UUIDs to retrieve",
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
				CustomType:  customfield.NewNestedObjectListType[ScheduledEventConfigurationsItemsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The UUID of the entity.",
							Computed:    true,
						},
						"entity": schema.StringAttribute{
							Description: "The referenced configuration or billing entity for which the desired scheduled Event will trigger.",
							Computed:    true,
						},
						"field": schema.StringAttribute{
							Description: "A DateTime field for which the desired scheduled Event will trigger - this must be a DateTime field on the referenced billing or configuration entity.",
							Computed:    true,
						},
						"name": schema.StringAttribute{
							Description: "The name of the custom Scheduled Event Configuration.",
							Computed:    true,
						},
						"offset": schema.Int64Attribute{
							Description: "The offset in days from the specified DateTime field on the referenced entity when the scheduled Event will trigger.",
							Computed:    true,
						},
						"version": schema.Int64Attribute{
							Description: "The version number:\n- **Create:** On initial Create to insert a new entity, the version is set at 1 in the response.\n- **Update:** On successful Update, the version is incremented by 1 in the response.",
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
					},
				},
			},
		},
	}
}

func (d *ScheduledEventConfigurationsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *ScheduledEventConfigurationsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
