package utils

import (
	"fmt"
	"reflect"
	"sync"
)

func PrintTypeAndKind[T any](val T) {
	t := reflect.TypeOf(val)
	fmt.Printf("Type: %v, Kind: %v\n", t, t.Kind())
	return
}

func ParallelTreatment(start int, stop int, fn func(<-chan int)) {
	count := stop - start
	if count < 1 {
		return
	}
	c := make(chan int, count)
	for i := start; i < stop; i++ {
		c <- i
	}
	close(c)
	var wg sync.WaitGroup
	for i := 0; i < count; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fn(c)
		}()
	}
	wg.Wait()
}
