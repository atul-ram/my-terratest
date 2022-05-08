output "resource_group_name" {
  value = data.azurerm_resource_group.resource_group.name
}

output "key_vault_name" {
  value = data.azurerm_key_vault.key_vault.name 
}

output "key_vault_secrets" {
  value = data.azurerm_key_vault_secrets.key_vault_secrets.names
}

output "subscriptionId" {
  value = data.azurerm_client_config.current.subscription_id
}