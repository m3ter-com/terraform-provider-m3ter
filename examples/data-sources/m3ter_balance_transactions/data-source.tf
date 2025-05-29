data "m3ter_balance_transactions" "example_balance_transactions" {
  org_id = "orgId"
  balance_id = "balanceId"
  entity_id = "entityId"
  entity_type = "BILL"
  transaction_type_id = "transactionTypeId"
}
