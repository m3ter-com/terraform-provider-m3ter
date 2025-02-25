data "m3ter_data_export_jobs" "example_data_export_jobs" {
  org_id = "orgId"
  date_created_end = "dateCreatedEnd"
  date_created_start = "dateCreatedStart"
  ids = ["string"]
  schedule_id = "scheduleId"
  status = "PENDING"
}
