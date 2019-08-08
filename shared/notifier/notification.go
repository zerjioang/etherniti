package notifier

// abstract internal notification implementation
type Notification struct {
	InternalNotifier
	trigger func() // must contain same definition as Emit method in InternalNotifier
}

func (n Notification) Emit() {
	if n.trigger != nil {
		n.trigger()
	}
}
