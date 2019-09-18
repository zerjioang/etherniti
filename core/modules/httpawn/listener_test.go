package httpawn_test

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"testing"
	"time"

	"github.com/pkg/profile"
	"github.com/zerjioang/etherniti/core/modules/httpawn"
)

// There are 7 places you can get profiles in the default webserver: the ones mentioned above
//
// http://localhost:6060/debug/pprof/
// http://localhost:6060/debug/pprof/goroutine
// http://localhost:6060/debug/pprof/heap
// http://localhost:6060/debug/pprof/threadcreate
// http://localhost:6060/debug/pprof/block
// http://localhost:6060/debug/pprof/mutex
//
// and also 2 more: the CPU profile and the CPU trace.
//
// http://localhost:6060/debug/pprof/profile?seconds=15
// http://localhost:6060/debug/pprof/trace?seconds=15
//
// run in the web
//
// go tool pprof -http=localhost:6061 profile.out
//
// go tool pprof  http://localhost:6060/debug/pprof/heap
// go tool pprof -png http://localhost:6060/debug/pprof/heap > out.png
// http://localhost:6060/debug/pprof/profile?seconds=30
// go tool pprof -http=:6060 heap.out
func TestServe(t *testing.T) {
	t.Run("serve-test", func(t *testing.T) {
		server := httpawn.New()
		server.GET("/", func(ctx *httpawn.Context) {
			ctx.String("Hello World!")
		})
		server.Start(":8080")
	})
	t.Run("serve-test-profiling-server", func(t *testing.T) {
		// create a locking channel for notification
		notif := make(chan bool)
		// we need a webserver to get the pprof webserver
		go func() {
			notif <- true
			log.Println(http.ListenAndServe("localhost:6060", nil))
		}()
		<-notif
		log.Println("running httpawn code with profiling enabled")
		server := httpawn.New()
		server.GET("/", func(ctx *httpawn.Context) {
			ctx.String("Hello World!")
		})
		server.Start(":8080")
	})
	t.Run("serve-test-profiling-cpu-file", func(t *testing.T) {
		// CPU Profile
		p := profile.Start()
		// run the code in a goroutine
		go func() {
			// run the code
			log.Println("running httpawn code with profiling enabled")
			server := httpawn.New()
			server.GET("/", func(ctx *httpawn.Context) {
				ctx.String("Hello World!")
			})
			server.Start(":8080")
		}()
		// wait 10 seconds to collect results
		time.Sleep(time.Second * 10)
		// p.Stop() must be called before the program exits to
		// ensure profiling information is written to disk.
		t.Log("stopping cpu profiler...")
		p.Stop()
		t.Log("stopped")
	})
	t.Run("serve-test-profiling-mem-file", func(t *testing.T) {
		// Memory Profile
		runtime.GC()
		p := profile.Start(profile.MemProfile, profile.ProfilePath("."), profile.NoShutdownHook, profile.MemProfileRate(100))
		// run the code in a goroutine
		go func() {
			// run the code
			log.Println("running httpawn code with profiling enabled")
			server := httpawn.New()
			server.GET("/", func(ctx *httpawn.Context) {
				ctx.String("Hello World!")
			})
			server.Start(":8080")
		}()
		// wait 10 seconds to collect results
		time.Sleep(time.Second * 10)
		// p.Stop() must be called before the program exits to
		// ensure profiling information is written to disk.
		t.Log("stopping mem profiler...")
		p.Stop()
		t.Log("stopped")
	})
}
