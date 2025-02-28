// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package external_mapping

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

var _ datasource.DataSourceWithConfigValidators = (*ExternalMappingsDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"org_id": schema.StringAttribute{
				Required: true,
			},
			"external_system_id": schema.StringAttribute{
				Description: "The name of the external system to use as a filter.\n\nFor example, if you want to list only those external mappings created for your Organization for the Salesforce external system, use:\n\n`?externalSystemId=Salesforce`",
				Optional:    true,
			},
			"integration_config_id": schema.StringAttribute{
				Description: "ID of the integration config",
				Optional:    true,
			},
			"m3ter_ids": schema.ListAttribute{
				Description: "IDs for m3ter entities",
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
				CustomType:  customfield.NewNestedObjectListType[ExternalMappingsItemsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The UUID of the entity.",
							Computed:    true,
						},
						"external_id": schema.StringAttribute{
							Description: "The unique identifier (UUID) of the entity in the external system.",
							Computed:    true,
						},
						"external_system": schema.StringAttribute{
							Description: "The name of the external system where the entity you are mapping resides.",
							Computed:    true,
						},
						"external_table": schema.StringAttribute{
							Description: "The name of the table in the external system where the entity resides.",
							Computed:    true,
						},
						"m3ter_entity": schema.StringAttribute{
							Description: `The name of the m3ter entity that is part of the External Mapping. For example, this could be "Account".`,
							Computed:    true,
						},
						"m3ter_id": schema.StringAttribute{
							Description: "The unique identifier (UUID) of the m3ter entity.",
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
						"integration_config_id": schema.StringAttribute{
							Description: "UUID of the configuration this mapping is for",
							Computed:    true,
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

func (d *ExternalMappingsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *ExternalMappingsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
