package utils

import (
	"fmt"
	"reflect"
)

func PrintTypeAndKind[T any](val T) {
	t := reflect.TypeOf(val)
	fmt.Printf("Type: %v, Kind: %v\n", t, t.Kind())
	return
}
