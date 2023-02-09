package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mezonhub/mezonhub/x/zvalidator/types"
)

// old val must exist && new val may exist or not exist in selectedZValidator
// pool will redelegate all delegation from old to new val
func (k Keeper) ProcessUpdateZValidatorProposal(ctx sdk.Context, p *types.UpdateZValidatorProposal) error {

	oldVal := types.ZValidator{
		Denom:       p.Denom,
		PoolAddress: p.PoolAddress,
		ValAddress:  p.OldAddress,
	}
	newVal := types.ZValidator{
		Denom:       p.Denom,
		PoolAddress: p.PoolAddress,
		ValAddress:  p.NewAddress,
	}
	if !k.HasSelectedZValidator(ctx, &oldVal) {
		return types.ErrZValidatorNotExist
	}
	if oldVal.ValAddress == newVal.ValAddress {
		return types.ErrOldEqualNewZValidator
	}

	if err := k.ZBankKeeper.CheckValAddress(ctx, p.Denom, p.OldAddress); err != nil {
		return err
	}
	if err := k.ZBankKeeper.CheckValAddress(ctx, p.Denom, p.NewAddress); err != nil {
		return err
	}
	if err := k.ZBankKeeper.CheckAccAddress(ctx, p.Denom, p.PoolAddress); err != nil {
		return err
	}
	cycleSeconds := k.GetCycleSeconds(ctx, p.Denom)
	if cycleSeconds.Version != p.Cycle.Version {
		return types.ErrCycleVersionNotMatch
	}

	latestVotedCycle := k.GetLatestVotedCycle(ctx, p.Denom, p.PoolAddress)
	if !(p.Cycle.Version > latestVotedCycle.Version || (p.Cycle.Version == latestVotedCycle.Version && p.Cycle.Number > latestVotedCycle.Number)) {
		return types.ErrCycleBehindLatestCycle
	}
	latestDealedCycle := k.GetLatestDealedCycle(ctx, p.Denom, p.PoolAddress)
	if latestDealedCycle.Number != latestVotedCycle.Number || latestDealedCycle.Version != latestVotedCycle.Version {
		return types.ErrLatestVotedCycleNotDealed
	}
	snapShots := k.ledgerKeeper.CurrentEraSnapshots(ctx, p.Denom)
	if len(snapShots.ShotIds) > 0 {
		return types.ErrLedgerIsBusyWithEra
	}
	chainEra, found := k.ledgerKeeper.GetChainEra(ctx, p.Denom)
	if !found {
		return types.ErrLedgerChainEraNotExist
	}

	k.SetDealingZValidator(ctx, &types.DealingZValidator{
		Denom:         p.Denom,
		PoolAddress:   p.PoolAddress,
		OldValAddress: p.OldAddress,
		NewValAddress: p.NewAddress,
	})
	k.SetLatestVotedCycle(ctx, p.Cycle)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeUpdateZValidator,
			sdk.NewAttribute(types.AttributeKeyDenom, p.Denom),
			sdk.NewAttribute(types.AttributeKeyPoolAddress, p.PoolAddress),
			sdk.NewAttribute(types.AttributeKeyChainEra, fmt.Sprintf("%d", chainEra.Era)),
			sdk.NewAttribute(types.AttributeKeyOldAddress, p.OldAddress),
			sdk.NewAttribute(types.AttributeKeyNewAddress, p.NewAddress),
			sdk.NewAttribute(types.AttributeKeyCycleVersion, fmt.Sprintf("%d", p.Cycle.Version)),
			sdk.NewAttribute(types.AttributeKeyCycleNumber, fmt.Sprintf("%d", p.Cycle.Number)),
			sdk.NewAttribute(types.AttributeKeyCycleSeconds, fmt.Sprintf("%d", cycleSeconds.Seconds)),
		),
	)

	return nil
}

func (k Keeper) ProcessUpdateZValidatorReportProposal(ctx sdk.Context, p *types.UpdateZValidatorReportProposal) error {
	latestVotedCycle := k.GetLatestVotedCycle(ctx, p.Denom, p.PoolAddress)
	if !(p.Cycle.Version == latestVotedCycle.Version && p.Cycle.Number == latestVotedCycle.Number) {
		return types.ErrReportCycleNotMatchLatestVotedCycle
	}
	dealingZValidator, found := k.GetDealingZValidator(ctx, p.Denom, p.PoolAddress)
	if !found {
		return types.ErrDealingZvalidatorNotFound
	}

	// should update zvalidator when redelegate success
	if p.Status == types.UpdateZValidatorStatusSuccess {
		k.RemoveSelectedZValidator(ctx, &types.ZValidator{
			Denom:       dealingZValidator.Denom,
			PoolAddress: dealingZValidator.PoolAddress,
			ValAddress:  dealingZValidator.OldValAddress,
		})
		k.AddSelectedZValidator(ctx, &types.ZValidator{
			Denom:       dealingZValidator.Denom,
			PoolAddress: dealingZValidator.PoolAddress,
			ValAddress:  dealingZValidator.NewValAddress,
		})
	}

	k.RemoveDealingZValidator(ctx, p.Denom, p.PoolAddress)
	k.SetLatestDealedCycle(ctx, p.Cycle)
	return nil
}
