resource "m3ter_notification_configuration" "example_notification_configuration" {
  org_id = "orgId"
  code = "commitment_under_10_percent"
  description = "Commitment amount fell below 10%"
  event_name = "configuration.commitment.updated"
  name = "Commitment has under 10% remaining"
  active = true
  always_fire_event = false
  calculation = <<EOT
  (new.amountSpent >= ((new.amount*90)/100)) 
  AND ((old.amountSpent <= ((old.amount*90)/100)) OR (old.amountSpent == null))
  EOT
}
