package cmd

import (
	"github.com/bakito/adguardhome-sync/pkg/log"
	"github.com/bakito/adguardhome-sync/pkg/sync"
	"github.com/bakito/adguardhome-sync/pkg/types"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// runCmd represents the run command
var doCmd = &cobra.Command{
	Use:   "run",
	Short: "Start a synchronisation from origin to replica",
	Long:  `Synchronizes the configuration form an origin instance to a replica`,
	Run: func(cmd *cobra.Command, args []string) {

		logger = log.GetLogger("root")
		cfg := &types.Config{}
		if err := viper.Unmarshal(cfg); err != nil {
			logger.Error(err)
			return
		}

		sync.Sync(cfg)
	},
}

func init() {
	rootCmd.AddCommand(doCmd)
	doCmd.PersistentFlags().String("cron", "", "The cron expression to run in daemon mode")
	_ = viper.BindPFlag(configCron, doCmd.PersistentFlags().Lookup("cron"))

	doCmd.PersistentFlags().String("origin-url", "", "Origin instance url")
	_ = viper.BindPFlag(configOriginURL, doCmd.PersistentFlags().Lookup("origin-url"))
	doCmd.PersistentFlags().String("origin-api-path", "/control", "Origin instance API path")
	_ = viper.BindPFlag(configOriginAPIPath, doCmd.PersistentFlags().Lookup("origin-api-path"))
	doCmd.PersistentFlags().String("origin-username", "", "Origin instance username")
	_ = viper.BindPFlag(configOriginUsername, doCmd.PersistentFlags().Lookup("origin-username"))
	doCmd.PersistentFlags().String("origin-password", "", "Origin instance password")
	_ = viper.BindPFlag(configOriginPassword, doCmd.PersistentFlags().Lookup("origin-password"))
	doCmd.PersistentFlags().String("origin-insecure-skip-verify", "", "Enable Origin instance InsecureSkipVerify")
	_ = viper.BindPFlag(configOriginInsecureSkipVerify, doCmd.PersistentFlags().Lookup("origin-insecure-skip-verify"))

	doCmd.PersistentFlags().String("replica-url", "", "Replica instance url")
	_ = viper.BindPFlag(configReplicaURL, doCmd.PersistentFlags().Lookup("replica-url"))
	doCmd.PersistentFlags().String("replica-api-path", "/control", "Replica instance API path")
	_ = viper.BindPFlag(configReplicaAPIPath, doCmd.PersistentFlags().Lookup("replica-api-path"))
	doCmd.PersistentFlags().String("replica-username", "", "Replica instance username")
	_ = viper.BindPFlag(configReplicaUsername, doCmd.PersistentFlags().Lookup("replica-username"))
	doCmd.PersistentFlags().String("replica-password", "", "Replica instance password")
	_ = viper.BindPFlag(configReplicaPassword, doCmd.PersistentFlags().Lookup("replica-password"))
	doCmd.PersistentFlags().String("replica-insecure-skip-verify", "", "Enable Replica instance InsecureSkipVerify")
	_ = viper.BindPFlag(configReplicaInsecureSkipVerify, doCmd.PersistentFlags().Lookup("replica-insecure-skip-verify"))
}
