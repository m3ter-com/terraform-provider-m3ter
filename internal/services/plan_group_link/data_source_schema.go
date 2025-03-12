// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package plan_group_link

import (
  "context"

  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework/datasource"
  "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = (*PlanGroupLinkDataSource)(nil)

func DataSourceSchema(ctx context.Context) (schema.Schema) {
  return schema.Schema{
    Attributes: map[string]schema.Attribute{
      "id": schema.StringAttribute{
        Required: true,
      },
      "org_id": schema.StringAttribute{
        Required: true,
      },
      "created_by": schema.StringAttribute{
        Description: "The id of the user who created this plan group link.",
        Computed: true,
      },
      "dt_created": schema.StringAttribute{
        Description: "The DateTime *(in ISO-8601 format)* when the plan group link was created.",
        Computed: true,
        CustomType: timetypes.RFC3339Type{

        },
      },
      "dt_last_modified": schema.StringAttribute{
        Description: "The DateTime *(in ISO-8601 format)* when the plan group link was last modified.",
        Computed: true,
        CustomType: timetypes.RFC3339Type{

        },
      },
      "last_modified_by": schema.StringAttribute{
        Description: "The id of the user who last modified this plan group link.",
        Computed: true,
      },
      "plan_group_id": schema.StringAttribute{
        Description: "ID of the linked PlanGroup",
        Computed: true,
      },
      "plan_id": schema.StringAttribute{
        Description: "ID of the linked Plan",
        Computed: true,
      },
      "version": schema.Int64Attribute{
        Description: "The version number:\n- **Create:** On initial Create to insert a new entity, the version is set at 1 in the response.\n- **Update:** On successful Update, the version is incremented by 1 in the response.",
        Computed: true,
      },
    },
  }
}

func (d *PlanGroupLinkDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
  resp.Schema = DataSourceSchema(ctx)
}

func (d *PlanGroupLinkDataSource) ConfigValidators(_ context.Context) ([]datasource.ConfigValidator) {
  return []datasource.ConfigValidator{
  }
}
