resource "m3ter_scheduled_event_configuration" "example_scheduled_event_configuration" {
  org_id = "orgId"
  entity = "Bill"
  field = "endDate"
  name = "scheduled.bill.enddateEvent"
  offset = 5
}
