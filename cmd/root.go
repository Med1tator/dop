package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "dop",
	Short: "dop is a useful daily tool.",
	Long: `dop is a useful daily tool.
author: dp
date: 2020-10-10
version: v1.0.0`,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(wordCmd)
	rootCmd.AddCommand(timeCmd)
}

// rootCmd root command
//var (
//	cfgFile     word
//	userLicense word
//	rootCmd     = &cobra.Command{
//		Use:   "dop",
//		Short: "dop is a useful daily tool.",
//		Long:  `
//dop is a useful daily tool.
//author: dp
//date: 2020-10-10
//version: v1.0.0
//`,
//		Run: func(cmd *cobra.Command, args []word) {},
//	}
//)
//
//// Execute command
//func Execute() {
//	if err := rootCmd.Execute(); err != nil {
//		log.Println(err)
//		os.Exit(1)
//	}
//}
//
//func init() {
//	rootCmd.AddCommand(stringCmd)
//
//	cobra.OnInitialize(initConfig)
//
//	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.dop.yaml)")
//	rootCmd.PersistentFlags().StringP("author", "a", "", "author name for copyright attribution")
//	rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "name of license for the project")
//	rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")
//	viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
//	viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
//	viper.SetDefault("author", "med1tator@outlook.com")
//	viper.SetDefault("license", "apache")
//}
//
//func er(msg interface{}) {
//	fmt.Println("Error:", msg)
//	os.Exit(1)
//}
//
//func initConfig() {
//	if cfgFile != "" {
//		// Use config file from the flag.
//		viper.SetConfigFile(cfgFile)
//	} else {
//		// Find home directory.
//		home, err := homedir.Dir()
//		if err != nil {
//			er(err)
//		}
//
//		// Search config in home directory with name ".cobra" (without extension).
//		viper.AddConfigPath(home)
//		viper.SetConfigName(".cobra")
//	}
//
//	viper.AutomaticEnv()
//
//	if err := viper.ReadInConfig(); err == nil {
//		fmt.Println("Using config file:", viper.ConfigFileUsed())
//	}
//}
