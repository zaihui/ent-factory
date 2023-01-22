package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/zaihui/ent-factory/constants"

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
	rootCmd.PersistentFlags().StringP("schemaFile", "f", "", "file which model schema defined")
	rootCmd.PersistentFlags().StringP("schemaPath", "p", "", "path of ent schema model file")
	rootCmd.PersistentFlags().StringP("outputPath", "o", "", "path to write factories")
	rootCmd.PersistentFlags().StringP("projectPath", "j", "", "the relative path of this project")
	rootCmd.PersistentFlags().BoolP("overwrite", "r", false, "whether overwrite exist files")
	rootCmd.PersistentFlags().StringP(
		"factoriesPath", "t", constants.DefaultFactoryPath, "the relative path of these factories located in this project")
	rootCmd.PersistentFlags().StringP("appPath", "a", constants.DefaultAppPath, "the relative path of app client")
	rootCmd.PersistentFlags().StringP("entClientName", "e", constants.DefaultEntClientName, "the name of ent client")
	rootCmd.PersistentFlags().StringP("modelPath", "m", constants.DefaultModelPath, "the path of the definition of models")
	rootCmd.PersistentFlags().BoolP("genImportFields", "i", false, "whether generate import fields")
}

func Fatal(msg string) {
	os.Exit(1)
}

func Fatalf(format string, v ...interface{}) {
	Fatal(fmt.Sprintf(format, v...))
}
