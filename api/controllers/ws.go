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
	Timestamp time.Time `json:"timestamp"`
	Filename  string    `json:"filename"`
	ReportId  string    `json:"reportId"`
}

// Ws godoc
//	@Summary		Open a websocket connection to receive notifications
//	@Description	Connect to this endpoint with the ws:// protocol to instantiate a websocket connection to get updates for new reports
//	@ID				ws
//	@Tags			reports
//	@Success		101
//	@Router			/api/reports [get]
func Ws(c *gin.Context) {
	// we have to handle the errors because the upgrade hijacks the response writer
	// so we cant use the context to write the response to the client
	conn, err := wsUpgrader.Upgrade(c.Writer, c.Request, nil)
	err = errors.SafeWrap(errors.CannotUpgradeWebsocketError, err, "cannot upgrade websocket connection")
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
		newFileName := <-watcher.Events
		report, err := validation.GetReport(newFileName)
		errors.HandleError(err)
		err = conn.WriteJSON(wsResponse{
			Timestamp: time.Now(),
			Filename:  newFileName,
			ReportId:  report.Id,
		})
		err = errors.SafeWrap(errors.CannotWriteWebsocketError, err, "cannot write to websocket")
		errors.HandleError(err)
	}
}
