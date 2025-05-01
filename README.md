# jngo
Basic Go data structures, algorithms and helper functions

map/filter/reduce (fold)

# Test
```
go build -v ./...
go test -v ./...
```

```
type Point struct {
    X, Y float64
}
```

```
hashmap := make(map[string]int)
hashmap["A"] = 25
value, exists := hashmap["A"]
isEmpty := len(hashmap) == 0
for key, value := range hashmap {
        fmt.Printf("%s -> %d\n", key, value)
}
toSlice := make([]int, 0, len(s.data))
    for key := range s.data {
        result = append(result, key)
}
delete(hashmap, "A")
```

RWMutex (instead of sync.Mutex) allows multiple readers to access the resource concurrently, while ensuring that only one writer can access it at any time. Use in read-heavy scenarios.
m.mu.RLock() allows multiple readers
m.mu.Lock() allows only one writer



