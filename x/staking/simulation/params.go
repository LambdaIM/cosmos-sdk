package simulation

// DONTCOVER

import (
	"fmt"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// ParamChanges defines the parameters that can be modified by param change proposals
// on the simulation
func ParamChanges(cdc *codec.Codec, r *rand.Rand) []simulation.ParamChange {
	return []simulation.ParamChange{
		simulation.NewSimParamChange("staking", "MaxValidators", "",
			func(r *rand.Rand) string {
				return fmt.Sprintf("%d", GenMaxValidators(cdc, r))
			},
		),
		simulation.NewSimParamChange("staking", "UnbondingTime", "",
			func(r *rand.Rand) string {
				return fmt.Sprintf("\"%d\"", GenUnbondingTime(cdc, r))
			},
		),
	}
}