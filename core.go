package e2s

import (
	"log"

	"github.com/getsentry/sentry-go"
)

func write(err error, logger *log.Logger) {
	sentry.CaptureException(err)
	printStackTraceWithError(err, 3, logger)
}
