// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package webhook

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = (*WebhookDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Required: true,
			},
			"org_id": schema.StringAttribute{
				Required: true,
			},
			"active": schema.BoolAttribute{
				Computed: true,
			},
			"code": schema.StringAttribute{
				Computed: true,
			},
			"created_by": schema.StringAttribute{
				Description: "The ID of the user who created this item.",
				Computed:    true,
			},
			"description": schema.StringAttribute{
				Computed: true,
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
				Computed: true,
			},
			"url": schema.StringAttribute{
				Description: "The URL to which webhook requests are sent. ",
				Computed:    true,
			},
			"version": schema.Int64Attribute{
				Description: "The version number:\n- **Create:** On initial Create to insert a new entity, the version is set at 1 in the response.\n- **Update:** On successful Update, the version is incremented by 1 in the response.",
				Computed:    true,
			},
			"credentials": schema.StringAttribute{
				Description: "This schema defines the credentials required for m3ter request signing. ",
				Computed:    true,
				CustomType:  jsontypes.NormalizedType{},
			},
		},
	}
}

func (d *WebhookDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *WebhookDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
