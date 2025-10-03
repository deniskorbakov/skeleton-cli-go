package form

import "github.com/charmbracelet/huh"

// Run Main function that runs an interactive form
func Run() (*Fields, error) {
	fields := &Fields{}

	err := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Example Title").
				Description("Example description input").
				Value(&fields.ExampleInput),
		),
	).WithShowHelp(true).Run()
	if err != nil {
		return nil, err
	}

	return fields, nil
}
