package main

import (
	"github.com/buonotti/apisense/cmd"
	"github.com/buonotti/apisense/config"
)

func main() {
	config.GetAsset = Asset
	cmd.Execute()
}