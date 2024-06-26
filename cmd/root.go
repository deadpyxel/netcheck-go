package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/deadpyxel/netcheck-go/internal/core"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// config values
var cfgFile string
var checkInterval int
var logFile string
var cfg core.Config

var rootCmd = &cobra.Command{
	Use:   "netcheck",
	Short: "netcheck is a simple CLI to monitor Internet connection",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	// add config flag
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/netcheck-go/config.yaml)")
	rootCmd.PersistentFlags().IntVar(&checkInterval, "interval", 5, "Interval for checking in seconds (dafaults to 5s)")
	rootCmd.PersistentFlags().StringVar(&logFile, "logFile", "netcheck-go.log", "Output logs location, used by monitoring and verification routines")
}

func initConfig() {

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home + "/.config/netcheck-go")
		viper.SetConfigType("yaml")
		viper.SetConfigName("config.yaml")
	}

	// Set default value for statePath. Will use your HOME if not set.
	viper.SetDefault("checkInterval", 5)

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}

	cfg = *core.NewConfig(checkInterval, logFile)
	err := cfg.Init()
	if err != nil {
		log.Fatalf("Error while initializing config: %v", err)
	}
}
