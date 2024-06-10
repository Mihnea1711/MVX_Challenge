package wallet

import (
	"github.com/multiversx/mx-chain-logger-go"
	"github.com/multiversx/mx-sdk-go/core"
	"github.com/multiversx/mx-sdk-go/interactors"
)

// GenerateAddressFromPrivateKey loads a private key from a JSON file and generates the corresponding address.
func GenerateAddressFromPrivateKey(filename, password string, log logger.Logger) (core.AddressHandler, error) {
	wallet := interactors.NewWallet()

	// Load private key from JSON file
	privateKey, err := wallet.LoadPrivateKeyFromJsonFile(filename, password)
	if err != nil {
		log.Error("Error loading private key", "error", err)
		return nil, err
	}

	// Generate address from private key
	address, err := wallet.GetAddressFromPrivateKey(privateKey)
	if err != nil {
		log.Error("Error generating address", "error", err)
		return nil, err
	}

	log.Info("Address generated successfully")

	log.Info(address.AddressAsBech32String())
	return address, nil
}
