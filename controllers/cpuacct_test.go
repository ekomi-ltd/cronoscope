package controllers

import (
	"testing"
)

func TestNewCPUAcctController(t *testing.T) {

	assertStringEqual := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("Got %s wanted %s", got, want)
		}
	}

	t.Run("Should initialize fields properly", func(t *testing.T){
		controller := NewCPUAcctController()
		assertStringEqual(t, controller.path, "/sys/fs/cgroup/cpuacct/cpuacct.stat")
		assertStringEqual(t, controller.metricType, "gauge")
	})
}
