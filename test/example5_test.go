package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/azure"
	"github.com/gruntwork-io/terratest/modules/terraform"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
	"github.com/stretchr/testify/assert"
)

func TestExample5(t *testing.T) {
	t.Parallel()

	workingDir := "../examples/example5"
	varFile := "../usecases/abcd02/example4.tfvars"

	// Deploy the app using Terraform
	test_structure.RunTestStage(t, "deploy_terraform", func() {
		deployUsingTerraform(t, workingDir, varFile)
	})

	test_structure.RunTestStage(t, "ValidateKeyVault", func() {
		validateKeyVault(t, workingDir, varFile)
	})

	// At the end of the test, undeploy the web app using Terraform
	defer test_structure.RunTestStage(t, "cleanup_terraform", func() {
		undeployUsingTerraform(t, workingDir)
	})

}

func deployUsingTerraform(t *testing.T, workingDir string, varFile string) {

	// Construct the terraform options with default retryable errors to handle the most common retryable errors in
	// terraform testing.
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: workingDir,
		// Variables to pass to our Terraform code using -var-file options
		VarFiles: []string{varFile},
		// Disable colors in Terraform commands so its easier to parse stdout/stderr
		NoColor: true,
	})

	// Save the Terraform Options struct, instance name, and instance text so future test stages can use it
	test_structure.SaveTerraformOptions(t, workingDir, terraformOptions)

	// This will run `terraform init` and `terraform apply` and fail the test if there are any errors
	terraform.InitAndApply(t, terraformOptions)
}

// Undeploy the terraform-packer-example using Terraform
func undeployUsingTerraform(t *testing.T, workingDir string) {
	// Load the Terraform Options saved by the earlier deploy_terraform stage
	terraformOptions := test_structure.LoadTerraformOptions(t, workingDir)

	terraform.Destroy(t, terraformOptions)
}

func validateKeyVault(t *testing.T, workingDir string, varFile string) {
	// Load the Terraform Options saved by the earlier deploy_terraform stage
	terraformOptions := test_structure.LoadTerraformOptions(t, workingDir)

	subscriptionId := terraform.Output(t, terraformOptions, "subscriptionId")
	resourceGroupName := terraform.Output(t, terraformOptions, "resource_group_name")

	keyVaultName := terraform.Output(t, terraformOptions, "key_vault_name")
	// Determine whether the keyvault exists
	keyVault := azure.GetKeyVault(t, resourceGroupName, keyVaultName, subscriptionId)
	assert.Equal(t, keyVaultName, *keyVault.Name)

}
