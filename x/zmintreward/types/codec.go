package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgAddMintRewardAct{}, "zmintreward/AddMintRewardAct", nil)
	cdc.RegisterConcrete(&MsgUpdateMintRewardAct{}, "zmintreward/UpdateMintRewardAct", nil)
	cdc.RegisterConcrete(&MsgClaimMintReward{}, "zmintreward/ClaimMintReward", nil)
	cdc.RegisterConcrete(&MsgProvideRewardToken{}, "zmintreward/ProvideRewardToken", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAddMintRewardAct{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateMintRewardAct{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgClaimMintReward{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgProvideRewardToken{},
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
