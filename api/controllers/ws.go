package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"

	"github.com/buonotti/apisense/errors"
	"github.com/buonotti/apisense/fs"
	"github.com/buonotti/apisense/validation"
)

var wsUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type wsResponse struct {
	Timestamp time.Time         `json:"timestamp"`
	Filename  string            `json:"filename"`
	ReportId  string            `json:"reportId"`
	Report    validation.Report `json:"report"`
}

func Ws(c *gin.Context) {
	conn, err := wsUpgrader.Upgrade(c.Writer, c.Request, nil)
	err = errors.SafeWrap(errors.CannotUpgradeWebsocket, err, "cannot upgrade websocket connection")
	errors.HandleError(err)

	defer func() {
		err := conn.Close()
		errors.HandleError(err)
	}()

	watcher, err := fs.NewDirectoryWatcherWithFiles(validation.ReportLocation())
	errors.HandleError(err)

	go func() {
		err := watcher.Start()
		errors.HandleError(err)
	}()

	for {
		newFile := <-watcher.Events
		report, err := validation.GetReport(newFile)
		errors.HandleError(err)
		err = conn.WriteJSON(wsResponse{
			Timestamp: time.Now(),
			Filename:  newFile,
			ReportId:  report.Id,
			Report:    *report,
		})
		err = errors.SafeWrap(errors.CannotWriteWebsocket, err, "cannot write to websocket")
		errors.HandleError(err)
	}
}
