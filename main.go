package main

import (
	"github.com/buonotti/odh-data-monitor/cmd"
	"github.com/buonotti/odh-data-monitor/config"
)

func main() {
	config.GetAsset = Asset
	cmd.Execute()
}
