package version

import (
	"github.com/edgexfoundry-holding/edgex-cli/pkg/cmd/version"

	"github.com/spf13/cobra"
)

var Version = "master"
// NewCommand returns the version command
func NewCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "version",
		Short: "Version command",
		Long:  `Outputs the current versions of EdgeX CLI and EdgeX Foundry.`,
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			cmd.Println("EdgeX CLI version: ", Version)
			edgeXVersion, err := version.GetEdgeXVersion("48080")
			if err != nil {
				return err
			}
			cmd.Println("EdgeX Foundry version: ", edgeXVersion.Version)
			return
		},
	}

	return cmd
}
