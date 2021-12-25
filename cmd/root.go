package cmd

import (
	"fmt"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	cfgFile string
	debug   bool
	// rootCmd represents the base command when called without any subcommands
	rootCmd = &cobra.Command{
		Use: "advent",
	}
)

type day struct {
	Function func()
}

// Execute cobra normal command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func setLogLevel() {
	if debug {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}
}

func addDaySubCommandToYearCommand(parentCmd *cobra.Command, days map[string]day) {
	for day := range days {
		// Run function context to ensure these right values are used.
		d := day
		dayfunc := days[day].Function

		parentCmd.AddCommand(&cobra.Command{
			Use: "day" + day,
			Aliases: []string{
				strings.ToLower(day),
				strings.TrimLeft(day, "0"),
				"day" + strings.TrimLeft(day, "0"),
			},
			Short: "Runs through the real puzzle data for the day" + day + ".",
			Run: func(cmd *cobra.Command, args []string) {
				setLogLevel()
				fmt.Printf("### Day %s\n", d)
				dayfunc()
			},
		})
	}
}

func printAllDaysOutput(cmd *cobra.Command, args []string) {
	setLogLevel()
	// Loops through the child commands (which cobra sorts) and executes the Run function.
	for _, c := range cmd.Commands() {
		c.Run(cmd, args)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "Enable debug output of logging")
}
