// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account_plan

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*AccountPlanResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "The UUID of the entity. ",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"org_id": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"account_id": schema.StringAttribute{
				Description: "The unique identifier (UUID) for the Account.",
				Required:    true,
			},
			"start_date": schema.StringAttribute{
				Description: "The start date *(in ISO-8601 format)* for the AccountPlan or AccountPlanGroup becoming active for the Account.",
				Required:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"bill_epoch": schema.StringAttribute{
				Description: "Optional setting to define a *billing cycle date*, which acts as a reference for when in the applied billing frequency period bills are created:\n* For example, if you attach a Plan to an Account where the Plan is configured for monthly billing frequency and you've defined the period the Plan will apply to the Account to be from January 1st, 2022 until January 1st, 2023. You then set a `billEpoch` date of February 15th, 2022. The first Bill will be created for the Account on February 15th, and subsequent Bills created on the 15th of the months following for the remainder of the billing period - March 15th, April 15th, and so on.\n* If not defined, then the `billEpoch` date set for the Account will be used instead.\n* The date is in ISO-8601 format.",
				Optional:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"child_billing_mode": schema.StringAttribute{
				Description: "If the Account is either a Parent or a Child Account, this specifies the Account hierarchy billing mode. The mode determines how billing will be handled and shown on bills for charges due on the Parent Account, and charges due on Child Accounts:\n\n* **Parent Breakdown** - a separate bill line item per Account. Default setting.\n\n* **Parent Summary** - single bill line item for all Accounts.\n\n* **Child** - the Child Account is billed.",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"PARENT_SUMMARY",
						"PARENT_BREAKDOWN",
						"CHILD",
					),
				},
			},
			"code": schema.StringAttribute{
				Description: "A unique short code for the AccountPlan or AccountPlanGroup.",
				Optional:    true,
			},
			"contract_id": schema.StringAttribute{
				Description: "The unique identifier (UUID) for a Contract to which you want to add the Plan or Plan Group being attached to the Account.",
				Optional:    true,
			},
			"end_date": schema.StringAttribute{
				Description: "The end date *(in ISO-8601 format)* for when the AccountPlan or AccountPlanGroup ceases to be active for the Account. If not specified, the AccountPlan or AccountPlanGroup remains active indefinitely.",
				Optional:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"plan_group_id": schema.StringAttribute{
				Description: "The unique identifier (UUID) of the PlanGroup to be attached to the Account to create an AccountPlanGroup.\n\n**Note:** Exclusive of the `planId` request parameter - exactly one of `planId` or `planGroupId` must be used per call.",
				Optional:    true,
			},
			"plan_id": schema.StringAttribute{
				Description: "The unique identifier (UUID) of the Plan to be attached to the Account to create an AccountPlan.\n\n**Note:** Exclusive of the `planGroupId` request parameter - exactly one of `planId` or `planGroupId` must be used per call.",
				Optional:    true,
			},
			"version": schema.Int64Attribute{
				Description: "The version number of the entity:\n- **Create entity:** Not valid for initial insertion of new entity - *do not use for Create*. On initial Create, version is set at 1 and listed in the response.\n- **Update Entity:**  On Update, version is required and must match the existing version because a check is performed to ensure sequential versioning is preserved. Version is incremented by 1 and listed in the response.",
				Optional:    true,
			},
			"custom_fields": schema.MapAttribute{
				Description: "User defined fields enabling you to attach custom data. The value for a custom field can be either a string or a number.\n\nIf `customFields` can also be defined for this entity at the Organizational level, `customField` values defined at individual level override values of `customFields` with the same name defined at Organization level.\n\nSee [Working with Custom Fields](https://www.m3ter.com/docs/guides/creating-and-managing-products/working-with-custom-fields) in the m3ter documentation for more information.",
				Optional:    true,
				ElementType: types.DynamicType,
			},
			"created_by": schema.StringAttribute{
				Description: "The unique identifier (UUID) for the user who created the AccountPlan or AccountPlanGroup.",
				Computed:    true,
			},
			"dt_created": schema.StringAttribute{
				Description: "The date and time *(in ISO 8601 format)* when the AccountPlan or AccountPlanGroup was first created.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"dt_last_modified": schema.StringAttribute{
				Description: "The date and time *(in ISO 8601 format)* when the AccountPlan or AccountPlanGroup was last modified.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"last_modified_by": schema.StringAttribute{
				Description: "The unique identifier (UUID) for the user who last modified the AccountPlan or AccountPlanGroup.",
				Computed:    true,
			},
			"product_id": schema.StringAttribute{
				Description: "The unique identifier (UUID) for the Product associated with the AccountPlan.\n\n**Note:** Not present in response for AccountPlanGroup - Plan Groups can contain multiple Plans belonging to different Products.",
				Computed:    true,
			},
		},
	}
}

func (r *AccountPlanResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *AccountPlanResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
