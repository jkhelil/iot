package server

import (
	"github.com/golang/glog"
	"github.com/spf13/cobra"
)

func RegisterCreateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "server",
		Short: "create server",
		Run: func(cmd *cobra.Command, args []string) {
			err := CreateServer(args)
			if err != nil {
				glog.Exitf("%s\n", err.Error())
			}
		},
	}

	cmd.Flags().IntVarP(&times, "times", "n", 1, "times to echo")

	return cmd
}
