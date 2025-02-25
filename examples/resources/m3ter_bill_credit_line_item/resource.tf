resource "m3ter_bill_credit_line_item" "example_bill_credit_line_item" {
  org_id = "orgId"
  bill_id = "billId"
  amount = 1
  description = "x"
  product_id = "productId"
  referenced_bill_id = "referencedBillId"
  referenced_line_item_id = "referencedLineItemId"
  service_period_end_date = "2019-12-27T18:11:19.117Z"
  service_period_start_date = "2019-12-27T18:11:19.117Z"
  credit_reason_id = "creditReasonId"
  line_item_type = "STANDING_CHARGE"
  reason_id = "reasonId"
  version = 0
}
