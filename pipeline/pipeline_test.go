package pipeline

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/jnsoft/jngo/randx"
	. "github.com/jnsoft/jngo/testhelper"
)

const FNAME = "delteme.txt"

type SampleData struct {
	Key   string
	Value float64
}

func TestManualPipeline(t *testing.T) {
	source := SliceSource([]int{1, 2, 3, 4, 5})

	Pipeline(source, []Stage[int]{
		TransformToStage(MapStage(func(n int) (int, error) {
			return n * 2, nil
		})),
		CollectStage(func(n int) { fmt.Println("Doubled:", n) }),
	}, func() {
		fmt.Println("Manual pipeline complete.")
	})
}

func TestSlicePipeline(t *testing.T) {
	input := []string{"a", "b", "c"}
	collected := []string{}

	pb := FromSource(SliceSource(input))

	pb2 := Then(pb, MapStage(func(s string) (string, error) {
		return strings.ToUpper(s), nil
	}))

	pb3 := Finally(pb2, func(result string) {
		collected = append(collected, result)
	})

	Run(pb3, nil)

	expected := []string{"A", "B", "C"}

	if len(collected) != len(expected) {
		t.Fatalf("expected %v items, got %v", len(expected), len(collected))
	}

	for i, v := range collected {
		if v != expected[i] {
			t.Errorf("expected %q, got %q", expected[i], v)
		}
	}
}

func TestUppercaseStrings(t *testing.T) {
	items := []string{"hello", "world", "golang"}

	pb := FromSource(SliceSource(items))

	pb2 := Then(pb, ParallelMapStage(4, func(s string) (string, error) {
		return strings.ToUpper(s), nil
	}))

	pb3 := Finally(pb2, func(result string) {
		fmt.Println("Uppercased:", result)
	})

	Run(pb3, func() {
		fmt.Println("Done with uppercase!")
	})
}

func TestFileLineCountPipeline(t *testing.T) {

	if err := generateTestFile(FNAME, 1_000); err != nil {
		panic(err)
	}
	defer func() {
		if err := os.Remove(FNAME); err != nil {
			t.Errorf("failed to delete test file: %v", err)
		}
	}()

	source := FileLineSource(FNAME)

	pb := FromSource(source)
	pb2 := Then(pb, MapStage(func(s string) (int, error) {
		return len(s), nil
	}))

	var totalLength int64
	var lines int64

	pb3 := ThenDo(pb2, func(n int) {
		atomic.AddInt64(&lines, 1)
	})
	pb4 := Finally(pb3, func(n int) {
		atomic.AddInt64(&totalLength, int64(n))
	})

	Run(pb4, func() {
		fmt.Printf("Total line length: %d\n", totalLength)
		fmt.Printf("File processing complete (%d lines).\n", lines)
	})

	AssertEqual(t, lines, 1000)

}

func TestPipelineBuilder(t *testing.T) {

	if err := generateTestFile(FNAME, 1_000); err != nil {
		t.Fatalf("failed to generate test file: %v", err)
	}
	defer func() {
		if err := os.Remove(FNAME); err != nil {
			t.Errorf("failed to delete test file: %v", err)
		}
	}()

	counter := make(map[string]int)
	var mu sync.Mutex

	start := time.Now()

	pb := FromSource(FileLineSource(FNAME))

	pb2 := Then(pb, ParallelMapStage[string, SampleData](8, parseLine))

	pb3 := ThenDo(pb2, func(data SampleData) {
		mu.Lock()
		counter[data.Key]++
		mu.Unlock()
	})

	Finally(pb3, func(data SampleData) {
		// drain if needed
	})

	Run(pb3, func() {
		fmt.Printf("Processed %d keys in %v\n", len(counter), time.Since(start))
	})

	if len(counter) == 0 {
		t.Fatal("expected non-zero keys in counter")
	}
}

// helpers

func generateTestFile(path string, lines int) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	for i := 0; i < lines; i++ {
		_, _ = fmt.Fprintf(f, "sensor-%d;%.2f\n", randx.Int(0, 999), randx.Float64()*100)
	}
	return nil
}

func parseLine(line string) (SampleData, error) {
	parts := strings.Split(line, ";")
	if len(parts) != 2 {
		return SampleData{}, fmt.Errorf("invalid line: %s", line)
	}
	val, err := strconv.ParseFloat(parts[1], 64)
	if err != nil {
		return SampleData{}, err
	}
	return SampleData{Key: parts[0], Value: val}, nil
}
