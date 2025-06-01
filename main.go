package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	_ "k8s.io/client-go/plugin/pkg/client/auth"

	"github.com/mohamedaminehamdi/kubepilot/cmd"
	"github.com/mohamedaminehamdi/kubepilot/cmd/audit"
	"github.com/mohamedaminehamdi/kubepilot/cmd/diagnose"
)

func getCommandName() string {
	if strings.HasPrefix(filepath.Base(os.Args[0]), "kubectl-") {
		// cobra will split on " " and take the first element
		return "kubectl\u2002kubepilot"
	}
	return "kubepilot"
}

func main() {
	opt := cmd.NewOptions()
	err := opt.Complete()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return
	}
	rootCmd := cobra.Command{
		Use: getCommandName(),
		Long: fmt.Sprintf(`You need to set TWO ENVs to run kubepilot.
Set %s to specify your token.
Set %s to specify the language. Valid options like Chinese, French, Spain, etc.
`, cmd.EnvkubepilotToken, cmd.EnvkubepilotLang),
	}
	rootCmd.AddCommand(
		audit.New(opt),
		diagnose.New(opt),
	)
	_ = rootCmd.Execute()
}
