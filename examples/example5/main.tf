data "azurerm_client_config" "current" {}

data "azurerm_resource_group" "resource_group" {
  name = var.resource_group_name
}


data "azurerm_key_vault" "key_vault"{
  name = "${var.usecase}-${var.environment}-01-kv"
  resource_group_name = data.azurerm_resource_group.resource_group.name
}

data "azurerm_key_vault_secrets" "key_vault_secrets"{
  key_vault_id = data.azurerm_key_vault.key_vault.id
}


