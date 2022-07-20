package app

import (
	cmdcfg "a4chain/cmd/a4chaind/config"
	sdk "github.com/cosmos/cosmos-sdk/types"
	ethcfg "github.com/tharsis/ethermint/cmd/config"
)

func SetConfig() {
	config := sdk.GetConfig()
	cmdcfg.SetBech32Prefixes(config)
	ethcfg.SetBip44CoinType(config)
	// Make sure address is compatible with ethereum
	config.SetAddressVerifier(VerifyAddressFormat)
	config.Seal()
}
