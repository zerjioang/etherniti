package notifier

type InternalNotifier interface {
	// a simple function with no arguments that
	// is uses to monitor internal components analytics
	// and usage
	Emit()
}
