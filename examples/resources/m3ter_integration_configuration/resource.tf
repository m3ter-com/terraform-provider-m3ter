resource "m3ter_integration_configuration" "example_integration_configuration" {
  org_id = "orgId"
  config_data = {
    foo = "bar"
  }
  credentials = {
    type = "HTTP_BASIC"
    destination = "WEBHOOK"
    empty = true
    name = "name"
    version = 0
  }
  destination = "destination"
  destination_id = "destinationId"
  entity_id = "entityId"
  entity_type = "entityType"
  integration_credentials_id = "integrationCredentialsId"
  name = "name"
}
