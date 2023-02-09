package types

import (
	"encoding/hex"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	zvotetypes "github.com/mezonhub/mezonhub/x/zvote/types"
	"github.com/tendermint/tendermint/crypto"
)

const TypeUpdateZValidatorProposal = "update_r_validator"
const TypeUpdateZValidatorReportProposal = "update_r_validator_report"

var _ sdk.Msg = &UpdateZValidatorProposal{}
var _ sdk.Msg = &UpdateZValidatorReportProposal{}
var _ zvotetypes.Content = &UpdateZValidatorProposal{}
var _ zvotetypes.Content = &UpdateZValidatorReportProposal{}

func init() {
	zvotetypes.RegisterProposalType(TypeUpdateZValidatorProposal)
	zvotetypes.RegisterProposalTypeCodec(&UpdateZValidatorProposal{}, "zvalidator/UpdateZValidator")
	zvotetypes.RegisterProposalType(TypeUpdateZValidatorReportProposal)
	zvotetypes.RegisterProposalTypeCodec(&UpdateZValidatorReportProposal{}, "zvalidator/UpdateZValidatorReport")
}

func NewUpdateZValidatorProposal(creator string, denom string, poolAddress, oldAddress string, newAddress string, cycle *Cycle) *UpdateZValidatorProposal {
	msg := UpdateZValidatorProposal{
		Denom:       denom,
		PoolAddress: poolAddress,
		OldAddress:  oldAddress,
		NewAddress:  newAddress,
		Cycle:       cycle,
	}
	msg.setPropId()

	msg.Creator = creator

	return &msg
}

func (p *UpdateZValidatorProposal) setPropId() {
	b, err := p.Marshal()
	if err != nil {
		panic(err)
	}

	p.PropId = hex.EncodeToString(crypto.Sha256(b))
}

func (p *UpdateZValidatorProposal) ProposalRoute() string {
	return ModuleName
}

func (p *UpdateZValidatorProposal) ProposalType() string {
	return TypeUpdateZValidatorProposal
}

func (p *UpdateZValidatorProposal) InFavour() bool {
	return true
}

func (msg *UpdateZValidatorProposal) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *UpdateZValidatorProposal) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *UpdateZValidatorProposal) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if !(msg.Denom == msg.Cycle.Denom && msg.PoolAddress == msg.Cycle.PoolAddress) {
		return fmt.Errorf("denom or pool address not equal")
	}
	return nil
}

func NewUpdateZValidatorReportProposal(creator string, denom string, poolAddress string, cycle *Cycle, status UpdateZValidatorStatus) *UpdateZValidatorReportProposal {
	msg := UpdateZValidatorReportProposal{
		Denom:       denom,
		PoolAddress: poolAddress,
		Cycle:       cycle,
		Status:      status,
	}
	msg.setPropId()

	msg.Creator = creator

	return &msg
}

func (p *UpdateZValidatorReportProposal) setPropId() {
	b, err := p.Marshal()
	if err != nil {
		panic(err)
	}

	p.PropId = hex.EncodeToString(crypto.Sha256(b))
}

func (p *UpdateZValidatorReportProposal) ProposalRoute() string {
	return ModuleName
}

func (p *UpdateZValidatorReportProposal) ProposalType() string {
	return TypeUpdateZValidatorReportProposal
}

func (p *UpdateZValidatorReportProposal) InFavour() bool {
	return true
}

func (msg *UpdateZValidatorReportProposal) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *UpdateZValidatorReportProposal) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *UpdateZValidatorReportProposal) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if !(msg.Denom == msg.Cycle.Denom && msg.PoolAddress == msg.Cycle.PoolAddress) {
		return fmt.Errorf("denom or pool address not equal")
	}
	return nil
}
