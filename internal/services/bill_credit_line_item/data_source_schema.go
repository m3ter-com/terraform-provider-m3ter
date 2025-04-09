// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package bill_credit_line_item

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = (*BillCreditLineItemDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"bill_id": schema.StringAttribute{
				Required: true,
			},
			"id": schema.StringAttribute{
				Required: true,
			},
			"org_id": schema.StringAttribute{
				Required: true,
			},
			"amount": schema.Float64Attribute{
				Description: "The amount for the line item.",
				Computed:    true,
			},
			"created_by": schema.StringAttribute{
				Description: "The id of the user who created this credit line item.",
				Computed:    true,
			},
			"credit_reason_id": schema.StringAttribute{
				Description: "The UUID of the credit reason for this credit line item.",
				Computed:    true,
			},
			"description": schema.StringAttribute{
				Description: "The description of the line item.",
				Computed:    true,
			},
			"dt_created": schema.StringAttribute{
				Description: "The DateTime when the credit line item was created *(in ISO-8601 format)*.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"dt_last_modified": schema.StringAttribute{
				Description: "The DateTime when the credit line item was last modified *(in ISO-8601 format)*.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"last_modified_by": schema.StringAttribute{
				Description: "The id of the user who last modified this credit line item.",
				Computed:    true,
			},
			"product_id": schema.StringAttribute{
				Description: "The UUID of the Product.",
				Computed:    true,
			},
			"referenced_bill_id": schema.StringAttribute{
				Description: "The UUID of the bill for the line item.",
				Computed:    true,
			},
			"referenced_line_item_id": schema.StringAttribute{
				Description: "The UUID of the line item.",
				Computed:    true,
			},
			"service_period_end_date": schema.StringAttribute{
				Description: "The service period end date in ISO-8601 format. *(exclusive of the ending date)*.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"service_period_start_date": schema.StringAttribute{
				Description: "The service period start date in ISO-8601 format. *(inclusive of the starting date)*.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"version": schema.Int64Attribute{
				Description: "The version number:\n- **Create:** On initial Create to insert a new entity, the version is set at 1 in the response.\n- **Update:** On successful Update, the version is incremented by 1 in the response.",
				Computed:    true,
			},
		},
	}
}

func (d *BillCreditLineItemDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *BillCreditLineItemDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
