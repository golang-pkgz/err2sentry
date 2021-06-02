# Handle errors with sentry
[![Go](https://github.com/golang-pkgz/err2sentry/actions/workflows/go.yml/badge.svg)](https://github.com/golang-pkgz/err2sentry/actions/)

[Go SDK](https://docs.sentry.io/platforms/go/)

## Install
```
go get github.com/golang-pkgz/err2sentry
```

## Usage
Instead of writing
```golang
err := someFunc()
if err != nil {
    log.Fatal(err)
}
```

You can write
```golang
<initialize sentry>

err := someFunc()
e2s.IfErrFatal(err)
```

## Docs
```shell
$ go doc -all
package e2s // import "github.com/golang-pkgz/err2sentry"


VARIABLES

var MaxStackDepth = 200
    Print stacktrace depth

var FatalLogger = log.Default()
    Logger for Fatal

var WarningLogger = log.Default()
    Logger for Warning


FUNCTIONS

func Fatal(err error)
    Send err to Sentry, log stacktrace, then exit

func Warning(err error)
    Send err to Sentry, log stacktrace

func IfErrFatal(err error)
    If err != nil: Send err to Sentry, log stacktrace, then exit

func IfErrWarning(err error)
    If err != nil: Send err to Sentry, log stacktrace

func NotOkFatal(ok bool, err error)
    If not ok: Send err to Sentry, log stacktrace, then exit

func NotOkWarning(ok bool, err error)
    If not ok: Send err to Sentry, log stacktrace

func OkFatal(ok bool, err error)
    If ok call: Send err to Sentry, log stacktrace, then exit

func OkWarning(ok bool, err error)
    If ok call: Send err to Sentry, log stacktrace

func SentryBeforeSend(event *sentry.Event, hint *sentry.EventHint) *sentry.Event
    Cleanup e2s's frames from sentry events
```

### Sentry integration
Use `e2s.SentryBeforeSend` hook to cleanup e2s's frames
```golang
err := sentry.Init(sentry.ClientOptions{
    ...
    BeforeSend:  e2s.SentryBeforeSend,
})
```
