// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package data_export_destination

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/m3ter-com/terraform-provider-m3ter/internal/apijson"
)

type DataExportDestinationModel struct {
	ID                  types.String      `tfsdk:"id" json:"id,computed"`
	OrgID               types.String      `tfsdk:"org_id" path:"orgId,optional"`
	BucketName          types.String      `tfsdk:"bucket_name" json:"bucketName,required"`
	IamRoleArn          types.String      `tfsdk:"iam_role_arn" json:"iamRoleArn,optional"`
	PartitionOrder      types.String      `tfsdk:"partition_order" json:"partitionOrder,optional"`
	PoolID              types.String      `tfsdk:"pool_id" json:"poolId,optional"`
	Prefix              types.String      `tfsdk:"prefix" json:"prefix,optional"`
	ProjectNumber       types.String      `tfsdk:"project_number" json:"projectNumber,optional"`
	ProviderID          types.String      `tfsdk:"provider_id" json:"providerId,optional"`
	ServiceAccountEmail types.String      `tfsdk:"service_account_email" json:"serviceAccountEmail,optional"`
	Version             types.Int64       `tfsdk:"version" json:"version,optional"`
	Code                types.String      `tfsdk:"code" json:"code,computed"`
	CreatedBy           types.String      `tfsdk:"created_by" json:"createdBy,computed"`
	DestinationType     types.String      `tfsdk:"destination_type" json:"destinationType,computed"`
	DtCreated           timetypes.RFC3339 `tfsdk:"dt_created" json:"dtCreated,computed" format:"date-time"`
	DtLastModified      timetypes.RFC3339 `tfsdk:"dt_last_modified" json:"dtLastModified,computed" format:"date-time"`
	LastModifiedBy      types.String      `tfsdk:"last_modified_by" json:"lastModifiedBy,computed"`
	Name                types.String      `tfsdk:"name" json:"name,computed"`
}

func (m DataExportDestinationModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m DataExportDestinationModel) MarshalJSONForUpdate(state DataExportDestinationModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
