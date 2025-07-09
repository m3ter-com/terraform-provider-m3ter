resource "m3ter_currency" "example_currency" {
  org_id = "orgId"
  name = "x"
  archived = true
  code = "code"
  max_decimal_places = 2
  rounding_mode = "UP"
}
