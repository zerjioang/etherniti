package notifier

import "github.com/zerjioang/etherniti/core/bus"

const (
	PowerOffEvent       = "poweroff"
	NewProxyRequest     = "new_proxy_request"
	NewDashboardAccount = "new_dashboard_account"
	NewProfileRequest   = "new_profile_request"
)

var (
	// implements InternalNotifier
	NewProxyRequestEvent InternalNotifier = Notification{
		trigger: func() {
			bus.Emit(NewProxyRequest)
		},
	}
	// implements InternalNotifier
	NewProfileRequestEvent InternalNotifier = Notification{
		trigger: func() {
			bus.Emit(NewProfileRequest)
		},
	}
	// implements InternalNotifier
	NewDashboardAccountEvent InternalNotifier = Notification{
		trigger: func() {
			bus.Emit(NewDashboardAccount)
		},
	}
)
