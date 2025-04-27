package main

import "github.com/steverusso/cli"

func main() {
	_ = cli.NewCmd("recqueue-server").
		Help("recommendation queue server").
		Build().
		ParseOrExit()
}
