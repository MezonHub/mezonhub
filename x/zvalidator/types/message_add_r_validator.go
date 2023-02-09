package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAddZValidator = "add_r_validator"

var _ sdk.Msg = &MsgAddZValidator{}

func NewMsgAddZValidator(creator string, denom string, poolAddress string, valAddress string) *MsgAddZValidator {
	return &MsgAddZValidator{
		Creator:     creator,
		Denom:       denom,
		PoolAddress: poolAddress,
		ValAddress:  valAddress,
	}
}

func (msg *MsgAddZValidator) Route() string {
	return RouterKey
}

func (msg *MsgAddZValidator) Type() string {
	return TypeMsgAddZValidator
}

func (msg *MsgAddZValidator) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAddZValidator) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddZValidator) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
