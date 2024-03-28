package testimpl

import (
	"os"
	"testing"

	"github.com/gruntwork-io/terratest/modules/azure"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/launchbynttdata/lcaf-component-terratest/types"
	"github.com/stretchr/testify/assert"
)

func TestComposableComplete(t *testing.T, ctx types.TestContext) {
	subscriptionId := os.Getenv("ARM_SUBSCRIPTION_ID")

	if len(subscriptionId) == 0 {
		t.Fatal("ARM_SUBSCRIPTION_ID environment variable is not set")
	}

	rgName := terraform.Output(t, ctx.TerratestTerraformOptions(), "name")
	rgLocation := terraform.Output(t, ctx.TerratestTerraformOptions(), "location")

	t.Run("TestAlwaysSucceeds", func(t *testing.T) {
		assert.Equal(t, "foo", "foo", "Should always be the same!")
		assert.NotEqual(t, "foo", "bar", "Should never be the same!")
	})

	t.Run("ResourceGroupWasCreated", func(t *testing.T) {
		assert.True(t, azure.ResourceGroupExists(t, rgName, subscriptionId), "Resource group didn't exist!")
		actualResourceGroup := azure.GetAResourceGroup(t, rgName, subscriptionId)
		assert.Equal(t, rgName, *actualResourceGroup.Name, "Resource group actual name didn't match expected")
		assert.Equal(t, rgLocation, *actualResourceGroup.Location, "Resource group actual location didn't match expected")
	})
}
