package daemon

import (
	"time"
)

// Monitor Structure for ensuring containers are alive.
type Monitor struct {
	Daemon *Daemon
	Container string
	Destroyed bool
}

func (m* Monitor) monitor() {
	for !m.Destroyed {
		time.Sleep(10 * time.Second)
		m.Daemon.Log.Debugf("Checking status of container '%s'", m.Container)
	}
}

// Start Start this monitor.
func (m* Monitor) Start() {
	go m.monitor()
}

// CreateMonitor Create a new monitor for the target container.
func (d* Daemon) CreateMonitor(container string) Monitor {
	monitor := Monitor {
		Daemon: d,
		Container: container,
		Destroyed: false,
	}
	return monitor
}
