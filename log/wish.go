package log

import (
	"time"

	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
)

// WishMiddleware returns a custom logging middleware that uses log.SSHLogger instead of the wish default logger
func WishMiddleware() wish.Middleware {
	return func(sh ssh.Handler) ssh.Handler {
		return func(s ssh.Session) {
			ct := time.Now()
			hpk := s.PublicKey() != nil
			pty, _, _ := s.Pty()
			SSHLogger.Infof("%s connect %s %v %v %s %v %v\n", s.User(), s.RemoteAddr().String(), hpk, s.Command(), pty.Term, pty.Window.Width, pty.Window.Height)
			sh(s)
			SSHLogger.Infof("%s disconnect %s\n", s.RemoteAddr().String(), time.Since(ct))
		}
	}
}
