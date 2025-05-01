package profiling

import (
	"fmt"
	"os"
	"runtime/pprof"
	"time"
)

func TimeFunction(label string, f func() (any, error)) {
	start := time.Now()
	result, err := f()
	elapsed := time.Since(start)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("%s: %v (%s)\n", label, result, elapsed)
}

// Usage: misc.ProfileFunction("Comment", <fname>, func() (any, error) { return MyFunc(), nil })
// View profile: go tool pprof -http 127.0.0.1:8080 ./<fname>
func ProfileFunction(label, prof_file string, f func() (any, error)) {
	pfile, err := os.Create(prof_file)
	if err != nil {
		panic(err)
	}
	defer pfile.Close()

	if err := pprof.StartCPUProfile(pfile); err != nil {
		panic(err)
	}
	defer pprof.StopCPUProfile()

	start := time.Now()
	result, err := f()
	elapsed := time.Since(start)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("%s: %v (%s)\n", label, result, elapsed)

}
