package bench

import (
	"fmt"
	"runtime"
	"sync"
	"testing"

	"github.com/zerjioang/etherniti/core/modules/monotonic"
)

func TestMonteCarlo(t *testing.T) {
	cores := runtime.NumCPU()
	runtime.GOMAXPROCS(cores)

	var wait sync.WaitGroup

	counts := make([]int, cores)

	// 30 million samples
	samples := 30000000

	start := monotonic.Now()
	wait.Add(cores)

	for i := 0; i < cores; i++ {
		go monteCarlo(100.0, samples/cores, &counts[i], &wait)
	}

	wait.Wait()

	total := 0
	for i := 0; i < cores; i++ {
		total += counts[i]
	}

	pi := (float64(total) / float64(samples)) * 4
	totalt := monotonic.Since(start)
	score := float64(samples) / totalt.Seconds()

	fmt.Println("Time: ", totalt)
	fmt.Println("Score: ", int64(score))
	fmt.Println("pi: ", pi)
	fmt.Println("")
}
