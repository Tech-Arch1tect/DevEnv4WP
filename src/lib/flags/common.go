package flags

import (
	"github.com/spf13/cobra"
)

var Debug bool
var Dev bool

func RegisterCommonFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().BoolVarP(&Debug, "debug", "d", false, "Enable debug mode")
	cmd.PersistentFlags().BoolVarP(&Dev, "dev", "D", false, "Enable dev mode")
}
