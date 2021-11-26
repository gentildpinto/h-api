package app

import (
	"github.com/gentildpinto/h-api/internal/tools/logger"
)

var (
	appPort = "80"
)

func Run(appVersion string) (err error) {
	if err = logger.Initialize(appVersion); err != nil {
		return
	}
	return
}
