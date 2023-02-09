package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreatePool{}, "zdex/CreatePool", nil)
	cdc.RegisterConcrete(&MsgAddLiquidity{}, "zdex/AddLiquidity", nil)
	cdc.RegisterConcrete(&MsgSwap{}, "zdex/Swap", nil)
	cdc.RegisterConcrete(&MsgRemoveLiquidity{}, "zdex/RemoveLiquidity", nil)
	cdc.RegisterConcrete(&MsgToggleProviderSwitch{}, "zdex/ToggleProviderSwitch", nil)
	cdc.RegisterConcrete(&MsgAddProvider{}, "zdex/AddProvider", nil)
	cdc.RegisterConcrete(&MsgRmProvider{}, "zdex/RmProvider", nil)
	cdc.RegisterConcrete(&MsgAddPoolCreator{}, "zdex/AddPoolCreator", nil)
	cdc.RegisterConcrete(&MsgRmPoolCreator{}, "zdex/RmPoolCreator", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreatePool{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAddLiquidity{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSwap{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRemoveLiquidity{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgToggleProviderSwitch{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAddProvider{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRmProvider{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAddPoolCreator{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRmPoolCreator{},
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
