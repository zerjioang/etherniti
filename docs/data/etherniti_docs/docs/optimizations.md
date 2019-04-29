# Optimizations

## A bunch of benchmarks for basic Go things

### Headlines

* Locking and unlocking an uncontended mutex takes ~25 ns
* Allocating some memory takes ~30 ns
* Using a sync.Pool takes ~25ns
* deferring a function takes ~90 ns
* Pushing a byte through a channel takes ~100 to 250 ns
* Checking if something is in a map, then adding it takes ~250 ns if the map has enough room
* Checking if something is in an unordered slice, then appending it takes ~222 ns up to around 100 items if the slice has cpacity
* A type assertion takes ~0.90 ns
* Storing a value in an interface type takes ~33 ns and involves an allocation, even for relatively small types
* Storing an existing pointer in an interface type takes ~8ns and no allocations.
* It's slightly faster to do a type assertion on an interface type then a call on the concrete type (1.1 ns), than to call a method on the interface type directly (2.3 ns)


To try and still break down a Go binary to its dependencies, we must use a Go-enlightened tool that can understand the Go binary format. Let’s find one.

```bash
go tool nm -sort size -size etherniti | head -n 20
```

## Gofat

There’s one last trick that will work. When you compile your Go binary, Go will generate interim binaries for each dependency, before statically linking these all up into the one binary you get in the end.

Introducing gofat — a shell script that’s a mix of Go and some Unix tools that analyzes a Go binary dependencies sizes:

```bash
eval `go build -work -a 2>&1` && find $WORK -type f -name "*.a" | xargs -I{} du -hxs "{}" | gsort -rh | sed -e s:${WORK}/::g
```

If you’re in a hurry, just copy or download the above shell script and set it to be executable (chmod +x). Then, run it in your project’s directory with no arguments in order to get that project’s dependency breakdown.

## alloc_space vs inuse_space

go tool pprof has the option to show you either allocation counts or in use memory. If you’re concerned with the amount of memory being used, you probably want the inuse metrics, but if you’re worried about time spent in garbage collection, look at allocations!

  -inuse_space      Display in-use memory size
  -inuse_objects    Display in-use object counts
  -alloc_space      Display allocated memory size
  -alloc_objects    Display allocated object counts

## Browser inspection

```bash
wget http://localhost:6060/debug/pprof/trace\?seconds\=5 && go tool trace pprof
```
