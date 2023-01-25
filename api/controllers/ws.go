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
	Timestamp time.Time
	Filename  string
	ReportId  string
	Report    validation.Report
}

func wsHandler(w http.ResponseWriter, r *http.Request) error {
	conn, err := wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		return errors.CannotUpgradeWebsocket.Wrap(err, "Cannot upgrade websocket")
	}

	defer func() {
		err := conn.Close()
		errors.HandleError(err)
	}()

	watcher, err := fs.NewDirectoryWatcherWithFiles(validation.ReportLocation())
	if err != nil {
		return err
	}

	go func() {
		err := watcher.Start()
		errors.HandleError(err)
	}()

	for {
		newFile := <-watcher.Events
		report, err := validation.GetReport(newFile)
		if err != nil {
			return err
		}
		err = conn.WriteJSON(wsResponse{
			Timestamp: time.Now(),
			Filename:  newFile,
			ReportId:  report.Id,
			Report:    *report,
		})
		if err != nil {
			return errors.CannotWriteWebsocket.Wrap(err, "cannot write to websocket")
		}
	}
}

func Ws(c *gin.Context) {
	err := wsHandler(c.Writer, c.Request)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}
}
