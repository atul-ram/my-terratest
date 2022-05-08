package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	// retryable errors in terraform testing.
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../examples/example1",
		// Variables to pass to our Terraform code using -var-file options
		VarFiles: []string{"../usecases/abcd02/example1.tfvars"},
		// Disable colors in Terraform commands so its easier to parse stdout/stderr
		NoColor: true,
	})

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	output := terraform.Output(t, terraformOptions, "hello_world")
	assert.Equal(t, "Hello, World! name of the variable is abcd02-d-rg", output)
}
