package main

import (
	"github.com/1995parham/mqtt-bench/internal/cmd"
	"github.com/pterm/pterm"
)

func main() {
	if err := pterm.DefaultBigText.WithLetters(
		pterm.NewLettersFromStringWithStyle("MQTT", pterm.NewStyle(pterm.FgCyan)),
		pterm.NewLettersFromStringWithStyle("-", pterm.NewStyle(pterm.FgLightMagenta)),
		pterm.NewLettersFromStringWithStyle("bench", pterm.NewStyle(pterm.FgLightRed)),
	).Render(); err != nil {
		_ = err
	}

	pterm.Description.Println("Sometimes we tell lies, sometimes we prove we don't lie. Let's prove ourselves.")

	cmd.Execute()
}
