package cmd

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "snip",
    Short: "snip is a fast CLI code snippet manager",
    Long:  `Save, search, and reuse your code snippets directly from the terminal.`,
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
