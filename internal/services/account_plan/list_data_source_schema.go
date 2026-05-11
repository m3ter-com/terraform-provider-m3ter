// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account_plan

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

var _ datasource.DataSourceWithConfigValidators = (*AccountPlansDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Endpoints for AccountPlan and AccountPlanGroup related operations such as creation, update, list and delete. \n\n**AccountPlans**\nAn Account represents one of your end-customer accounts. To create an AccountPlan, you attach a Product Plan to an Account. The AccountPlan then determines the charges incurred at billing by your end customer for consuming the Product the Plan is for:\n* **AccountPlan Active/Inactive**. Set start and end dates to define the period the AccountPlan is active for the Account.\n* **AccountPlan per Product**. If an end customer consumes multiple Products, create separate AccountPlans to charge for each Product.\n\n**AccountPlan Constraints:**\n* Only one AccountPlan per Product can be active at any one time for an Account.\n* If you create a Plan as a custom Plan for a specific Account, you can only use it to create an AccountPlan for that Account.\n\n**AccountPlanGroups**\nPlan Groups are used when you want to apply a minimum spend amount at billing across several of your Products each of which are priced separately - when you create the Plan Group, you define an overall minimum spend and then add any priced Plans you want to include in the Group. To create an AccounPlanGroup, you can attach a Plan Group to an Account that consumes the separate Products which are priced using the included Plans. At billing, the minimum spend you've defined for the Plan Group is applied:\n* **Active AccountPlanGroup**. Set the start and end dates to define the period for which the Plan Group will be active for the Account.\n\n**Plan Group Notes:**\n* You can only add *one Plan for the same Product* to a Plan Group. See the [Plan Group](https://www.m3ter.com/docs/api#tag/PlanGroup) in this API Reference for more details on creating Plan Groups.\n* You can create a *custom Plan Group* for an Account, which means the Plan Group can only be attached to that Account to create an AccountPlanGroup.\n\n**AcountPlanGroup - Notes and Constraints:**\n* **AccountPlanGroup is type of AccountPlan** When you attach a Plan Group to an Account, this creates an AccountPlanGroup. However, the m3ter data model *does not support a separate AccountPlanGroup entity*, and an AccountPlanGroup is a type of AccountPlan where a `planGroupId` is used instead of a `planId` when it's created. See the [Create AccountPlan](https://www.m3ter.com/docs/api#tag/AccountPlan/operation/PostAccountPlan) call in this section and [Attaching Plan Groups to an Account](https://www.m3ter.com/docs/guides/end-customer-accounts/attaching-plan-groups-to-an-account) in our main User Documentation.\n* **Multiple AccountPlan Groups:** You can attach more than one Plan Group to an Account to create multiple AccountPlanGroups, but the rule that *only one attached Plan per Product can be active at any one time for an Account* is preserved:\n\t* Multiple attached Plan Groups on an Account can have overlapping dates only if none of the Plan Groups contain a Plan belonging to the same Product. If you try to attach a Plan Group to an Account with Plan Groups already attached and:\n\t\t* The new Plan Group contains a Product Plan that also belongs to a Plan Group already attached to the Account.\n\t\t* The dates for these \"matched Plan\" Plan Groups being active for the Account would overlap.\n\t\t* Then you'll receive an error and the attachment will be blocked.\n",
		Attributes: map[string]schema.Attribute{
			"org_id": schema.StringAttribute{
				Optional:           true,
				DeprecationMessage: "the org id should be set at the client level instead",
			},
			"account": schema.StringAttribute{
				Description: "The unique identifier (UUID) for the Account whose AccountPlans and AccountPlanGroups you want to retrieve.\n\n**NOTE:** Only returns the currently active AccountPlans and AccountPlanGroups for the specified Account. Use in combination with the `includeall` query parameter to return both active and inactive.",
				Optional:    true,
			},
			"contract": schema.StringAttribute{
				Description: "The unique identifier (UUID) of the Contract which the AccountPlans you want to retrieve have been linked to.\n\n**NOTE:** Does not return AccountPlanGroups that have been linked to the Contract.",
				Optional:    true,
			},
			"date": schema.StringAttribute{
				Description: "The specific date for which you want to retrieve AccountPlans and AccountPlanGroups.\n\n**NOTE:** Returns both active and inactive AccountPlans and AccountPlanGroups for the specified date.",
				Optional:    true,
			},
			"includeall": schema.BoolAttribute{
				Description: "A Boolean flag that specifies whether to include both active and inactive AccountPlans and AccountPlanGroups in the list.\n\n* **TRUE** - both active and inactive AccountPlans and AccountPlanGroups are included in the list.\n* **FALSE** - only active AccountPlans and AccountPlanGroups are retrieved in the list.*(Default)*\n\n**NOTE:** Only operative if you also have one of `account`, `plan` or `contract` as a query parameter.",
				Optional:    true,
			},
			"plan": schema.StringAttribute{
				Description: "The unique identifier (UUID) for the Plan whose associated AccountPlans you want to retrieve.\n\n**NOTE:** Does not return AccountPlanGroups if you use a `planGroupId`.",
				Optional:    true,
			},
			"product": schema.StringAttribute{
				Description: "The unique identifier (UUID) for the Product whose associated AccountPlans you want to retrieve.\n\n**NOTE:** You cannot use the `product` query parameter as a single filter condition, but must always use it in combination with the `account` query parameter.",
				Optional:    true,
			},
			"ids": schema.ListAttribute{
				Description: "A list of unique identifiers (UUIDs) for specific AccountPlans and AccountPlanGroups you want to retrieve.",
				Optional:    true,
				ElementType: types.StringType,
			},
			"max_items": schema.Int64Attribute{
				Description: "Max items to fetch, default: 1000",
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
			},
			"items": schema.DynamicAttribute{
				Description: "The items returned by the data source",
				Computed:    true,
				CustomType:  customfield.NormalizedDynamicType{},
			},
		},
	}
}

func (d *AccountPlansDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *AccountPlansDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
