// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account_plan

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/customfield"
)

var _ datasource.DataSourceWithConfigValidators = (*AccountPlansDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
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
				Description: "A Boolean flag that specifies whether to include both active and inactive AccountPlans and AccountPlanGroups in the list.\n\n* **TRUE** - both active and inactive AccountPlans and AccountPlanGroups are included in the list.\n* **FALSE** - only active AccountPlans and AccountPlanGroups are retrieved in the list.",
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
			"items": schema.ListNestedAttribute{
				Description: "The items returned by the data source",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[AccountPlansItemsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The UUID of the entity.",
							Computed:    true,
						},
						"account_id": schema.StringAttribute{
							Description: "The unique identifier (UUID) for the Account to which the AccountPlan or AccounPlanGroup is attached.",
							Computed:    true,
						},
						"bill_epoch": schema.StringAttribute{
							Description: "The initial date for creating the first bill against the Account for charges due under the AccountPlan or AccountPlanGroup. All subsequent bill creation dates are calculated from this date. If left empty, the first bill date definedfor the Account is used. The date is in ISO-8601 format.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"child_billing_mode": schema.StringAttribute{
							Description: "If the Account is either a Parent or a Child Account, this specifies the Account hierarchy billing mode. The mode determines how billing will be handled and shown on bills for charges due on the Parent Account, and charges due on Child Accounts:\n\n* **Parent Breakdown** - a separate bill line item per Account. Default setting.\n\n* **Parent Summary** - single bill line item for all Accounts.\n\n* **Child** - the Child Account is billed.\nAvailable values: \"PARENT_SUMMARY\", \"PARENT_BREAKDOWN\", \"CHILD\".",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"PARENT_SUMMARY",
									"PARENT_BREAKDOWN",
									"CHILD",
								),
							},
						},
						"code": schema.StringAttribute{
							Description: "The unique short code of the AccountPlan or AccountPlanGroup.",
							Computed:    true,
						},
						"contract_id": schema.StringAttribute{
							Description: "The unique identifier (UUID) for the Contract to which the Plan or Plan Group  attached to the Account has been added.",
							Computed:    true,
						},
						"custom_fields": schema.DynamicAttribute{
							Description: "User defined fields enabling you to attach custom data. The value for a custom field can be either a string or a number.\n\nIf `customFields` can also be defined for this entity at the Organizational level,`customField` values defined at individual level override values of `customFields` with the same name defined at Organization level.\n\nSee [Working with Custom Fields](https://www.m3ter.com/docs/guides/creating-and-managing-products/working-with-custom-fields) in the m3ter documentation for more information.",
							Computed:    true,
						},
						"end_date": schema.StringAttribute{
							Description: "The end date *(in ISO-8601 format)* for when the AccountPlan or AccountPlanGroup ceases to be active for the Account. If not specified, the AccountPlan or AccountPlanGroup remains active indefinitely.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"plan_group_id": schema.StringAttribute{
							Description: "The unique identifier (UUID) of the Plan Group that has been attached to the Account to create the AccountPlanGroup.",
							Computed:    true,
						},
						"plan_id": schema.StringAttribute{
							Description: "The unique identifier (UUID) of the Plan that has been attached to the Account to create the AccountPlan.",
							Computed:    true,
						},
						"product_id": schema.StringAttribute{
							Description: "The unique identifier (UUID) for the Product associated with the AccountPlan.\n\n**Note:** Not present in response for AccountPlanGroup - Plan Groups can contain multiple Plans belonging to different Products.",
							Computed:    true,
						},
						"start_date": schema.StringAttribute{
							Description: "The start date *(in ISO-8601 format)* for the when the AccountPlan or AccountPlanGroup starts to be active for the Account.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"version": schema.Int64Attribute{
							Description: "The version number:\n- **Create:** On initial Create to insert a new entity, the version is set at 1 in the response.\n- **Update:** On successful Update, the version is incremented by 1 in the response.",
							Computed:    true,
						},
					},
				},
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
