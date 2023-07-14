package validator

import (
	"math/big"
	"testing"

	"github.com/0xPolygon/polygon-edge/types"
	"github.com/stretchr/testify/require"
)

func TestValidatorSet_HasQuorum(t *testing.T) {
	t.Parallel()

	// enough signers for quorum (2/3 super-majority of validators are signers)
	validators := NewTestValidatorsWithAliases(t, []string{"A", "B", "C", "D", "E", "F", "G"})
	vs := validators.ToValidatorSet()

	signers := make(map[types.Address]struct{})

	validators.IterAcct([]string{"A", "B", "C", "D", "E"}, func(v *TestValidator) {
		signers[v.Address()] = struct{}{}
	})

	require.True(t, vs.HasQuorum(signers))

	// not enough signers for quorum (less than 2/3 super-majority of validators are signers)
	signers = make(map[types.Address]struct{})

	validators.IterAcct([]string{"A", "B", "C", "D"}, func(v *TestValidator) {
		signers[v.Address()] = struct{}{}
	})
	require.False(t, vs.HasQuorum(signers))
}

func TestValidatorSet_getQuorumSize(t *testing.T) {
	t.Parallel()

	cases := []struct {
		totalVotingPower   int64
		expectedQuorumSize int64
	}{
		{9, 9},
		{10, 7},
		{12, 8},
		{13, 8},
		{50, 31},
		{100, 62},
		{100000, 61401},
		{2500000, 1535001},
		{7528364981, 4622416099},
		{10000000000, 6140000001},
	}

	for _, c := range cases {
		quorumSize := getQuorumSize(big.NewInt(c.totalVotingPower))
		require.Equal(t, c.expectedQuorumSize, quorumSize.Int64())
	}
}
