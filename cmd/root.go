package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func New() *cobra.Command {

	var rootCmd = &cobra.Command{
		Use:   "hello-cobra",
		Short: "This is hello cobra",
		//Long:  `All software has versions. This is Hugo's`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("hello cobra is running")
		},
	}

	viper.SetConfigName("config")         // name of config file (without extension)
	viper.SetConfigType("yaml")           // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("/etc/appname/")  // path to look for the config file in
	viper.AddConfigPath("$HOME/.appname") // call multiple times to add many search paths
	viper.AddConfigPath(".")              // optionally look for config in the working directory
	err := viper.ReadInConfig()           // Find and read the config file
	if err != nil {                       // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	param := viper.Get("parameter1")
	param2 := viper.Get("parameter2")

	fmt.Println(param)
	fmt.Println(param2)

	rootCmd.AddCommand(versionCmd)

	return versionCmd
}
