// Go-uptime returns a named structure, containing the current time, system uptime, and load average

package uptime

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

const (
	CurTime int = 0
	User        = 1
	Hours       = 2
	Min         = 3
)

type Up struct {
	CurTime  string    // current time
	Up       string    // system’s uptime
	User     string    // number of logged-in users
	Label    string    // load average label
	LoadAver []float64 // current load average
}

func (up Up) String() string {
	s1 := strconv.FormatFloat(up.LoadAver[0], 'f', -1, 64)
	s2 := strconv.FormatFloat(up.LoadAver[1], 'f', -1, 64)
	s3 := strconv.FormatFloat(up.LoadAver[2], 'f', -1, 64)
	return " " + up.CurTime + " " + up.Up + up.User + " " + up.Label + " " + s1 + ", " + s2 + ", " + s3
}

func Uptime() (ut Up, err error) {
	path, err := exec.LookPath("uptime")
	if err != nil {
		err = fmt.Errorf("exec.LookPath: %v", err)
		return
	}

	out, err := exec.Command(path).Output()
	if err != nil {
		err = fmt.Errorf("exec.Command: %v", err)
		return
	}

	ss := strings.Split(string(out), ",")
	timeUptime := strings.TrimSpace(ss[0])

	x := strings.Split(timeUptime, " ")

	curTime, err := time.Parse("15:04:05", x[CurTime])
	if err != nil {
		fmt.Printf("time.Parse: %v\n", err)
		os.Exit(1)
	}

	ut.CurTime = curTime.Format("15:04:05")
	ut.Up = "up " + x[Hours] + " " + x[Min] + ","
	ut.User = strings.TrimSpace(ss[User] + ",")

	trd := strings.TrimSpace(ss[2])
	v := strings.Split(trd, " ")

	la1 := v[len(v)-1]
	la2 := strings.TrimSpace(ss[3])
	la3 := strings.TrimSpace(ss[4])

	v1, err := strconv.ParseFloat(la1, 64)
	v2, err := strconv.ParseFloat(la2, 64)
	v3, err := strconv.ParseFloat(la3, 64)

	ut.Label = "load average:"
	ut.LoadAver = append(ut.LoadAver, v1, v2, v3)

	return
}
