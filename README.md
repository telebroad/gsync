# GSYNC
## gsync is a generic version of the sync.Map in golang

#### Usage
```go
package main

func main()  {
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

```