package gsync

import (
	"fmt"
	"sync"
	"testing"
)

func TestMap(t *testing.T) {
	SyncedMap := Map[string, int]{}
	wg := sync.WaitGroup{}
	for i := range 100 {
		wg.Add(1)
		go func(i int) {
			SyncedMap.Store(fmt.Sprint(i), i)
			wg.Done()
		}(i)
	}
	wg.Wait()
	t.Logf(SyncedMap.String())
	SyncedMap.Clear()
	t.Logf(SyncedMap.String())
}

func ExampleMap() {
	SyncedMap := Map[string, int]{}
	wg := sync.WaitGroup{}
	for i := range 100 {
		wg.Add(1)
		go func(i int) {
			SyncedMap.Store(fmt.Sprint(i), i)
			wg.Done()
		}(i)
	}
	wg.Wait()
	// Output:
	fmt.Println(SyncedMap.String())
	SyncedMap.Clear()
	fmt.Println(SyncedMap.String())
}
