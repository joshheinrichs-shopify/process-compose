package cmd

import (
	"fmt"

	"github.com/f1bonacc1/process-compose/src/types"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var name string
var namespace string

// startCmd represents the start command
var addCmd = &cobra.Command{
	Use:   "add [PROCESS]",
	Short: "Start a process",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		command := args[0]

		processConfig := types.ProcessConfig{
			Name:      name,
			Namespace: namespace,
			Command:   command,
		}

		err := getClient().AddProcess(processConfig)
		if err != nil {
			log.Fatal().Err(err).Msgf("failed to add process %s", name)

		}
		fmt.Printf("Process %s added\n", name)
	},
}

func init() {
	processCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&name, "name", "", "", "")
	addCmd.Flags().StringVarP(&namespace, "namespace", "", "default", "")
}
