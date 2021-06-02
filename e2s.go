package e2s

import (
	"os"
)

// Send err to Sentry, log stacktrace, then exit
func Fatal(err error) {
	write(err, FatalLogger)
	os.Exit(1)
}

// Send err to Sentry, log stacktrace
func Warning(err error) {
	write(err, WarningLogger)
}

// If err != nil: Send err to Sentry, log stacktrace, then exit
func IfErrFatal(err error) {
	if err != nil {
		write(err, FatalLogger)
		os.Exit(1)
	}
}

// If err != nil: Send err to Sentry, log stacktrace
func IfErrWarning(err error) {
	if err != nil {
		write(err, WarningLogger)
	}
}

// If not ok: Send err to Sentry, log stacktrace, then exit
func NotOkFatal(ok bool, err error) {
	if !ok {
		write(err, FatalLogger)
		os.Exit(1)
	}
}

// If not ok: Send err to Sentry, log stacktrace
func NotOkWarning(ok bool, err error) {
	if !ok {
		write(err, WarningLogger)
	}
}

// If ok call: Send err to Sentry, log stacktrace, then exit
func OkFatal(ok bool, err error) {
	if ok {
		write(err, FatalLogger)
		os.Exit(1)
	}
}

// If ok call: Send err to Sentry, log stacktrace
func OkWarning(ok bool, err error) {
	if ok {
		write(err, WarningLogger)
	}
}
