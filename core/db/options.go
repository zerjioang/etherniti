package db

import (
	"time"

	"github.com/dgraph-io/badger"
)

// Options contains all the configuration used to open the Badger db
type Options struct {
	// Path is the directory path to the Badger db to use.
	Path string

	// BadgerOptions contains any specific Badger options you might
	// want to specify.
	BadgerOptions *badger.Options

	// NoSync causes the database to skip fsync calls after each
	// write to the log. This is unsafe, so it should be used
	// with caution.
	NoSync bool

	// ValueLogGC enables a periodic goroutine that does a garbage
	// collection of the value log while the underlying Badger is online.
	ValueLogGC bool

	// GCInterval is the interval between conditionally running the garbage
	// collection process, based on the size of the vlog. By default, runs every 1m.
	GCInterval time.Duration

	// GCInterval is the interval between mandatory running the garbage
	// collection process. By default, runs every 10m.
	MandatoryGCInterval time.Duration

	// GCThreshold sets threshold in bytes for the vlog size to be included in the
	// garbage collection cycle. By default, 1GB.
	GCThreshold int64
}

// New uses the supplied options to open the Badger db and prepare it for
// use as a raft backend.
func NewBadgerStorageGc(options *Options) (*BadgerStorage, error) {

	// build badger options
	if options.BadgerOptions == nil {
		defaultOpts := badger.DefaultOptions
		options.BadgerOptions = &defaultOpts
	}
	options.BadgerOptions.Dir = options.Path
	options.BadgerOptions.ValueDir = options.Path
	options.BadgerOptions.SyncWrites = !options.NoSync

	// try to create new database handler
	storage, err := NewCollection("")
	if err != nil {
		return nil, err
	}
	storage.options = options

	// Start GC routine
	if options.ValueLogGC {

		var gcInterval time.Duration
		var mandatoryGCInterval time.Duration
		var threshold int64

		if gcInterval = 1 * time.Minute; options.GCInterval != 0 {
			gcInterval = options.GCInterval
		}
		if mandatoryGCInterval = 10 * time.Minute; options.MandatoryGCInterval != 0 {
			mandatoryGCInterval = options.MandatoryGCInterval
		}
		if threshold = int64(1 << 30); options.GCThreshold != 0 {
			threshold = options.GCThreshold
		}

		storage.vlogTicker = time.NewTicker(gcInterval)
		storage.mandatoryVlogTicker = time.NewTicker(mandatoryGCInterval)
		go storage.runVlogGC(storage.instance, threshold)
	}

	return storage, nil
}
