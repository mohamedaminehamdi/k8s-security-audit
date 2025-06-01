package diagnose

import (
	"github.com/spf13/cobra"

	"github.com/mohamedaminehamdi/kubepilot/cmd"
)

func New(commonOpts cmd.Options) *cobra.Command {
	opts := Options{
		common: commonOpts,
	}
	c := &cobra.Command{
		Use:          "diagnose TYPE NAME",
		Short:        "Diagnose a resource",
		SilenceUsage: true,
		Args:         cobra.ExactArgs(2),
		RunE:         opts.Run,
	}
	opts.AddFlags(c.Flags())
	return c
}
