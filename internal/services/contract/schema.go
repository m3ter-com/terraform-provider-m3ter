// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package contract

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*ContractResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
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
			"custom_fields": schema.MapAttribute{
				Description: "User defined fields enabling you to attach custom data. The value for a custom field can be either a string or a number.\n\nIf `customFields` can also be defined for this entity at the Organizational level, `customField` values defined at individual level override values of `customFields` with the same name defined at Organization level.\n\nSee [Working with Custom Fields](https://www.m3ter.com/docs/guides/creating-and-managing-products/working-with-custom-fields) in the m3ter documentation for more information.",
				Optional:    true,
				ElementType: types.DynamicType,
			},
			"created_by": schema.StringAttribute{
				Description: "The unique identifier (UUID) of the user who created this Contract.",
				Computed:    true,
			},
			"dt_created": schema.StringAttribute{
				Description: "The date and time *(in ISO-8601 format)* when the Contract was created.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"dt_last_modified": schema.StringAttribute{
				Description: "The date and time *(in ISO-8601 format)* when the Contract was last modified.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"last_modified_by": schema.StringAttribute{
				Description: "The unique identifier (UUID) of the user who last modified this Contract.",
				Computed:    true,
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
