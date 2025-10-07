// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package external_mapping

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*ExternalMappingDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"org_id": schema.StringAttribute{
				Required:           true,
				DeprecationMessage: "the org id should be set at the client level instead",
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
			"integration_config_id": schema.StringAttribute{
				Description: "UUID of the configuration this mapping is for",
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
			"find_one_by": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
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
				},
			},
		},
	}
}

func (d *ExternalMappingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *ExternalMappingDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("id"), path.MatchRoot("find_one_by")),
	}
}
