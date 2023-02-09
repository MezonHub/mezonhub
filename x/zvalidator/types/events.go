package types

// sudo module event types
const (
	EventTypeInitZValidator   = "init_zvalidator"
	EventTypeAddZValidator    = "add_zvalidator"
	EventTypeUpdateZValidator = "update_zvalidator"

	AttributeKeyDenom        = "denom"
	AttributeKeyAddresses    = "addresses"
	AttributeKeyAddress      = "address"
	AttributeKeyNewAddress   = "new_address"
	AttributeKeyOldAddress   = "old_address"
	AttributeKeyAddedAddress = "added_address"
	AttributeKeyPoolAddress  = "pool_address"
	AttributeKeyChainEra     = "chain_era"
	AttributeKeyCycleVersion = "cycle_version"
	AttributeKeyCycleNumber  = "cycle_number"
	AttributeKeyCycleSeconds = "cycle_seconds"
)
