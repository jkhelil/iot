package cmd

import (
	"github.com/jkhelil/iot/cluster"
	"github.com/jkhelil/iot/server"
	"github.com/spf13/cobra"
)

func RegisterCreateCommand() *cobra.Command {
	create := &cobra.Command{
		Use:   "create",
		Short: "Create cluster/node/etc",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Usage()
		},
	}
	create.AddCommand(server.RegisterCreateCommand())
	create.AddCommand(cluster.RegisterCreateCommand())

	return create
}
