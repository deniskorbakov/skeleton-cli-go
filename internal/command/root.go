package command

import (
	"context"
	"os"

	"github.com/charmbracelet/fang"
	"github.com/deniskorbakov/skeleton-cli-go/configs/constname"
	"github.com/deniskorbakov/skeleton-cli-go/internal/component/output"
	"github.com/deniskorbakov/skeleton-cli-go/internal/version"
	"github.com/spf13/cobra"
)

// Run Start app with cobra cmd
func Run() {
	cmd := &cobra.Command{
		Use:     constname.UseRootCmd,
		Long:    constname.LongRootCmd,
		Example: constname.ExampleRootCmd,
	}

	// Disable default options cmd
	cmd.Root().CompletionOptions.DisableDefaultCmd = true

	// Add all command in your app
	cmd.AddCommand(
		exampleCmd,
	)

	if err := fang.Execute(
		context.Background(),
		cmd,
		fang.WithVersion(version.Get()),
	); err != nil {
		output.Red("The app is broken")
		os.Exit(1)
	}
}
