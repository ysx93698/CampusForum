package main

import (
	"github.com/ysx93698/CampusForum/core"
	"github.com/ysx93698/CampusForum/global"
)

func main() {
	println("Start")
	global.LOG = core.Zap()
	global.LOG.Info("Testing")
	println("End")
}
