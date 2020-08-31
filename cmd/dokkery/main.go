package main

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	appname = "dokkery"
	cfgFile string
	version string
	debug   bool
)

var rootCmd = &cobra.Command{
	Use:     appname,
	Version: version,
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", ``)
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, ``)
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		curDir, _ := os.Getwd()
		viper.AddConfigPath(curDir)
		viper.SetConfigName("config")
	}

	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		_, ok := err.(viper.ConfigFileNotFoundError)
		if !ok {
			log.Fatalf("error reading in config: %v", err)
		}
	}
}

func main() {
	rootCmd.SetVersionTemplate(rootCmd.Version + "\n")
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
