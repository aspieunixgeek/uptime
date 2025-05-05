package uptime

import (
	"testing"
	"time"
)

func TestUptime(t *testing.T) {
	now := time.Now()
	u, err := Uptime()
	if err != nil {
		t.Errorf("Uptime: %v\n", err)
	}

	if u.CurTime != time.Now().Format("15:04:05") {
		t.Fatalf("u.CurTime(%s) != time.Now(%s)", u.CurTime, now)
	}
	if len(u.Up) == 0 {
		t.Fatal("len(u.Up) == 0\n")
	}

	if len(u.LoadAver) == 0 {
		t.Fatal("len(u.LoadAver) == 0\n")
	}

	curTime := u.CurTime     // current time
	up := u.Up               // uptime
	loadAverLabel := u.Label // load average label
	loadAver := u.LoadAver   // load average

	t.Logf("curTime: %v\n", curTime)
	t.Logf("uptime: %v\n", up)
	t.Logf("%v %v\n", loadAverLabel, loadAver)
}
