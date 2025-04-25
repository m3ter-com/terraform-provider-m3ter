resource "m3ter_data_export_schedule" "example_data_export_schedule" {
  org_id = "orgId"
  operational_data_types = ["BILLS"]
  source_type = "OPERATIONAL"
  version = 0
}
