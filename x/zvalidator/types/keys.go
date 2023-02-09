package types

const (
	// ModuleName defines the module name
	ModuleName = "zvalidator"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_zvalidator"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

var (
	DefaultCycleSeconds   = uint64(60 * 10)
	DefaultShuffleSeconds = uint64(60 * 60 * 24 * 14)
)

var (
	SelectedZValidatorStoreKeyPrefix = []byte{0x01}
	LatestVotedCycleStoreKeyPrefix   = []byte{0x02}
	CycleSecondsStoreKeyPrefix       = []byte{0x03}
	ShuffleSecondsStoreKeyPrefix     = []byte{0x04}
	LatestDealedCycleStoreKeyPrefix  = []byte{0x05}
	DealingZValidatorStoreKeyPrefix  = []byte{0x06}
)

// prefix + denomLen + denom + poolAddressLen + poolAddress + zValidatorAddressLen + zValidatorAddress
func SelectedRValdidatorStoreKey(denom, poolAddress, zValidatorAddress string) []byte {
	denomLen := len([]byte(denom))
	poolAddressLen := len([]byte(poolAddress))
	zValidatorAddressLen := len([]byte(zValidatorAddress))

	key := make([]byte, 1+1+denomLen+1+poolAddressLen+1+zValidatorAddressLen)
	copy(key[0:], SelectedZValidatorStoreKeyPrefix)
	key[1] = byte(denomLen)
	copy(key[2:], []byte(denom))
	key[2+denomLen] = byte(poolAddressLen)
	copy(key[2+denomLen+1:], []byte(poolAddress))
	key[2+denomLen+1+poolAddressLen] = byte(zValidatorAddressLen)
	copy(key[2+denomLen+1+poolAddressLen+1:], []byte(zValidatorAddress))
	return key
}

// prefix + denomLen + denom + poolAddressLen + poolAddress
func LatestVotedCycleStoreKey(denom, poolAddress string) []byte {
	denomLen := len([]byte(denom))
	poolAddressLen := len([]byte(poolAddress))

	key := make([]byte, 1+1+denomLen+1+poolAddressLen)
	copy(key[0:], LatestVotedCycleStoreKeyPrefix)
	key[1] = byte(denomLen)
	copy(key[2:], []byte(denom))
	key[2+denomLen] = byte(poolAddressLen)
	copy(key[2+denomLen+1:], []byte(poolAddress))

	return key
}

// prefix + denomLen + denom + poolAddressLen + poolAddress
func LatestDealedCycleStoreKey(denom, poolAddress string) []byte {
	denomLen := len([]byte(denom))
	poolAddressLen := len([]byte(poolAddress))

	key := make([]byte, 1+1+denomLen+1+poolAddressLen)
	copy(key[0:], LatestDealedCycleStoreKeyPrefix)
	key[1] = byte(denomLen)
	copy(key[2:], []byte(denom))
	key[2+denomLen] = byte(poolAddressLen)
	copy(key[2+denomLen+1:], []byte(poolAddress))

	return key
}

func CycleSecondsStoreKey(denom string) []byte {
	return append(CycleSecondsStoreKeyPrefix, []byte(denom)...)
}

func ShuffleSecondsStoreKey(denom string) []byte {
	return append(ShuffleSecondsStoreKeyPrefix, []byte(denom)...)
}

// prefix + denomLen + denom + poolAddressLen + poolAddress
func DealingZValidatorStoreKey(denom, poolAddress string) []byte {
	denomLen := len([]byte(denom))
	poolAddressLen := len([]byte(poolAddress))

	key := make([]byte, 1+1+denomLen+1+poolAddressLen)
	copy(key[0:], DealingZValidatorStoreKeyPrefix)
	key[1] = byte(denomLen)
	copy(key[2:], []byte(denom))
	key[2+denomLen] = byte(poolAddressLen)
	copy(key[2+denomLen+1:], []byte(poolAddress))

	return key
}
