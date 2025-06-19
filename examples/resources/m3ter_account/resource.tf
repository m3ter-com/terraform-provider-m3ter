resource "m3ter_account" "example_account" {
  org_id = "orgId"
  code = "JS!?Q0]r] ]$]"
  email_address = "dev@stainless.com"
  name = "x"
  address = {
    address_line1 = "addressLine1"
    address_line2 = "addressLine2"
    address_line3 = "addressLine3"
    address_line4 = "addressLine4"
    country = "country"
    locality = "locality"
    post_code = "postCode"
    region = "region"
  }
  auto_generate_statement_mode = "NONE"
  bill_epoch = "2019-12-27"
  config_data = {
    foo = "bar"
  }
  credit_application_order = ["PREPAYMENT"]
  currency = "USD"
  custom_fields = {
    foo = "string"
  }
  days_before_bill_due = 1
  parent_account_id = "parentAccountId"
  purchase_order_number = "purchaseOrderNumber"
  statement_definition_id = "statementDefinitionId"
}
