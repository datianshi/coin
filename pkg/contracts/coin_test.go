package contracts

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
)

func TestDeployInbox(t *testing.T) {

	//Setup simulated block chain
	key, _ := crypto.GenerateKey()
	auth := bind.NewKeyedTransactor(key)
	auth.GasLimit = 65964
	auth.GasPrice = big.NewInt(984375000)

	key, _ = crypto.GenerateKey()
	to := bind.NewKeyedTransactor(key)

	alloc := make(core.GenesisAlloc)
	alloc[auth.From] = core.GenesisAccount{Balance: big.NewInt(875000000000000)}
	alloc[to.From] = core.GenesisAccount{Balance: big.NewInt(10000000000)}
	blockchain := backends.NewSimulatedBackend(alloc, 65964)

	//Deploy contract
	address, _, coin, err := DeployCoin(
		auth,
		blockchain,
	)
	if err != nil {
		t.Fatalf("Failed to deploy the Inbox contract: %v", err)
	}
	blockchain.Commit()

	if len(address.Bytes()) == 0 {
		t.Error("Expected a valid deployment address. Received empty address byte array instead")
	}

	tx, err := coin.Mint(auth, to.From, big.NewInt(200))

	if err != nil {
		t.Fatalf("Failed to mint: %v", err)
	}

	data, err := tx.MarshalJSON()
	if err != nil {
		t.Fatalf("Failed to marshar transaction: %v", err)
	}

	fmt.Println(string(data))

}
