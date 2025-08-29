resource "m3ter_integration_configuration" "example_integration_configuration" {
  org_id = "orgId"
  destination = "Stripe"
  entity_type = "Bill"
  config_data = {
    foo = "bar"
  }
  credentials = {
    type = "HTTP_BASIC"
    destination = "WEBHOOK"
    empty = true
    name = "Integration Credentials"
    version = 0
  }
  destination_id = "00000000-0000-0000-0000-000000000000"
  entity_id = "00000000-0000-0000-0000-000000000000"
  integration_credentials_id = "00000000-0000-0000-0000-000000000000"
  name = "My Integration"
}
