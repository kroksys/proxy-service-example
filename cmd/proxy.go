package cmd

import (
	"fmt"
	"os"

	"github.com/kroksys/proxy-service-example/src/proxy"
	"github.com/kroksys/proxy-service-example/src/utils"
	"github.com/spf13/cobra"
)

var (
	configFile string
	config     = &utils.Config{}
)

var proxyCmd = &cobra.Command{
	Use:   "proxy",
	Short: "A very simple HTTP proxy which forwards all the requests to a given domain.",
	Run: func(cmd *cobra.Command, args []string) {
		var err error

		if configFile != "" {
			config, err = utils.ReadConfig(configFile)
			if err != nil {
				fmt.Printf("Could not read provided config file: %s. Error: %s\n", configFile, err.Error())
				os.Exit(0)
			}
		}

		fmt.Printf("\nProxy server started\n")
		fmt.Printf("Listen: %s\n", config.Proxy.Listen)
		fmt.Printf("Target: %s\n", config.Proxy.Target)
		fmt.Printf("Log: %s\n\n", config.Proxy.Log)

		err = proxy.ListenAndServe(config)
		if err != nil {
			fmt.Printf("Proxy.ListenAndServe error: %s\n", err.Error())
			os.Exit(0)
		}
	},
}

func init() {
	rootCmd.AddCommand(proxyCmd)

	proxyCmd.Flags().StringVarP(&config.Proxy.Listen, "listen", "l", "0.0.0.0:5000", "The IP and port the proxy binds to.")
	proxyCmd.Flags().StringVarP(&config.Proxy.Target, "target", "t", "http://httpforever.com", "Target where all requests will be forwarded.")
	proxyCmd.Flags().StringVarP(&config.Proxy.Log, "log", "g", "log.txt", "Log file.")
	proxyCmd.Flags().StringVarP(&configFile, "configFile", "c", "", "Configuration file in .yaml or .toml or .json format.")
}
