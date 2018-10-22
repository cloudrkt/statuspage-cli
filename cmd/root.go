package cmd

import (
	"errors"
	"fmt"
	"os"

	statuspage "github.com/cloudrkt/go-statuspage-api"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// AppConfig represents the CLI application
type AppConfig struct {
	Client  *statuspage.Client
	Config  *viper.Viper
	CfgFile string
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

	app.Config = viper.New()

	if app.CfgFile != "" {
		app.Config.SetConfigFile(app.CfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		app.Config.AddConfigPath(home)
		app.Config.AddConfigPath("/etc/statuspage/")
		app.Config.AddConfigPath(".")
		app.Config.SetConfigName(".statuspage")
	}

	app.Config.SetEnvPrefix("sp")
	app.Config.AutomaticEnv()

	if err := app.Config.ReadInConfig(); err == nil {
		if app.Debug {
			fmt.Println("Using config file:", app.Config.ConfigFileUsed())
		}
	}

	err := app.Config.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}

	err = configureClient()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func configureClient() error {
	if !app.Config.IsSet("apikey") {
		return errors.New("apikey not set")
	}

	if !app.Config.IsSet("pageid") {
		return errors.New("pageid not set")
	}

	var err error
	app.Client, err = statuspage.NewClient(
		app.Config.GetString("apikey"),
		app.Config.GetString("pageid"),
	)
	return err
}
