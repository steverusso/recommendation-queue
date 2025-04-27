package main

import (
	"github.com/steverusso/cli"
	rq "github.com/steverusso/recommendation-queue"
)

var _ = rq.Recommendation

func main() {
	_ = cli.NewCmd("recqueue-cli").
		Help("recommendation queue cli").
		Build().
		ParseOrExit()
}
