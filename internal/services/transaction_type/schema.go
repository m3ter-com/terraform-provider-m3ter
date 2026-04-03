// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package transaction_type

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

var _ resource.ResourceWithConfigValidators = (*TransactionTypeResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Endpoints for TransactionType operations such as creation, update, list, retrieve, and delete. \n\nYou can create TransactionTypes for your Organization, which can then be used when creating and updating Balances. Example TransactionTypes: \"Balance Amount\" or \"Add Funds\".\n\nFor details on creating a Transaction amount for a Balance using a TransactionType you've created for your Organization, see the [Create Balance Transaction](https://www.m3ter.com/docs/api#tag/Balances/operation/PostBalanceTransaction) call in the [Balances](https://www.m3ter.com/docs/api#tag/Balances) section of this API Reference.",
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
				Description: "The name of the entity.",
				Required:    true,
			},
			"archived": schema.BoolAttribute{
				Description: "A Boolean TRUE / FALSE flag indicating whether the entity is archived. An entity can be archived if it is obsolete.\n\n* TRUE - the entity is in the archived state.\n* FALSE - the entity is not in the archived state.",
				Optional:    true,
			},
			"code": schema.StringAttribute{
				Description: "The short code for the entity.",
				Optional:    true,
			},
			"version": schema.Int64Attribute{
				Description: "The version number:\n- **Create:** On initial Create to insert a new entity, the version is set at 1 in the response.\n- **Update:** On successful Update, the version is incremented by 1 in the response.",
				Computed:    true,
			},
		},
	}
}

func (r *TransactionTypeResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *TransactionTypeResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
