// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package contract

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

var _ resource.ResourceWithConfigValidators = (*ContractResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Description: "Endpoints for Contract related operations such as creation, update, list and delete.\n\nContracts are created for Accounts, which are your end-user customers. Contracts can be used for:\n* **Accounts Reporting**. To serve your general accounting operations and processes, you can report on total Contract values for an Account.\n* **Contract Billing**. Various billing entities associated with an Account can  be linked to Contracts on the Account to meet your specific Contract billing use cases. ",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "The UUID of the entity.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"org_id": schema.StringAttribute{
				Optional:           true,
				DeprecationMessage: "the org id should be set at the client level instead",
				PlanModifiers:      []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"account_id": schema.StringAttribute{
				Description: "The unique identifier (UUID) of the Account associated with this Contract.",
				Required:    true,
			},
			"end_date": schema.StringAttribute{
				Description: "The exclusive end date of the Contract *(in ISO-8601 format)*. This means the Contract is active until midnight on the day ***before*** this date.",
				Required:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"name": schema.StringAttribute{
				Description: "The name of the Contract.",
				Required:    true,
			},
			"start_date": schema.StringAttribute{
				Description: "The start date for the Contract *(in ISO-8601 format)*. This date is inclusive, meaning the Contract is active from this date onward.",
				Required:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"apply_contract_period_limits": schema.BoolAttribute{
				Description: "For Contract billing, a boolean setting for restricting the charges billed to the period defined for the Contract:\n* **TRUE** - Contract billing for the Account will be restricted to charge amounts that fall within the defined Contract period.\n* **FALSE** - The period for amounts billed under the Contract will be determined by the Account Plan attached to the Account and linked to the Contract.(*Default*)",
				Optional:    true,
			},
			"bill_grouping_key_id": schema.StringAttribute{
				Description: "The ID of the Bill Grouping Key assigned to the Contract.\n\nIf you are implementing Contract Billing for an Account, use `billGroupingKey` to control how charges linked to Contracts on the Account will be billed:\n* **Independent Contract billing**. Assign an *exclusive* Bill Grouping Key to the Contract - only charges due against the Account and linked to the single Contract will appear on a separate Bill.\n* **Collective Contract billing**. Assign the same *non-exclusive* Bill Grouping Key to multiple Contracts - all charges due against the Account and linked to the multiple Contracts will appear together on a single Bill.",
				Optional:    true,
			},
			"code": schema.StringAttribute{
				Description: "The short code of the Contract.",
				Optional:    true,
			},
			"description": schema.StringAttribute{
				Description: "The description of the Contract, which provides context and information.",
				Optional:    true,
			},
			"purchase_order_number": schema.StringAttribute{
				Description: "The Purchase Order Number associated with the Contract.",
				Optional:    true,
			},
			"usage_filters": schema.ListNestedAttribute{
				Description: "Use `usageFilters` to control Contract billing and charge at billing only for usage where Product Meter dimensions equal specific defined values:\n* Define Usage filters to either *include* or *exclude* charges for usage associated with specific Meter dimensions.\n* The Meter dimensions must be present in the data field schema of the Meter used to submit usage data measurements.",
				Optional:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"dimension_code": schema.StringAttribute{
							Required: true,
						},
						"mode": schema.StringAttribute{
							Description: `Available values: "INCLUDE", "EXCLUDE".`,
							Required:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("INCLUDE", "EXCLUDE"),
							},
						},
						"value": schema.StringAttribute{
							Required: true,
						},
					},
				},
			},
			"custom_fields": schema.DynamicAttribute{
				Description:   "User defined fields enabling you to attach custom data. The value for a custom field can be either a string or a number.\n\nIf `customFields` can also be defined for this entity at the Organizational level, `customField` values defined at individual level override values of `customFields` with the same name defined at Organization level.\n\nSee [Working with Custom Fields](https://www.m3ter.com/docs/guides/creating-and-managing-products/working-with-custom-fields) in the m3ter documentation for more information.",
				Optional:      true,
				CustomType:    customfield.NormalizedDynamicType{},
				PlanModifiers: []planmodifier.Dynamic{customfield.NormalizeDynamicPlanModifier()},
			},
			"version": schema.Int64Attribute{
				Description: "The version number:\n- **Create:** On initial Create to insert a new entity, the version is set at 1 in the response.\n- **Update:** On successful Update, the version is incremented by 1 in the response.",
				Computed:    true,
			},
		},
	}
}

func (r *ContractResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *ContractResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
