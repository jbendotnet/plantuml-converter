package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile, PlantUmlServer, FilePattern string

var rootCmd = &cobra.Command{
	Use:   "plantuml-converter",
	Short: "scan readme and place plantuml link",
	Long:  "scan readme and place plantuml link",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("hello world")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.plantuml-converter.yaml)")
	rootCmd.PersistentFlags().StringVarP(&PlantUmlServer, "server", "s", "http://www.plantuml.com", "plantUML server address")
	rootCmd.PersistentFlags().StringVarP(&FilePattern, "pattern", "p", "*.md", "plantUML server address")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".plantuml-converter" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".plantuml-converter")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
