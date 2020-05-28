// Package benchmark helper methods
package benchmark

import (
	"reflect"
	"runtime"
)

// GetFunctionName get the running function name of a objectBenchmark.
// ref: https://stackoverflow.com/a/7053871
func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
