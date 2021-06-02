package e2s

import "log"

// Logger for Fatal
var FatalLogger = log.Default()

// Logger for Warning
var WarningLogger = log.Default()

// Print stacktrace depth
var MaxStackDepth = 200
