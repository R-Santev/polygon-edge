package polybft

import (
	// H_MODIFY: Registration module is moved to sidechain
	"github.com/0xPolygon/polygon-edge/command/sidechain/rewards"
	"github.com/0xPolygon/polygon-edge/command/sidechain/unstaking"
	"github.com/0xPolygon/polygon-edge/command/sidechain/withdraw"
	sidechainWithdraw "github.com/0xPolygon/polygon-edge/command/sidechain/withdraw"
	"github.com/spf13/cobra"
)

func GetCommand() *cobra.Command {
	polybftCmd := &cobra.Command{
		Use:   "polybft",
		Short: "Polybft command",
	}

	// H_MODIFY: TODO: Modify commands
	polybftCmd.AddCommand(
		// sidechain (validator set) command to unstake on child chain
		unstaking.GetCommand(),
		// sidechain (validator set) command to withdraw stake on child chain
		sidechainWithdraw.GetCommand(),
		// sidechain (reward pool) command to withdraw pending rewards
		rewards.GetCommand(),
		// rootchain (stake manager) command to withdraw stake
		withdraw.GetCommand(),
	)

	return polybftCmd
}
