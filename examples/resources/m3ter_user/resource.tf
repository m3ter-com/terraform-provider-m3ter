resource "m3ter_user" "example_user" {
  org_id = "orgId"
  id = "id"
  dt_end_access = "2019-12-27T18:11:19.117Z"
  permission_policy = [{
    action = ["ALL"]
    effect = "ALLOW"
    resource = ["string"]
  }]
  version = 0
}
