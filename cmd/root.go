package cmd

import (
	"fmt"
	"github.com/sepuka/focalism/def"
	"os"

	"github.com/spf13/cobra"
)

var (
	configFile string

	rootCmd = &cobra.Command{
		Use:  `focalism`,
		Args: cobra.MinimumNArgs(1),
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return def.Build(configFile)
		},
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "", "/path/to/config.yml")
	_ = rootCmd.MarkPersistentFlagRequired("config")
}
