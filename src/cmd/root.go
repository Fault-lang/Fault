package root

// import (
// 	"fmt"
// 	"os"

// 	"github.com/mitchellh/go-homedir"
// 	"github.com/spf13/viper"
// )

// var (
// 	rootCmd = &cobra.Command{
// 		Use:   "fault",
// 		Short: "Fault is a language for modeling system complexity",
// 		Long: `Build compartmentalize models of software systems
// 				  using stock/flow abstractions.
// 				  Complete documentation is available at https://fault-lang.com`,
// 	}
// )

// func Execute() {
// 	if err := rootCmd.Execute(); err != nil {
// 		fmt.Fprintln(os.Stderr, err)
// 		os.Exit(1)
// 	}
// }

// func init() {
// 	rootCmd.AddCommand(versionCmd)
// 	// cobra.OnInitialize(initConfig)

// 	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")
// 	// rootCmd.PersistentFlags().StringP("author", "a", "YOUR NAME", "author name for copyright attribution")
// 	// rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "name of license for the project")
// 	// rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")
// 	// viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
// 	// viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
// 	// viper.SetDefault("author", "NAME HERE <EMAIL ADDRESS>")
// 	// viper.SetDefault("license", "apache")

// 	// rootCmd.AddCommand(addCmd)
// 	// rootCmd.AddCommand(initCmd)
// }

// func initConfig() {
// 	if cfgFile != "" {
// 		// Use config file from the flag.
// 		viper.SetConfigFile(cfgFile)
// 	} else {
// 		// Find home directory.
// 		home, err := homedir.Dir()
// 		cobra.CheckErr(err)

// 		// Search config in home directory with name ".cobra" (without extension).
// 		viper.AddConfigPath(home)
// 		viper.SetConfigName(".cobra")
// 	}

// 	viper.AutomaticEnv()

// 	if err := viper.ReadInConfig(); err == nil {
// 		fmt.Println("Using config file:", viper.ConfigFileUsed())
// 	}
// }
