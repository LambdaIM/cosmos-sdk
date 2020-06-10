package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/params"
)

// type declaration for parameters
func ParamKeyTable() params.KeyTable {
	return params.NewKeyTable(
		ParamStoreKeyCommunityTax, sdk.Dec{},
		ParamStoreKeyBaseProposerReward, sdk.Dec{},
		ParamStoreKeyBonusProposerReward, sdk.Dec{},
		ParamStoreKeyPdpReward, sdk.Dec{},
		ParamStoreKeyPdpProposerReward, sdk.Dec{},
		ParamStoreKeyWithdrawAddrEnabled, false,
		ParamStoreKeyPickedAssetMinerReward, sdk.Dec{},
		ParamStoreKeyRewardSlashFraction, sdk.Dec{},
		ParamStoreKeyMaxRewardSlashFraction, sdk.Dec{},
		ParamStoreKeyRewardSlashPeriod, int64(0),
	)
}

// returns the current CommunityTax rate from the global param store
// nolint: errcheck
func (k Keeper) GetCommunityTax(ctx sdk.Context) sdk.Dec {
	var percent sdk.Dec
	k.paramSpace.Get(ctx, ParamStoreKeyCommunityTax, &percent)
	return percent
}

// nolint: errcheck
func (k Keeper) SetCommunityTax(ctx sdk.Context, percent sdk.Dec) {
	k.paramSpace.Set(ctx, ParamStoreKeyCommunityTax, &percent)
}

// returns the current BaseProposerReward rate from the global param store
// nolint: errcheck
func (k Keeper) GetBaseProposerReward(ctx sdk.Context) sdk.Dec {
	var percent sdk.Dec
	k.paramSpace.Get(ctx, ParamStoreKeyBaseProposerReward, &percent)
	return percent
}

// nolint: errcheck
func (k Keeper) SetBaseProposerReward(ctx sdk.Context, percent sdk.Dec) {
	k.paramSpace.Set(ctx, ParamStoreKeyBaseProposerReward, &percent)
}

// returns the current BaseProposerReward rate from the global param store
// nolint: errcheck
func (k Keeper) GetBonusProposerReward(ctx sdk.Context) sdk.Dec {
	var percent sdk.Dec
	k.paramSpace.Get(ctx, ParamStoreKeyBonusProposerReward, &percent)
	return percent
}

// nolint: errcheck
func (k Keeper) SetBonusProposerReward(ctx sdk.Context, percent sdk.Dec) {
	k.paramSpace.Set(ctx, ParamStoreKeyBonusProposerReward, &percent)
}

// returns the current WithdrawAddrEnabled
// nolint: errcheck
func (k Keeper) GetWithdrawAddrEnabled(ctx sdk.Context) bool {
	var enabled bool
	k.paramSpace.Get(ctx, ParamStoreKeyWithdrawAddrEnabled, &enabled)
	return enabled
}

// nolint: errcheck
func (k Keeper) SetWithdrawAddrEnabled(ctx sdk.Context, enabled bool) {
	k.paramSpace.Set(ctx, ParamStoreKeyWithdrawAddrEnabled, &enabled)
}

func (k Keeper) SetPdpProposerReward(ctx sdk.Context, percent sdk.Dec) {
	k.paramSpace.Set(ctx, ParamStoreKeyPdpProposerReward, &percent)
}

func (k Keeper) GetPdpProposerReward(ctx sdk.Context) sdk.Dec {
	var percent sdk.Dec
	k.paramSpace.Get(ctx, ParamStoreKeyPdpProposerReward, &percent)
	return percent
}

func (k Keeper) SetPdpReward(ctx sdk.Context, percent sdk.Dec) {
	k.paramSpace.Set(ctx, ParamStoreKeyPdpReward, &percent)
}

func (k Keeper) GetPdpReward(ctx sdk.Context) sdk.Dec {
	var percent sdk.Dec
	k.paramSpace.Get(ctx, ParamStoreKeyPdpReward, &percent)
	return percent
}

func (k Keeper) SetPickedAssetMinerReward(ctx sdk.Context, percent sdk.Dec) {
	k.paramSpace.Set(ctx, ParamStoreKeyPickedAssetMinerReward, &percent)
}

func (k Keeper) GetPickedAssetMinerReward(ctx sdk.Context) sdk.Dec {
	var percent sdk.Dec
	k.paramSpace.Get(ctx, ParamStoreKeyPickedAssetMinerReward, &percent)
	return percent
}

func (k Keeper) SetRewardSlashFraction(ctx sdk.Context, percent sdk.Dec) {
	k.paramSpace.Set(ctx, ParamStoreKeyRewardSlashFraction, &percent)
}

func (k Keeper) GetRewardSlashFraction(ctx sdk.Context) sdk.Dec {
	var percent sdk.Dec
	k.paramSpace.Get(ctx, ParamStoreKeyRewardSlashFraction, &percent)
	return percent
}

func (k Keeper) SetMaxRewardSlashFraction(ctx sdk.Context, percent sdk.Dec) {
	k.paramSpace.Set(ctx, ParamStoreKeyMaxRewardSlashFraction, &percent)
}

func (k Keeper) GetMaxRewardSlashFraction(ctx sdk.Context) sdk.Dec {
	var percent sdk.Dec
	k.paramSpace.Get(ctx, ParamStoreKeyMaxRewardSlashFraction, &percent)
	return percent
}

func (k Keeper) SetRewardSlashPeriod(ctx sdk.Context, period int64) {
	k.paramSpace.Set(ctx, ParamStoreKeyRewardSlashPeriod, period)
}

func (k Keeper) GetRewardSlashPeriod(ctx sdk.Context) int64 {
	var period int64
	k.paramSpace.Get(ctx, ParamStoreKeyRewardSlashPeriod, &period)
	return period
}
