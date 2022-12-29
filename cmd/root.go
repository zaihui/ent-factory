package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Use:   "ent-factory",
	Short: "a code generation tool for factoring entity of ent model schema",
	Long:  ``,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ent-factory.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.PersistentFlags().StringP("schemaFile", "f", "", "file which model schema defined")
	rootCmd.PersistentFlags().StringP("schemaPath", "p", "", "path of ent schema model file")
	rootCmd.PersistentFlags().StringP("outputPath", "o", "", "path to write factories")
	rootCmd.PersistentFlags().StringP("projectPath", "j", "", "the relative path of this project")
	rootCmd.PersistentFlags().StringP(
		"factoriesPath", "t", "", "the relative path of these factories located in this project")
	rootCmd.PersistentFlags().StringP("appPath", "a", "", "the relative path of app client")
	rootCmd.PersistentFlags().StringP("entClientName", "e", "", "the name of ent client")
}

func Fatal(msg string) {
	os.Exit(1)
}

func Fatalf(format string, v ...interface{}) {
	Fatal(fmt.Sprintf(format, v...))
}
