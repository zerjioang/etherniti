package bench

import (
	"github.com/zerjioang/etherniti/core/modules/monotonic"
	"math"
	"math/rand"
	"runtime"
	"sync"
	"time"

	"github.com/zerjioang/etherniti/core/logger"
)

const (
	// 30 milion samples
	samples = 30000000
)

var (
	// runtime benchmark execution score
	runScore int64
	// total runtime
	totalTime time.Duration
	// calculated flag
	calculated bool
)

func init() {
	logger.Debug("loading internal p.o.s.t bencharking")
	calculateScore()
}

func calculateScore() {
	if !calculated {
		logger.Debug("benchmarking current server configuration")
		logger.Debug("getting benchmark (multicore) score")

		cores := runtime.NumCPU()
		runtime.GOMAXPROCS(cores)

		logger.Debug("calculating score using all server cores: ", cores)

		var wait sync.WaitGroup

		counts := make([]int, cores)

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

		//pi := (float64(total) / float64(samples)) * 4
		totalTime = monotonic.Since(start)
		score := float64(samples) / totalTime.Seconds()
		runScore = int64(score)
		calculated = true
	}
}

func GetScore() int64 {
	return runScore
}

func GetBenchTime() time.Duration {
	return totalTime
}

func monteCarlo(radius float64, reps int, result *int, wait *sync.WaitGroup) {
	var x, y float64
	count := 0
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)

	for i := 0; i < reps; i++ {
		x = random.Float64() * radius
		y = random.Float64() * radius

		if num := math.Sqrt(x*x + y*y); num < radius {
			count++
		}
	}
	*result = count
	wait.Done()
}
