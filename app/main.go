package main

import (
	"fmt"

	"github.com/example/workspace-vendor-lib"
	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{
		Use: "app",
		Run: func(cmd *cobra.Command, args []string) {
			lib.Print("Hello from workspace-vendor app!")
			fmt.Println("Done")
		},
	}
	cmd.Execute()
}