package common

import (
	"fmt"
	"runtime/debug"
	"strings"
)

func GetKey(prefix string, items ...interface{}) string {
	format := prefix + strings.Repeat(":%v", len(items))
	return fmt.Sprintf(format, items...)
}

func CatchPanic() {
	if err := recover(); err != nil {
		Logger.Errorf("catch panic | %s\n%s", err, debug.Stack())
	}
}
