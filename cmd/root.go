package cmd

import (
	"fmt"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "timestamp-converter",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// GET type of conversion

		typeConv, err := cmd.Flags().GetString("type")
		if err != nil {
			return
		}
		value, err := cmd.Flags().GetString("value")
		if err != nil {
			return
		}

		var timestamp time.Time

		switch typeConv {
		case "SECONDS":
			cv, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				return
			}

			timestamp = time.Unix(cv, 0)
			break
		case "MILLISECONDS":
			cv, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				return
			}
			cv *= 1000

			timestamp = time.Unix(cv, 0)
			break
		case "RFC3339":
			timestamp, err = time.Parse(time.RFC3339, value)
			if err != nil {
				return
			}
			break
		}

		fmt.Printf("Seconds: %d\nMilliseconds: %d\nRFC3339: %s\n", timestamp.Unix(), timestamp.Unix()*1000, timestamp.Format(time.RFC3339))
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.Flags().String("type", "SECONDS", "Type of conversion. Possible values: SECONDS, MILLISECONDS, RFC3339")
	rootCmd.MarkFlagRequired("type")
	rootCmd.Flags().String("value", "0", "Value to parse")
	rootCmd.MarkFlagRequired("value")
}
