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
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

var _ resource.ResourceWithConfigValidators = (*AccountPlanResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Description: "Endpoints for AccountPlan and AccountPlanGroup related operations such as creation, update, list and delete. \n\n**AccountPlans**\nAn Account represents one of your end-customer accounts. To create an AccountPlan, you attach a Product Plan to an Account. The AccountPlan then determines the charges incurred at billing by your end customer for consuming the Product the Plan is for:\n* **AccountPlan Active/Inactive**. Set start and end dates to define the period the AccountPlan is active for the Account.\n* **AccountPlan per Product**. If an end customer consumes multiple Products, create separate AccountPlans to charge for each Product.\n\n**AccountPlan Constraints:**\n* Only one AccountPlan per Product can be active at any one time for an Account.\n* If you create a Plan as a custom Plan for a specific Account, you can only use it to create an AccountPlan for that Account.\n\n**AccountPlanGroups**\nPlan Groups are used when you want to apply a minimum spend amount at billing across several of your Products each of which are priced separately - when you create the Plan Group, you define an overall minimum spend and then add any priced Plans you want to include in the Group. To create an AccounPlanGroup, you can attach a Plan Group to an Account that consumes the separate Products which are priced using the included Plans. At billing, the minimum spend you've defined for the Plan Group is applied:\n* **Active AccountPlanGroup**. Set the start and end dates to define the period for which the Plan Group will be active for the Account.\n\n**Plan Group Notes:**\n* You can only add *one Plan for the same Product* to a Plan Group. See the [Plan Group](https://www.m3ter.com/docs/api#tag/PlanGroup) in this API Reference for more details on creating Plan Groups.\n* You can create a *custom Plan Group* for an Account, which means the Plan Group can only be attached to that Account to create an AccountPlanGroup.\n\n**AcountPlanGroup - Notes and Constraints:**\n* **AccountPlanGroup is type of AccountPlan** When you attach a Plan Group to an Account, this creates an AccountPlanGroup. However, the m3ter data model *does not support a separate AccountPlanGroup entity*, and an AccountPlanGroup is a type of AccountPlan where a `planGroupId` is used instead of a `planId` when it's created. See the [Create AccountPlan](https://www.m3ter.com/docs/api#tag/AccountPlan/operation/PostAccountPlan) call in this section and [Attaching Plan Groups to an Account](https://www.m3ter.com/docs/guides/end-customer-accounts/attaching-plan-groups-to-an-account) in our main User Documentation.\n* **Multiple AccountPlan Groups:** You can attach more than one Plan Group to an Account to create multiple AccountPlanGroups, but the rule that *only one attached Plan per Product can be active at any one time for an Account* is preserved:\n\t* Multiple attached Plan Groups on an Account can have overlapping dates only if none of the Plan Groups contain a Plan belonging to the same Product. If you try to attach a Plan Group to an Account with Plan Groups already attached and:\n\t\t* The new Plan Group contains a Product Plan that also belongs to a Plan Group already attached to the Account.\n\t\t* The dates for these \"matched Plan\" Plan Groups being active for the Account would overlap.\n\t\t* Then you'll receive an error and the attachment will be blocked.\n",
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
				Description: "If the Account is either a Parent or a Child Account, this specifies the Account hierarchy billing mode. The mode determines how billing will be handled and shown on bills for charges due on the Parent Account, and charges due on Child Accounts:\n\n* **Parent Breakdown** - a separate bill line item per Account. Default setting.\n\n* **Parent Summary** - single bill line item for all Accounts.\n\n* **Child** - the Child Account is billed.\nAvailable values: \"PARENT_SUMMARY\", \"PARENT_BREAKDOWN\", \"CHILD\".",
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
			"custom_fields": schema.DynamicAttribute{
				Description:   "User defined fields enabling you to attach custom data. The value for a custom field can be either a string or a number.\n\nIf `customFields` can also be defined for this entity at the Organizational level, `customField` values defined at individual level override values of `customFields` with the same name defined at Organization level.\n\nSee [Working with Custom Fields](https://www.m3ter.com/docs/guides/creating-and-managing-products/working-with-custom-fields) in the m3ter documentation for more information.",
				Optional:      true,
				CustomType:    customfield.NormalizedDynamicType{},
				PlanModifiers: []planmodifier.Dynamic{customfield.NormalizeDynamicPlanModifier()},
			},
			"product_id": schema.StringAttribute{
				Description: "The unique identifier (UUID) for the Product associated with the AccountPlan.\n\n**Note:** Not present in response for AccountPlanGroup - Plan Groups can contain multiple Plans belonging to different Products.",
				Computed:    true,
			},
			"version": schema.Int64Attribute{
				Description: "The version number:\n- **Create:** On initial Create to insert a new entity, the version is set at 1 in the response.\n- **Update:** On successful Update, the version is incremented by 1 in the response.",
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
