// Copyright © 2017 Google LLC
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/golang/glog"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/rspier/go-ecobee/ecobee"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile    string
	thermostat string
	appID      string
)

const (
	authCacheFile = ".go-ecobee-authcache"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "go-ecobee",
	Short: "A CLI tool to control Ecobee thermostats",
	Long:  `A command line tool to control Ecobee thermostats.  It can view status, set temperatures, turn the fan on and off, send messages, and more!`,
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {

	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-ecobee.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	RootCmd.PersistentFlags().StringP("thermostat", "t", "", "thermostat id")
	RootCmd.PersistentFlags().StringP("appid", "i", "", "app id")
	RootCmd.PersistentFlags().StringP("authcache", "", "", "auth cache file")

	// This is a little messy... is there a nicer way to do this?
	ck := func(err error) {
		if err != nil {
			log.Fatal("unexpected error setting up flag parsing!")
		}
	}
	ck(viper.BindPFlag("thermostat", RootCmd.PersistentFlags().Lookup("thermostat")))
	ck(viper.BindPFlag("appid", RootCmd.PersistentFlags().Lookup("appid")))
	ck(viper.BindPFlag("authcache", RootCmd.PersistentFlags().Lookup("authcache")))
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
			glog.Exitf("error retrieving homedir: %v", err)
		}

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".go-ecobee")
	}

	viper.SetEnvPrefix("ecobee")
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		glog.Exitf("Using config file: %s", viper.ConfigFileUsed())
	}

	// load important configs into global variables
	thermostat = viper.GetString("thermostat")
	appID = viper.GetString("appid")

}

func client() *ecobee.Client {

	ac := viper.GetString("authcache")
	if ac == "" {
		home, err := homedir.Dir()
		if err != nil {
			glog.Exitf("error retrieving homedir: %v", err)
		}
		ac = path.Join(home, authCacheFile)
	}
	glog.V(1).Infof("authCache: %s", ac)
	// replace authCacheFile with authCache flag
	return ecobee.NewClient(appID, ac)
}
