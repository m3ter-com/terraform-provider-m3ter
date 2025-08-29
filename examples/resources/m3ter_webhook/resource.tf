resource "m3ter_webhook" "example_webhook" {
  org_id = "orgId"
  credentials = {
    api_key = "api key"
    secret = "api secret"
    type = "M3TER_SIGNED_REQUEST"
    empty = true
    version = 0
  }
  description = "x"
  name = "x"
  url = "x"
  active = true
  code = "code"
}
