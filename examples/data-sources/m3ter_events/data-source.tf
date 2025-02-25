data "m3ter_events" "example_events" {
  org_id = "orgId"
  account_id = "accountId"
  event_name = "eventName"
  event_type = "eventType"
  ids = ["string"]
  include_actioned = true
  notification_code = "notificationCode"
  notification_id = "notificationId"
  resource_id = "resourceId"
}
