package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestExample2(t *testing.T) {
	// retryable errors in terraform testing.
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../examples/example2",
		// Variables to pass to our Terraform code using -var-file options
		VarFiles: []string{"../usecases/abcd02/example2.tfvars"},
		// Disable colors in Terraform commands so its easier to parse stdout/stderr
		NoColor: true,
	})

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	output := terraform.Output(t, terraformOptions, "resource_group_name")
	assert.Equal(t, "abcd02-d-rg", output)
}
