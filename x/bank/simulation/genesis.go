package simulation

// DONTCOVER

import (
	"fmt"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/bank/internal/types"
)

// Simulation parameter constants
const (
	SendEnabled = "send_enabled"
)

// GenSendEnabled randomized SendEnabled
func GenSendEnabled(r *rand.Rand) bool {
	return r.Int63n(2) == 0
}

// RandomizedGenState generates a random GenesisState for bank
func RandomizedGenState(input *module.GeneratorInput) {

	var sendEnabled bool

	input.AppParams.GetOrGenerate(input.Cdc, SendEnabled, &sendEnabled, input.R,
		func(r *rand.Rand) { sendEnabled = GenSendEnabled(input.R) })

	bankGenesis := types.NewGenesisState(sendEnabled)

	fmt.Printf("Selected randomly generated bank parameters:\n%s\n", codec.MustMarshalJSONIndent(input.Cdc, bankGenesis))
	input.GenState[types.ModuleName] = input.Cdc.MustMarshalJSON(bankGenesis)
}