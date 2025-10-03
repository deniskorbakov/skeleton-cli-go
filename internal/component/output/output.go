package output

import (
	"fmt"

	"github.com/charmbracelet/lipgloss/v2"
)

var (
	green = lipgloss.Color("#04B575")
	red   = lipgloss.Color("#D4634C")
)

func output(colorText string) {
	fmt.Println(colorText)
}

func Green(str ...string) {
	output(lipgloss.NewStyle().Foreground(green).Render(str...))
}

func Red(str ...string) {
	output(lipgloss.NewStyle().Foreground(red).Render(str...))
}
