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

// the winding number algorithm. 
// This algorithm counts how many times the polygon winds around the origin. 
// If the winding number is non-zero, the origin is inside the polygon.
```
func containsOrigo(points []Point) bool {
    windingNumber := 0

    for i := 0; i < len(points); i++ {
        next := (i + 1) % len(points)
        if points[i].Y <= 0 {
            if points[next].Y > 0 && isLeft(points[i], points[next], Point{0, 0}) > 0 {
                windingNumber++
            }
        } else {
            if points[next].Y <= 0 && isLeft(points[i], points[next], Point{0, 0}) < 0 {
                windingNumber--
            }
        }
    }

    return windingNumber != 0
}
```

func isLeft(p1, p2, p Point) float64 {
    return (p2.X-p1.X)*(p.Y-p1.Y) - (p.X-p1.X)*(p2.Y-p1.Y)
}
