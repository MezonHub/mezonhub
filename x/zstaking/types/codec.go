package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgSetInflationBase{}, "zstaking/SetInflationBase", nil)
	cdc.RegisterConcrete(&MsgAddValToWhitelist{}, "zstaking/AddValToWhitelist", nil)
	cdc.RegisterConcrete(&MsgToggleValidatorWhitelistSwitch{}, "zstaking/ToggleValidatorWhitelistSwitch", nil)
	cdc.RegisterConcrete(&MsgWithdraw{}, "zstaking/Withdraw", nil)
	cdc.RegisterConcrete(&MsgAddDelegatorToWhitelist{}, "zstaking/AddDelegatorToWhitelist", nil)
	cdc.RegisterConcrete(&MsgToggleDelegatorWhitelistSwitch{}, "zstaking/ToggleDelegatorWhitelistSwitch", nil)
	cdc.RegisterConcrete(&MsgProvideToken{}, "zstaking/ProvideToken", nil)
	cdc.RegisterConcrete(&MsgRmDelegatorFromWhitelist{}, "zstaking/RmDelegatorFromWhitelist", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSetInflationBase{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAddValToWhitelist{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgToggleValidatorWhitelistSwitch{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgWithdraw{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAddDelegatorToWhitelist{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgToggleDelegatorWhitelistSwitch{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgProvideToken{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRmDelegatorFromWhitelist{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(Amino)
)

func init() {
	RegisterCodec(Amino)
	cryptocodec.RegisterCrypto(Amino)
	sdk.RegisterLegacyAminoCodec(Amino)
	Amino.Seal()
}
