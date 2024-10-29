package alerting

import (
	"fmt"
	"net/smtp"

	"github.com/buonotti/apisense/util"
	"github.com/spf13/viper"
)

// Information used to send an alert
type AlertData struct {
	Time        util.ApisenseTime
	ErrorAmount uint
}

// SendAlert sends an alert to the receiver specified in the config. It uses the AlertData to know what and if to send an alert
func SendAlert(data AlertData) error {
	emailUser := viper.GetString("daemon.notification.username")
	emailPass := viper.GetString("daemon.notification.password")
	sendOnErrors := viper.GetBool("daemon.notification.only_on_error")
	sender := viper.GetString("daemon.notification.sender")
	receiver := viper.GetString("daemon.notification.receiver")
	server := viper.GetString("daemon.notification.smtp_server")
	port := viper.GetUint("daemon.notification.smtp_port")

	auth := smtp.PlainAuth("", emailUser, emailPass, server)
	if data.ErrorAmount > 0 {
		msg := fmt.Sprintf("Subject: Apisense report alert\r\n\r\nThe report generated at %s has %d failing test cases.\nConsult your apisense instance for further information.", data.Time.String(), data.ErrorAmount)
		err := smtp.SendMail(fmt.Sprintf("%s:%d", server, port), auth, sender, []string{receiver}, []byte(msg))
		if err != nil {
			return err
		}
	} else if !sendOnErrors {
		msg := fmt.Sprintf("Subject: New apisense report available\r\n\r\nA new apisense report without errors has been generated at %s", data.Time.String())
		err := smtp.SendMail(fmt.Sprintf("%s:%d", server, port), auth, sender, []string{receiver}, []byte(msg))
		if err != nil {
			return err
		}
	}
	return nil
}
