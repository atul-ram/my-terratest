package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/azure"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestExample4(t *testing.T) {
	// retryable errors in terraform testing.
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../examples/example4",
		// Variables to pass to our Terraform code using -var-file options
		VarFiles: []string{"../usecases/abcd02/example4.tfvars"},
		// Disable colors in Terraform commands so its easier to parse stdout/stderr
		NoColor: true,
	})

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	subscriptionId := terraform.Output(t, terraformOptions, "subscriptionId")
	resourceGroupName := terraform.Output(t, terraformOptions, "resource_group_name")
	//assert.Equal(t, "abcd02-d-rg", output)
	keyVaultName := terraform.Output(t, terraformOptions, "key_vault_name")
	// website::tag::4:: Determine whether the keyvault exists
	keyVault := azure.GetKeyVault(t, resourceGroupName, keyVaultName, subscriptionId)
	assert.Equal(t, keyVaultName, *keyVault.Name)
}
