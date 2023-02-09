package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRmZValidator = "rm_r_validator"

var _ sdk.Msg = &MsgRmZValidator{}

func NewMsgRmZValidator(creator string, denom string, poolAddress string, oldAddress string, newAddress string) *MsgRmZValidator {
	return &MsgRmZValidator{
		Creator:     creator,
		Denom:       denom,
		PoolAddress: poolAddress,
		OldAddress:  oldAddress,
		NewAddress:  newAddress,
	}
}

func (msg *MsgRmZValidator) Route() string {
	return RouterKey
}

func (msg *MsgRmZValidator) Type() string {
	return TypeMsgRmZValidator
}

func (msg *MsgRmZValidator) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRmZValidator) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRmZValidator) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
