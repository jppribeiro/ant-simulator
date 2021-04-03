package main

import (
	"os"

	"github.com/jppribeiro/ant-simulator/internal/config"
	"github.com/jppribeiro/ant-simulator/internal/simulator"
)

var wTitle string = "Ant Simulator c.JoaoRibeiro"
var wWidth, wHeight int32 = config.WINDOW.X, config.WINDOW.Y

func main() {
	os.Exit(simulator.Run(wTitle, wWidth, wHeight))
}
