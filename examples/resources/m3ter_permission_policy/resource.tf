resource "m3ter_permission_policy" "example_permission_policy" {
  org_id = "orgId"
  name = "x"
  permission_policy = [{
    action = ["ALL"]
    effect = "ALLOW"
    resource = ["string"]
  }]
}
