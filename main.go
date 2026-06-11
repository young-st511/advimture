package main

import (
	"embed"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/young-st511/advimture/internal/app"
)

//go:embed content
var contentFS embed.FS

var version = "dev"

func main() {
	if shouldPrintVersion(os.Args[1:]) {
		fmt.Printf("advimture %s\n", version)
		return
	}

	p := tea.NewProgram(
		app.NewWithOptions(app.Options{ContentFS: contentFS}),
		tea.WithAltScreen(),
	)
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "[Advimture 실행 실패] %v\n", err)
		os.Exit(1)
	}
}

func shouldPrintVersion(args []string) bool {
	if len(args) != 1 {
		return false
	}
	return args[0] == "version" || args[0] == "--version" || args[0] == "-v"
}
