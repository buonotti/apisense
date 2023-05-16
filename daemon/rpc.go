package daemon

import (
	"net"
	"net/http"
	"net/rpc"

	"github.com/buonotti/apisense/log"
)

type RpcDaemonManager struct {
	daemon *daemon
}

func (mgr *RpcDaemonManager) ReloadDaemon(retries *int, reply *int) error {
	var err error
	for i := 0; i <= *retries; i++ {
		err = mgr.daemon.Pipeline.Reload()
		if err == nil {
			break
		}
	}
	if err != nil {
		log.DaemonLogger.WithError(err).Error("cannot reload daemon")
		*reply = 1
		return err
	}

	*reply = 0
	return nil
}

func startRpcServer(daemon *daemon) error {
	manager := &RpcDaemonManager{daemon: daemon}
	err := rpc.Register(manager)
	if err != nil {
		log.DaemonLogger.WithError(err).Error("cannot register daemon rpc server")
		return err
	}
	rpc.HandleHTTP()
	l, err := net.Listen("tcp", "127.0.0.1:42069")
	if err != nil {
		if err != http.ErrServerClosed {
			log.DaemonLogger.WithError(err).Error("cannot start daemon rpc server")
		} else {
			log.DaemonLogger.Info("daemon rpc server stopped")
		}
	}
	log.DaemonLogger.WithField("address", "127.0.0.1:42069").Info("daemon rpc server started")
	return http.Serve(l, nil)
}
