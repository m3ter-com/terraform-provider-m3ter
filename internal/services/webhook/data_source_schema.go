// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package webhook

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

var _ datasource.DataSourceWithConfigValidators = (*WebhookDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"org_id": schema.StringAttribute{
				Optional:           true,
				DeprecationMessage: "the org id should be set at the client level instead",
			},
			"active": schema.BoolAttribute{
				Computed: true,
			},
			"code": schema.StringAttribute{
				Computed: true,
			},
			"description": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Computed: true,
			},
			"url": schema.StringAttribute{
				Description: "The URL to which webhook requests are sent.",
				Computed:    true,
			},
			"version": schema.Int64Attribute{
				Description: "The version number:\n- **Create:** On initial Create to insert a new entity, the version is set at 1 in the response.\n- **Update:** On successful Update, the version is incremented by 1 in the response.",
				Computed:    true,
			},
			"credentials": schema.SingleNestedAttribute{
				Description: "Response representing a set of credentials used for signing m3ter requests.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[WebhookCredentialsDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Description: "The UUID of the entity.",
						Computed:    true,
					},
					"destination": schema.StringAttribute{
						Description: "the system the integration is for",
						Computed:    true,
					},
					"type": schema.StringAttribute{
						Description: "the type of credentials",
						Computed:    true,
					},
					"api_key": schema.StringAttribute{
						Description: "The API key provided by m3ter. This key is part of the credential set required for signing requests and authenticating with m3ter services.",
						Computed:    true,
					},
					"destination_id": schema.StringAttribute{
						Description: "the destinationId the integration is for",
						Computed:    true,
					},
					"name": schema.StringAttribute{
						Description: "the name of the credentials",
						Computed:    true,
					},
					"secret": schema.StringAttribute{
						Description: "The secret associated with the API key. This secret is used in conjunction with the API key to generate a signature for secure authentication.",
						Computed:    true,
					},
					"version": schema.Int64Attribute{
						Description: "The version number:\n- **Create:** On initial Create to insert a new entity, the version is set at 1 in the response.\n- **Update:** On successful Update, the version is incremented by 1 in the response.",
						Computed:    true,
					},
				},
			},
			"find_one_by": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"ids": schema.ListAttribute{
						Optional:    true,
						ElementType: types.StringType,
					},
				},
			},
		},
	}
}

func (d *WebhookDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *WebhookDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("id"), path.MatchRoot("find_one_by")),
	}
}
