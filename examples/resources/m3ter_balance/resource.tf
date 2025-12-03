resource "m3ter_balance" "example_balance" {
  org_id = "orgId"
  account_id = "x"
  code = "S?oC\"$]C] ]]]]]5]"
  currency = "x"
  end_date = "2019-12-27T18:11:19.117Z"
  name = "x"
  start_date = "2019-12-27T18:11:19.117Z"
  balance_draw_down_description = "balanceDrawDownDescription"
  consumptions_accounting_product_id = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
  contract_id = "contractId"
  custom_fields = {
    foo = "string"
  }
  description = "description"
  fees_accounting_product_id = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
  line_item_types = ["STANDING_CHARGE"]
  overage_description = "overageDescription"
  overage_surcharge_percent = 0
  product_ids = ["string"]
  rollover_amount = 0
  rollover_end_date = "2019-12-27T18:11:19.117Z"
}
