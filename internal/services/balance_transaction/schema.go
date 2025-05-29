// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package balance_transaction

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = (*BalanceTransactionResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "The UUID of the entity.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"balance_id": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"org_id": schema.StringAttribute{
				Optional:           true,
				DeprecationMessage: "the org id should be set at the client level instead",
				PlanModifiers:      []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"amount": schema.Float64Attribute{
				Description:   "The financial value of the transaction.",
				Required:      true,
				PlanModifiers: []planmodifier.Float64{float64planmodifier.RequiresReplace()},
			},
			"applied_date": schema.StringAttribute{
				Description:   "The date *(in ISO 8601 format)* when the Balance transaction was applied.",
				Optional:      true,
				CustomType:    timetypes.RFC3339Type{},
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"currency_paid": schema.StringAttribute{
				Description:   "The currency code of the payment if it differs from the Balance currency. For example: USD, GBP or EUR.",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"description": schema.StringAttribute{
				Description:   "A brief description explaining the purpose and context of the transaction.",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"paid": schema.Float64Attribute{
				Description:   "The payment amount if the payment currency differs from the Balance currency.",
				Optional:      true,
				PlanModifiers: []planmodifier.Float64{float64planmodifier.RequiresReplace()},
			},
			"transaction_date": schema.StringAttribute{
				Description:   "The date *(in ISO 8601 format)* when the transaction occurred.",
				Optional:      true,
				CustomType:    timetypes.RFC3339Type{},
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"transaction_type_id": schema.StringAttribute{
				Description:   "The unique identifier (UUID) of the transaction type. This is obtained from the list of created Transaction Types within the Organization Configuration.",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"version": schema.Int64Attribute{
				Description:   "The version number of the entity:\n- **Create entity:** Not valid for initial insertion of new entity - *do not use for Create*. On initial Create, version is set at 1 and listed in the response.\n- **Update Entity:**  On Update, version is required and must match the existing version because a check is performed to ensure sequential versioning is preserved. Version is incremented by 1 and listed in the response.",
				Optional:      true,
				PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
			},
			"created_by": schema.StringAttribute{
				Description: "The unique identifier (UUID) for the user who created the balance transaction.",
				Computed:    true,
			},
			"dt_created": schema.StringAttribute{
				Description: "The date and time *(in ISO 8601 format)* when the balance transaction was first created.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"dt_last_modified": schema.StringAttribute{
				Description: "The date and time *(in ISO 8601 format)* when the balance transaction was last modified.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"entity_id": schema.StringAttribute{
				Description: "The unique identifier (UUID) for the entity associated with the Transaction, as specified by the `entityType`.",
				Computed:    true,
			},
			"entity_type": schema.StringAttribute{
				Description: "The type of entity associated with the Transaction - identifies who or what was responsible for the Transaction being added to the Balance - such as a **User**, a **Service User**, or a **Bill**.\nAvailable values: \"BILL\", \"COMMITMENT\", \"USER\", \"SERVICE_USER\", \"SCHEDULER\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"BILL",
						"COMMITMENT",
						"USER",
						"SERVICE_USER",
						"SCHEDULER",
					),
				},
			},
			"last_modified_by": schema.StringAttribute{
				Description: "The unique identifier (UUID) for the user who last modified the balance transaction.",
				Computed:    true,
			},
		},
	}
}

func (r *BalanceTransactionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *BalanceTransactionResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
