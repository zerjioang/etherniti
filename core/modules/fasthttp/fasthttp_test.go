package fasthttp

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"testing"
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
		_ = Serve("localhost:3333")
	})
	t.Run("serve-test-profile", func(t *testing.T) {
		// create a locking channel for notification
		notif := make(chan bool)
		// we need a webserver to get the pprof webserver
		go func() {
			notif <- true
			log.Println(http.ListenAndServe("localhost:6060", nil))
		}()
		<- notif
		log.Println("running fasthttp code with profiling enabled")
		_ = Serve("localhost:3333")
	})
}
