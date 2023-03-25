package cmd

import (
	"net/rpc"

	"github.com/buonotti/apisense/errors"
	"github.com/buonotti/apisense/log"
	"github.com/spf13/cobra"
)

var daemonReloadCmd = &cobra.Command{
	Use:   "reload",
	Short: "Reload the daemon",
	Long:  `Reload the daemon`, // TODO: Add more info
	Run: func(cmd *cobra.Command, args []string) {
		client, err := rpc.DialHTTP("tcp", "127.0.0.1:1234")
		errors.CheckErr(err)
		var reply int
		err = client.Call("RpcDaemonManager.ReloadDaemon", 0, &reply)
		errors.CheckErr(err)
		if reply == 0 {
			log.CliLogger.Info("reloaded daemon pipeline items")
		} else {
			log.CliLogger.Error("cannot reload daemon pipeline items")
		}
	},
}

func init() {
	daemonCmd.AddCommand(daemonReloadCmd)
}
