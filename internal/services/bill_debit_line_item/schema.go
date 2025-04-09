// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package bill_debit_line_item

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = (*BillDebitLineItemResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "The UUID of the entity.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"bill_id": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"org_id": schema.StringAttribute{
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"amount": schema.Float64Attribute{
				Description: "The amount for the line item.",
				Required:    true,
				Validators: []validator.Float64{
					float64validator.AtLeast(0),
				},
			},
			"description": schema.StringAttribute{
				Description: "The description for the line item.",
				Required:    true,
			},
			"product_id": schema.StringAttribute{
				Description: "The UUID of the Product.",
				Required:    true,
			},
			"referenced_bill_id": schema.StringAttribute{
				Description: "The UUID of the bill for the line item.",
				Required:    true,
			},
			"referenced_line_item_id": schema.StringAttribute{
				Description: "The UUID of the line item.",
				Required:    true,
			},
			"service_period_end_date": schema.StringAttribute{
				Description: "The service period end date in ISO-8601 format.*(exclusive of the ending date)*.",
				Required:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"service_period_start_date": schema.StringAttribute{
				Description: "The service period start date in ISO-8601 format. *(inclusive of the starting date)*.",
				Required:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"debit_reason_id": schema.StringAttribute{
				Description: "The ID of the Debit Reason given for this debit line item.",
				Optional:    true,
			},
			"line_item_type": schema.StringAttribute{
				Description: `Available values: "STANDING_CHARGE", "USAGE", "COUNTER_RUNNING_TOTAL_CHARGE", "COUNTER_ADJUSTMENT_DEBIT", "COUNTER_ADJUSTMENT_CREDIT", "USAGE_CREDIT", "MINIMUM_SPEND", "MINIMUM_SPEND_REFUND", "CREDIT_DEDUCTION", "MANUAL_ADJUSTMENT", "CREDIT_MEMO", "DEBIT_MEMO", "COMMITMENT_CONSUMED", "COMMITMENT_FEE", "OVERAGE_SURCHARGE", "OVERAGE_USAGE", "BALANCE_CONSUMED", "BALANCE_FEE".`,
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"STANDING_CHARGE",
						"USAGE",
						"COUNTER_RUNNING_TOTAL_CHARGE",
						"COUNTER_ADJUSTMENT_DEBIT",
						"COUNTER_ADJUSTMENT_CREDIT",
						"USAGE_CREDIT",
						"MINIMUM_SPEND",
						"MINIMUM_SPEND_REFUND",
						"CREDIT_DEDUCTION",
						"MANUAL_ADJUSTMENT",
						"CREDIT_MEMO",
						"DEBIT_MEMO",
						"COMMITMENT_CONSUMED",
						"COMMITMENT_FEE",
						"OVERAGE_SURCHARGE",
						"OVERAGE_USAGE",
						"BALANCE_CONSUMED",
						"BALANCE_FEE",
					),
				},
			},
			"reason_id": schema.StringAttribute{
				Description: "The UUID of the line item reason.",
				Optional:    true,
			},
			"version": schema.Int64Attribute{
				Description: "The version number of the entity:\n- **Create entity:** Not valid for initial insertion of new entity - *do not use for Create*. On initial Create, version is set at 1 and listed in the response.\n- **Update Entity:**  On Update, version is required and must match the existing version because a check is performed to ensure sequential versioning is preserved. Version is incremented by 1 and listed in the response.",
				Optional:    true,
			},
			"created_by": schema.StringAttribute{
				Description: "The id of the user who created this debit line item.",
				Computed:    true,
			},
			"dt_created": schema.StringAttribute{
				Description: "The DateTime when the debit line item was created *(in ISO-8601 format)*.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"dt_last_modified": schema.StringAttribute{
				Description: "The DateTime when the debit line item was last modified *(in ISO-8601 format)*.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"last_modified_by": schema.StringAttribute{
				Description: "The id of the user who last modified this debit line item.",
				Computed:    true,
			},
		},
	}
}

func (r *BillDebitLineItemResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *BillDebitLineItemResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
