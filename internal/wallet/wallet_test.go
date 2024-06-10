package wallet

import (
	"MultiversX_Challenge/internal/logging"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateAddressFromPrivateKey(t *testing.T) {
	// Initialize the logger
	logger := logging.GetLogger()

	// Hardcoded key paths and passwords
	key1Path := "../../keys/erd136ag3fktmudhc52ssh4a46f2av0rtnywh89qlf39nsjgh6nqyzlqj56m9d.json"
	key1Password := "SomeSecretPass@2000"
	key2Path := "../../keys/erd1wr5ffdllrmuq75x7stntkg27t8dk2quncd7dcdjd5ja0k4xjuegquq0tak.json"
	key2Password := "SomeOtherSecretPass@2000"

	// Call the function to generate the address for key 1
	address1, err := GenerateAddressFromPrivateKey(key1Path, key1Password, logger)
	if err != nil {
		t.Errorf("Error generating address for key 1: %v", err)
	}

	// Assert that the address for key 1 is not empty
	assert.NotEmpty(t, address1)

	// Call the function to generate the address for key 2
	address2, err := GenerateAddressFromPrivateKey(key2Path, key2Password, logger)
	if err != nil {
		t.Errorf("Error generating address for key 2: %v", err)
	}

	// Assert that the address for key 2 is not empty
	assert.NotEmpty(t, address2)
}
