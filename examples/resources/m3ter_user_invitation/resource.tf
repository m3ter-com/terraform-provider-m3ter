resource "m3ter_user_invitation" "example_user_invitation" {
  org_id = "orgId"
  email = "dev@stainless.com"
  first_name = "x"
  last_name = "x"
  contact_number = "contactNumber"
  dt_end_access = "2019-12-27T18:11:19.117Z"
  dt_expiry = "2019-12-27T18:11:19.117Z"
  m3ter_user = true
  permission_policy_ids = ["string"]
  version = 0
}
