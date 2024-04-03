package cmd

import (
	"fmt"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/spf13/cobra"
)

var tokenCmd = &cobra.Command{
	Use:   "token",
	Short: "Create an Application token",
	Example: heredoc.Doc(`
		Create an Application token 
		$ gh-runner-util app token --key $KEY --id <app-id> 
	`),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hi there from the token command!")
	},
}

func init() {
	appCmd.AddCommand(tokenCmd)
}
