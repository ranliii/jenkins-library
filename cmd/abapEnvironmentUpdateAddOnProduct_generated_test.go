//go:build unit
// +build unit

package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAbapEnvironmentUpdateAddOnProductCommand(t *testing.T) {
	t.Parallel()

	testCmd := AbapEnvironmentUpdateAddOnProductCommand()

	// only high level testing performed - details are tested in step generation procedure
	assert.Equal(t, "abapEnvironmentUpdateAddOnProduct", testCmd.Use, "command name incorrect")

}