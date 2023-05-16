package main

import (
	"github.com/buonotti/apisense/cmd"
)

//	@title			Apisense API
//	@version		1.0
//	@description	Api specification for the Apisense API

//	@contact.name	buonotti
//	@contact.url	https://github.com/buonotti/apisense/issues

//	@license.name	MIT
//	@license.url	https://opensource.org/licenses/MIT

//	@host						localhost:8080
//	@BasePath					/api
//
//	@securityDefinitions.bearer	ApiKeyAuth
//	@in							header
//	@name						api_key
func main() {
	cmd.Execute()
}
