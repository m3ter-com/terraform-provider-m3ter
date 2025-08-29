resource "m3ter_external_mapping" "example_external_mapping" {
  org_id = "orgId"
  external_id = "cus_00000000000000"
  external_system = "Stripe"
  external_table = "Customer"
  m3ter_entity = "Account"
  m3ter_id = "00000000-0000-0000-0000-000000000000"
  integration_config_id = "00000000-0000-0000-0000-000000000000"
}
