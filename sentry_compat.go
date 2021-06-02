package e2s

import (
	"github.com/getsentry/sentry-go"
)

func isE2sFrame(frame sentry.Frame) bool {
	return frame.Module == "git.dev.a1fred.com/golibs/err2sentry%2egit"
}

// Cleanup e2s's frames from sentry events
func SentryBeforeSend(event *sentry.Event, hint *sentry.EventHint) *sentry.Event {
	for _, exc := range event.Exception {
		newFrames := make([]sentry.Frame, 0)
		if exc.Stacktrace != nil {
			for _, frame := range exc.Stacktrace.Frames {
				if !isE2sFrame(frame) {
					newFrames = append(newFrames, frame)
				}
			}
			exc.Stacktrace.Frames = newFrames
		}
	}
	return event
}
