// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package counter

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

var _ resource.ResourceWithConfigValidators = (*CounterResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Endpoints for listing, creating, retrieving, updating, or deleting Counters.\n\nYou can create Counters for your m3ter Organization, which can then be used as pricing metrics to apply a unit-based [CounterPricing](https://www.m3ter.com/docs/api#tag/CounterPricing) to Product Plans or Plan Templates for recurring subscription charges on Accounts.\n\nCounters can then be used to post [CounterAdjustments](https://www.m3ter.com/docs/api#tag/CounterAdjustments) on your end-customer Accounts.\n\nAccounts are then billed in accordance with the CounterPricing on Plans attached to the Accounts and for the actual Counter quantities Accounts subscribe to. See [Recurring Charges: Counters](https://www.m3ter.com/docs/guides/recurring-charges-counters) in our main user documentation for more details.",
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
			"name": schema.StringAttribute{
				Description: "Descriptive name for the Counter.",
				Required:    true,
			},
			"unit": schema.StringAttribute{
				Description: "User defined label for units shown on Bill line items, and indicating to your customers what they are being charged for.",
				Required:    true,
			},
			"code": schema.StringAttribute{
				Description: "Code for the Counter. A unique short code to identify the Counter.",
				Optional:    true,
			},
			"product_id": schema.StringAttribute{
				Description: "UUID of the product the Counter belongs to. *(Optional)* - if left blank, the Counter is Global. A Global Counter can be used to price Plans or Plan Templates belonging to any Product.",
				Optional:    true,
			},
			"version": schema.Int64Attribute{
				Description: "The version number:\n- **Create:** On initial Create to insert a new entity, the version is set at 1 in the response.\n- **Update:** On successful Update, the version is incremented by 1 in the response.",
				Computed:    true,
			},
		},
	}
}

func (r *CounterResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *CounterResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
