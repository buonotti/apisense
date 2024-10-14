package controllers

import (
	"github.com/buonotti/apisense/log"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"

	"github.com/buonotti/apisense/errors"
	"github.com/buonotti/apisense/filesystem"
	"github.com/buonotti/apisense/filesystem/locations/directories"
	"github.com/buonotti/apisense/validation/pipeline"
)

var wsUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type wsResponse struct {
	Timestamp time.Time `json:"timestamp"`
	Filename  string    `json:"filename"`
	ReportId  string    `json:"reportId"`
}

// Ws godoc
//
//	@Summary		Open a websocket connection to receive notifications
//	@Description	Connect to this endpoint with the ws:// protocol to instantiate a websocket connection to get updates for new reports
//	@ID				ws
//	@Tags			reports
//	@Success		101
//	@Router			/ws [get]
func Ws(c *gin.Context) {
	// we have to handle the errors because the upgrade hijacks the response writer
	// so we cant use the context to write the response to the client
	conn, err := wsUpgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.ApiLogger().Error(errors.CannotUpgradeWebsocketError.Wrap(err, "cannot upgrade websocket connection"))
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	defer func() {
		closeErr := conn.Close()
		if closeErr != nil {
			log.ApiLogger().Error(closeErr)
		}
	}()

	watcher, err := filesystem.NewDirectoryWatcherWithFiles(filepath.FromSlash(directories.ReportsDirectory()))
	if err != nil {
		log.ApiLogger().Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	go func() {
		watcherStartErr := watcher.Start()
		if watcherStartErr != nil {
			log.ApiLogger().Error(watcherStartErr)
		}
	}()

	for {
		newFileName := <-watcher.Events
		report, err := pipeline.GetReport(newFileName)
		if err != nil {
			log.ApiLogger().Error(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		err = conn.WriteJSON(wsResponse{
			Timestamp: time.Now(),
			Filename:  newFileName,
			ReportId:  report.Id,
		})
		if err != nil {
			log.ApiLogger().Error(errors.CannotWriteWebsocketError.Wrap(err, "cannot write to websocket"))
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
	}
}
