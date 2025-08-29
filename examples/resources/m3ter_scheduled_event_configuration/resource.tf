resource "m3ter_scheduled_event_configuration" "example_scheduled_event_configuration" {
  org_id = "orgId"
  entity = "Bill"
  field = "dueDate"
  name = "10 Days After Bill Due Date"
  offset = 10
}
