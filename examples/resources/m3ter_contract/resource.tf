resource "m3ter_contract" "example_contract" {
  org_id = "orgId"
  account_id = "x"
  end_date = "2019-12-27"
  name = "x"
  start_date = "2019-12-27"
  code = "S?oC\"$]C] ]]]]]5]"
  custom_fields = {
    foo = "string"
  }
  description = "description"
  purchase_order_number = "purchaseOrderNumber"
}
