package logger

import (
	"fmt"
	"github.com/bugsnag/bugsnag-go"
	"runtime"
	"strings"
	"time"
)

const LEVEL = 3

func init() {

}

func Fatal(err error) {
	fmt.Printf("[%s][%s] > FATAL: %s\n", getDateFormat(), getRuntimeLocal(), err.Error())
	bugsnag.Notify(err)
}

func getDateFormat() string {
	return time.Now().Format("02-01-2006 15:04:05")
}

func getRuntimeLocal() string {
	pc := make([]uintptr, 10)
	runtime.Callers(LEVEL, pc)

	funcDirty := runtime.FuncForPC(pc[0])
	f := strings.Split(funcDirty.Name(), "/")

	return f[len(f)-1]
}
