package tests

import (
	"testing"

	"gotest.tools/v3/assert"

	"github.com/gcarreno/cobra-cli-ng/utils"
)

var (
	// Sample inputs
	inputs = []string{
		"",
		"1Some-Other_stuff",
		"SOme-other_stuff__",
		"sOme--OTHER__StUfF",
		"cli_CMD-EXAMPLE",
		"CLI-CMD-EXAMPLE",
		"123hello-world",
		"____",
		"FOO_barBaz",
		"foo_BARbaz",
	}

	// Expected outputs for Sanitize
	outputs = []string{
		"error",
		"error",
		"SomeOtherStuff",
		"sOmeOtherStUfF",
		"cliCmdExample",
		"CliCmdExample",
		"error",
		"error",
		"FooBarBaz",
		"fooBarbaz",
	}

	// Expected outputs for SanitizeStrict
	outputsStrict = []string{
		"error",
		"SomeOtherStuff",
		"SomeOtherStuff",
		"sOmeOtherStUfF",
		"cliCmdExample",
		"CliCmdExample",
		"helloWorld",
		"error",
		"FooBarBaz",
		"fooBarbaz",
	}

	// Error messages for Sanitize
	outErrors = []string{
		"command cannot be empty",
		"command name cannot start with a digit",
		"",
		"",
		"",
		"",
		"command name cannot start with a digit",
		"no valid characters in command name",
		"",
		"",
	}

	// Error messages for SanitizeStrict
	outErrorsStrict = []string{
		"command cannot be empty",
		"command name cannot start with a digit",
		"",
		"",
		"",
		"",
		"command name cannot start with a digit",
		"no valid characters in command name",
		"",
		"",
	}
)

func TestSanitize(t *testing.T) {
	// Make test run in parallel
	t.Parallel()

	// Run through our test cases
	for index, input := range inputs {
		output, err := utils.Sanitize(input)
		if outputs[index] == "error" {
			assert.Equal(t, "", output)
			assert.Error(t, err, outErrors[index])
		} else {
			assert.Equal(t, output, outputs[index])
			assert.NilError(t, err)
		}
	}
}

func TestSanitizeStrict(t *testing.T) {
	// Make test run in parallel
	t.Parallel()

	// Run through our test cases
	for index, input := range inputs {
		output, err := utils.SanitizeStrict(input)
		if outputsStrict[index] == "error" {
			assert.Equal(t, "", output)
			assert.Error(t, err, outErrorsStrict[index])
		} else {
			assert.Equal(t, output, outputsStrict[index])
			assert.NilError(t, err)
		}
	}
}
