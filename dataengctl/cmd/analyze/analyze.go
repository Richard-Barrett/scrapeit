package analyze

import (
	"github.com/Mirantis/dataeng/pkg/client"
	"github.com/Mirantis/dataeng/cmd/analyze/jira"

	"github.com/spf13/cobra"
)

// Wrapper for Jira Client
func NewAnalyzeCmd(dataClient *client.DataClient) *cobra.Command {
	//salesforcecfg := &sfconfig.Config{}
	//jiracfg := &jiracfg.Config
	cmd := &cobra.Command{
		Use:   "analyze",
		Short: "Analyze something",
	}
	
	


	// Wrapper for Output Flag to All Base Commands
	/*func OutPutCmd() *cobra.Command {
		var output string
		output := &output{}
		cmd := &cobra.Command{
			Use:	"output"
		},
	}*/

	// Universal Flag Commands
	//cmd.PersistentFlags().StringVar(&output), "output", "", "csv,dataframe,json,yaml")

	// Universal Base Commands
	cmd.AddCommand(jira.NewAnalyzeJiraCommand(dataClient))

	return cmd
}
