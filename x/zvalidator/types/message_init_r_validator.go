package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgInitZValidator = "init_r_validator"

var _ sdk.Msg = &MsgInitZValidator{}

func NewMsgInitZValidator(creator string, denom, poolAddress string, addressList []string) *MsgInitZValidator {
	return &MsgInitZValidator{
		Creator:        creator,
		Denom:          denom,
		PoolAddress:    poolAddress,
		ValAddressList: addressList,
	}
}

func (msg *MsgInitZValidator) Route() string {
	return RouterKey
}

func (msg *MsgInitZValidator) Type() string {
	return TypeMsgInitZValidator
}

func (msg *MsgInitZValidator) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgInitZValidator) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgInitZValidator) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if err := sdk.ValidateDenom(msg.Denom); err != nil {
		return err
	}

	if len(msg.ValAddressList) == 0 {
		return fmt.Errorf("address list is empty")
	}
	return nil
}
