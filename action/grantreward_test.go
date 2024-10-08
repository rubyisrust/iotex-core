package action

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGrandReward(t *testing.T) {
	require := require.New(t)
	tests := []struct {
		rewardType int
		height     uint64
	}{
		{BlockReward, 100},
		{EpochReward, 200},
	}
	for _, test := range tests {
		g := &GrantReward{
			rewardType: test.rewardType,
			height:     test.height,
		}
		require.Equal(test.rewardType, g.RewardType())
		require.Equal(test.height, g.Height())
		require.NoError(g.SanityCheck())
		require.NoError(g.LoadProto(g.Proto()))
		intrinsicGas, err := g.IntrinsicGas()
		require.NoError(err)
		require.Equal(uint64(0), intrinsicGas)
		cost, err := g.Cost()
		require.NoError(err)
		require.Equal(big.NewInt(0), cost)
	}
}
