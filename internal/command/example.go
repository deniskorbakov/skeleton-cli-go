package command

import (
	"github.com/deniskorbakov/skeleton-cli-go/configs/constname"
	"github.com/deniskorbakov/skeleton-cli-go/internal/component/form"
	"github.com/deniskorbakov/skeleton-cli-go/internal/component/output"
	"github.com/spf13/cobra"
)

// exampleCmd Command for build pipeline
var exampleCmd = &cobra.Command{
	Use:   constname.UseExampleCmd,
	Short: constname.ShortExampleCmd,
	Long:  constname.LongExampleCmd,
	RunE: func(cmd *cobra.Command, args []string) error {
		fields, err := form.Run()
		if err != nil {
			return err
		}

		output.Green("Success green output: ", fields.ExampleInput)
		output.Red("This command will be run successfully")

		return nil
	},
}
