package cmd

import (
	"fmt"
	"os"

	statuspage "github.com/cloudrkt/go-statuspage-api"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// AppConfig represents the CLI application
type AppConfig struct {
	CfgFile string
	Client  *statuspage.Client
	Debug   bool
}

var app *AppConfig

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Version: "0.1",
	Use:     "statuspage",
	Short:   "The Statuspage cli tool",
	Long:    `This commandline tool interacts with the statuspage.io API`,
}

// Execute root cmd
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	app = &AppConfig{}

	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().StringVar(&app.CfgFile, "config", "", "config file (default is $HOME/.statuspage.yaml)")
	RootCmd.PersistentFlags().BoolVar(&app.Debug, "debug", false, "debug mode")
}

func initConfig() {
	if app.CfgFile != "" {
		viper.SetConfigFile(app.CfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.AddConfigPath(home)
		viper.AddConfigPath("/etc/statuspage/")
		viper.AddConfigPath(".")
		viper.SetConfigName(".statuspage")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		if app.Debug {
			fmt.Println("Using config file:", viper.ConfigFileUsed())
		}
	}

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}

	configureClient(viper.GetString("sp_apikey"), viper.GetString("sp_pageid"))
}

func configureClient(apikey, pageid string) {
	var err error

	app.Client, err = statuspage.NewClient(apikey, pageid)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
