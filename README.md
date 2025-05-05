#### go-uptime

Go-uptime returns a named structure, containing the current time, system uptime, and load average.

Usage example:

---

Install dependencies:
```bash
$ go get github.com/aspieunixgeek/uptime
go: downloading github.com/aspieunixgeek/uptime v0.2.3
```

---

Import into the app:
```go

package main

import(
	"fmt"
	"os"
	"github.com/aspieunixgeek/uptime"
)

func main() {
	u, err := uptime.Uptime()
	if err != nil {
		fmt.Fprintf(os.Stderr, "uptime.Uptime: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("%v\n", u)
	// Example output:
	//
	// 12:09:18 up 34 min,  1 user,  load average: 0.46, 0.73, 0.76
}
```

```go

package main

import(
	"fmt"
	"os"
	"github.com/aspieunixgeek/uptime"
)

func main() {
	u, err := uptime.Uptime()
	if err != nil {
		fmt.Fprintf(os.Stderr, "uptime.Uptime: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("%v\n", u)
	// Example output:
	//
	// 12:09:18 up 34 min,  1 user,  load average: 0.46, 0.73, 0.76

	fmt.Printf("%s\n", u.CurTime)
	fmt.Printf("%s\n", u.Up)
	fmt.Printf("%s\n", u.User)
	fmt.Printf("%s %.2f %.2f %.2f\n", u.Label, u.LoadAver[0], u.LoadAver[1], u.LoadAver[2])
}
```

---

Unit tests output:

```bash
$ go test -v
=== RUN   TestUptime
    uptime_test.go:31: 14:08:50
    uptime_test.go:32: up 30 min,
    uptime_test.go:33: load average: [0.6 0.67 0.65]
--- PASS: TestUptime (0.00s)
PASS
ok      github.com/aspieunixgeek/uptime 0.004s
```