package daemon

import (
	"errors"
	"net"
	"net/http"
	"net/rpc"

	"github.com/buonotti/apisense/log"
)

type RpcDaemonManager struct {
	daemon *daemon
}

func startRpcServer(daemon *daemon) error {
	manager := &RpcDaemonManager{daemon: daemon}
	err := rpc.Register(manager)
	if err != nil {
		log.DaemonLogger().Error("Cannot register daemon rpc server", "reason", err.Error())
		return err
	}
	rpc.HandleHTTP()
	l, err := net.Listen("tcp", "127.0.0.1:42069")
	if err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			log.DaemonLogger().Error("Cannot start daemon rpc server", "reason", err.Error())
		} else {
			log.DaemonLogger().Info("Daemon rpc server stopped")
		}
	}
	log.DaemonLogger().Info("Daemon rpc server started", "address", "127.0.0.1:42069")
	return http.Serve(l, nil)
}
