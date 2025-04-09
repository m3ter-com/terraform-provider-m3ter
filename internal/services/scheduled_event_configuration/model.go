// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package scheduled_event_configuration

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/apijson"
)

type ScheduledEventConfigurationModel struct {
	ID             types.String      `tfsdk:"id" json:"id,computed"`
	OrgID          types.String      `tfsdk:"org_id" path:"orgId,optional"`
	Entity         types.String      `tfsdk:"entity" json:"entity,required"`
	Field          types.String      `tfsdk:"field" json:"field,required"`
	Name           types.String      `tfsdk:"name" json:"name,required"`
	Offset         types.Int64       `tfsdk:"offset" json:"offset,required"`
	Version        types.Int64       `tfsdk:"version" json:"version,optional"`
	CreatedBy      types.String      `tfsdk:"created_by" json:"createdBy,computed"`
	DtCreated      timetypes.RFC3339 `tfsdk:"dt_created" json:"dtCreated,computed" format:"date-time"`
	DtLastModified timetypes.RFC3339 `tfsdk:"dt_last_modified" json:"dtLastModified,computed" format:"date-time"`
	LastModifiedBy types.String      `tfsdk:"last_modified_by" json:"lastModifiedBy,computed"`
}

func (m ScheduledEventConfigurationModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ScheduledEventConfigurationModel) MarshalJSONForUpdate(state ScheduledEventConfigurationModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
