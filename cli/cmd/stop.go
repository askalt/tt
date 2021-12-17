package cmd

import (
	"github.com/apex/log"
	"github.com/spf13/cobra"
	"github.com/tarantool/tt/cli/context"
	"github.com/tarantool/tt/cli/modules"
	"github.com/tarantool/tt/cli/running"
)

// NewStopCmd creates stop command.
func NewStopCmd() *cobra.Command {
	var stopCmd = &cobra.Command{
		Use:   "stop [INSTANCE_NAME]",
		Short: "Stop tarantool instance",
		Run: func(cmd *cobra.Command, args []string) {
			err := modules.RunCmd(&ctx, cmd.Name(), &modulesInfo, internalStopModule, args)
			if err != nil {
				log.Fatalf(err.Error())
			}
		},
	}

	return stopCmd
}

// internalStopModule is a default stop module.
func internalStopModule(ctx *context.Ctx, args []string) error {
	cliOpts, err := modules.GetCliOpts(ctx.Cli.ConfigPath)
	if err != nil {
		return err
	}

	if err = running.FillCtx(cliOpts, ctx, args); err != nil {
		return err
	}

	if err = running.Stop(ctx); err != nil {
		return err
	}

	return nil
}
