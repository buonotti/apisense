package log

import (
	"fmt"
	"time"

	"github.com/apex/log"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	"github.com/gin-gonic/gin"
)

// GinMiddleware returns a custom logging middleware that uses log.ApiLogger instead of the gin default logger
func GinMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		url := c.FullPath()
		if url == "" {
			url = "404"
		}
		t := time.Now()
		c.Next()
		elapsed := time.Since(t)
		ApiLogger.WithFields(log.Fields{
			"time":   fmt.Sprintf("%dms", elapsed.Milliseconds()),
			"status": c.Writer.Status(),
			"method": c.Request.Method,
			"ip":     c.ClientIP(),
		}).Info(url)
	}
}

// WishMiddleware returns a custom logging middleware that uses log.SSHLogger instead of the wish default logger
func WishMiddleware() wish.Middleware {
	return func(sh ssh.Handler) ssh.Handler {
		return func(s ssh.Session) {
			ct := time.Now()
			hpk := s.PublicKey() != nil
			pty, _, _ := s.Pty()
			SSHLogger.WithFields(log.Fields{
				"user":   s.User(),
				"ip":     s.RemoteAddr().String(),
				"hpk":    hpk,
				"cmd":    s.Command(),
				"term":   pty.Term,
				"width":  pty.Window.Width,
				"height": pty.Window.Height,
			}).Infof("%s connect", s.User())
			sh(s)
			SSHLogger.WithField("ip", s.RemoteAddr().String()).WithField("time", time.Since(ct)).Info("disconnected")
		}
	}
}
