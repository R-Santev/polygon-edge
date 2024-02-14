package contractsapi

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestArtifactNotEmpty(t *testing.T) {
	require.NotEmpty(t, BLS.Bytecode)
	require.NotEmpty(t, BLS.DeployedBytecode)
	require.NotEmpty(t, BLS.Abi)

	require.NotEmpty(t, RewardPool.Bytecode)
	require.NotEmpty(t, RewardPool.DeployedBytecode)
	require.NotEmpty(t, RewardPool.Abi)

	require.NotEmpty(t, ValidatorSet.Bytecode)
	require.NotEmpty(t, ValidatorSet.DeployedBytecode)
	require.NotEmpty(t, ValidatorSet.Abi)

	require.NotEmpty(t, LiquidityToken.Bytecode)
	require.NotEmpty(t, LiquidityToken.DeployedBytecode)
	require.NotEmpty(t, LiquidityToken.Abi)

	require.NotEmpty(t, GenesisProxy.Bytecode)
	require.NotEmpty(t, GenesisProxy.DeployedBytecode)
	require.NotEmpty(t, GenesisProxy.Abi)

	require.NotEmpty(t, TransparentUpgradeableProxy.Bytecode)
	require.NotEmpty(t, TransparentUpgradeableProxy.DeployedBytecode)
	require.NotEmpty(t, TransparentUpgradeableProxy.Abi)
}
