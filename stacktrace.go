package e2s

import (
	"fmt"
	"log"
	"net/url"
	"runtime"
	"strings"
)

func packageAndName(fn *runtime.Func) (string, string) {
	name := fn.Name()
	pkg := ""

	if lastslash := strings.LastIndex(name, "/"); lastslash >= 0 {
		pkg += name[:lastslash] + "/"
		name = name[lastslash+1:]
	}
	if period := strings.Index(name, "."); period >= 0 {
		pkg += name[:period]
		name = name[period+1:]
	}

	name = strings.Replace(name, "·", ".", -1)

	decodedPkg, err := url.QueryUnescape(pkg)
	if err != nil {
		fmt.Printf("decodeerr: %v", err.Error())
		return pkg, name
	}

	return decodedPkg, name
}

func printStackTraceWithError(err error, skip int, logger *log.Logger) {

	stack := make([]uintptr, MaxStackDepth)
	_ = runtime.Callers(1+skip, stack[:])

	for i := len(stack) - 1; i >= 0; i-- {
		if stack[i] != 0 {
			funcForPC := runtime.FuncForPC(stack[i])

			packageName, name := packageAndName(funcForPC)
			file, lineNumber := funcForPC.FileLine(stack[i] - 1)

			log.Printf("%v.%v\n\tat %s:%d (0x%x)", packageName, name, file, lineNumber, stack[i])
		}
	}

	log.Printf("Error: %v", err.Error())
}
