package command

import (
	"github.com/deniskorbakov/skeleton-cli-go/configs/constname"
	"github.com/deniskorbakov/skeleton-cli-go/internal/component/form"
	"github.com/deniskorbakov/skeleton-cli-go/internal/component/output"
	"github.com/spf13/cobra"
)

// buildCmd Command for build pipeline
var buildCmd = &cobra.Command{
	Use:   constname.UseExampleCmd,
	Short: constname.UseExampleCmd,
	Long:  constname.LongExampleCmd,
	RunE: func(cmd *cobra.Command, args []string) error {
		fields, err := form.Run()
		if err != nil {
			return err
		}

		output.Green("Success green output: ", fields.RemoteUser)
		output.Red("This command will be run successfully")

		return nil
	},
}
