package udp

import (
	"fmt"

	"github.com/fsvxavier/unico/pkg/logger"
)

func ErrorCheck(err error, where string, kill bool) {
	var logger logger.GenericLogger
	if err != nil {
		if kill {
			logger.LogIt("FATAL", "Script Terminated...", nil)
		} else {
			logger.LogIt("WARN", fmt.Sprintf("@ %s\n", where), nil)
		}
	}
}
