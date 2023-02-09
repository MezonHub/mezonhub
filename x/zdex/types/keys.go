package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// ModuleName defines the module name
	ModuleName = "zdex"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_zdex"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

var (
	SwitchStateClose = []byte{0x00}
	SwitchStateOpen  = []byte{0x01}
)

var (
	SwapPoolStoreKeyPrefix    = []byte{0x01}
	ProviderStoreKeyPrefix    = []byte{0x02}
	ProviderSwitchStoreKey    = []byte{0x03}
	PoolCreatorStoreKeyPrefix = []byte{0x04}
	SwapPoolIndexStoreKey     = []byte{0x05}
)

func SwapPoolStoreKey(denom string) []byte {
	return append(SwapPoolStoreKeyPrefix, []byte(denom)...)
}

func GetLpTokenDenom(index uint32) string {
	return fmt.Sprintf("zdexlp/%d", index)
}

func ProviderStoreKey(addr sdk.AccAddress) []byte {
	return append(ProviderStoreKeyPrefix, addr.Bytes()...)
}

func PoolCreatorStoreKey(addr sdk.AccAddress) []byte {
	return append(PoolCreatorStoreKeyPrefix, addr.Bytes()...)
}
