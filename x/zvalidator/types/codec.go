package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	zvotetypes "github.com/mezonhub/mezonhub/x/zvote/types"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgInitZValidator{}, "zvalidator/InitZValidator", nil)
	cdc.RegisterConcrete(&UpdateZValidatorProposal{}, "zvalidator/UpdateZValidator", nil)
	cdc.RegisterConcrete(&UpdateZValidatorReportProposal{}, "zvalidator/UpdateZValidatorReport", nil)
	cdc.RegisterConcrete(&MsgSetCycleSeconds{}, "zvalidator/SetCycleSeconds", nil)
	cdc.RegisterConcrete(&MsgSetShuffleSeconds{}, "zvalidator/SetShuffleSeconds", nil)
	cdc.RegisterConcrete(&MsgAddZValidator{}, "zvalidator/AddZValidator", nil)
	cdc.RegisterConcrete(&MsgRmZValidator{}, "zvalidator/RmZValidator", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgInitZValidator{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&UpdateZValidatorProposal{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSetCycleSeconds{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSetShuffleSeconds{},
	)
	registry.RegisterImplementations(
		(*zvotetypes.Content)(nil),
		&UpdateZValidatorProposal{},
	)
	registry.RegisterImplementations(
		(*zvotetypes.Content)(nil),
		&UpdateZValidatorReportProposal{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAddZValidator{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRmZValidator{},
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
