resource "m3ter_contract" "example_contract" {
  org_id = "orgId"
  account_id = "x"
  end_date = "2019-12-27"
  name = "x"
  start_date = "2019-12-27"
  apply_contract_period_limits = true
  bill_grouping_key_id = "billGroupingKeyId"
  code = "S?oC\"$]C] ]]]]]5]"
  custom_fields = {
    foo = "string"
  }
  description = "description"
  purchase_order_number = "purchaseOrderNumber"
  usage_filters = [{
    dimension_code = "x"
    mode = "INCLUDE"
    value = "x"
  }]
}
