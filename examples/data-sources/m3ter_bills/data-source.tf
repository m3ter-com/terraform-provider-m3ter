data "m3ter_bills" "example_bills" {
  org_id = "orgId"
  account_id = "accountId"
  bill_date = "billDate"
  bill_date_end = "billDateEnd"
  bill_date_start = "billDateStart"
  billing_frequency = "billingFrequency"
  exclude_line_items = true
  external_invoice_date_end = "externalInvoiceDateEnd"
  external_invoice_date_start = "externalInvoiceDateStart"
  ids = ["string"]
  include_bill_total = true
  locked = true
  status = "PENDING"
}
