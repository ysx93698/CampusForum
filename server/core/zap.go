package core

import (
	"github.com/ysx93698/CampusForum/global"
	"go.uber.org/zap"
)

func Zap() (logger *zap.Logger) {
	global.LOG, _ = zap.NewProduction()
	defer global.LOG.Sync()
	return global.LOG
}
