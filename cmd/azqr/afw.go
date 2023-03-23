package azqr

import (
	"github.com/cmendible/azqr/internal/scanners"
	"github.com/cmendible/azqr/internal/scanners/afw"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(afwCmd)
}

var afwCmd = &cobra.Command{
	Use:   "afw",
	Short: "Scan Azure Firewall",
	Long:  "Scan Azure Firewall",
	Run: func(cmd *cobra.Command, args []string) {
		serviceScanners := []scanners.IAzureScanner{
			&afw.FirewallScanner{},
		}

		scan(cmd, serviceScanners)
	},
}