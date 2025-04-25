resource "m3ter_data_export_destination" "example_data_export_destination" {
  org_id = "orgId"
  bucket_name = "xxx"
  iam_role_arn = "arn:aws:iam::321669910225:role/z"
  destination_type = "S3"
  partition_order = "TYPE_FIRST"
  prefix = "prefix"
  version = 0
}
