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
hashMap := make(map[string]int)
myMap["A"] = 25
value, exists := hashMap["A"]
isEmpty := len(hashMap) == 0
for key, value := range myMap {
        fmt.Printf("%s -> %d\n", key, value)
}
toSlice := make([]int, 0, len(s.data))
    for key := range s.data {
        result = append(result, key)
}
delete(hashMap, "A")


```
